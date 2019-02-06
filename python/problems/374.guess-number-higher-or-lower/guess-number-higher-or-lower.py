# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num):
import random

n, pick = 10,6

print(pick)

def guess(g):
    if g > pick:
        return 1
    if g < pick:
        return -1
    return 0


class Solution(object):
    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        l = 0
        h = n
        res = 0
        while l < h:
            mid = l + (h - l) // 2
            res = guess(mid)
            if res == 0:
                break
            elif res == -1:
                l = mid + 1
            else:
                h = mid - 1
        if guess(l) == 0:
            return l
        return mid


s = Solution()

l = s.guessNumber(n)
print(l)
