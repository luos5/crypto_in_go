# Golang学习笔记6
切片slice：
* slice是引用类型
* slice不是数组，slice指向已经存在的数组（底层的数组）
```
    a := []int{0,1,2,3,4,5,6,7,8,9}
    slice1 := a[2:5]
    slice2 := a[3:5]
    fmt.Println(slice1)
    println(slice1)
    fmt.Println(slice2)
    println(slice2)
//运行结果
    [0 1 2 3 4 5 6 7 8 9]
    [2 3 4]
    //3是存数的元素个数，8是容量，0xc0000200b0第一个元素起始地址
    [3/8]0xc0000200b0
    [3 4]
    [2/7]0xc0000200b8
```
* slice可以直接创建，或从一个已经存在的数组（底层数组）获取生成
* 方法一新建一个slice：
```
    //声明一个空的slice
    var slice1 []int
    println(slice1)
//运行结果
    []
```
* 方法二在数组上建slice：
```
    //一个底层数组
    a := [10]int{0,1,2,3,4,5,6,7,8,9}
    //声明一个slice，取数组a的局部位置，a[m:n]指的是取a[m]到a[n-1]，包含m不包含n
    slice1 := a[5:10]
    slice2 := a[5:len(a)]
    //取数组第六个值及之后
    slice3 := a[5:]
    //取数组前五个值
    slice4 := a[:5]
    println(a)
    println(slice1)
    println(slice2)
    println(slice3)
    println(slice4)
//运行结果
    [0,1,2,3,4,5,6,7,8,9]
    [5,6,7,8,9]
    [5,6,7,8,9]
    [5,6,7,8,9]
    [0,1,2,3,4]
```
* 一般使用make()创建，make([]T, len, cap)，其中cap可以省略，默认与len相同，len代表存数的元素个数，cap代表容量,可以通过len()获取元素个数，cap()获取容量
```
    //三个参数，数组内值类型、存数的元素个数、容量
    slice1 := make([]int,3,10)
    println(len(slice1),cap(slice1))
    //省略cap，相对于默认与len相同
    slice2 := make([]int,3)
    println(len(slice2),cap(slice2))
    //增加slice超出容量的时候，会自动翻倍容量，重新分配
//运行结果
    3 10
    3 3
```
* 如果多个slice指向同一个底层数组，其中一个的值改变会影响全部slice
二次切片reslic：
* 对slice进行切片
* reslice时索引以被slice的切片为准
* 索引不可以超过被slice的切片的容量cap()值
* 索引越界不会导致底层数组的重新分配，会引发错误
```
    //一个底层数组
    a := [10]int{0,1,2,3,4,5,6,7,8,9}
    //一个slice
    slice1 := a[2:5]
    //reslice，对slice1取slice，取slice1的第二个值到第三个值
    slice2 := slice1[1:3]
    fmt.Println(a)
    fmt.Println(slice1)
    fmt.Println(slice2)
//运行结果
    [0 1 2 3 4 5 6 7 8 9]
    [2 3 4]
    [3 4]
```
**attention:** 可以代替变长数组，可以关联底层数组的局部或者全部

追加append：
* 可以在slice尾部追加元素
* 可以把一个slice追加在另一个slice尾部，追加要有“...”，否则会报错
* 如果追加后最终长度未超过被追加的slice的容量，返回原始slice
* 如果超过被追加的slice的容量，将重新分配新的底层数组并拷贝原始数据，原本底层数组的改变将不影响新的切片
```
    //一个底层数组
    a := [10]int{0,1,2,3,4,5,6,7,8,9}
    //一个slice
    slice1 := make([]int,3,6)
    println(slice1)
    fmt.Println(slice1)
    //追加三个值，1,2,3
    slice1 = append(slice1,1,2,3)
    println(slice1)
    fmt.Println(slice1)
    //另一个slice
    slice2 := a[2:5]
    //slice1追加slice2，必须在切片后面加“...”
    slice1=append(slice1,slice2...)
    println(slice1)
    fmt.Println(slice1)
//运行结果
    [0 0 0]
    [3/6]0xc00001a120
    [0 0 0 1 2 3]
    [6/6]0xc00001a120
    [0 0 0 1 2 3 2 3 4]
    [9/12]0xc00001c0c0
```
复制copy:
```
    slice1 := []int{1,2,3,4,5,6}
    slice2 := []int{7,8,9}
    //把后者拷贝到前者上
    copy(slice1,slice2)
    fmt.Println(slice1)
    slice3 := []int{1,2,3,4,5,6}
    slice4 := []int{7,8,9}
    //把前者拷贝到后者上
    copy(slice4,slice3)
    fmt.Println(slice4)
    slice5 := []int{1,2,3,4,5,6}
    slice6 := []int{7,8,9}
    //部分索引拷贝
    copy(slice5[2:4],slice6[1:3])
    fmt.Println(slice5)
//运行结果
    [7 8 9 4 5 6]
    [1 2 3]
    [1 2 8 9 5 6]
```
