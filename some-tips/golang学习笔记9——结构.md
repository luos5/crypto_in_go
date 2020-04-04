# Golang学习笔记9
结构struct：
* Go没有class
* 使用type <变量名称> struct{}定义结果，名称遵循可见性规则
```
    //定义了一个空的结构
    type test struct{}
    func main(){
        a := test{}
        fmt.Println(test)
    }
//运行结果
    {}
```
```
    //定义了一个结构
    type student struct{
        //首字母大写表示是public可调用的
        Name string
        Age int
    }
    func main(){
        a := student{}
        fmt.Println(student)
        //Go没有指针运算，使用“ . ”来操作指针目标对象的成员
        a.Name = "Ann"
        a.Age = 20
        fmt.Println(student)
    }
//运行结果
    //字符串的默认值是空字符串，int默认值是0
    {  0}
    {Ann 20}
```
* 可以使用字面值对结构进行初始化，复合结构内的结构不能字面值初始化，要用多个“ . ”来初始化
```
    type student struct{
        Name string
        Age int
    }
    func main(){
        //字面值初始化，每个属性<varName> : <value>,记住每个属性赋值完要加逗号间隔
        a := student{
            Name :"Ann",
            Age :20,
        }
        b := student{
            Name :"Bob",
        }
        fmt.Println(a)
        fmt.Println(b)
    }
//运行结果
    {Ann 20}
    {Bob 0}
```
* 支持指向自身的指针类型成员
```
    //复合结构，结构也可以是结构的成员
   type student struct{
        Name string
        Age int
        Contact struct{
            Phone, City string
        }
    }
    func main(){
        a := student{
            Name :"Ann",
            Age :20,
        }
        a.Contact.Phone="12345678912"
        a.Contact.City="Beijing"
        fmt.Println(a)
    }
//运行结果
    {Ann 20 {12345678912 Beijing}}
```
* 支持匿名结构，可用作成员或定义成员变量
```
    func main(){
        //没有对结构命名，也可以在struct前面加上&，就取地址了，更方便后面使用
        a := struct{
            Name string
            Age int
        }{
            Name : "Ann",
            Age : 20
        }
        fmt.Println(a)
    }
//运行结果
    {Ann 20}
```
* 匿名结构也可以用于map的值
```
    
```

* 允许直接通过指针来读写结构成员
```
    type student struct{
        Name string
        Age int
    }
    func main(){
        a := student{
            Name :"Ann",
            Age :20,
        }
        //可以用“ . ”读写结构成员
        a.Name="Bob"
        fmt.Println(a)
    }
//运行结果
    {Bob 20}
```
```
    type student struct{
        Name string
        Age int
    }
    func main(){
        //初始化的时候就取地址，之后传递给其他函数就直接是拷贝地址，方便直接操作
        a := &student{
            Name :"Ann",
            Age :20,
        }
        fmt.Println("a1:",a)
        A(a)
        fmt.Println("a2:",a)
    }
    //输入类型是指针，需要写“ * ”
    func A(s *student){
        //指针也可以使用“ . ”来直接读写结构对象
        s.Name="Bob"
        fmt.Println("s:",s)
    }
//运行结果
    a1: {Ann 20}
    s: {Bob 20}
    a2: {Bob 20}
```
* 相同类型的成员可进行直接拷贝赋值
```
    type student struct{
        Name string
        Age int
    }
    func main(){
        a := student{
            Name :"Ann",
            Age :20,
        }
        //另外定义一个同类型的变量
        var b student
        //可以直接拷贝赋值
        b = a
        fmt.Println(a)
        fmt.Println(b)
    }
//运行结果
    {Ann 20}
    {Ann 20}
```
* 支持==与!=比较运算符，结构名称不一样也是不一样的类型，不能比较，否则会报错，不支持<和>
```
    type student struct{
        Name string
        Age int
    }
    type student2 struct{
        Name string
        Age int
    }
    func main(){
        a := student{
            Name :"Ann",
            Age :20,
        }
        var b student
        b = a
        c :=student{
            Name :"Bob",
            Age :20,
        }
        d :=student2{
            Name :"Bob",
            Age :20,
        }
        //d不能和abc进行比较
        fmt.Println(a==b)
        fmt.Println(c==b)
    }
//运行结果
    true
    false
```
* 支持匿名字段，本质上是定义了以某个类型名为名称的字段
```
    //匿名字段，字段没有名称
    type student struct{
        string
        int
    }
    func main(){
        //按顺序赋值，否则会报错或赋值不正确
        a := student{
            "Ann",20,
        }
        fmt.Println(a)
    }
//运行结果
    {Ann 20}
```
* 嵌入结构作为匿名字段不是继承
* 可以使用匿名字段指针
```
    type human struct{
        Sex int
    }
    type student struct{
        //嵌入了结构human，达到类似继承的效果
        human
        Name string
        Age int
    }
    func main(){
        a := student{
            Name :"Ann",
            Age :20,
            //嵌入的要这样定义
            human: human{Sex :0},
        }
        fmt.Println(a)
        //两种改嵌入的字段的方法，一种直接“ . ”最终的字段，默认嵌入结构的字段都给了外层结构
        a.Sex=1
        fmt.Println(a)
        //两种改嵌入的字段的方法，一种“ . ”结构名称再“ . ”最终的字段，可以预防名称冲突
        a.human.Sex=0
        fmt.Println(a)
    }
//运行结果
    {{0} Ann 20}
    {{1} Ann 20}
    {{0} Ann 20}
```


