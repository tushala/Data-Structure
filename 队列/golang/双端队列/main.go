package main

import (
	"errors"
	"fmt"
)


var (
	Empty_Queue_Error = errors.New("Empty Queue")
)


type dataObject interface{}

type DoubleEndedQueue struct {
	top int
	arr []dataObject
}

func Constructor() *DoubleEndedQueue {
	arr := make([]dataObject, 0)
	dq := &DoubleEndedQueue{
		top: 0,
		arr: arr,
	}
	return dq
}
// 元素的数量
func(self *DoubleEndedQueue) Size() int{
	return self.top
}

func(self *DoubleEndedQueue) isEmpty() bool {
	return self.top == 0
}

func (self *DoubleEndedQueue) Clear(){
	self.top = 0
	arr := make([]dataObject, 0)
	self.arr = arr
}
//头插
func (self *DoubleEndedQueue) HeadPush(x interface{}) {
	arr := make([]dataObject,1)
	arr[0] = x
	self.arr = append(arr, self.arr...)
	self.top++
}
//尾插
func (self *DoubleEndedQueue) TailPush(x interface{}) {
	self.arr = append(self.arr, x)
	self.top++
}

func (self *DoubleEndedQueue) HeadPop()(data interface{}, err error) {
	if !self.isEmpty(){
		h := self.arr[0]
		self.arr = self.arr[1:]
		self.top --
		return h, nil
	}else{
		panic(Empty_Queue_Error)
	}
}

func (self *DoubleEndedQueue) TailPop()(data interface{}, err error) {
	if !self.isEmpty() {
		t := self.arr[self.top-1]
		self.arr = self.arr[:self.top-1]
		self.top--
		return t, nil
	}else{
		panic(Empty_Queue_Error)
	}

}

// 获取头结点
func (self *DoubleEndedQueue) Peek()(data interface{}, err error) {
	if !self.isEmpty(){
		return self.arr[0], err
	}else{
		panic(Empty_Queue_Error)
	}

}

// 获取尾结点
func (self *DoubleEndedQueue) Rear() (data interface{}, err error) {
	if !self.isEmpty(){
		return self.arr[self.top-1], err
	}else{
		panic(Empty_Queue_Error)
	}
}

func main() {
	d := Constructor()
	d.HeadPush(1)
	d.HeadPush(2)
	d.HeadPush(3)
	d.TailPush(4)
	fmt.Println(d.HeadPop())
	fmt.Println(d.HeadPop())
	fmt.Println(d.HeadPop())
	fmt.Println(d.HeadPop())
	fmt.Println(d.isEmpty())
}