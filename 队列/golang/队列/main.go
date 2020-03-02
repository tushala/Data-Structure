package 队列

import "fmt"

// leetcode 232
//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。

type MyQueue struct {
	top int
	arr []int
}

func Constructor() MyQueue {
	arr := make([]int, 0)
	m := MyQueue{
		top: 0,
		arr: arr,
	}
	return m
}

//将一个元素放入队列的尾部。
func (self *MyQueue) Push(x int) {
	self.arr = append(self.arr, x)
	self.top++
}

//从队列首部移除元素
func (self *MyQueue) Pop() int {
	p := self.arr[0]
	self.arr = self.arr[1:]
	self.top--
	return p
}

//返回队列首部的元素
func (self *MyQueue) Peek() int {
	return self.arr[0]
}

//返回队列是否为空
func (self *MyQueue) isEmpty() bool {
	return self.top == 0
}

// 清空
func (self *MyQueue) Clear() {
	self.top = 0
	arr := make([]int, 0)
	self.arr = arr
}


func main() {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	fmt.Println(m.Peek())
	fmt.Println(m.Pop())
	fmt.Println(m.isEmpty())
}
