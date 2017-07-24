// 练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的
// 话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对
// 应英尺和米，重量单位可以对应磅和公斤等。
package lengthconv

func MToF(m Meter) Foot { return Foot(m / 0.3048) }

func FToM(f Foot) Meter { return Meter(f * 0.3048) }
