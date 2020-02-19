"""双向链表"""
class IndexOutRange(Exception):
    def __init__(self, info):
        self.info = info
    def __str__(self):
        return self.info

class ListNode(object):
    def __init__(self, val):
        self.val = val
        self.prev = None  # 上节点
        self.next = None  # 下节点


class DoubleLinKList(object):
    def __init__(self):
        self._head :ListNode = None

    def is_empty(self):
        """判断链表是否为空"""
        return self._head == None

    def length(self):
        """返回链表的长度"""
        cur = self._head
        ans = 0
        while cur is not None:
            ans += 1
            cur = cur.next
        return ans

    def travel(self):
        """遍历链表"""
        print("链表包含元素有:  ", end="")
        cur = self._head
        while cur is not None:
            print(cur.val, end=" ")
            cur = cur.next
        print()

    def head_insert(self, val):
        """头插"""
        node = ListNode(val)
        if self.is_empty():
            self._head = node
        else:
           node.next = self._head
           self._head.prev = node
           self._head = node

    def tail_insert(self, val):
        """尾插"""
        node = ListNode(val)
        if self.is_empty():
            self._head = node
        else:
            cur = self._head
            while cur.next != None:
                cur = cur.next
            cur.next = node
            node.prev = cur

    def search(self, val):
        """查找"""
        cur = self._head
        while cur is not None:
            if cur.val == val:
                return True
            cur = cur.next
        return False

    def insert(self, pos, val):
        """在指定位置添加节点"""
        if pos <=0 or pos > self.length() + 1:
            raise IndexOutRange("插入节点位置有误")
        elif pos == 1:
            self.head_insert(val)
        elif pos == self.length() + 1:
            self.tail_insert(val)
        else:
            node = ListNode(val)
            cur = self._head
            count = 1
            while count < pos - 1:
                count + 1
                cur = cur.next
            node.prev = cur
            node.next = cur.next
            cur.next.prev = node
            cur.next = node

    def remove(self, val):
        """刪除"""
        if self.is_empty():
            return
        else:
            cur = self._head
            if cur.val == val:
                if cur.next == None:
                    self._head = None
                else:
                    cur.next.prev = None
                    self._head = cur.next
                return

            while cur is not None:
                if cur.val == val:
                    cur.prev.next = cur.next
                    cur.next.prev = cur.prev
                    break
                cur = cur.next