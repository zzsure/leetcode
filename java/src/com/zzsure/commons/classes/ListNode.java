package com.zzsure.commons.classes;

import java.util.Objects;
import java.util.List;

/**
 * 链表结构
 * @author zhangzongshuo
 * @date 2020/10/14
 */
public class ListNode {
    public int val;
    public ListNode next;
    public ListNode() {};
    public ListNode(int val) {
        this.val = val;
    }
    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (!(o instanceof ListNode)) {
            return false;
        }

        ListNode node = (ListNode)o;
        if (this.val != node.val) {
            return false;
        }

        return Objects.equals(this.next, node.next);
    }

    @Override
    public String toString() {
        return "ListNode{"+"val="+val+", next="+next+'}';
    }

    public static ListNode createSinglyLinkedList(List<Integer> numList) {
        if (null == numList || 0 == numList.size()) {
            throw new IllegalArgumentException("please input valid array list");
        }
        ListNode head = new ListNode(numList.get(0));
        ListNode tmp = head;
        for (int i = 1; i < numList.size(); i++) {
            ListNode next = new ListNode(numList.get(i));
            tmp.next = next;
            tmp = next;
        }
        printList(head);
        return head;
    }

    public static void printList(ListNode head) {
        ListNode tmp = head;
        System.out.println();
        while (null != tmp) {
            System.out.print(tmp.val + "\t");
            tmp = tmp.next;
        }
    }
}