class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None
        self.parent = None  # Node


class BinarySearchTree(object):
    def add(self, root, val, parent=None):
        if root is None:
            root = Node(val)
            root.parent = parent
        elif val <= root.val:
            parent = root
            root.left = self.add(root.left, val, parent)
        elif val > root.val:
            parent = root
            root.right = self.add(root.right, val, parent)
        return root

    def query(self, root, val):
        if root is None:
            print(val, '不存在')
            return False
        if root.val == val:
            print(val, '存在')
            return root
        elif val < root.val:
            return self.query(root.left, val)
        elif val > root.val:
            return self.query(root.right, val)

    def findMin(self, root):
        if root.left:
            return self.findMin(root.left)
        else:
            return root

    def findMax(self, root):
        if root.right:
            return self.findMax(root.right)
        else:
            return root

    def delNode(self, root, val):
        if root is None:
            return
        if val < root.val:
            root.left = self.delNode(root.left, val)
        elif val > root.val:
            root.right = self.delNode(root.right, val)
        # 当val == root.val时，分为三种情况：只有左子树或者只有右子树、有左右子树、即无左子树又无右子树
        else:
            if root.left and root.right:
                # 既有左子树又有右子树，则需找到右子树中最小值节点
                temp = self.findMin(root.right)
                root.val = temp.val
                root.right = self.delNode(root.right, temp.val)
            elif root.right is None and root.left is None:
                # 左右子树都为空
                root = None
            elif root.right is None:
                # 只有左子树
                root = root.left
            elif root.left is None:
                root = root.right
        return root

    def proNode(self, node: Node):
        if node.left is not None:
            node = node.left
            while node.right:
                node = node.right
            return node.val
        else:
            if node.parent is not None:
                cur_val = node.val
                while node.parent:
                    node = node.parent
                    if node.val <= cur_val:
                        return node.val
            else:
                print('没有前驱节点')
                return None

    def backNode(self, node: Node):
        if node.right is not None:
            node = node.right
            while node.left:
                node = node.left
            return node.val
        else:
            if node.parent is not None:
                cur_val = node.val
                while node.parent:
                    node = node.parent
                    if node.val >= cur_val:
                        return node.val
            else:
                print('没有后继节点')
                return None

    def get_pro(self, root):
        # 前序遍历
        if not root:
            return []
        return [root.val] + self.get_pro(root.left) + self.get_pro(root.right)

    def get_midd(self, root):

        if not root:
            return []
        return self.get_midd(root.left) + [root.val] + self.get_midd(root.right)

    def get_back(self, root):
        # 后序遍历
        if not root:
            return []
        return self.get_back(root.left) + self.get_back(root.right) + [root.val]

    def printTree(self, root):
        if root is None:
            return
        self.printTree(root.left)
        print(root.val, end=' ')
        self.printTree(root.right)


if __name__ == '__main__':
    List = [17, 5, 34, 2, 11, 35, 29, 38, 9, 16, 8]
    # List = [5, 20, 33, 34]
    root = None
    bt = BinarySearchTree()
    for val in List:
        root = bt.add(root, val)
    print('中序打印二叉搜索树：', end=' ')
    bt.printTree(root)
    print('')
    print('根节点的值为：', root.val)
    print('树中最大值为:', bt.findMax(root).val)
    print('树中最小值为:', bt.findMin(root).val)
    print('查询树中值为5的节点:', bt.query(root, 5))
    print('查询树中值为100的节点:', bt.query(root, 100))
    print('删除树中值为16的节点:', end=' ')
    root = bt.delNode(root, 16)
    bt.printTree(root)
    print('')
    print('删除树中值为5的节点:', end=' ')
    root = bt.delNode(root, 5)
    bt.printTree(root)
    print('')

    Node_9 = bt.query(root, 9)
    if Node_9:
        print(9, '的前驱节点是: ', bt.proNode(Node_9))

    Node_35 = bt.query(root, 35)
    if Node_35:
        print(35, '的前驱节点是: ', bt.backNode(Node_35))
    print('前序遍历：', bt.get_pro(root))
    print('中序遍历：', bt.get_midd(root))
    print('后序遍历：', bt.get_back(root))

