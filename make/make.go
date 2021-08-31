package main

import "fmt"

//与new的区别 make 只能作用于map slice chan,返回类型的本身不是类型指针。
func main(){
	// 声明和分配切片,len为1，cap为10的slice
	s := make([]int32, 1, 10)
	fmt.Printf("make([]int32, 1, 10)'type:%T\n", s)
	// map
	m := make(map[int]int)
	fmt.Printf("make(map[int]int)'type:%T\n", m)
	// chan 的类型是int，通道内最多存10个元素
	c := make(chan int, 10)
	fmt.Printf("make(chan int, 10)'type:%T\n", c)
}