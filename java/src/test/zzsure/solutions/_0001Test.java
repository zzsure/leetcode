package test.zzsure.solutions;

import com.zzsure.solutions._0001;
import org.junit.BeforeClass;
import org.junit.Test;

import static org.junit.Assert.assertArrayEquals;

public class _0001Test {
    private static _0001.Solution solution;

    @BeforeClass
    public static void setup() {
        solution = new _0001.Solution();
    }

    @Test
    public void test() {
        assertArrayEquals(new int[]{0, 1}, solution.twoSum(new int[]{2, 7, 11, 15}, 9));
    }
}
