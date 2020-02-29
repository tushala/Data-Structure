package main

import "fmt"
func PrintJosephus(l *LinkList, size int){
	p := l.Head
	for i:=0;i<size+1;i++{
		fmt.Print(p,"  ")
		p = p.next
	}
}
// 整合环
func MakeJosephus(link *LinkList, head * Node, size int){
	p := link.Head
	for i:=0;i<size-1;i++{
		p = p.next
	}
	p.next = head
}
// 约瑟夫环
func JosephusSolve(num int, step int) {
	var link LinkList
	for i := 1; i < num+1; i++ {
		node := &Node{
			data: i,
			next: nil,
		}
		err := link.AddNode(node, i-1)
		if err != nil {
			return
		}
	}
	p := link.Head
	for p.next != nil {
		p = p.next
	}
	MakeJosephus(&link,link.Head,link.Size)
	//PrintJosephus(&link, link.Size)
	p = p.next
	for link.Size != step {
		for i := 0; i < step; i++ {
			p = p.next
		}
		err := link.DelNode(p)
		MakeJosephus(&link, link.Head, link.Size)
		p = p.next
		//PrintJosephus(p, link.Size)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// 输出
	p = link.Head
	for i:=0;i<step;i++{
		fmt.Println(p.data)
		p = p.next
	}
}
