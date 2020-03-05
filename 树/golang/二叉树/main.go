package main

import "fmt"

type TreeNode struct {
	Val         int
	left, right *TreeNode
}

type Tree struct {
	RootNode *TreeNode
}

//添加节点
func add(t *TreeNode, value int) *TreeNode {
	if t == nil {
		t = new(TreeNode)
		t.Val = value
		return t
	}
	if value < t.Val {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// 查找最小值
func (self *Tree) FindMin() int {
	p := self.RootNode
	for p.left != nil {
		p = p.left
	}
	return p.Val
}

// 查找最大值
func (self *Tree) FindMax() int {
	p := self.RootNode
	for p.left != nil {
		p = p.right
	}
	return p.Val
}

//前序遍历
func (self *Tree) PreorderTravel(node *TreeNode) {
	if node == nil{
		return
	}
	fmt.Print(node.Val,"\t")
	self.PreorderTravel(node.left)

	self.PreorderTravel(node.right)
}

//中序遍历
func (self *Tree) InorderTravel(node *TreeNode) {
	if node == nil{
		return
	}
	self.InorderTravel(node.left)
	fmt.Print(node.Val,"\t")
	self.InorderTravel(node.right)
}

//后序遍历
func (self *Tree) PostorderTravel(node *TreeNode) {
	if node == nil{
		return
	}
	self.PostorderTravel(node.left)
	self.PostorderTravel(node.right)
	fmt.Print(node.Val,"\t")
}

//层序遍历
func (self *Tree) LevelorderTravel(node *TreeNode) {
	if node == nil{
		return
	}
	fmt.Println(node.Val)

}

func main() {
	tree := &Tree{}
	insert_slice := []int{17, 5, 34, 2, 11, 35, 29, 38, 9, 16, 8}
	for _, i := range insert_slice {
		tree.RootNode = add(tree.RootNode, i)
	}
	fmt.Println(tree.FindMin())
	fmt.Println(tree.FindMax())
	tree.InorderTravel(tree.RootNode)
}
