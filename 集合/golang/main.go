package main

//  https://www.cnblogs.com/golove/p/5840145.html 抄写
//  用于多例程
import (
	"fmt"
	"sync"
)

type Set struct {
	sync.RWMutex
	m map[int]bool
}

func New(items ...int) *Set {
	s := &Set{
		m: make(map[int]bool, len(items)),
	}
	s.Add(items...)
	return s
}
func (self Set) Size() int {
	self.Lock()
	defer self.Unlock()
	return len(self.m)
}

// 添加元素
func (self Set) Add(items ...int) {
	self.Lock()
	defer self.Unlock()
	for _, v := range (items) {
		self.m[v] = true
	}
}
func (self Set) Remove(items ...int) {
	self.Lock()
	defer self.Unlock()
	for _, v := range (items) {
		delete(self.m, v)
	}
}
func (self *Set) Has(items ...int) bool {
	self.Lock()
	defer self.Unlock()
	for _, v := range items {
		if _, ok := self.m[v]; !ok {
			return false
		}
	}
	return true
}
func (self *Set) Clear() {
	self.Lock()
	defer self.Unlock()
	self.m = map[int]bool{}
}
func (self *Set) IsEmpty() bool {
	self.Lock()
	defer self.Unlock()
	return self.Size() == 0
}
func (self *Set) Union(sets ...*Set){
	// 为了防止多例程死锁，不能同时锁定两个集合
	// 所以这里没有锁定 s，而是创建了一个临时集合
	r := self.Duplicate()
	// 获取并集
	for _, set := range sets {
		set.Lock()
		for e := range set.m {
			r.m[e] = true
		}
		set.Unlock()
	}
	// 将结果转入 self
	self.Lock()
	defer self.Unlock()
	self.m = map[int]bool{}
	for e := range r.m {
		self.m[e] = true
	}
}
func (self *Set) Minus(sets ...*Set){
	r := self.Duplicate()
	for _, set := range sets{
		set.Lock()
		for e := range set.m {
			if _, ok := r.m[e];ok{
				delete(r.m, e)
			}
		}
		set.Unlock()
	}
	// 将结果转入 self
	self.Lock()
	defer self.Unlock()

}
// 创建副本
func (self *Set) Duplicate()*Set{
	self.Lock()
	defer self.Unlock()
	r := &Set{
		m: make(map[int]bool, len(self.m)),
	}
	for e := range self.m {
		r.m[e] = true
	}
	return r
}
// 并集
func Union(sets ...*Set) *Set{
	if len(sets) == 0{
		return New()
	}else if len(sets) == 1 {
		return sets[0]
	}
	res := sets[0].m
	for _, set := range sets[1:]{
		for e := range set.m{
			if _, ok := res[e]; !ok{
				res[e] = true
			}
			//fmt.Println(123, res)
		}

	}
	return &Set{m:res}
}
// 差集
func Minus(sets ...*Set) *Set{
	if len(sets) == 0{
		return New()
	}else if len(sets) == 1 {
		return sets[0]
	}
	res := sets[0].m
	for _, set := range sets[1:]{
		for e := range set.m{
			delete(res, e)
		}
	}
	return &Set{m:res}
}
func main() {
	s1 := New(1, 2, 3, 4, 5, 6, 7, 8)
	s2 := New(3, 4, 5, 6)
	s3 := New(1, 2, 5, 6, 8, 9)
	//r1 := Union(s1, s2, s3)     // 获取并集
	//r2 := Minus(s1, s2, s3)     // 获取差集
	//fmt.Println(r1)
	//fmt.Println(r2)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			for i := 0; i < 100; i++ {
				s2.Union(s1) // 获取并集
				fmt.Printf("%v\n", s2.m)

				s2.Minus(s3) // 获取差集
				fmt.Printf("%v\n", s2.m)
				//s2.Intersect(s1) // 获取交集
				//fmt.Printf("%2v：s2 * %v = %v\n", n, s1.SortedList(), s2.SortedList())
				//
				//s2.Complement(s1) // 获取 s2 相对于 s1 的补集
				//fmt.Printf("%2v：%v / s2 = %v\n", n, s1.SortedList(), s2.SortedList())
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
