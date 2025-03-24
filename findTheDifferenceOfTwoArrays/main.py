class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        set1 = list_to_set(nums1)
        set2 = list_to_set(nums2)

        return [comp_sets(set1, set2), comp_sets(set2, set1)]

# convert List to set
def list_to_set(num: List[int]):
    s = set()
    for n in num:
        s.add(n)

    return s

# check if each value in set1 is in set2, return missing values
def comp_sets(set1, set2):
    l = []
    for val in set1:
        if val not in set2:
            l.append(val)
    return l