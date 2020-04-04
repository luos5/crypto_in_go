# Golang学习笔记7
map：
* map类似哈希表或者字典，以key-value的形式存储数据
* key必须是指出==或!=比较运算的类型，不能是函数、map、slice
* map查找比线性搜索快，但比使用索引访问数据的类型（slice、数组）慢
* 变量var创建map，var <变量名称> map[keyType]valueType
```
    var m map[int]string
    m = map[int]string{}
    println(m)
    fmt.Println(m)
//运行结果
    0xc000078150
    map[]
```
* 使用make()创建map，make([keyType]valueType, cap)，cap表示容量，可以省略
```
    var m map[int]string
    m = make(map[int]string, 10)
    println(m)
    fmt.Println(m)
//运行结果
    0xc000078150
    map[]
```
* 简写创建map
```
    m := make(map[int]string)
    println(m)
    fmt.Println(m)
//运行结果
    0xc000078150
    map[]
```
* 超出容量会自动翻倍扩容，建议提供一个合理的初始值
* 可以使用len()获取元素个数
* 键值对不存在时自动添加，delete()可以删除指定键值对
```
    m := make(map[int]string)
    //不存在就自动添加
    m[1] = "first"
    println(m)
    fmt.Println(m)
    //删除指定键值对，以key为检索
    delete(m,1)
//运行结果
    0xc000076150
    map[1:first]
    0xc000078150
    map[]
```
* 复合map，valueType是map的情况，有几层就得make()几次，否则会编译异常
```
    var m map[int]map[int]string
    m = make(map[int]map[int]string)
    //要新建内部的map
    m[1] = make(map[int]string)
    //再新建内部的键值对
    m[1][1]="first"
    println(m)
    fmt.Println(m)
//运行结果
    0xc000080150
    map[1:map[1:first]]    
```
* 判断键值对是否有值，可以将索引的返回值赋给两个值，前者会返回该索引指向的值，后者返回该指向对象是否存在
```
     var m map[int]map[int]string
    m = make(map[int]map[int]string)
    m[1] = make(map[int]string)
    m[1][1]="first"
    //a将返回m[1][1]指向的值，b将返回该索引是否存在（true/false）
    a, b := m[1][1]
    println(m)
    fmt.Println(m)
    println(a,b)
//运行结果
    0xc000080150
    map[1:map[1:first]]
    first true
```
* 使用for range对map和slice进行迭代操作
```
    slice_m := make([]map[int]string,5)
    //i相当于索引，v相当于value，但这些是拷贝值，对v直接操作不会对真实的slice_m有影响
    for i,v := range slice_m{
        v := make(map[int]string,1)
        v[1]="first"
        fmt.Println(i)
        fmt.Println(v)
    }
    fmt.Println(m)
//运行结果
    0
    map[1:first]
    1
    map[1:first]
    2
    map[1:first]
    3
    map[1:first]
    4
    map[1:first]
    [map[] map[] map[] map[] map[]]
```
```
    slice_m := make([]map[int]string,5)
    //i相当于索引，v相当于value
    for i:= range slice_m{
        //取索引后对slice_m操作才能改变真实的slice
        slice_m[i] := make(map[int]string,1)
        slice_m[i][1]="first"
        fmt.Println(i)
        fmt.Println(slice_m[i])
    }
    fmt.Println(m)
//运行结果
    0
    map[1:first]
    1
    map[1:first]
    2
    map[1:first]
    3
    map[1:first]
    4
    map[1:first]
    [map[1:first] map[1:first] map[1:first] map[1:first] map[1:first]]
```
```
    m := make(map[int]string,5)
    m[1] = "first"
    m[5] = "five"
    for i,v := range m{
        fmt.Println(i)
        fmt.Println(v)
        v="zero"
    }
    fmt.Println(m)
//运行结果
    1
    first
    5
    five
    map[1:first 5:five]
```