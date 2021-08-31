package init_test

import (
	"fmt"
	"os"
	"testing"
)


func TestValues(t *testing.T) {
	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)

	fmt.Printf("home:%v\n", home)
	fmt.Printf("user:%v\n", user)
	fmt.Printf("gopath:%v\n", gopath)

}

// init() 主要用于1. 在使用程序前, 校验和修正程序 2. 幅值 3. 运行只执行一次的程序
func TestInit(t *testing.T){

}