package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const count = 10

type newUnionSet struct {
	numNode []int
	set     []int
}

func NewUnionSet(size int) *newUnionSet {
	buf1 := make([]int, size)
	for i := 0; i < size; i++ {
		buf1[i] = i
	}
	buf2 := make([]int, size)
	for i := 0; i < size; i++ {
		buf2[i] = 1 // 初始节点数量均为1
	}
	return &newUnionSet{
		numNode: buf2,
		set:     buf1,
	}
}
func (self *newUnionSet) GetSize() int {
	return len(self.set)
}

func (self *newUnionSet) GetID(p int) (int, error) {
	if p < 0 || p > len(self.set) {
		return 0, errors.New(
			"failed to get ID,index is illegal.")
	}
	return self.getRoot(p), nil
}

func (self *newUnionSet) getRoot(p int) int {
	for p != self.set[p] {
		self.set[p] = self.set[self.set[p]] // 路径压缩
		p = self.set[p]
	}
	return p
}

func (self *newUnionSet) IsConnected(p, q int) (bool, error) {
	if p < 0 || p > len(self.set) || q < 0 || q > len(self.set) {
		return false, errors.New(
			"error: index is illegal.")
	}
	return self.getRoot(self.set[p]) == self.getRoot(self.set[q]), nil
}

func (self *newUnionSet) Union(p, q int) error {
	if p < 0 || p > len(self.set) || q < 0 || q > len(self.set) {
		return errors.New(
			"error: index is illegal.")
	}
	pRoot := self.getRoot(p)
	qRoot := self.getRoot(q)

	if pRoot != qRoot {
		if self.numNode[pRoot] < self.numNode[qRoot] {
			self.set[pRoot] = qRoot
			self.numNode[qRoot] += self.numNode[pRoot]
		} else {
			self.set[qRoot] = pRoot
			self.numNode[pRoot] += self.numNode[qRoot]
		}
	}
	return nil
}

func test3(nums []int) {
	s := NewUnionSet(count)
	fmt.Println(s)
	t := time.Now()
	for i := 1; i < 2*count; i++ {
		s.GetID(nums[i])
		s.Union(nums[i-1], nums[i])
	}
	fmt.Println(s)
	fmt.Println("第三版时间: ", t.Sub(time.Now()))
}

func main() {

	nums := []int{}
	for i := 0; i < 2*count; i++ {
		nums = append(nums, rand.Intn(count))
	}

	test3(nums)
}
