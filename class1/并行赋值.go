//该代码是对第一次学习golang涉及的知识点的总结，包含了并行赋值、iota、init函数
package main

import (
	"class1/trans"
	"fmt"
)

var a = 1
var b = 2

const (
	x  = "as" //97
	y  = iota //1
	z  = 'b'  //98
	z1        //98
)
const (
	zz = iota //0
)

func main() {
	fmt.Printf("hello,world\n")
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(z1)
	fmt.Printf("%v", zz)
	fmt.Printf("%v", trans.Pi)
	fmt.Println("###############################################")
	f1()
	fmt.Println(a == z1)
}
func f1() {
	a := 4
	fmt.Println(a)
	f2()
}

func f2() {
	fmt.Println(a)
}
