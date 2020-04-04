# Golang学习笔记1
一个.go文件的内容结构：
* 非注释的第一行标识所属包（package）
* 引用的包import导入并用" "引号框住
* 常量const
* 全局变量var
* 类型声明type xxx int
* 结构声明type xxx struct{}，尾巴必须跟着大括号，里面有内容
* 接口声明type xxx interface{}，尾巴必须跟着大括号，里面有内容
* 接着为函数部分
```
package main
import "tmp"
const a=1
var b="abc"
type mytype int
type mystruct struct{}
type myinterface interface{}
func main(){
    tmp.Println("Hello,world!")
}
```
其他注意事项：
* 只能有一个package main，也只能有一个main函数,main函数必须在packet main里面，必须有且仅能有一个main函数
* import多个包可以用括号框住，一次性导入，同理适用于常量声明、全局变量声明、一般类型声明。结构和接口声明不行，函数体内变量声明不行。
```
import (
    "tmp"
    "time"
    "string"
    "io"
)
```
* 调用包内的函数，packagename.functionname
```
fmt.Println("Hello,world!")
```
* 别名，有些包名称相近或相同的时候，可以用别名标识
```
import std "tmp"
std.Println
```
* 大小写标识常量、变量、类型、结构、接口、函数是否可以被外部包调用
* 大写是public
* 小写是private