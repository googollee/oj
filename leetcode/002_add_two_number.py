class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        ret = ListNode(0)
        result = ret
        addition = 0
        while l1 != None and l2 != None:
            v = l1.val+l2.val+addition
            result.next = ListNode(v % 10)
            addition = v / 10
            result = result.next
            l1 = l1.next
            l2 = l2.next
        rest = l1
        if l2 != None:
            rest = l2
        while rest != None:
            v = rest.val+addition
            result.next = ListNode(v % 10)
            addition = v / 10
            result = result.next
            rest = rest.next
        if addition == 1:
            result.next = ListNode(1)
        return ret.next

l1 = ListNode(9)
l1.next = ListNode(4)
l1.next.next = ListNode(9)
l2 = ListNode(9)
l2.next = ListNode(6)
l2.next.next = ListNode(4)
s = Solution()
ret = s.addTwoNumbers(l1, l2)
print ret.val, ret.next.val, ret.next.next.val, ret.next.next.next.val
