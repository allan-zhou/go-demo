package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// ok, _ := regexp.Match(pat, []byte(searchIn))
// ok, _ := regexp.MatchString(pat, searchIn)

func main() {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern := "[0-9]+.[0-9]+" //正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v, 'f', 1, 32)
	}

	if ok, _ := regexp.Match(pattern, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	reg, _ := regexp.Compile(pattern)

	// 将匹配到的部分替换为 “##.#”
	str := reg.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	// 参数为函数
	str2 := reg.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
