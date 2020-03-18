package main

/*平衡二叉B树

// 没有平衡因子概念  所以节点不设置 height 概念
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
	/*性质二：根节点是黑色；
	性质四：每个红色节点的两个子节点都是黑色的（也就是说不存在两个连续的红色节点）；
	节点和其父点必定不能够同时为红色节点
	*/
	//如果是父节点直接设置成黑色节点，退出
	if self.RootNode == node || self.RootNode == node.parent {
		self.RootNode.set_Black_Node()
		return
	}
	//如果父节点是黑色节点，直接退出
	if node.parent.isBlack() {
		return
	}
	//1.如果父节点的兄弟节点也是红色节点
	//node node.parent node.parent.parent 不同边
	grand := node.parent.parent
	if grand == nil {
		self.check_node(node.parent)
		return
	}
	if grand.left != nil && grand.left.isRed() && grand.right != nil && grand.right.isRed(){
		grand.left.set_Black_Node()
		grand.right.set_Black_Node()
		grand.set_Red_Node()
		self.check_node(grand)
		return
	}

	//2.如果父节点的兄弟节点也是黑色节点
	parent := node.parent
	if parent.left == node && grand.right == node.parent{
		self.right_rotate(node.parent)
		self.check_node(parent)
		return
	}
	if parent.right == node && grand.left == node.parent{
		self.left_rotate(node.parent)
		self.check_node(parent)
		return
	}

	//node node.parent node.parent.parent 不同边\

	//parent.set_Black_Node()
	//grand.set_Red_Node()
	if parent.left == node && grand.left == node.parent{
		self.right_rotate(grand)
		return
	}
	if parent.right == node && grand.right == node.parent{
		self.left_rotate(grand)
		return
	}
}
//func (self *RBTree) check_node(node *RBNode) {
//	if self.RootNode == node || self.RootNode == node.parent {
//		self.RootNode.set_Black_Node()
//		return
//	}
//	parent := node.parent
//	grand  := parent.parent
//	if parent.isBlack() {
//		return
//	} else {
//		//父亲节点为红色则要处理
//		uncle := node.getUncle()
//		if uncle != nil && uncle.isRed() {
//			// 上溢
//			uncle.set_Black_Node()
//			parent.set_Black_Node()
//			grand.set_Red_Node()
//			self.check_node(grand)
//		}else{
//			if grand.left == parent{
//				if parent.right == node{
//					// LR
//					self.left_rotate(parent)
//					self.right_rotate(grand)
//				}else{
//					// LL
//					self.right_rotate(parent)
//				}
//			}else{
//				if parent.right == node{
//					// RR
//					self.left_rotate(parent)
//				}else{
//					// RL
//					self.right_rotate(parent)
//					self.left_rotate(grand)
//				}
//			}
//			grand.set_Red_Node()
//			parent.set_Black_Node()
//			self.check_node(parent)
//		}
//	}
//}
func (self *RBTree) left_rotate(node *RBNode) {
	//* 左旋示意图：对节点x进行左旋
	//*     parent               parent
	//*    /                       /
	//*  node                   right
	//*  / \                     / \
	//* ln  right   ----->     node  ry
	//*    / \                 / \
	//*   ly ry               ln ly
	//
	//* 1. 将right的左子节点ly赋给node的右子节点,并将node赋给right左子节点ly的父节点(ly非空时)
	//* 2. 将right的左子节点设为node，将node的父节点设为right
	//* 3. 将node的父节点parent(非空时)赋给right的父节点，同时更新parent的子节点为right(左或右)
	parent := node.parent
	right := node.right
	// 1
	if node.left != nil {
		node.left.parent = node
	}
	// 2
	right.left = node
	node.parent = right

	// 3
	right.parent = parent
	if parent == nil {
		self.RootNode = right
	} else {
		if parent.left == node {
			parent.left = right
		} else {
			parent.right = right
		}
	}
}
func (self *RBTree) right_rotate(node *RBNode) {
	//* 左旋示意图：对节点y进行右旋
	//*        parent           parent
	//*       /                   /
	//	*      node                left
	//*     /    \               / \
	//*    left  ry   ----->   ln  node
	//*   / \                     / \
	//* ln  rn                   rn ry
	//1. 将left的右子节点rn赋给node的左子节点,并将node赋给rn右子节点的父节点(left右子节点非空时)
	//2. 将left的右子节点设为node，将node的父节点设为left
	//3. 将node的父节点parent(非空时)赋给left的父节点，同时更新parent的子节点为left(左或右)

	parent := node.parent
	left := node.left

	// 1
	node.left = left.right
	if node.left != nil {
		node.left.parent = node
	}
	// 2
	left.right = node
	node.parent = left

	// 3
	left.parent = parent
	if parent == nil {
		self.RootNode = left
	} else {
		if parent.left == node {
			parent.left = left
		} else {
			parent.right = left
		}
	}
}
func (self *RBTree) Put(node *RBNode) {
	self.insert_node(node)
	self.check_node(node)
}
