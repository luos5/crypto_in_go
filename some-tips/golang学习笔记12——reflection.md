# Golang学习笔记12
反射reflection：
* reflection可以提供程序的灵活性，使interface{}有更大的发挥余地
* 反射使用TyoeOf和ValueOf函数从接口中获取目标对象信息
```
    import (
        "fmt"
        "reflect"
    )
    type User struct{
        Id int
        Name string
        Age int
    }
    func (u User) Hello(){
        fmt.Println("Hello world.")
    }
    func main(){
        u :=User{
            Id:1,
            Name:"Ann",
            Age:20,
        }
        Info(u)
    }
    //可以输入任何类型
    func Info(o interface{}){
        t := reflect.TypeOf(o)
        fmt.Println("Type:",t.Name())
        //若传入是一个地址会引发报错
        if k :=t.Kind(); k!=reflect.Struct{
            fmt.Println("not reflect struct")
            return
        }
        v := reflect.ValueOf(o)
        fmt.Println("Fields:")
        //取出结构的每一个字段
        for i:=0;i<t.NumField();i++{
            f := t.Field(i)
            val := v.Field(i).Interface()
            fmt.Println("%6s:%v = %v\n",f.Name,f.Type,val)
        }
    }
//运行结果
    Type: User
    Fields:
        Id: int = 1
      Name: string = Ann
       Age: int = 20
```
* 反射会将匿名字段作为独立字段
```
    import (
        "fmt"
        "reflect"
    )
    type User struct{
        Id int
        Name string
        Age int
    }
    type Manager struct{
        //嵌入字段，匿名字段
        User
        title string
    }
    func main(){
        m :=Manager{User:User{Id:1,Name:"Ann",Age:20,},title:"123"
        }
        t := reflect.TypeOf(m)
        fmt.Printf("%#v\n",t.Field(0))
        fmt.Printf("%#v\n",t.Field(1))
        //打印第一个字段的第二个元素，就是User的Name
        fmt.Printf("%#v\n",t.FieldByIndex([]int{0,1}))
    }
//运行结果
    reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x4b0d20), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
    reflect.StructField{Name:"title", PkgPath:"main", Type:(*reflect.rtype)(0x4a02e0), Tag:"", Offset:0x20, Index:[]int{1}, Anonymous:false}
    reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4a02e0), Tag:"", Offset:0x8, Index:[]int{1}, Anonymous:false}
```
* 想要利用反射修改对象状态，前提是interface.data是settable，即pointer-interface
```
    import (
        "fmt"
        "reflect"
    )
    func main(){
        //类型要settable
        x := 123
        //pointer-interface，要传入指针给reflect
        v := reflect.ValueOf(&x)
        v.Elen().setInt(999)

        fmt.Println(x)
    }
///运行结果
    999
```
```
    import (
        "fmt"
        "reflect"
    )
    type User struct{
        Id int
        Name string
        Age int
    }
    func main(){
        u := User{1,"Ann",20}
        Set(&u)
        fmt.Println(u)
    }
    func Set(o interface{}){
        v:=reflect.ValueOf(o)
        //判断redlect的值不是指针类型，同时是可以被修改的
        if v.Kind()==reflect.Ptr && !v.Elem().CanSet(){
            fmt.Println("Error")
            return
        }else{
            v=v.Elem()
        }
        f:=v.FieldByName("Name")
        //如果没有找到会返回一个空值，这样判断可以确定是否真的找到了Name
        if !f.IsValid(){
            fmt.Println("BAD")
            return
        }
        if f.Kind()==reflect.String{
            f.SetString("BYE")
        }
    }
///运行结果
    {a BYE 20}
```
* 通过反射可以“动态”调用方法
```
    import (
        "fmt"
        "reflect"
    )
    type User struct{
        Id int
        Name string
        Age int
    }
    func (u User) Hello(name string){
        fmtPrintln("Hello",name,"! I am",u.Name)
    }
    func main(){
        u := User{1,"Ann",20}
        //正常调用method
        u.Hello("Bob")
        //reflect调用method
        v := reflect.ValueOf(u)
        mv := v.MethodByName("Hello")
        args := []reflect.Value{reflect.ValueOf("Cindy")}
        mv.Call(args)
    }
//运行结果
    Hello Bob ! I am Ann
    Hello Cindy ! I am Ann
```