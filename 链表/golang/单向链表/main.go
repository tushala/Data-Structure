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
}

func (self *LinkList) AddNode(node *Node, index int) (err error) {
	if self.Head == nil {
		self.Head = node
		return
	}
	linklist_Size := self.Size()
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

	return
}
func (self *LinkList) DelNode(index int) (err error) {
	linklist_Size := self.Size()
	if index < 0 || index > linklist_Size {
		err = fmt.Errorf("Wrong Index Error")
		return
	}
	p := self.Head
	for i := 0; i < index; i++ {
		p = p.next
	}
	p.next = p.next.next
	return
}
func (self *LinkList) Size() int {
	i := 0
	p := self.Head

	for p != nil {
		i += 1
		p = p.next
	}

	return i
}

func (self *LinkList) PrintLinkList() {
	p := self.Head
	for p != nil {
		fmt.Printf("%v\t", p.data)
		p = p.next
	}
}
func (self *LinkList) Index(index int) (node *Node, err error) {
	linklist_Size := self.Size()
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
func main() {
	var link LinkList
	for i := 0; i < 4; i++ {
		node := &Node{
			data: fmt.Sprintf("node%d", i),
			next: nil,
		}
		err := link.AddNode(node, i)
		if err != nil {
			fmt.Println("link.AddNode err", err)
			return
		}
	}

	link.PrintLinkList()
	err := link.DelNode(2)
	if err != nil {
		fmt.Println("link.DelNode err", err)
		return
	}
	fmt.Println()
	link.PrintLinkList()
	fmt.Println(link.Index(0))
}
