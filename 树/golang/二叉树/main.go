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

// todo 删除节点
//func delete(t *TreeNode, value int) *TreeNode{
//	if t == nil{
//		return nil
//	}
//}
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
	if node == nil {
		return
	}
	fmt.Print(node.Val, "\t")
	self.PreorderTravel(node.left)

	self.PreorderTravel(node.right)
}

//中序遍历
func (self *Tree) InorderTravel(node *TreeNode) {
	if node == nil {
		return
	}
	self.InorderTravel(node.left)
	fmt.Print(node.Val, "\t")
	self.InorderTravel(node.right)
}

//后序遍历
func (self *Tree) PostorderTravel(node *TreeNode) {
	if node == nil {
		return
	}
	self.PostorderTravel(node.left)
	self.PostorderTravel(node.right)
	fmt.Print(node.Val, "\t")
}

//层序遍历
func (self *Tree) LevelorderTravel(node *TreeNode) {
	if node == nil {
		return
	}

	stack := make([]*TreeNode, 0, 1000)
	stack = append(stack, node)
	for len(stack) > 0 {
		node := stack[0]
		fmt.Print(node.Val, "\t")
		if node.left != nil {
			stack = append(stack, node.left)
		}
		if node.right != nil {
			stack = append(stack, node.right)
		}
		stack = stack[1:]
	}
}

//func (self *Tree)SuccessorNode(node *TreeNode)*TreeNode{
//	if node.right != nil{
//		curr := node.right
//		for curr.left != nil{
//			curr = curr.left
//		}
//		return curr
//	}else if node.right == nil && node.
//}
// 根据前序遍历重构二叉树（之一解）
func PreorderTravelTOTree(r []int) (node *TreeNode) {
	if len(r) == 0 {
		return nil
	}
	top := 0
	stack := make([]*TreeNode, 0, len(r))
	for i, v := range r {
		cur_node := &TreeNode{Val: v}
		if top == 0 || v < stack[top-1].Val {
			if top > 0 {
				back_node := stack[top-1]
				back_node.left = cur_node
			}

			top++
		} else {
			if stack[top-1].left!=nil{
				top -= 1
			}else{
				top -= 2
			}
			back_node := stack[top]
			back_node.right = cur_node
			//}
		}

		stack = append(stack, cur_node)
		if top == 0 && len(stack) > 1 {
			// 重回根节点
			rnode := PreorderTravelTOTree(r[i:])
			stack[0].right = rnode
			break
		}
	}
	return stack[0]
}
func main() {
	tree := &Tree{}
	insert_slice := []int{17, 5, 34, 2, 11, 35, 29, 38, 9, 16, 8}
	for _, i := range insert_slice {
		tree.RootNode = add(tree.RootNode, i)
	}
	//fmt.Println(tree.FindMin())
	//fmt.Println(tree.FindMax())
	//tree.InorderTravel(tree.RootNode)
	//fmt.Println()
	//tree.LevelorderTravel(tree.RootNode)
	s1 := []int{7, 4, 2, 1, 3, 5, 9, 8, 11, 10, 12}
	node := PreorderTravelTOTree(s1)
	fmt.Println(node.left)
	fmt.Println(node.left.left)
	fmt.Println(node.right)
	fmt.Println(node.right.right)
	fmt.Println(node.right.right.right)
}
