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
	var reverseAndPrint func(node *TreeNode)
	reverseAndPrint = func(node *TreeNode) {
		if node == nil {
			return
		}

		cur := node

		var b *TreeNode
		i := 0
		for cur != nil {
			p := cur.right
			if i == 0 {
				cur.right = nil
				//if node.Val == 4{
				//	fmt.Println(11, node.Val,node.right.Val)
				//	panic("1")
				//}
			} else {
				cur.right = b
			}

			b = cur
			cur = p
			i++
		}
		node = b
		//fmt.Println(123, node)
		//fmt.Println(456, node.right)
		for node != nil {
			fmt.Printf("%d\t", node.Val)
			node = node.right
		}
	}

	if node == nil {
		return
	}
	p := node
	for p != nil {
		if p.left != nil {
			curnode := p.left
			for curnode.right != nil && curnode.right != p {
				curnode = curnode.right
			}
			if curnode.right == nil {
				curnode.right = p
				p = p.left
				continue
			} else if curnode.right == p {
				curnode.right = nil
				reverseAndPrint(p.left)
			}
		}
		p = p.right
	}
	reverseAndPrint(node)
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
	tree.Postorder_Travel(tree.RootNode)
	fmt.Println()
}
