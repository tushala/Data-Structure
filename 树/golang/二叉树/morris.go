// morris遍历  非
package main

import "fmt"

// 前序遍历
func (self *Tree) Preorder_Travel(node *TreeNode) {
	if node == nil {
		return
	}
	p := node
	for p != nil{
		if p.left != nil{
			curnode := p.left
			for curnode.right != nil && curnode.right != p{
				curnode = curnode.right
			}
			if curnode.right == nil{
				fmt.Printf("%d\t", p.Val)
				curnode.right = p
				p = p.left
				continue
			}else if curnode.right == p{
				curnode.right = nil
			}
		}else{
			fmt.Printf("%d\t", p.Val)
		}
		p = p.right
	}
}

// 中序遍历
func (self *Tree) Inorder_Travel(node *TreeNode) {
	if node == nil {
		return
	}
	p := node
	for p != nil {
		if p.left != nil {
			curnode := p.left
			for curnode.right != nil && curnode.right != p{
				curnode = curnode.right
			}
			if curnode.right == nil {
				curnode.right = p
				p = p.left
				//fmt.Printf("%d\t", p.Val)
				//continue
			}else if curnode.right == p{

				curnode.right = nil
			}
		} else {
			fmt.Printf("%d\t", p.Val)
		}
		p = p.right
	}
}

// 后序遍历
func (self *Tree) Postorder_Travel(node *TreeNode) {

}

func main() {
	t0 := &TreeNode{Val: 7}
	t10 := &TreeNode{Val: 17}
	t1 := &TreeNode{Val: 4}
	t2 := &TreeNode{Val: 1}
	t3 := &TreeNode{Val: 6}
	t0.left = t1
	t0.right = t10
	t1.left = t2
	t1.right = t3
	tree := &Tree{RootNode: t0}
	tree.Preorder_Travel(tree.RootNode)
	fmt.Println()
	//tree.Inorder_Travel(tree.RootNode)
	//fmt.Println()
}
