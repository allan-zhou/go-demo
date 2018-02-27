# error and testing

Go 没有像 Java 和 .NET 那样的 try/catch 异常机制：不能执行抛异常操作。但是有一套 defer-panic-and-recover 机制

Go 的设计者觉得 try/catch 机制的使用太泛滥了，而且从底层向更高的层级抛异常太耗费资源。他们给 Go 设计的机制也可以 “捕捉” 异常，但是更轻量，并且只应该作为（处理错误的）最后的手段。

Go 是怎么处理普通错误的呢？  
通过在函数和方法中返回错误对象作为它们的唯一或最后一个返回值——如果返回 nil，则没有错误发生——并且主调（calling）函数总是应该检查收到的错误。

**Go 检查和报告错误条件的惯有方式**：

- 产生错误的函数会返回两个变量，一个值和一个错误码；如果后者是 nil 就是成功，非 nil 就是发生了错误。
- 为了防止发生错误时正在执行的函数（如果有必要的话甚至会是整个程序）被中止，在调用函数后必须检查错误。

## 错误处理

Go 有一个预先定义的 error 接口类型

```go
type error interface {
    Error() string
}
```

任何时候当你需要一个新的错误类型，都可以用 errors（必须先 import）包的 errors.New 函数接收合适的错误信息来创建，像下面这样：

```go
err := errors.New(“math - square root of negative number”)
```

## 运行时异常和panic

当发生像数组下标越界或类型断言失败这样的运行错误时，Go 运行时会触发运行时 panic，伴随着程序的崩溃抛出一个 runtime.Error 接口类型的值。这个错误值有个 RuntimeError() 方法用于区别普通错误。

panic 可以直接从代码初始化：当错误条件（我们所测试的代码）很严苛且不可恢复，程序不能继续运行时，可以使用 panic 函数产生一个中止程序的运行时错误。  
panic 接收一个任意类型的参数，通常是字符串，在程序死亡时被打印出来。Go 运行时负责中止程序并给出调试信息。  
[panic](./panic/panic.go)

## Go panicking：

在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。  
这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是 panicking。  
[panicking](./panicking/panicking.go)

## 从 panic 中恢复（Recover）

这个（recover）内建函数被用于从 panic 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。  
recover 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil，且没有其它效果。

总结：
panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止。
defer-panic-recover 在某种意义上也是一种像 if，for 这样的控制流机制。  
[panic_recover](./panic_recover/panic_recover.go)

## 自定义包中的错误处理和 panicking

这是所有自定义包实现者应该遵守的最佳实践：  
1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()  
2）向包的调用者返回错误值（而不是 panic）。  

在包内部，特别是在非导出函数中有很深层次的嵌套调用时，对主调函数来说，用 panic 来表示应该被翻译成错误的错误场景是很有用的（并且提高了代码可读性）。  
[panic_package](./panic_package/panic_package.go)

## Go 中的单元测试和基准测试

名为 testing 的包被专门用来进行自动化测试，日志和错误报告。并且还包含一些基准测试函数的功能。

对一个包做（单元）测试，需要写一些可以频繁（每次更新后）执行的小块测试单元来检查代码的正确性。于是我们必须写一些 Go 源文件来测试代码。  
测试程序必须属于被测试的包，并且文件名满足这种形式 *_test.go，所以测试代码和包中的业务代码是分开的。

_test 程序不会被普通的 Go 编译器编译，所以当放应用部署到生产环境时它们不会被部署；  
只有 gotest 会编译所有的程序：普通程序和测试程序。

测试文件中必须导入 "testing" 包，并写一些名字以 TestZzz 打头的全局函数，这里的 Zzz 是被测试函数的字母描述，如 TestFmtInterface，TestPayEmployees 等。

测试函数必须有这种形式的头部：

```go
func TestAbcde(t *testing.T)
```

T 是传给测试函数的结构类型，用来管理测试状态，支持格式化测试日志，如 t.Log，t.Error，t.ErrorF 等。  
在函数的结尾把输出跟想要的结果对比，如果不等就打印一个错误。成功的测试则直接返回。
[oddeventest](./oddeventest/oddeven.go)

## 性能调试：分析并优化 Go 程序

### 时间和内存消耗

可以用这个便捷脚本 xtime 来测量：

```shell
#!/bin/sh
/usr/bin/time -f '%Uu %Ss %er %MkB %C' "$@"
```

在 Unix 命令行中像这样使用 xtime goprogexec，这里的 progexec 是一个 Go 可执行程序，  
这句命令行输出类似：56.63u 0.26s 56.92r 1642640kB progexec，分别对应用户时间，系统时间，实际时间和最大内存占用。