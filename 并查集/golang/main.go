package main

import (
	"errors"
)

type unionSet struct {
	set []int
}

func New(size int) *unionSet {
	buf := make([]int, size)
	for i := 0; i < size; i++ {
		buf[i] = i // 初始时，所有节点均指向自己
	}
	return &unionSet{set: buf}
}

func (self *unionSet) GetSize() int {
	return len(self.set)
}
func (self *unionSet) GetID(p int) (int, error) {
	if p < 0 || p >= len(self.set) {
		return 0, errors.New("Wrong index")
	}
	return self.set[p], nil
}

func (self *unionSet) IsConnected(p, q int) (bool, error) {
	if p < 0 || p > len(self.set) || q < 0 || q > len(self.set) {
		return false, errors.New(
			"failed to get ID,index is illegal.")
	}
	return self.set[p] == self.set[q], nil
}
func (self *unionSet) Union(v1, v2 int) error {
	b, err := self.IsConnected(v1, v2)
	if err != nil {
		return err
	}
	if !b {
		pid := self.set[v1]
		qid := self.set[v2]
		for k, v := range self.set {
			if v == pid {
				self.set[k] = qid
			}
		}
	}
	return nil
}
