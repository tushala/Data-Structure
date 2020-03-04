package main

type TreeNode struct {
	Val         int
	left, right *TreeNode
}

type Tree struct {
	RootNode *TreeNode
}

//添加节点
func (self *Tree) AddNode(val int) {
	t := &TreeNode{
		Val:   val,
		left:  nil,
		right: nil,
	}
	p := self.RootNode
	if p == nil{
		p = t
		return
	}
	for
}


//前序遍历
func (this *Tree) PreorderTravel() {

}

//中序遍历
func (this *Tree) InorderTravel() {

}

//后序遍历
func (this *Tree) PostorderTravel() {

}

//层序遍历
func (this *Tree) LevelorderTravel() {

}

func main() {
	tree := Tree{}
	tree.AddNode(0)
