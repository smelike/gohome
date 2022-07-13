package router

import (
	"net/http"
	api "sulink/app/api"

	"github.com/gin-gonic/gin"
)

var noNeedLogin = []string{
	"/user/getAuthUrl",
	"/user/wxLogin",
	"/user/notify",
	"/board/get",
	"/activity/getOrder",
}

//跨域访问：cross  origin resource share
func crossHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//接收客户端发送的origin （重要！）
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//允许跨域设置可以返回其他子段，可以自定义字段
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length")
		// 允许浏览器（客户端）可以解析的头部 （重要）
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		//设置缓存时间
		c.Header("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}
		c.Next()
	}
}

/*
//请求时间间隔限制
//duration有效期，单位秒
func requestLimit(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		action := c.FullPath()

		//获取上次请求时间
		userId := user.(*controllers.Claims).UserId
		key := strconv.Itoa(int(userId)) + action

		_, ok := services.CacheService.Get(key)
		if ok {
			c.AbortWithStatusJSON(400, gin.H{
				"code": 400,
				"data": nil,
				"msg":  "数据处理中，请稍后",
			})
			return
		}
		services.CacheService.Set(key, time.Now(), duration*time.Second)
		c.Next()
	}
}

// 定义中间
func checkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		if utils.InArray(path, noNeedLogin) {
			c.Next()
		} else {
			token := c.GetHeader("Authorization")
			isLogin := true
			if token == "" {
				isLogin = false
			}
			claims := &controllers.Claims{}
			_, err := controllers.CheckToken(token, claims)
			if err != nil {
				isLogin = false
			}

			if isLogin == false || claims.UserId == 0 {
				c.AbortWithStatusJSON(400, gin.H{
					"code": 400,
					"data": nil,
					"msg":  "请重新登录",
				})
				return
			}
			//保存当前登录用户的id到上下文
			c.Set("user", claims)
		}
	}
}

//微信公众号初始化
func wechat(c *gin.Context) {
	uniacidStr := c.Query("uniacid")
	if uniacidStr == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "缺少必要的参数uniacid",
		})
		return
	}
	uniacid, _ := strconv.Atoi(uniacidStr)

	account := &models.Account{
		Uniacid: uint(uniacid),
	}
	account.Get()
	if account.Key == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "找不到对应的公众号信息",
		})
		return
	}
	services.SetKey(account.Key)

	//获取公众号 access_token
	accessToken := utils.Get(viper.GetString("site.url") + "/go.php?uniacid=" + uniacidStr + "&do=get")
	//accessToken := "44_syPLO2W-rOee3pTriw_MQ5O1MH6HJ_Gae5AA9u-lJSvC7o17FPOnt_t5h9YI27S4IsdxqOOhW_BSSc49pU3L-aeTv-Ma5DXZJPcd3-_G_wgjOag08KWO5PkR0wQlewpt2PnlhQyaIgxwjgJBCTAeAGAVMT"
	if accessToken == "Array" {
		c.AbortWithStatusJSON(500, gin.H{
			"code": 500,
			"data": nil,
			"msg":  "微信授权接口异常",
		})
		return
	}
	services.SetToken(accessToken)

	c.Set("uniacid", uniacid)

	//记录邀请人id
	inviteUserId := c.Query("invite_id")
	c.Set("invite_id", inviteUserId)

	//渠道
	from := c.Query("from")
	c.Set("from", from)

	c.Next()
}*/

func Register(port string) {
	r := gin.Default()
	//跨域设置
	r.Use(crossHandler())

	//中间件
	//r.Use(checkAuth())
	// 创建路由
	user := r.Group("/user")
	{
		userApi := &api.User{}

		user.GET("/get", userApi.Get)

		user.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	business := r.Group("/business")
	{
		businessApi := &api.Business{}
		business.GET("/get", businessApi.Get)
	}

	r.Run(":" + port)
}
