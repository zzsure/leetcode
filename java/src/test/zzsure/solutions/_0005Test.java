package test.zzsure.solutions;

import com.zzsure.solutions._0005;
import org.junit.Assert;
import org.junit.BeforeClass;
import org.junit.Test;

public class _0005Test {
    private static _0005.Solution solution;

    @BeforeClass
    public static void setup() {
        solution = new _0005.Solution();
    }

    @Test
    public void test1() {
        String s = "babad";
        Assert.assertEquals("bab", solution.longestPalindrome(s));
    }

    @Test
    public void test2() {
        String s = "cbbd";
        Assert.assertEquals("bb", solution.longestPalindrome(s));
    }

    @Test
    public void test3() {
        String s = "a";
        Assert.assertEquals("a", solution.longestPalindrome(s));
    }

    @Test
    public void test4() {
        String s = "ac";
        Assert.assertEquals("a", solution.longestPalindrome(s));
    }

    @Test
    public void test5() {
        String s = "ccc";
        Assert.assertEquals("ccc", solution.longestPalindrome(s));
    }
}
