package main

import "fmt"

// 双向链表

type DataObject interface{}
type Node struct {
	Val       DataObject
	Next, Pre *Node
}

type DoubleLinkList struct {
	Head, Tail *Node
}

func (self *DoubleLinkList) Add(node *Node, index int)(err error) {
	if self.Head == nil {
		self.Head = node
		return
	}
	linklist_Size := self.Size()
	if index < 0 || index > linklist_Size -1 {
		err = fmt.Errorf("Wrong Index Error")
		return
	}
	p := self.Head
	for i := 0; i < index-1; i++ {
		p = p.Next
	}
	node.Next = p.Next
	node.Pre = p
	p.Next = node
	
	return
}
func (self *DoubleLinkList) Size() int {
	i := 0
	p := self.Head
	for p != nil {
		i += 1
		p = p.Next
	}
	return i
}
func (self *DoubleLinkList) PrintDoubleLinkList() {
	p := self.Head
	for p != nil {
		fmt.Printf("%v\t", p.Val)
		p = p.Next
	}

}

func (self *DoubleLinkList) Index(index int)(node *Node,err error){
	linklist_Size := self.Size()
	if index < 0 || index > linklist_Size -1 {
		err = fmt.Errorf("Wrong Index Error")
		return
	}
	p := self.Head
	for i := 0; i < index-1; i++ {
		p = p.Next
	}
	node = p.Next
	return
}