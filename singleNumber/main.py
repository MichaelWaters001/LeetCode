class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        nums.sort()
        for i in range(len(nums) -1, 2): # range from 0 to end, step 2 at a time
            pre = nums[i]
            cur = nums[i+1]
            if pre != cur:
                return pre
            
        return nums[-1] #catch for largest number being the single entry 