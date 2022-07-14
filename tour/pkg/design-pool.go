package pool

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrClosed = errors.New("Pool is closed")
)

/*

梳理需求

	问题#1： 连接池有什么功能和接口？
		Get 获取一个连接
		Put 连接放回去
		Close 释放/关闭连接

*/

type Pool interface {
	Get() (interface{}, error)
	Put(interface{}) error
	Close(interface{}) error
	Release()
	Len() int
}

var (
	ErrMaxActiveConnReached = errors.New("MaxActiveConnReached")
)

// 问题#2： 连接池的连接从哪来

type ConnectionFactory interface {
	Factory() (interface{}, error)
	Close(interface{}) error
	Ping(interface{}) error
}

// 连接池相关配置
type PoolConfig struct {
	InitialCap  int               // 连接池的最小连接数
	MaxCap      int               // 最大并发存活连接数
	MaxIdle     int               // 最大空闲连接
	Factory     ConnectionFactory // 工厂
	IdleTimeout time.Duration     // 连接最大空闲时间，超过该事件则将失效
}

type connReq struct {
	idleConn *idleConn
}

// channel Pool 存放连接信息
type channelPool struct {
	mu                       sync.RWMutex
	conns                    chan *idleConn    // 存储最大空闲连接
	factory                  ConnectionFactory // 连接制造工厂
	idleTimeout, waitTimeOut time.Duration     // 空闲超时和等待超时
	maxActive                int               // 最大连接数
	openingConns             int               // 活跃的连接数
	connReqs                 []chan connReq    // 连接请求缓冲区，如无法从 conns 取到连接，则在该缓冲区创建一个新的元素，之后连接放回去时先填充这个缓冲区
}

type idleConn struct {
	conn interface{}
	t    time.Time
}

// NewChannelPool 初始化连接
func NewChannelPool(poolConfig *PoolConfig) (Pool, error) {

	// 初始化连接数、空闲连接数、最大连接数、最大空闲数
	if !(poolConfig.InitialCap <= poolConfig.MaxIdle && poolConfig.MaxCap >= poolConfig.MaxIdle && poolConfig.InitialCap >= 0) {
		return nil, errors.New("invalid capacity settings")
	}

	// 校验参数，连接工厂不能为空
	if poolConfig.Factory == nil {
		return nil, errors.New("iunvalid factory interface settings")
	}

	// 连接池
	c := &channelPool{
		conns:        make(chan *idleConn, poolConfig.MaxIdle),
		factory:      poolConfig.Factory,
		idleTimeout:  poolConfig.IdleTimeout,
		maxActive:    poolConfig.MaxCap,
		openingConns: poolConfig.InitialCap,
	}

	// 初始化连接，放入 channel 中
	for i := 0; i < poolConfig.InitialCap; i++ {
		conn, err := c.factory.Factory()
		if err != nil {
			c.Release()
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		c.conns <- &idleConn{conn: conn, t: time.Now()}
	}
	return c, nil
}

// getConns 获取所有连接
func (c *channelPool) getConns() chan *idleConn {
	c.mu.Lock()
	conns := c.conns
	c.mu.Unlock()
	return conns
}

// Get 从 pool 中取一个连接
func (c *channelPool) Get() (interface{}, error) {
	conns := c.getConns()
	if conns == nil {
		return nil, ErrClosed
	}
	for {
		select {
		case wrapConn := <-conns:
			if wrapConn == nil {
				return nil, ErrClosed
			}
			// 判断是否超时，超时则丢弃
			if timeout := c.idleTimeout; timeout > 0 {
				if wrapConn.t.Add(timeout).Before(time.Now()) {
					// 丢弃并关闭连接
					_ = c.Close(wrapConn.conn)
					continue
				}
			}
			return wrapConn.conn, nil
		default:
			c.mu.Lock()
			log.Printf("openConn %v %v", c.openingConns, c.maxActive)
			if c.openingConns >= c.maxActive {
				req := make(chan connReq, 1)
				c.connReqs = append(c.connReqs, req)
				c.mu.Unlock()
				ret, ok := <-req
				if !ok {
					return nil, ErrMaxActiveConnReached
				}
				if timeout := c.idleTimeout; timeout > 0 {
					if ret.idleConn.t.Add(timeout).Before(time.Now()) {
						_ = c.Close(ret.idleConn.conn)
						continue
					}
				}
				return ret.idleConn.conn, nil
			}
			if c.factory == nil {
				c.mu.Unlock()
				return nil, ErrClosed
			}
			conn, err := c.factory.Factory()
			if err != nil {
				c.mu.Unlock()
				return nil, err
			}
			c.openingConns++
			c.mu.Unlock()
			return conn, nil
		}
	}
}

// Put 将连接放回 pool 中，这里会影响到 NewChannelPool 中的 retun c, nil
func (c *channelPool) Put(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conns == nil {
		return c.Close(conn)
	}

	if l := len(c.connReqs); l > 0 {
		req := c.connReqs[0]
		copy(c.connReqs, c.connReqs[1:])
		c.connReqs = c.connReqs[:l-1]
		req <- connReq{
			idleConn: &idleConn{conn: conn, t: time.Now()},
		}
		return nil
	}
	select {
	case c.conns <- &idleConn{conn: conn, t: time.Now()}:
		return nil
	default:
		return c.Close(conn)
	}
}

// Close 关闭单条连接
func (c *channelPool) Close(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.openingConns--
	return c.factory.Close(conn)
}

// Ping 检查单条连接是否有效
func (c *channelPool) Ping(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	return c.factory.Ping(conn)
}

// 释放连接
func (c *channelPool) Release() {
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.mu.Unlock()
	defer func() {
		c.factory = nil
	}()

	if conns == nil {
		return
	}

	// clsoe a channel
	close(conns)

	for wrapConn := range conns {
		_ = c.factory.Close(wrapConn.conn)
	}
}

// Len 连接池中已有的连接
func (c *channelPool) Len() int {
	return len(c.getConns())
}
