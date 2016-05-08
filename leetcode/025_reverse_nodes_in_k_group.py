# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if k < 2:
            return head
        p = None
        c = head
        while True:
            print getList(head)
            begin = c
            end = c
            i = 0
            while end!=None and i<k:
                i += 1
                end = end.next
            if i != k:
                if p != None:
                    p.next = begin
                return head
            newHead = self.reverse(begin, end)
            if p == None:
                head = newHead
            else:
                p.next = newHead
            p = begin
            c = end

    def reverse(self, head, end):
        p = None
        c = head
        while c != end:
            n = c.next
            c.next = p
            p = c
            c = n
        return p

def mkList(a):
    if len(a) == 0:
        return None
    ret = ListNode(0)
    c = ret
    for i in a:
        c.next = ListNode(i)
        c = c.next
    return ret.next

def getList(head):
    ret = []
    while head != None:
        ret += [head.val]
        head = head.next
    return ret

l = mkList([1,2,3,4,5])
s = Solution()
h = s.reverseKGroup(l, 6)
a = getList(h)
print a
