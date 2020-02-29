package main

import (
	"container/list" // 环形链表
	"fmt"
)

func main() {
	l := list.New() // 创建一个新list
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) //输出list的值,01234
	}
	fmt.Println("")
	fmt.Println(l.Front().Value) //输出首部元素的值,0
	fmt.Println(l.Back().Value)  //输出首部元素的值,4

	l.InsertAfter(6, l.Front())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //061234
	}
	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换
	fmt.Println()
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //601234
	}
	fmt.Println()
	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //460123
	}
	fmt.Println()

	l2 := list.New()
	l2.PushBack(1)
	fmt.Println(l2.Len())
	l2.Init() // 清空
	fmt.Println(l2.Len())

}
