#CAS(compare-and-swap)
## 一. 什么是CAS
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)

假设被操作的值与旧值相等，并一旦确定这个假设的真实性就立即进行值替换

### 正常模式

#### 正常模式怎么工作

`锁的等待者会按FIFO的顺序获取锁。`

`但是刚被唤起的 Goroutine 与新创建的 Goroutine 竞争时，大概率会获取不到锁（因为已经在cpu上跑了，避免切换消耗性能）`

#### 何时切换到饥饿模式
`一旦队列里的waiter Goroutine 超过 1ms 没有获取到锁，它就会将当前互斥锁切换饥饿模式，防止部分 Goroutine 被『饿死』。`


### 饥饿模式
#### 为什么需要饥饿模式
`在 Go 语言 1.9 版本引入的优化1，引入的目的是保证互斥锁的公平性（Fairness）。`

#### 饥饿模式下Goroutine会怎么执行
`互斥锁会直接交给等待队列最前面的 Goroutine。`

`新的 Goroutine 会被追加到队列的末尾等待。`

#### 切换回正常模式的时机
`* owner Goroutine 是队列中最后一个`

`* owner Goroutine等待的时间少于 1ms`

### 为什么需要两种模式
`相比于饥饿模式，正常模式下的互斥锁能够提供更好地性能，饥饿模式的能避免 Goroutine 由于陷入等待无法获取锁而造成的高尾延时。`

