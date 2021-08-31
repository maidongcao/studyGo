package main

import "fmt"

// 数组的长度和元素类型是数据类型的一部分
// [3]bool 和[4]bool 不是相同的类型,不能直接比较
func main(){
	a1 := [3]bool{true, false, true}
	a2 := [4]bool{true, false, true, false}
	a10 := [...]int{0,1,2,3, 4} // 自动推断数据的大小
	fmt.Printf("a1:%T, a2:%T, a10:%T", a1, a2, a10)
}