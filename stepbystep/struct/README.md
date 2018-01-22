# struct and method

Go 通过类型别名（alias types）和结构体的形式支持用户自定义类型，或者叫定制类型。
一个带属性的结构体试图表示一个现实世界中的实体。

结构体是复合类型（composite types），当需要定义一个类型，它由一系列属性组成，每个属性都有自己的类型和值的时候，就应该使用结构体，它把数据聚集在一起。然后可以访问这些数据，就好像它是一个独立实体的一部分。

结构体也是值类型，因此可以通过 new 函数来创建。

组成结构体类型的那些数据称为**字段（fields）**。
每个字段都有一个类型和一个名字；在一个结构体中，字段名字必须是唯一的。

结构体的概念在软件工程上旧的术语叫 ADT（抽象数据类型：Abstract Data Type），在一些老的编程语言中叫 记录（Record），比如 Cobol，在 C 家族的编程语言中它也存在，并且名字也是 struct，在面向对象的编程语言中，跟一个无方法的轻量级类一样。

不过因为**Go 语言中没有类的概念**，因此在 Go 中结构体有着更为重要的地位。

## 定义

```go
type identifier struct {
    field1 type1
    field2 type2
    ...
}
```

结构体的字段可以是任何类型，甚至是结构体本身

数组可以看作是一种结构体类型，不过它使用下标而不是具名的字段。

## 使用new

```go
var t *T = new(T)

// 可以放在不同的行
var t *T
t = new(T)
// OR
t := new(T)

t.name = "name" //赋值

```

变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值。

声明 var t T 也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型T 。
t 通常被称做类型 T 的一个实例（instance）或对象（object）。

## 赋值&获取字段的值

使用**点号符**给字段赋值：structname.fieldname = value

同样的，使用点号符可以获取结构体字段的值：structname.fieldname

 Go 语言中这叫**选择器（selector）**

## 初始化

```go
// 字面量初始化实例 struct-literal
t := &type{"name", 20, "beijing" }
// 或者,混合字面量 composite literal syntax
var t2 type
t = type{"name", 20, "beijing" }
```

字面量写法，底层仍然调用new()。 new(T) 与 &T{} 等价

初始化方式：

```go
type Interval struct {
    start int
    end   int
}

intr := Interval{0, 3}            (A)
intr := Interval{end:5, start:1}  (B)
intr := Interval{end:5}           (C)
```

（A）中，值必须以字段在结构体定义时的顺序给出，& 不是必须的。
（B）显示了另一种方式，字段名加一个冒号放在值的前面，这种情况下值的顺序不必一致，并且某些字段还可以被忽略掉，就像（C）中那样。

结构体类型实例和一个指向它的指针的内存布局,如下图
![](./images/init-memory.png)

## 结构体内存布局

Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。
不像 Java 中的引用类型，一个对象和它里面包含的对象可能会在不同的内存空间中，这点和 Go 语言中的指针很像。
![](./images/struct-memory-layout.png)

## 递归结构体

- 链表

![](./images/linked-list.png)

- 二叉树

![](./images/binary-tree.png)

## 结构体转换

Go 中的类型转换遵循严格的规则。当为结构体定义了一个 alias 类型时，此结构体类型和它的 alias 类型都有相同的底层类型，可以转换。

## 结构体工厂

按惯例，工厂的名字以 new 或 New 开头。

```go
// struct
type File struct {
    fd      int     // 文件描述符
    name    string  // 文件名
}

// 工厂方法
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }

    return &File{fd, name}
}

//调用
f := NewFile(10, "./test.txt")
```

在 Go 语言中常常像上面这样在工厂方法里使用初始化来简便的实现构造函数

## 带标签的结构体

结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记。

标签的内容不可以在一般的编程中使用，只有包 reflect 能获取它。

## 匿名字段

结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。

匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。

可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。Go 语言中的继承是通过内嵌或组合来实现的，所以可以说，在 Go 语言中，相比较于继承，组合更受青睐。

> 在一个结构体中对于每一种数据类型只能有一个匿名字段。

## 方法

Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。

接收者类型可以是（几乎）任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，但是接收者不能是一个接口类型，因为接口是一个抽象定义，但是方法却是具体实现；

接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。

> 一个类型加上它的方法等价于面向对象中的一个类。
> 一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的。

## 函数和方法的区别

函数将变量作为参数：Function1(recv)

方法在变量上被调用：recv.Method1()

方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

## 多重继承

多重继承指的是类型获得多个父类型行为的能力，它在传统的面向对象语言中通常是不被实现的（C++ 和 Python 例外）。因为在类继承层次中，多重继承会给编译器引入额外的复杂度。但是在 Go 语言中，通过在类型中嵌入所有必要的父类型，可以很简单的实现多重继承。

## 面向对象编程总结

在 Go 中，类型就是类（数据和关联的方法）。Go 不知道类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。

在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫 组件编程（Component Programming）。

## 垃圾回收和 SetFinalizer

Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，即垃圾收集器（GC），会处理这些事情，它搜索不再使用的变量然后释放它们的内存。可以通过 runtime 包访问 GC 进程。

通过调用 runtime.GC() 函数可以显式的触发 GC，但这只在某些罕见的场景下才有用，比如当内存资源不足时调用 runtime.GC()，它会在此函数执行的点上立即释放一大片内存，此时程序可能会有短时的性能下降（因为 GC 进程在执行）