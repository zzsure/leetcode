package com.zzsure.solutions;
/**
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
 *
 * A region is captured by flipping all 'O's into 'X's in that surrounded region.
 *
 * Example:
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 * After running your function, the board should be:
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 * Explanation:
 *
 * Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
 *
 */
public class _0130 {
    public static void main(String []args) {
        char[][] board = new char[][]{{'O', 'X'},{'X', 'O'}};
        new Solution().solve(board);
    }

    public static class Solution {
        public void solve(char[][] board) {
            if (null == board || 0 == board.length || 0 == board[0].length) {
                return;
            }
            // 是否连通
            Boolean[][] connArr = new Boolean[board.length][board[0].length];
            for (int i = 0; i < board.length; i++) {
                for (int j = 0; j < board[0].length; j++) {
                    connArr[i][j] = recuConn(board, connArr, i, j);
                }
            }
            for (int i = 0; i < board.length; i++) {
                for (int j = 0; j < board[0].length; j++) {
                    if ('O' == board[i][j] && !connArr[i][j]) {
                        board[i][j] = 'X';
                    }
                }
            }
        }

        private Boolean recuConn(char [][]board, Boolean[][] connArr, int i, int j) {
            if (i < 0 || j < 0 || i == board.length || j == board[0].length) {
                return Boolean.FALSE;
            }
            if ('O' != board[i][j]) {
                return Boolean.FALSE;
            }
            if (0 == i || 0 == j || board.length-1 == i || board[0].length-1 == j) {
                return Boolean.TRUE;
            }
            // 不为null代表遍历过了
            if (null != connArr[i][j]) {
                return connArr[i][j];
            }
            Boolean left = connArr[i][j-1];
            Boolean right = connArr[i][j+1];
            Boolean top = connArr[i-1][j];
            Boolean bottom = connArr[i+1][j];
            if (Boolean.TRUE.equals(left) || Boolean.TRUE.equals(right) || Boolean.TRUE.equals(top) || Boolean.TRUE.equals(bottom)) {
                return Boolean.TRUE;
            }
            return Boolean.FALSE;
        }
    }
}
