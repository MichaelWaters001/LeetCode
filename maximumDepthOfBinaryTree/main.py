# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
            def depth(dep, root):
                if root is None:
                    return dep
                l = depth(dep+1, root.left)
                r = depth(dep+1, root.right)
                return max(l,r)
            return depth(0, root)      
            
            
