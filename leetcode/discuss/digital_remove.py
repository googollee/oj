# https://leetcode.com/discuss/88166/given-integer-find-digit-equal-either-adjacent-digits-remove
def digital_remove(n):
    if n <= 0:
        return 0
    last = -1
    left = n
    right = 0
    base = 1
    ret = 0
    while left > 0:
        p = left % 10
        left = left / 10
        if p == last:
            n = left*base + right
            if n > ret:
                ret = n
        last = p
        right = base*p + right
        base = base * 10
    return ret

print digital_remove(-2)
