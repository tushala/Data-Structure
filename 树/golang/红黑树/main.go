package main

/*平衡二叉B树
1. 节点是 RED 或者 BLACK
2. 根节点是 BLACK
3. 叶子节点 （外部节点，空）都是 BLACK
4. RED 节点的子都是 BLACK
4.1 RED 节点的 parent 都是 BLACK
4.2 从根节点到 叶子节点 的所有 路径上不能2 个连续的 RED 节点
5. 从任一节点到 叶子节点 的 所有路径都包含相同数目BLACK
*/

type RBTree struct {
	RootNode *RBNode
}

func (self *RBTree) insert_node(node *RBNode) {
	if self.RootNode == nil {
		self.RootNode = node
		return
	}
	cur := self.RootNode
	for cur != nil {
		if cur.val < node.val {
			if cur.right == nil {
				node.parent = cur
				cur.right = node
				break
			}
			cur = cur.right
		} else {
			if cur.left == nil {
				node.parent = cur
				cur.left = node
				break
			}
			cur = cur.left
		}
	}
}
func (self *RBTree) check_node(node *RBNode) {

}
func (self *RBTree) Put(node *RBNode) {
	self.insert_node(node)
	self.check_node(node)
}
