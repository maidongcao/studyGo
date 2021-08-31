package main

import "fmt"

func main() {
	type T struct {
		a int
		b float32
		c string
	}
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%g\n", t.b) // %g 自动选择合适的输出，可以去除小数点后多余的0
	fmt.Printf("%f\n", t.b) // %f 按照浮点输出，可能多余的0，不会去掉
	fmt.Printf("%e\n", t.b) // %e 按照科学计数法输出

	fmt.Printf("%v\n", t)  // 输出value
	fmt.Printf("%+v\n", t) // 输出变量名+value
	fmt.Printf("%#v\n", t) // 输出类型+变量名+value
}
