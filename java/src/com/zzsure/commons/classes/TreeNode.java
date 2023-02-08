package com.zzsure.commons.classes;

import java.util.List;
import java.util.Objects;

/**
 * 树节点
 * @author zhangzongshuo
 * @date 2021/01/06
 */
public class TreeNode {
    public int val;
    public TreeNode left;
    public TreeNode right;

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (!(o instanceof TreeNode)) {
            return false;
        }

        TreeNode treeNode = (TreeNode) o;

        if (val != treeNode.val) {
            return false;
        }

        if (left != null ? !left.equals(treeNode.left) : treeNode.left != null) {
            return false;
        }

        return right != null ? !right.equals(treeNode.right) : treeNode.right != null;
    }

    /*@Override
    public int hashCode() {
        int result = val;
        result = 31 * result + (null != left ? left.hashCode() : 0);
        result = 31 * result + (null != right ? right.hashCode() : 0);
        return result;
    }*/

    @Override
    public String toString() {
        return "TreeNode{" + val + ", left=" + left + ", right=" + right + '}';
    }

    public TreeNode(int x) {
        this.val = x;
    }

    public TreeNode(TreeNode left, int val, TreeNode right) {
        this.left = left;
        this.val = val;
        this.right = right;
    }
}