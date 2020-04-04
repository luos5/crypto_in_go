# Golang学习笔记4
Go的指针：
* 不支持指针运算，不支持“ -> ”运算符
* 通过“ . ”选择符来操作指针目标对象的成员
* 操作符“ & ”取变量地址
* 用“ * ”间接访问目标对象
* 默认值是nil
```
    a := 1
    b := 1.5
    // *int就是指向int的指针
    // &a 就是取a的变量地址
    var p *int = &a
    var q *float64 = &b
    println(p)
    // *p 就是通过指针p间接访问目标对象a
    println(*p)
    println(q)
    println(*q)
//运行效果
    0xc000018108
    1
    0xc000018120
    1.5
```
递增递减语句：
* ++和--是语句，不是表达式
* 不允许 a := a++，因为这里作为表达式使用了
* 允许单独成行的 a++，这里作为语句使用
```
    a := 1
    a++
    println(a)
    b := 1
    b--
    println(b)
//运行结果
    2
    0
```
判断语句if：
* 条件表达式没有括号，加上括号会编译错误
```
//if (a>b)会编译错误
    if a>b{
        println(a)
    }
```
* 支持一个初始化表达式，可以是并行方式
```
    if a,b:=1,2;a>0{
        println(a,b)
    }
//运行结果
    1 2
```
* 初始化语句中的变量为block级别，只有if这个语句块能用，对外是undefined的，同时隐藏外部同名变量，可以初始化一个跟外面同名的变量
* if语句没有初始化同名变量的话，if语句块内部的操作会作用在外部的那个变量上，如果有初始化同名变量，作用在初始化的同名变量上
```
    a := 1
    b := 2
    if a:=3;a>0{
        //作用内部a
        a++
        //没有同名，作用外部b
        b++
        //打印出内部的a
        println(a)
    }
    //内部a已释放，打印外部a,b
    print(a,b)
//运行效果
    4
    1 3
```
* 左大括号“ { ”必须和条件语句或else在同一行，否则会编译错误
```
    if a>b{
        println(a)
    }
    else{
        println(b)
    }
```
* 支持单行模式
```
    if a:=1; a>0{println(a)}
```
循环语句for：
* 关键字只有for
* 初始化和步进表达式可以是多个值
* 每次循环都会检查条件语句，不建议在条件语句中使用函数，建议提前计算好条件并以变量或常量代替（函数表达式）
* 左大括号“ { ”必须和条件语句在同一行
```
//do while模式 
    a :=1
    for{
        //操作
        a++
        //条件语句
        if a>3{
            break
        }
    }
```
```
//while模式
    a :=1
    //条件语句
    for a<=3{
        //操作
        a++
    }
```
```
//for模式
    a :=1
    //条件语句
    for i:=0;i<3;i++{
        //操作
        a++
    }
```
选择语句switch：
* 左大括号“ { ”必须和条件语句在同一行
* 可以用任何类型或者表达式作为条件语句
```
    a := 1
    switch{
    //条件语句可以是任何类型
    case 0:
        println("0")
    //条件语句可以是表达式
    case a>0:
        println(a)
    }
//运行结果
    1
```
* 不需要break，一旦条件符合自动终止，其他或者默认的关键词是default
```
    a := 1
    switch{
    case 0:
        println("0")
    default:
        println("notzero")
    }
```
* 如果希望继续执行下一个case，就是已经符合还想继续往下检查，要用fallthrough语句
```
    a := 0
    switch{
    case 0:
        a++
        //符合也继续执行下一个case
        fallthrough
    case a<5:
        println(a)
    }
//运行结果
    1
```
* 支持一个初始化表达式，可以是并行方式，右侧要跟分号，离开switch语句块，初始化变量将undefined
```
    //初始化表达式，可并行，右侧必须有分号
    switch a,b:=1,2;{
    case 0<a:
        println(a)
    case a<=0:
        println(b)
    }
//运行结果
    1
```
跳转语句goto, break, continue：
* 可以配合标签使用
* 标签名区分大小写，写了标签又不适用会造成编译错误
* break和continue配合标签可以用于多层循环的跳出
```
LABEL1:
    for{
        //标签写在这行是不能跳出的
        for i :=1; i<=3;i++{
            if i>3{
                //将break到跳出与标签同级的循环
                break LABEL1
            }
        }
    }
    println("pass")
//运行结果
    pass
```
```
LABEL1:
    for i,a:=0,1;i<=3;i++{
        for{
            a++
            if a>3{
                //继续执行标签同级的循环，即外层的有限循环
                println(i,a)
                continue LABEL1
            }
            println(a)
        }
    }
    println("pass")
//运行结果
    2
    3
    0 4
    1 5
    2 6
    3 7
    pass
```
* goto是调整执行位置，所以为了配合跳出循环的标签建议放在循环（外）的后面
```
    for{
        for i :=1; i<=3;i++{
            if i>3{
                //将跳到标签位置所在行
                goto LABEL1
            }
        }
    }
    LABEL1:
    println("pass")
//运行结果，没有死循环
    pass
```
