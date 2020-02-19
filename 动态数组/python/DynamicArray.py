"""动态数组"""


class Dynamicarray(object):
    def __init__(self, capacity=10):
        self._capacity = capacity
        self._n = 0
        self._data = [None] * self._capacity

    def __getitem__(self, item):
        return self._data[item]

    def __len__(self):
        return self._n

    def isEmpty(self):
        # 判断是否为空
        return self._n == 0

    def add_last(self, val):
        # 添加元素到最后面
        if self._data[-1] is not None:
            self._data += [None] * self._capacity
            self._data[-1] = val
        else:
            idx = self._data.index(None)
            self._data = self._data[:idx] + self._data[idx + 1:] + [val]
        self._n += 1

    def add(self, index, val):
        # 添加元素到index位值为val
        if index < 0:  # 插入的位置无效
            raise Exception('插入的位置无效')
        elif index > len(self._data):
            self._data += [None] * (index + 1 - len(self._data))
        self._data[index] = val
        self._n += 1

    def contains(self, val):
        # 是否包含元素val
        for i in range(self._n):
            if self._data[i] == val:
                return True
        return False

    def get(self, index):
        # 获取index位的值
        if index < 0 or index >= self._n:
            raise Exception('Get failed. Index is illegal.')
        return self._data[index]

    def set(self, index, val):
        # 设置index位的值为val
        if index < 0 or index >= self._n:
            raise Exception('Get failed. Index is illegal.')
        self._data[index] = val
        if index > self._n:
            self._n += 1

    def remove(self, index):
        if index < 0 or index >= self._n:
            raise Exception('Get failed. Index is illegal.')

        self._data[index] = None
        if index < self._n:
            self._n -= 1

    def clear(self):
        # 清空
        self._data = [None] * self._capacity
        self._n = 0
