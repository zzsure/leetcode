package test.zzsure.solutions;

import com.zzsure.solutions._0003;
import org.junit.Assert;
import org.junit.BeforeClass;
import org.junit.Test;

public class _0003Test {
    private static _0003.Solution solution;

    @BeforeClass
    public static void setup() {
        solution = new _0003.Solution();
    }

    @Test
    public void test1() {
        String s = "abcabcbb";
        Assert.assertEquals(3, solution.lengthOfLongestSubstring(s));
    }

    @Test
    public void test2() {
        String s = "bbbbb";
        Assert.assertEquals(1, solution.lengthOfLongestSubstring(s));
    }

    @Test
    public void test3() {
        String s = "pwwkew";
        Assert.assertEquals(3, solution.lengthOfLongestSubstring(s));
    }

    @Test
    public void test4() {
        String s = "";
        Assert.assertEquals(0, solution.lengthOfLongestSubstring(s));
    }
}
