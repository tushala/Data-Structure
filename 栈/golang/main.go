package main

import (
	"errors"
	"fmt"
)

// 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int     //最大可以存放数个数
	Top    int     //栈顶 (长度)
	arr    [20]int // 数组模拟栈
}

//入栈
func (self *Stack) Push(val int) (err error) {
	//先判断栈是否满了
	if self.Top == self.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	self.Top++
	self.arr[self.Top] = val
	return
}

// 出栈
func (self *Stack) Pop() (val int, err error) {
	if self.Top == -1 {
		fmt.Println("stack empty!")
		return 0, errors.New("stack empty")
	}
	val = self.arr[self.Top]
	self.Top--
	return val, nil
}

// 遍历栈
func (self *Stack) PrintList() {
	if self.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	for i := self.Top; i == 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, self.arr[i])
	}
}
