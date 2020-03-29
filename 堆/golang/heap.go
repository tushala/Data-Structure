package main

import "fmt"

// 小顶堆
type Heap struct {
	s []int
}

func get_child(i int) (child int) {
	child = (i - 1) / 2
	return child
}
func (self *Heap) size() int {
	return len(self.s)
}
func (self *Heap) get() int {
	if self.size() > 0 {
		return self.s[0]
	} else {
		panic("empty heap")
	}
}
func (self *Heap) swap(i, j int) {
	self.s[i], self.s[j] = self.s[j], self.s[i]
}
func (self *Heap) less(i, j int) bool {
	return self.s[i] < self.s[j]
}
func New() *Heap {
	s := make([]int, 0)
	return &Heap{s: s}
}
func (self *Heap) Put(x int) {
	if len(self.s) == 0 {
		self.s = append(self.s, x)
		return
	} else {
		self.put(x, 0, len(self.s))
	}
}
func (self *Heap) put(x, i0, n int) {
	self.s = append(self.s, x)
	i := n
	for i > 0 {
		child := get_child(i)
		if self.less(child, i) {
			break
		} else {
			self.swap(child, i)
		}
		i = child
	}
}
func (self *Heap) Replace(x int) {
	// 代替堆顶元素
	if self.size() < 2 {
		self.s = append(self.s, x)
	} else {
		self.replace(x, 0, len(self.s))
	}
}
func (self *Heap) replace(x, i0, n int){
	self.s[0] = x
	i := i0
	for i < n {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		if !self.less(i, j1){
			self.swap(i,j1)
			i = j1
		}
		if j2 := j1+1; j2 < self.size() && !self.less(i, j2){
			self.swap(i,j2)
			i = j2
		}
		i  = j1
	}
}
func (self *Heap) Remove(){
	// 移除头
	x := self.Pop()
	self.Replace(x)
}
func (self *Heap) Pop()int{
	if self.size() > 0 {
		s := self.s[self.size()-1]
		self.s = self.s[:self.size()-1]
		return s
	} else {
		panic("empty heap")
	}

}
func TopK(k int, input []int) {
	h := New()
	for i := 0; i < len(input); i++ {
		if h.size() < k {
			h.Put(input[i])
		} else {
			if input[i] > h.get() {
				h.Replace(input[i])
			}
		}
	}
	fmt.Println(h.s)

}
func main() {
	//h := New()
	testSlice := []int{2, 3, 1, 4, 5, 6, 2, 1, 23, 43, 1, 32, 3, 5, 2, 1, 8, 54, 4, 0}
	//testSlice := []int{4, 5, 3, 2, 1}
	//testSlice := []int{2, 3, 1}
	//for _, t := range testSlice {
	//	h.Put(t)
	//}
	//fmt.Println(h.s)
	//h.Replace(0)
	//fmt.Println(h.s)
	TopK(5, testSlice)
}
