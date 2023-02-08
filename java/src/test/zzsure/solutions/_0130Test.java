package test.zzsure.solutions;

import com.zzsure.solutions._0130;
import org.junit.Assert;
import org.junit.BeforeClass;
import org.junit.Test;

import java.util.Arrays;

public class _0130Test {
    private static _0130.Solution solution;

    @BeforeClass
    public static void setup() {
        solution = new _0130.Solution();
    }

    @Test
    public void test1() {
        char[][] board = new char[][]{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}};
        char[][] expectBoard = new char[][]{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}};
        solution.solve(board);
        Assert.assertArrayEquals(expectBoard, board);
    }

    @Test
    public void test2() {
        // [["O","X","X","O","X"],["X","O","O","X","O"],["X","O","X","O","X"],["O","X","O","O","O"],["X","X","O","X","O"]]
        // [["O","X","X","O","X"],["X","X","X","X","O"],["X","X","X","X","X"],["O","X","X","X","O"],["X","X","O","X","O"]]
        // [["O","X","X","O","X"],["X","X","X","X","O"],["X","X","X","O","X"],["O","X","O","O","O"],["X","X","O","X","O"]]
        char[][] board = new char[][]{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}};
        char[][] expectBoard = new char[][]{{'O', 'X', 'X', 'O', 'X'}, {'X', 'X', 'X', 'X', 'O'}, {'X', 'X', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}};
        solution.solve(board);
        System.out.println(Arrays.deepToString(board));
        Assert.assertArrayEquals(expectBoard, board);
    }
}