# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
    
        cur = head
        pre = None

        while cur is not None:
            # save next
            nex = cur.next
            #revers pointer from next to previous
            cur.next = pre

            # move pointers for next node
            pre = cur
            cur = nex
            
        #final head will be none, use last ptr saved as previous 
        return pre 

