# read data

## 读取用户输入

最简单的办法是使用 fmt 包提供的 Scan 和 Sscan 开头的函数

也可以使用 bufio 包提供的缓冲读取（buffered reader）来读取数据

## 文件读写

在 Go 语言中，文件使用指向 os.File 类型的指针来表示的，也叫做文件句柄。
