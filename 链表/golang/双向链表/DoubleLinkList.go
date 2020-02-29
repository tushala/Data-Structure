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

func (self *DoubleLinkList) AddNode(node *Node, index int)(err error) {
	if self.Head == nil {
		self.Head = node
		return
	}
	linklist_Size := self.Size()
	//fmt.Println(12345, linklist_Size)
	if index < 0 || index > linklist_Size {
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
		//fmt.Println(666, p)
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
	if index < 0 || index > linklist_Size {
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
// 差 del

func main() {
	var link DoubleLinkList
	for i := 0; i < 4; i++ {
		node := &Node{
			Val: fmt.Sprintf("node%d", i),
			Next: nil,
			Pre:nil,
		}
		err := link.AddNode(node, i)
		if err != nil {
			fmt.Println("link.AddNode err", err)
			return
		}
	}
	link.PrintDoubleLinkList()
}