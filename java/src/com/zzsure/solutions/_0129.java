package com.zzsure.solutions;

import com.zzsure.commons.classes.TreeNode;

/**
 * https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
 *
 * Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.
 * An example is the root-to-leaf path 1->2->3 which represents the number 123.
 * Find the total sum of all root-to-leaf numbers.
 * Note: A leaf is a node with no children.
 * Example:
 * Input: [1,2,3]
 *     1
 *    / \
 *   2   3
 * Output: 25
 * Explanation:
 * The root-to-leaf path 1->2 represents the number 12.
 * The root-to-leaf path 1->3 represents the number 13.
 * Therefore, sum = 12 + 13 = 25.
 * Example 2:
 * Input: [4,9,0,5,1]
 *     4
 *    / \
 *   9   0
 *  / \
 * 5   1
 * Output: 1026
 * Explanation:
 * The root-to-leaf path 4->9->5 represents the number 495.
 * The root-to-leaf path 4->9->1 represents the number 491.
 * The root-to-leaf path 4->0 represents the number 40.
 * Therefore, sum = 495 + 491 + 40 = 1026.
 *
 * @author zzsure
 * @date 2021/01/06
 */

public class _0129 {
    public static class Solution {
        public class RecurSum {
            public int preSum;
            public int sum;
        }

        public int sumNumbers(TreeNode root) {
            RecurSum recurSum = new RecurSum();
            preOrder(root, recurSum);
            return recurSum.sum;
        }

        public void preOrder(TreeNode root, RecurSum recurSum) {
            if (null == root) {
                return;
            }
            recurSum.preSum = recurSum.preSum * 10 + root.val;
            if (null == root.left && null == root.right) {
                recurSum.sum += recurSum.preSum;
                System.out.println("preSum: " + recurSum.preSum + ", sum: " + recurSum.sum);
                recurSum.preSum = (recurSum.preSum - root.val) / 10;
                return;
            }
            preOrder(root.left, recurSum);
            preOrder(root.right, recurSum);
            recurSum.preSum = (recurSum.preSum - root.val) / 10;
        }
    }
}
