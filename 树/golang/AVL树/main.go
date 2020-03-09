package main

// AVl树
type TreeNode struct {
	Val         int
	left, right *TreeNode
	height      int
}

type AVLTree struct {
	RootNode *TreeNode
}
func max(n1,n2 int) int{
	if n1 > n2{
		return n1
	}
	return n2
}
func (self *AVLTree) Put(key int) {
	//if self.RootNode != nil {
	//	self.RootNode = &TreeNode{Val: key}
	//} else {
	//	self.RootNode = self.put(key, self.RootNode)
	//}
	self.RootNode = self.put(key, self.RootNode)
}
func (self *AVLTree) height(node *TreeNode)int{
	if node == nil{
		return -1
	}
	return node.height
}
func (self *AVLTree) put(key int, node *TreeNode) (root *TreeNode) {
	if node == nil{
		root = &TreeNode{Val: key}
	}else if key < node.Val{

	}else if key > node.Val{

	}
	root.height = max(self.height(root.left), self.height(root.right)) + 1
	return
}
