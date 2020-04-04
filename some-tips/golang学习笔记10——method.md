# Golang学习笔记10
方法method：
* Go没有class，但有method
* 通过显示说明receiver来实现与某个类型的组合
```
    type A struct{
        Name string
    }
    func main(){
        a := A()
        a.Print()
    }
    //函数名称前面有一个receiver，根据receiver的类型来判断是哪个结构的方法
    func (a A)Print(){
        fmt.Println("A")
    }
```
* 只能为同一个包中的类型定义方法
* Receiver可以是类型的值或者指针
* 可以使用值或者指针来调用方法，编译器会自动完成转换
```
    type A struct{
        Name string
    }
    type B struct{
        Name string
    }
    func main(){
        a := A()
        a.Print()
        fmt.Println(a.Name)
        b := B()
        b.Print()
        fmt.Println(b.Name)
    }
    //可以是类型的指针
    func (a *A)Print(){
        //操作拷贝地址是有效的
        a.Name="AA"
        fmt.Println("A")
    }
    //可以是类型的值
    func (b B)Print(){
        //操作拷贝值是无效的
        b.Name="AA"
        fmt.Println("B")
    }
//运行结果
    A
    AA
    B
     
```
* 不存在方法重载，但是方法可以跟类型绑定，达到类似效果
```
    type A struct{
        Name string
    }
    type B struct{
        Name string
    }
    func main(){
        a := A()
        a.Print()
        b := B()
        b.Print()
    }
    //函数名称前面有一个receiver，根据receiver的类型来判断是哪个结构的方法
    func (a A)Print(){
        fmt.Println("A")
    }
    func (b B)Print(){
        fmt.Println("B")
    }
//运行结果
    A
    B
```
* 从某种意义上来说，方法是函数的语法糖？receiver就是方法所接收的第一个参数？
```
    type TZ int
    func main(){
        var a TZ
        //以下两种方式都可以调用到方法，method_value,method_expression
        a.Print()
        (*TZ).Print(&a)
    }
    //函数名称前面有一个receiver，根据receiver的类型来判断是哪个结构的方法
    func (a A)Print(){
        fmt.Println("TZ")
    }
//运行结果
    TZ
    TZ
```
* 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法，同级重名会导致编译错误
* 类型别名不会拥有底层类型所附带的方法
* 方法可以调用结构中的非公开字段，因为实际上还是在一个包里面
```
    type A struct{
        //非公开字段name
        name string
    }
    func main(){
        a := A()
        a.Print()
        fmt.Println(a.name)
    }
    //函数名称前面有一个receiver，根据receiver的类型来判断是哪个结构的方法
    func (a *A)Print(){
        //操作结构的私有字段
        a.name="123"
        fmt.Println(a.name)
    }
//运行结果
    123
    123
```