package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsolutZeroC Celsius = -273.15 // 绝对零度
	FreezingC    Celsius = 0       // 结冰温度
	BoilingC     Celsius = 100     // 沸水温度
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C\n", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F\n", f)
}
