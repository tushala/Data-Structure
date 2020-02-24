"""链表"""


class IndexOutRange(Exception):
    def __init__(self, info):
        self.info = info

    def __str__(self):
        return self.info


class ListNode(object):
    def __init__(self, val, idx, next=None):
        self.val = val
        self.idx = idx
        self.next = next


class LinkList(object):
    def __init__(self, node_list: list):
        self.linklist = self.makelist(node_list)

    def makelist(self, node_list: list):
        """构建链表"""
        if not node_list:
            return
        Node = ListNode(node_list[-1], 0)
        for i in reversed(range(len(node_list) - 1)):
            val = node_list[i]
            Node = ListNode(val, len(node_list) - 1 - i, Node)
        return Node

    def clear(self):
        self.linklist = None

    def add(self, idx, val):
        if idx > self.linklist.idx + 1:
            raise IndexOutRange("插入的index值{}大于链表长度+1".format(idx))
        curnode = self.linklist
        if idx == curnode.idx + 1:
            self.linklist = ListNode(val, idx, curnode)
        else:
            while curnode:
                if idx == curnode.idx:
                    Node = ListNode(val, idx, curnode.next)
                    curnode.next = Node
                    break
                curnode.idx += 1
                curnode = curnode.next

    def remove(self, idx):
        if idx > self.linklist.idx + 1:
            raise IndexOutRange("移除的index值{}大于链表长度+1".format(idx))
        curnode = self.linklist
        if idx == curnode.idx + 1:
            self.linklist = curnode.next
        else:
            while curnode:
                if idx == curnode.idx:
                    curnode.next = curnode.next.next
                curnode.idx -= 1
                curnode = curnode.next

    def printlist(self):
        curnode = self.linklist
        while curnode.next:
            print(curnode.val, "-->", end="")
            curnode = curnode.next
        print(curnode.val)


if __name__ == '__main__':
    l = LinkList(list(range(3)))
    l.printlist()
    l.add(3, 6)
    l.printlist()
    # l.remove(2)
    # l.printlist()

