class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        cost.append(0) # final step is free (top)
        for i in range(len(cost)-4,-1,-1): #total len - top and bottom range, start at end, move left
            cost[i] += min(cost[i+1], cost[i+2]) # replace each with cost from that point forward

        return min(cost[0],cost[1])

