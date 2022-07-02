package main

import "fmt"

var a int = 1
var b int = 2

const (
	x  = 'a'  //97
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
	fmt.Println(zz)
}
