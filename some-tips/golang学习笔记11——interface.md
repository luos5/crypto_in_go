# Golang学习笔记11
接口interface：
* 接口是一个或多个方法签名的集合
* 只要某个类型拥有该接口的所有方法签名，就算实现该接口，无需显示声明实现了哪个接口，Structural Typing
```
    type USB interface{
        //两个方法，一个返回USB的名称，一个连接的方法
        Name() string
        Connect()
    }
    type PhoneConnecter struct{
        name string
    }
    func (pc PhoneConnecter) Name() string{
        return pc.name
    }
    func (pc PhoneConnecter) Connect(){
        fmt.Println("Connected:",pc.name)
    }
    func main(){
        var a USB
        //PhoneConnecter成功实现了USB这个接口，就可以作为USB类型来用了
        a=PhoneConnecter{
            name:"iPhone",
        }
        a.Connect()
        //PhoneConnecter成功实现了USB这个接口，就可以作为USB类型来用了
        Dissconnect(a)
    }
    func Dissconnect(usb USB){
        fmt.Println("Dissconnected.")
    }
//运行结果
    Connected: iPhone
    Dissconnected.
```
* 接口只有方法声明，没有实现，没有数据字段
* 接口可以匿名嵌入其他接口，或嵌入到结构中
```
    type USB interface{
        Name() string
        //嵌入一个接口
        Connecter
    }
    type Connecter interface{
        Connect()
    }
    type PhoneConnecter struct{
        name string
    }
    func (pc PhoneConnecter) Name() string{
        return pc.name
    }
    func (pc PhoneConnecter) Connect(){
        fmt.Println("Connected:",pc.name)
    }
    func main(){
        var a USB
        //PhoneConnecter成功实现了USB这个接口，就可以作为USB类型来用了
        a=PhoneConnecter{
            name:"iPhone",
        }
        a.Connect()
        //PhoneConnecter成功实现了USB这个接口，就可以作为USB类型来用了
        Dissconnect(a)
    }
    func Dissconnect(usb USB){
        //类型断点，可以用来判断类型
        if pc, ok :=usb.(PhoneConnecter);ok{
            fmt.Println("Dissconnected:",pc.name)
            return
        }
        fmt.Println("Unknown device.")
    }
//运行结果
    Connected: iPhone
    Dissconnected: iPhone
```
* 将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
```
    type USB interface{
        Name() string
        //嵌入一个接口
        Connecter
    }
    type Connecter interface{
        Connect()
    }
    type PhoneConnecter struct{
        name string
    }
    func (pc PhoneConnecter) Name() string{
        return pc.name
    }
    func (pc PhoneConnecter) Connect(){
        fmt.Println("Connected:",pc.name)
    }
    func main(){
        pc := PhoneConnecter{"PhoneConnecter"}
        var a Connecter
        a = Connecter(pc)
        a.Connect()
        //完全没用
        pc.name="pc"
    }
//运行结果
    Connected: PhoneConnecter
```
* 只有当接口存储的类型和对象都为nil时，接口才等于nil
```
    func main(){
        var a interface{}
        fmt.Println(a==nil)

        var p *int =nil
        a=p
        fmt.Println(a==nil)
    }
//运行结果
    true
    false
```
```
    type USB interface{
        Name() string
        //嵌入一个接口
        Connecter
    }
    type Connecter interface{
        Connect()
    }
    type PhoneConnecter struct{
        name string
    }
    func (pc PhoneConnecter) Name() string{
        return pc.name
    }
    func (pc PhoneConnecter) Connect(){
        fmt.Println("Connected:",pc.name)
    }
    type TVConnecter struct{
        name string
    }
    func (tv TVConnecter) Connect(){
        fmt.Println("Connected:",tv.name)
    }
    func main(){
        tv := TVConnecter{"TVConnnecter"}
        var a USB
        //强制转换类型
        a = USB(tv)
        a.Connect()
    }
//运行结果
    报错，1，Name是一个field而不是一个方法2，TVConnecter没有实现接口USB，不能转换成USB
    但是如果是pc转成connecter是可以的，没有障碍的
```
* 接口调用不会做receiver的自动转换，跟调用结构不一样
* 接口同样支持匿名字段方法
* 接口也可以实现类似OOP中的多态
* 空接口可以作为任何类型数据的容器
```
    //函数输入类型是空接口
    func Dissconnect(usb interface()){
        //类型断点，可以用来判断类型,type switch
        switch v:=usb.(type){
            case PhoneConnecter:
                fmt.Println("Dissconnected:",v.name)
            default:
                fmt.Println("Unknown device.")
        }
    }
//运行结果
    Connected: iPhone
    Dissconnected: iPhone
```
