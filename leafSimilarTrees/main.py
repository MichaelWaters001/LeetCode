from typing import Optional
#Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Helper function to create a binary tree from a list
def create_tree(nodes, index=0):
    if index < len(nodes) and nodes[index] is not None:
        node = TreeNode(nodes[index])
        node.left = create_tree(nodes, 2 * index + 1)
        node.right = create_tree(nodes, 2 * index + 2)
        return node
    return None

# t1 = create_tree([1])
# t2 = create_tree([1])
t1 = create_tree([1, 2, 3])
t2 = create_tree([1, 3, 2])

class Solution:
    def leafSimilar(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
       
        def get_leafs(root: Optional[TreeNode], leafs):
            if root is None:
                return
            if root.left == None and root.right == None: # if leaf
                leafs.append(root.val)
                return
            
            get_leafs(root.left, leafs)
            get_leafs(root.right, leafs)

        leafs1 = []
        leafs2 = []

        get_leafs(root1, leafs1)
        get_leafs(root2, leafs2)

        if len(leafs1) != len(leafs2):
            return False
        
        for i in range(len(leafs1)):
            if leafs1[i] != leafs2[i]:
                return False
            
        return True
    
sol = Solution()
print(sol.leafSimilar(t1, t2))