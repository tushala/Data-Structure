package main

import "fmt"

// 单向列表

type DataObject interface{}
type Node struct {
	data DataObject
	next *Node
}

type LinkList struct {
	Head *Node
	Size int
}

func (self *LinkList) AddNode(node *Node, index int) (err error) {

	if self.Head == nil {
		self.Head = node
		self.Size++
		return
	}
	linklist_Size := self.Size
	if index < 0 || index > linklist_Size {
		err = fmt.Errorf("Wrong Index Error")
		return
	}
	p := self.Head
	for i := 0; i < index-1; i++ {
		p = p.next
	}
	node.next = p.next
	p.next = node
	self.Size++
	return
}
func (self *LinkList) DelNodeByIndex(index int) (err error) {
	linklist_Size := self.Size
	if index < 0 || index > linklist_Size {
		err = fmt.Errorf("Wrong Index Error")
		return
	}
	p := self.Head
	for i := 0; i < index; i++ {
		p = p.next
	}
	self.Size--
	p.next = p.next.next
	return
}
func (self *LinkList) DelNode(node *Node) (err error) {
	p := self.Head
	if node == self.Head {
		self.Head = p.next
		self.Size--
		return
	}
	hasnode := false
	for i := 0; i < self.Size; i++ {
		if p.next == node {
			hasnode = true
			p.next = p.next.next
			self.Size--
		} else {
			p = p.next
		}

	}
	if !hasnode {
		err = fmt.Errorf("Wrong Node Del")
		return
	}
	return
}

func (self *LinkList) PrintLinkList() {
	p := self.Head
	for p != nil {
		fmt.Printf("%v\t", p.data)
		p = p.next
	}
}
func (self *LinkList) Index(index int) (node *Node, err error) {
	linklist_Size := self.Size
	if index < 0 || index > linklist_Size {
		err = fmt.Errorf("Wrong Index Error")
		return
	}
	p := self.Head
	for i := 0; i < index-1; i++ {
		p = p.next
	}
	node = p.next
	return
}

//func main() {
//	var link LinkList
//	for i := 0; i < 4; i++ {
//		node := &Node{
//			data: fmt.Sprintf("node%d", i),
//			next: nil,
//		}
//		err := link.AddNode(node, i)
//		if err != nil {
//			fmt.Println("link.AddNode err", err)
//			return
//		}
//	}
//
//	link.PrintLinkList()
//	err := link.DelNode(2)
//	if err != nil {
//		fmt.Println("link.DelNode err", err)
//		return
//	}
//	fmt.Println()
//	link.PrintLinkList()
//	fmt.Println(link.Index(0))
//}
