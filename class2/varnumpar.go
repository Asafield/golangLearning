package class2

import "fmt"

func ParameterMulti() {
	x := min(1, 2, 3, 4)
	fmt.Printf("The minimum is : %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is : %d\n", x)
	fmt.Println("---------------------------------------")
	defera()
	fmt.Println("---------------------------------------")
	b()
	fmt.Println("---------------------------------------")
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Println("The results is : ", result)
	}
	fmt.Println("---------------------------------------")
	//printTen(10)
	fmt.Println("---------------------------------------")
}
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func defera() (v int) {
	i := 0
	defer func() {
		fmt.Println("the defer value is: ", i)
	}()
	i = 4
	return 4
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
func printTen(n int) (res int) {
	fmt.Println(n)
	if n == 0 {
		return
	}
	res = printTen(n - 1)
	fmt.Println(n)
	return
}
func multiNum(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = n * multiNum(n-1)
	}
	return
}
