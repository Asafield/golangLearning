package class2

func f() int {
	i := 5
	defer func() {
		i++
	}()

	return i
}

func f1() (result int) {
	defer func() {
		result++
	}()

	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()

	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)

	return 1
}

func Defer() {
	a := f()
	b := f1()
	c := f2()
	d := f3()
	println(a)
	println(b)
	println(c)
	println(d)
}
func Add_to_num(x int, y int) (r int) {
	return x + y
}
