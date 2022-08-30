package class2

import (
	"fmt"
	"unsafe"
)

func farray(a [3]int) { fmt.Println(a) }
func fp(a *[3]int)    { fmt.Println(a) }

func Array() {
	var ar [3]int
	var ar2 = new([3]int)
	a := []int{1, 23, 4}
	farray(*ar2) // passes a copy of ar
	fp(&ar)      // passes a pointer to ar
	fmt.Println(a)
}
func ArrayTest() {
	a := [5]int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

	c := new([50]int)[0:50]
	b = a[2:4]
	b[1] = 1
	c = (&a)[0:5]
	b = b[:cap(b)]

	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
func ArrayItems() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2

	}
	fmt.Println(items)
}
func SliceExpand() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))

	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d  \n", i, slice1[i])
	}
}
func CopyAppendSlice() {
	sl_from := []int{1, 2, 3}
	sl_from = append(sl_from, 4, 5)
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1, 2, 3}
	fmt.Println(cap(sl3))
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(cap(sl3))
	fmt.Println(sl3)

}
func StringsUnicode() {
	s := "我是谁,那么这是为什么呢？asafield"
	for i, c := range s {
		fmt.Printf("%d:%c", i, c)
	}
}
