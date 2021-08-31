package main

import (
	"fmt"
	"sync"
)

type user struct {
	lock sync.Mutex
	name string
	age int
}

func main() {
	// panic栗子
	// var i *int64 // i 是一个nil指针，没有分配空间
	// *i = 10 // 取nil指针地址，会保存。可以使用i:= new(int64) 分配一个值为0的地址
	p := new(user) // new 会把结构体的值都设为默认值，然后返回一个指针，不需要逐个指定。
	// 等价于&new{},其实new()不常用
	fmt.Println(p)
	var v user
	fmt.Println(v)

	// 返回任意类型的指针
	m := new(map[int]string)
	fmt.Printf("new(map[int]string)'type:%T", m)
}
