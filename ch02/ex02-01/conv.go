// 练习 2.1： 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝
// 对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。
package tempconv


func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FToK(f Fahrenheit) Kelvin { return Kelvin((f - 32) * 5 / 9  + AbsoluteZeroC) }

func KToC(k Kelvin) Celsius { return Celsius(k - AbsoluteZeroC) }

func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k - AbsoluteZeroC)*9/5 + 32) }

