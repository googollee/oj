class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        l = len(nums1) + len(nums2)
        if l % 2 == 0:
            m1 = self.findKth(nums1, nums2, l/2)
            m2 = self.findKth(nums1, nums2, l/2+1)
            return (m1+m2)/2
        return self.findKth(nums1, nums2, l/2+1)

    def findKth(self, nums1, nums2, k):
        if len(nums1) == 0:
            return float(nums2[k-1])
        if len(nums2) == 0:
            return float(nums1[k-1])
        if k == 1:
            if nums1[0] < nums2[0]:
                return float(nums1[0])
            return float(nums2[0])
        al = k/2
        bl = al
        if len(nums1) < al:
            al = len(nums1)
        am = nums1[al-1]
        if len(nums2) < bl:
            bl = len(nums2)
        bm = nums2[bl-1]
        if am > bm:
            return self.findKth(nums1, nums2[bl:], k-bl)
        if am < bm:
            return self.findKth(nums1[al:], nums2, k-al)
        if k-al-bl>0:
            return self.findKth(nums1[al:], nums2[bl:], k-al-bl)
        return float(am)

s = Solution()
n1 = []
n2 = [2,3]
print s.findMedianSortedArrays(n1, n2)
