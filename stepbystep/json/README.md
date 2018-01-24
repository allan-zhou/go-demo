# JSON 处理

## 解析

```go
func Unmarshal(data []byte, v interface{}) error
```

- 解析到结构体
- 解析到interface

## 生成JSON

```go
func Marshal(v interface{}) ([]byte, error)
```

针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:

字段的tag是"-"，那么这个字段不会输出到JSON
tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中
tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串

## 编码和解码流

json 包提供 Decoder 和 Encoder 类型来支持常用 JSON 数据流读写
