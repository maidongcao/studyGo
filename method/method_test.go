package method_test

import "fmt"
import "testing"


//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}


type MyInt int

// 自定义类型 也可以绑定方法
func (MyInt) Hello(){
	fmt.Println("我是自定义int")
}

// 方法必须和接收者在同一个包内，不能给包外的结构体绑定方法
// func (int) Hello(){
// 	fmt.Println("我是int")
// }

func TestPerson(t *testing.T){
	p1 := NewPerson("hezhuhui", 25)
	p1.Dream()

}

func TestSelfType(t *testing.T){
	myInt := new(MyInt)
	myInt.Hello()
}
