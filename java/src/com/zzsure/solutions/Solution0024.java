package com.zzsure.solutions;

import com.zzsure.commons.classes.ListNode;

import java.util.Arrays;

/**
 * Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
 * Input: head = [1,2,3,4]
 * Output: [2,1,4,3]
 * Example 2:
 *
 * Input: head = []
 * Output: []
 * Example 3:
 *
 * Input: head = [1]
 * Output: [1]
 *
 * Constraints:
 *
 * The number of nodes in the list is in the range [0, 100].
 * 0 <= Node.val <= 100
 *
 * @author zhangzongshuo
 */
public class Solution0024 {
    public static void main(String []args) {
        ListNode listNode = ListNode.createSinglyLinkedList(Arrays.asList(1, 2, 3, 4));
        listNode = new Solution().swapPairs(listNode);
        ListNode.printList(listNode);
    }

    public static class Solution {
        public ListNode swapPairs(ListNode head) {
            if (null == head) {
                return null;
            }
            ListNode next = head.next;
            if (null == next) {
                return head;
            }
            ListNode nextNext = next.next;
            next.next = head;
            head.next = swapPairs(nextNext);
            return next;
        }
    }
}
