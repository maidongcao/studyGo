package interface_test

import (
	"fmt"
	"sort"
	"testing")


type Sequence []int

// 实现sort.Interface 的三个方法 可以实现排序
func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}

func TestSort(t *testing.T){
	s := Sequence([]int{9,1,3,4,5})
	str := s.String()
	fmt.Println(str)
}
