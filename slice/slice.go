package main

import "fmt"

func main() {
	var s1 []int
	s1 = []int{1, 2, 3}
	var s2 []string
	s2 = []string{"fdsa", "fafa", "fdaf"} // 底层造了一个数组 然后baozh

	// 2 由数组得到切片
	a1 := [...]int{1,2,3,4,5,6,7,8,9}
	s3 := a1[0:4]
	fmt.Printf("s1 len:%d cap:%d, s2 len:%d s2 cap:%d, s3 len:%d s3 cap:%d\n",len(s1), cap(s1),len(s2), cap(s2),len(s3), cap(s3))
	s4 :=s3[3:6] // 切片的cap是切片开始的第一个元素到底层数组的最后一个元素的长度
	fmt.Printf("s4:%T, s4 len:%d cap:%d val:%v\n", s4, len(s4), cap(s4), s4)
}