package test.zzsure.solutions;

import com.zzsure.commons.classes.TreeNode;
import com.zzsure.solutions._0129;
import org.junit.Assert;
import org.junit.BeforeClass;
import org.junit.Test;

public class _0129Test {
    private static _0129.Solution solution;

    @BeforeClass
    public static void setup() {
        solution = new _0129.Solution();
    }

    @Test
    public void test1() {
        TreeNode root = new TreeNode(1);
        TreeNode left = new TreeNode(2);
        TreeNode right = new TreeNode(3);
        root.left = left;
        root.right = right;
        Assert.assertEquals(25, solution.sumNumbers(root));
    }

    @Test
    public void test2() {
        TreeNode root = new TreeNode(4);
        TreeNode node1 = new TreeNode(9);
        TreeNode node2 = new TreeNode(5);
        TreeNode node3 = new TreeNode(1);
        TreeNode node4 = new TreeNode(0);

        root.left = node1;
        node1.left = node2;
        node1.right = node3;
        root.right = node4;

        Assert.assertEquals(1026, solution.sumNumbers(root));
    }
}
