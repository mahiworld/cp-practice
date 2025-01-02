class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
def maxDepth(root: TreeNode) -> int:
    if not root:
        return 0
    elif not root.left and not root.right:
        return 1
    else:
        return 1 + max(maxDepth(root.left), maxDepth(root.right))
    

# Example: 
#       3
#     /   \
#    9     20
#         /  \
#        15   7    
root = TreeNode(3)
root.left = TreeNode(9)
root.right = TreeNode(20)
root.right.left = TreeNode(15)
root.right.right = TreeNode(7)
root.right.right.right = TreeNode(10)

# Call the maxDepth function and print the result
print(maxDepth(root))