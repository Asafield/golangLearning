//用于学习golang中的map相关的内容
package class3

import (
	"fmt"
)

//map的声明
//var map1 map[keytype]valuetype
var map1 map[string]int

//简单使用map
func MapTest() {
	var mapLit map[string]int
	var mapAssigned map[string]float32 = make(map[string]float32)
	//动态的结构是可以在初始化之外也能这样直接给整个map赋值吗
	mapLit = map[string]int{"test": 1, "test2": 2}
	mapAssigned0 := map[string]float32{"test": 1.3, "test2": 2}
	mapAssigned = map[string]float32{"test": 1.3, "tset2": 2}

	fmt.Printf("第一个map的值：%v\n", mapLit)
	fmt.Printf("第二个map的值：%v\n", mapAssigned)
	fmt.Printf("第三个map的值：%v\n", mapAssigned0)
}
