//该代码是对第一次学习golang涉及的知识点的总结，包含了并行赋值、iota、init函数
package main

import (
	"class1/trans"
	"fmt"
	"time"
)

var a float64 = 1.0
var b float64 = 1.0
var s string = "the world"

const (
	x     = "as" //97
	y     = iota //1
	z     = 'b'  //98
	z1           //98
	compl = complex(1.0, 1.0)
)
const (
	zz = iota //0
)

func main() {
	fmt.Printf("hello,world\n")
	//a, b = b, a
	s = "test"
	s = "fsfs"
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(z1)
	fmt.Printf("%v", zz)
	fmt.Printf("%v\n", trans.Pi)
	fmt.Println("###############################################")
	f1()
	fmt.Println(compl)
	trans.Display()
	fmt.Println("###############################################")
	trans.SliceDisplay()
	fmt.Println("###############################################")
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now())
	fmt.Println("###############################################")
	trans.BreakContinue()
}
func f1() {
	a := 4
	fmt.Println(a)
	f2()
}

func f2() {
	fmt.Println(a)
}
