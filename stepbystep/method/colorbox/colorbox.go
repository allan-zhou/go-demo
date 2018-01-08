package main

import (
	"fmt"
	"strconv"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b Box) ToString() string {
	width := strconv.FormatFloat(b.width, 'g', -1, 64)
	height := strconv.FormatFloat(b.height, 'g', -1, 64)
	depth := strconv.FormatFloat(b.depth, 'g', -1, 64)
	color := b.color.ToString()

	return "width:" + width + " height:" + height + " depth:" + depth + " color:" + color
}

func (c Color) ToString() string {
	strs := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strs[c]
}

func (bl BoxList) BiggestBox() Box {
	v := 0.00
	var result Box
	for _, box := range bl {
		if bv := box.Volume(); bv > v {
			v = box.Volume()
			result = box
		}
	}

	return result
}

func (b *Box) setColor(c Color) {
	b.color = c
}

func (bl BoxList) setAllColor(c Color) {
	for i,_ := range bl {
		bl[i].setColor(c)
	}

	// 使用以下代码运行结果错误
	// for _,v := range bl {
	// 	v.setColor(c)
	// }
}

func printBox(boxes BoxList) {
	for i, v := range boxes {
		fmt.Printf("box'%d volume is: %v ", i, v.Volume())
		fmt.Printf("  color is: %v \n", v.color.ToString())
	}
}

func main() {
	boxes := BoxList{
		Box{3.1, 3, 3, RED},
		Box{4.3, 4, 4, BLACK},
		Box{6.3, 6, 6, WHITE},
		Box{5.1, 5, 5, BLUE},
	}

	fmt.Printf("total boxes: %d \n", len(boxes))
	printBox(boxes)

	biggest := boxes.BiggestBox()
	fmt.Printf("biggest box is:%v \n", biggest.ToString())
	biggest.setColor(BLACK)
	fmt.Printf("biggest box is:%v \n", biggest.ToString())

	boxes.setAllColor(RED)
	printBox(boxes)
}
