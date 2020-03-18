package main

const (
	RED   string = "R"
	BLACK string = "B"
)

type RBNode struct {
	val                 int
	left, right, parent *RBNode
	height              int
	color               string
}

func NewRBNode(val int) *RBNode {
	rbNode := &RBNode{
		color:  RED,
		val:    val,
		parent: nil,
		left:   nil,
		right:  nil,
	}
	return rbNode
}
func (self *RBNode) isBlack() bool {
	return self.color == BLACK
}

func (self *RBNode) isRed() bool {
	return self.color == RED
}
func (self *RBNode) set_Black_Node() {
	self.color = BLACK
}
func (self *RBNode) set_Red_Node() {
	self.color = RED
}
func (self *RBNode) getUncle() *RBNode {
	parent := self.parent
	if parent != nil {
		if self == parent.left {
			return parent.right
		} else {
			return parent.left
		}
	} else {
		return nil
	}
}
