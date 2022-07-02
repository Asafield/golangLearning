package main

import "fmt"

var a int = 1
var b int = 2
var c int = 3

func main() {
	fmt.Printf("hello,world\n")
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}
