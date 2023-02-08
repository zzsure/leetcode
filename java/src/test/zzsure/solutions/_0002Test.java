package test.zzsure.solutions;

import com.zzsure.commons.classes.ListNode;
import com.zzsure.solutions._0002;
import org.junit.Assert;
import org.junit.BeforeClass;
import org.junit.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class _0002Test {
    private static _0002.Solution solution;

    @BeforeClass
    public static void setup() {
        solution = new _0002.Solution();
    }

    @Test
    public void test1() {
        List<Integer> l1 = Arrays.asList(2, 4, 3);
        List<Integer> l2 = Arrays.asList(5, 6, 4);
        List<Integer> expectedList = Arrays.asList(7, 0, 8);
        ListNode ln1 = ListNode.createSinglyLinkedList(l1);
        ListNode ln2 = ListNode.createSinglyLinkedList(l2);
        ListNode expectedListNode = ListNode.createSinglyLinkedList(expectedList);
        ListNode resultListNode = solution.addTwoNumbers(ln1, ln2);
        Assert.assertEquals(expectedListNode, resultListNode);
    }

    @Test
    public void test2() {
        List<Integer> l1 = Arrays.asList(0);
        List<Integer> l2 = Arrays.asList(0);
        List<Integer> expectedList = Arrays.asList(0);
        ListNode ln1 = ListNode.createSinglyLinkedList(l1);
        ListNode ln2 = ListNode.createSinglyLinkedList(l2);
        ListNode expectedListNode = ListNode.createSinglyLinkedList(expectedList);
        ListNode resultListNode = solution.addTwoNumbers(ln1, ln2);
        Assert.assertEquals(expectedListNode, resultListNode);
    }

    @Test
    public void test3() {
        List<Integer> l1 = Arrays.asList(9, 9, 9, 9, 9, 9, 9);
        List<Integer> l2 = Arrays.asList(9, 9, 9, 9);
        List<Integer> expectedList = Arrays.asList(8, 9, 9, 9, 0, 0, 0, 1) ;
        ListNode ln1 = ListNode.createSinglyLinkedList(l1);
        ListNode ln2 = ListNode.createSinglyLinkedList(l2);
        ListNode expectedListNode = ListNode.createSinglyLinkedList(expectedList);
        ListNode resultListNode = solution.addTwoNumbers(ln1, ln2);
        Assert.assertEquals(expectedListNode, resultListNode);
    }
}
