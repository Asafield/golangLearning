package class2

import (
	"fmt"
	"strings"
)

func Bibao() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder() func(delta int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
func fibonacciB() func() int {
	x := 0
	y := 1
	return func() int {
		defer func() {
			x, y = y, x+y
		}()
		return x
	}
}

func DisplayFibonacci() {
	f := fibonacciB()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
