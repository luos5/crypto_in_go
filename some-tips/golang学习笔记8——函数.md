# Golang学习笔记8
函数function：
* 关键字func
* 左大括号必须和关键词同行
* Go函数不支持嵌套、重载、默认参数
* 支持无需声明原型
```
    func A(a int, b string)(int, string, int){

    }
```
* 支持不定长度变参，但必须作为参数列最后一个参数来使用
```
    func main(){
        A(1,2,3,4,5,6)
    }
    //未知输入参数有几个，用“...”不定长变参，以slice形式获取输入
    func A(x int,a ... int){
        fmt.Println(x)
        fmt.Println(a)
    }
//运行结果
    1
    [2 3 4 5 6]
```
```
    func main(){
        a, b, c := 99,999,9999
        A(a,b,c)
        fmt.Println("final of a,b,c:",a,b,c)
    }
    //不定长度变参获取输入，普通类型参数输入，只是拷贝值，操作不对原数据造成影响。拷贝到地址了，例如真正输入一个slice作为参数，slice1 []int，才会对原slice1造成影响
    func A(x int,y ... int){
        fmt.Println(x)
        fmt.Println(y)
        y[0]=100
        y[1]=101
        fmt.Println(y)
    }
//运行结果
    99
    [999 9999]
    [100 101]
    final of a,b,c: 99 999 9999
```
* 支持多返回值
```
//第二个括号内的是返回值的类型，可以多个
    func A(a int, b string)(int, string, int){

    }
//只有一个返回值的话，可以不用括号
    func B(a int, b string) int {

    }    
```
* 支持命名返回值参数，想要简写多个同类型的返回值必须命名，如果想简写“return”必须命名
```
//可以命名返回型，先写名称再写类型
    func A()(a int, b string, c int){

    }
//返回类型相同，可以省略前面几个的类型
    func A()(a, b, c int){
        //在函数就命名了返回值名称，就不需要“ := ”来声明变量了，应当直接用“ = ”赋值
        a, b, c = 1,2,3
        //简写return，没有指出具体返回值，必须命名返回值参数
        return
    }
```
* 支持匿名函数
```
    func main(){
        //这个代码块实现了一个函数，但是没有名称，赋值给了“a”可以通过a()来运行，称为匿名函数
        a := func(){
            fmt.Println("function")
        }
        a()
    }
//运行结果
    function
```
* 支持闭包
```
    func main(){
        f := closure(10)
        fmt.Println(f(1))
        fmt.Println(f(2))
    }
    //将返回一个匿名函数
    func closure(x int)(func(int) int){
        //x是一样的，都是直接指向同一个地址，不是拷贝值
        println(&x)
        return func (y int) int{
            //x是一样的，都是直接指向同一个地址，不是拷贝值
            println(&x)
            return x+y
        }
    }
//运行结果
    0xc000096010
    0xc000096010
    11
    0xc000096010
    12
```
* 函数可以作为一种类型使用
```
    func main(){
        a := A
        //相当于运行了A()
        a()
    }
    func A(){
        fmt.Println("function A")
    }
//运行结果
    function A
```
defer:
* defer的执行方式类似析构函数，在函数体执行结束后按照调用顺序的**相反顺序**逐个执行
```
    func main(){
        fmt.Println("a")
        defer fmt.Println("b")
        //defer逆序调用，所以会先打印出“c”
        defer fmt.Println("c")
    }
//运行结果
    a
    c
    b
```
* defer支持匿名函数的调用
* 如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时已经获得拷贝，否则将引用该变量的地址
```
    func main(){
        for i :=0; i<3;i++{
            //此处i是作为参数拷贝
            defer fmt.Println("print_i:",i)
            //相当于defer a()
            defer func (){
                //此处i是作为地址引用
                fmt.Println("func_1:",i)
            }()
        }
    }
//运行结果
    func_1: 3
    print_i: 2
    func_1: 3
    print_i: 1
    func_1: 3
    print_i: 0
```
* 即使函数发生**严重错误**也会执行
* defer常用于资源清理、文件关闭、解锁，以及记录时间等操作
* 通过与匿名函数配合可在return之后**修改**函数计算结果
* Go没有异常捕获机制（没有try catch），但有panic/recover模式来处理错误
* panickey在任何地方引发，recover只有在defer调用的函数中有效
```
    func main(){
        A()
        B()
        C()
    }
    func A(){
        fmt,Println("func A")
    }
    func B(){
        //在panic之前先定义defer才能有效
        //defer一定会执行
        defer func(){
            //执行recover()就能从panic或者异常里恢复过来
            if err := recover();err != nil{
                fmt.Println("recover in B")
            }
        }()
        //panic会中断所有的函数，但之后defer也还是会执行
        panic("panic in B")
    }
    func C(){
        fmt.Println("func C")
    }
```

