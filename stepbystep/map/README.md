# map

- map 是引用类型

> 不要使用 new，永远用 make 来构造 map

```go
var map1 map[keytype]valuetype
var map1 map[string]int

var mapLit map[string]int
mapLit = map[string]int{"one": 1, "two": 2}

var mapAssigned map[string]int
mapAssigned = mapLit

mapCreated := make(map[string]float32)
//相当于 mapCreated := map[string]float32{}

```

> 在声明的时候不需要知道 map 的长度，map 是可以动态增长的。
>
> 未初始化的 map 的值是 nil。

- key 可以是任意可以用 == 或者 != 操作符比较的类型,比如 string、int、float

> 指针和接口类型可以
>
> 所以数组、切片和结构体不能作为 key

- value 可以是任意类型

- map 传递给函数的代价很小

> 在 32 位机器上占 4 个字节，64 位机器上占 8 个字节, 无论实际上存储了多少数据
>
> 通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果你很在乎性能的话还是建议用切片来解决问题。

- map 也可以用函数作为自己的值

- map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
