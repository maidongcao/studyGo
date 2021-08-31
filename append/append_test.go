package append_test

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...) // y的类型是[]int, ...的意思是取y的元素 逐个输入
	fmt.Println(x)
	z := append(y, 1, 2, 3) // 可以传入多个数值
	fmt.Println(z)
}

func TestTriblePoint(t *testing.T) {
	y := []int{4, 5, 6}
	triblePoint(y...)
}

func triblePoint(y ...int) {
	fmt.Printf("y:%v\n, y:%T\n", y, y)
}
