package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func index(num int, s []int) int {
	// 默认存在
	for i, v := range s{
		if  v == num{
			return i
		}
	}
	return 0
}

// 105 从前序与中序遍历序列构造二叉树
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(inorder) == 0 {
//		return nil
//	}
//	rootval := preorder[0]
//	root := &TreeNode{Val: rootval}
//	mid := index(rootval, inorder)
//	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
//	root.Right = buildTree(preorder[mid+1:],inorder[mid+1:])
//	return root
//
//}
// 106. 从中序与后序遍历序列构造二叉树
//func buildTree(inorder []int, postorder []int) *TreeNode {
//	if len(inorder) == 0 {
//		return nil
//	}
//	rootval := postorder[len(postorder)-1]
//	root := &TreeNode{Val: rootval}
//	mid := index(rootval, inorder)
//	root.Left = buildTree(inorder[:mid], postorder[:mid])
//	root.Right = buildTree(inorder[mid+1:], postorder[mid:len(postorder)-1])
//	return root
//}

// 889. 根据前序和后序遍历构造二叉树
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(post) == 0 {
		return nil
	}

	rootval := post[len(post)-1]
	root := &TreeNode{Val: rootval}
	if len(post) == 1{
		return root
	}
	rval := post[len(post)-2]
	mid := index(rval, pre)
	root.Left = constructFromPrePost(pre[1:mid], post[:mid-1])
	root.Right= constructFromPrePost(pre[mid:], post[mid-1:len(post)-1])
	return root
}

func main() {
	pre:= []int{1,2,4,5,3,6,7}
	post := []int{4,5,2,6,7,3,1}
	root := constructFromPrePost(pre, post)
	fmt.Println(root.Val)
	fmt.Println(root.Left.Val)
	fmt.Println(root.Right.Val)
	fmt.Println(root.Left.Left.Val)
	fmt.Println(root.Left.Right.Val)
	fmt.Println(root.Right.Val)

}