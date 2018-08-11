# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None
        
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype
        """
        start = ListNode(-1)
        current_position = start
        pos = 0
        while l1 or l2 or pos:
            l1_value = l1.val if l1 else 0
            l2_value = l2.val if l2 else 0
            pos, new_value = divmod(l1_value + l2_value + pos, 10)
            current_position.next = ListNode(new_value)
            current_position = current_position.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        return start.next