# Golang学习笔记2
Go的基本类型：
* 布尔型
```
//关键词bool，长度1字节，取值范围true/false，不能用数字代替
	var b bool
	println("Size of it is:",unsafe.Sizeof(b),"Byte")
	println("Default value:",b)
	b=true
	println("The new value:",b)
//运行效果
Size of it is: 1 Byte
Default value: false
The new value: true
```
**tip:**
```
变量声明格式：var <变量名称> <变量类型>，例如：var b bool
变量赋值格式：<变量名称> = <表达式>，例如：b = true
声明同时赋值：var <变量名称> <变量类型> = <表达式>，例如：var b bool = true
```
* 整型
```
//整型可以用math库的函数Min+整型名/Max+整型名知道该整型的取值范围
//关键词int/uint，32位或64位，4字节或8字节
	var b int
	println("Size of it is:",unsafe.Sizeof(b),"Byte")
	println("Default value:",b)
	b=math.MaxInt64
	println("The new value:",b)
//运行效果
Size of it is: 8 Byte
Default value: 0
The new value: 9223372036854775807
```
* 8位整型
```
//关键词int8/uint8，1字节，取值范围-128~127/0~255
	var b int8
	println("Size of it is:",unsafe.Sizeof(b),"Byte")
	println("Default value:",b)
	b=127
	println("The new value:",b)
//运行效果
Size of it is: 1 Byte
Default value: 0
The new value: 127
```
* 字节型（uint8）
```
//关键词byte，1字节，取值范围0~255
	var b byte
	println("Size of it is:",unsafe.Sizeof(b),"Byte")
	println("Default value:",b)
	b=127
	println("The new value:",b)
//运行效果
Size of it is: 1 Byte
Default value: 0
The new value: 127
```
* 16位整型
```
//关键词int16/uint16，2字节
```
* 32位整型
```
//关键词int32（rune）/uint32，4字节
```
* 64位整型
```
//关键词int64/uint64，8字节
```
* 浮点型
```
///关键词float32/float64,4/8字节，精确到7/15小数位
	var b float32
	println("Size of it is:",unsafe.Sizeof(b),"Byte")
	println("Default value:",b)
	b=1.54321543215432154321
	println("The new value:",b)
//运行效果
Size of it is: 4 Byte
Default value: +0.000000e+000
The new value: +1.543215e+000
```
* 复数
```
//关键词complex64/complex128,8/16字节
	var b complex64
	println("Size of it is:",unsafe.Sizeof(b),"Byte")
	println("Default value:",b)
	b=1+2i
	println("The new value:",b)
//运行效果
Size of it is: 8 Byte
Default value: (+0.000000e+000+0.000000e+000i)
The new value: (+1.000000e+000+2.000000e+000i)
```
* 可以保存指针的整数型
```
//关键词uintptr，32位或64位，就是整数型加上一个指针的长度
//具体使用后面再说
```
* 数组
```
//关键词array
    var a [0]bool
    println(a)
    var b []int
    println(b)
    var c [0]int
    println(c)
//运行效果
[false]
[]
[0]
```
* 结构
```
//关键词struct
```
* 字符串
```
//关键词string
```
* 引用类型
```
//关键词slice、map、chan，切片，地图（哈希表），通道
```
* 接口类型
```
//关键词interface
```
* 函数类型
```
//关键词func
```

如何跟简略地声明变量并赋值：
* 方法一：
变量声明格式：var <变量名称> <变量类型>，例如：var b bool
变量赋值格式：<变量名称> = <表达式>，例如：b = true
* 方法二：
声明同时赋值：var <变量名称> <变量类型> = <表达式>，例如：var b bool = true
* 方法三：
系统自动推断：var <变量名称> = <表达式>，例如：var b = true
* 方法四：
系统自动推断：<变量名称> := <表达式>，例如：b := true
**tip:**
当初始赋值容易产生歧义时，不宜使用自动推断，例如开始为“123”，后续运算后为“61.5”

如何同时对多个变量进行声明与赋值：
* 多个全局变量：var()，例如：
```
var(
    a = "lala"
    b, c = 1, 2
)
```
* 多个局部变量：并列，例如：
```
var a, b, c, d int
a, b, c, d = 1,2,3,4
var a, b, c, d int= 1,2,3,4
var a, b, c, d = 1,2,3,4
a, b, c, d := 1,2,3,4
```
如何进行类型转换：
* 必须显式转换
* 只能在两种互相兼容的类型之间转换，例如：浮点型和整型可以转换，整型和字符串可以转换，布尔型和整型不能转换
* 转换格式：<变量名称a> := <变量类型a>(<变量名称b>)
```
var a float32 = 1.2
b := int(a)
```

