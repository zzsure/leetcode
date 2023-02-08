package com.zzsure.solutions;

import com.zzsure.commons.classes.ListNode;

public class _0002 {
    public static class Solution {
        public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
            ListNode dump = new ListNode();
            ListNode c = dump;
            int carry = 0;
            while(null != l1 && null != l2) {
                int num = l1.val + l2.val + carry;
                carry = num / 10;
                num = num % 10;
                ListNode n = new ListNode(num);
                c.next = n;
                c = n;
                l1 = l1.next;
                l2 = l2.next;
            }

            ListNode last = (null == l1) ? l2 : l1;
            while (null != last) {
                int num = last.val + carry;
                carry = num / 10;
                num = num % 10;
                ListNode n = new ListNode(num);
                c.next = n;
                c = n;
                last = last.next;
            }

            if (0 != carry) {
                ListNode n = new ListNode(carry);
                c.next = n;
            }

            return dump.next;
        }
    }
}
