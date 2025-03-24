class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        m = defaultdict(int) #set default type for the map (dictionary)
        for num in arr:
            m[num]+= 1 # track counts of all ints

        s = set()
        for val in m.values():
            if val in s: # check if count seen before
                return False
            s.add(val) # track counts seen (avoids sort)
        
        return True


        