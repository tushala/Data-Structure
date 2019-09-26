class Node(object):
    def __init__(self, key):
        self.key = key
        self.left: Node = None
        self.right: Node = None
        self.height = 0


class AVLTree(object):
    
    def __init__(self):
        self.root = None
    
    def height(self, node: Node):
        if node is None:
            return -1
        else:
            return node.height
    
    def find(self, key: int):
        if self.root is None:
            return None
        else:
            self._find(key, self.root)
    
    def _find(self, key: int, node: Node):
        if node is None:
            return None
        elif key < node.key:
            return self._find(key, node.left)
        elif key > node.key:
            return self._find(key, node.right)
        else:
            return node
    
    def findMin(self):
        if self.root is None:
            return None
        else:
            return self._findMin(self.root)
    
    def _findMin(self, node):
        if node.left:
            return self._findMin(node.left)
        else:
            return node
    
    def findMax(self):
        if self.root is None:
            return None
        else:
            return self._findMax(self.root)
    
    def _findMax(self, node):
        if node.right:
            return self._findMax(node.right)
        else:
            return node
    
    def singleLeftRotate(self, node: Node):
        # 左左 右旋
        #              node                                 x
        #            /    \                              /     \
        #          x       T4                         z         node
        #        /   \            ----------->      /   \      /     \
        #      z      T3                          T1     T2  T3       T4
        #    /   \
        #  T1     T2
        x = node.left
        node.left = x.right  # T3
        x.right = node
        node.height = max(self.height(node.left), self.height(node.right)) + 1
        x.height = max(self.height(x.left), node.height) + 1
        return x
    
    def singleRightRotate(self, node: Node):
        # 右右 左旋
        #             node                                  x
        #            /    \                              /     \
        #          T1      x                          node       z
        #                /  \     ----------->       /   \      / \
        #              T2    z                     T1    T2    T3  T4
        #                  /   \
        #                T3    T4
        x = node.right
        node.right = x.left  # T2
        x.left = node
        node.height = max(self.height(node.left), self.height(node.right)) + 1
        x.height = max(self.height(x.left), node.height) + 1
        return x
    
    def doubleLeftRotate(self, node: Node):
        # 右左 先右旋后左旋
        #       8                     8                     8
        #     /   \                 /   \                 /   \
        #   7      10   ---->     7      10     ---->   7     11
        #           \                      \                  /  \
        #            13                     11               10   13
        #           /                         \
        #         11                          13
        node.left = self.singleRightRotate(node.left)
        return self.singleLeftRotate(node)
    
    def doubleRightRotate(self, node: Node):
        # 左右 先左旋后右旋
        #       8                     8                     8
        #     /   \                 /   \                 /   \
        #   7      10   ---->     7      10     ---->   6      10
        #  /                     /                     / \
        # 5                    6                     5    7
        #  \                 /
        #   6               5
        node.right = self.singleLeftRotate(node.right)
        return self.singleRightRotate(node)
    
    def put(self, key: int):
        if not self.root:
            self.root = Node(key)
        else:
            self.root = self._put(key, self.root)
    
    def show(self):
        return self._show(self.root)
    
    def _show(self, root: Node):
        if root is not None:
            print(root.key, end='\t')
            self._show(root.left)
            self._show(root.right)
        else:
            return
    
    def _put(self, key, node: Node):
        if node is None:
            node = Node(key)
        elif key < node.key:
            node.left = self._put(key, node.left)
            if self.height(node.left) - self.height(node.right) == 2:
                if key < node.left.key:
                    node = self.singleLeftRotate(node)
                else:
                    node = self.doubleLeftRotate(node)
        
        elif key > node.key:
            node.right = self._put(key, node.right)
            if self.height(node.right) - self.height(node.left) == 2:
                if key < node.right.key:
                    node = self.doubleRightRotate(node)
                else:
                    node = self.singleRightRotate(node)
        
        node.height = max(self.height(node.right), self.height(node.left)) + 1
        return node
    
    def delete(self, key):
        self.root = self.remove(key, self.root)

    def remove(self, key, node:Node):
        if node is None:
            raise Exception('Error,key not in tree')
        elif key < node.key:
            node.left = self.remove(key, node.left)
            if self.height(node.right) - self.height(node.left) == 2:
                if self.height(node.right.right) >= self.height(node.right.left):
                    node = self.singleRightRotate(node)
                else:
                    node = self.doubleRightRotate(node)
            node.height = max(self.height(node.left), self.height(node.right)) + 1

        elif key > node.key:
            node.right = self.remove(key, node.right)
            if (self.height(node.left) - self.height(node.right)) == 2:
                if self.height(node.left.left) >= self.height(node.left.right):
                    node = self.singleLeftRotate(node)
                else:
                    node = self.doubleLeftRotate(node)
            node.height = max(self.height(node.left), self.height(node.right)) + 1

        elif node.left and node.right:
            if node.left.height <= node.right.height:
                minNode = self._findMin(node.right)
                node.key = minNode.key
                node.right = self.remove(node.key,node.right)
            else:
                maxNode = self._findMax(node.left)
                node.key = maxNode.key
                node.left = self.remove(node.key, node.left)
            node.height = max(self.height(node.left), self.height(node.right)) + 1

        else:
            if node.right:
                node = node.right
            else:
                node = node.left
        return node


if __name__ == '__main__':
    avl = AVLTree()
    L = [5, 1, 2, 3, 4]
    # L = [5, 3]
    for i in L:
        avl.put(i)
    # print(avl.root.key)
    print('最大值: ', avl.findMax().key)
    print('最小值: ', avl.findMin().key)
    avl.show()
    avl.delete(2)
    print()
    avl.show()
