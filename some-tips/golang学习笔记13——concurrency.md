# Golang学习笔记13
并发concurrency：
* 每个实例4-5KB的栈内存占用，实现机制大幅减少创建和销毁开销，高并发，goroutine简单易用
* 并发主要由切换时间片来实现“**同时**”运行；并行是直接利用多核实现多线程的运行。Go可以设置使用核数，以发挥多核计算机的能力。
```
    import (
        "fmt"
        "time"
    )
    func main(){
        //go关键字就是goroutine
        go Go()
        //sleep的时候Go()运行了打印了东西
        time Sleep()(2 * time.Second)
    }
    func Go(){
        fmt.Println("Go")
    }
```
* Goroutine通过通信来共享内存，而不是共享内存来通信
channel：
* channel是goroutine沟通的桥梁，大多是阻塞同步的
* 通过make创建，close关闭
* channel是引用类型，传入就可以直接操作，不是拷贝值
* 可以使用for range来迭代不断操作channel
* 可以设置单向或双向通道
* 可以设置缓存大小，在未填满前不会发生阻塞
```
```
select：
* 可处理一个或多个channel的发送和接收
* 同时有多个可用的channel时按随机顺序处理
* 可用空的select来阻塞main函数
* 可设置超时

