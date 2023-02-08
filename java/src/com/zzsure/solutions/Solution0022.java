package com.zzsure.solutions;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

/**
 * Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
 *
 *
 *
 * Example 1:
 *
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 * Example 2:
 *
 * Input: n = 1
 * Output: ["()"]
 *
 * Constraints:
 *
 * 1 <= n <= 8
 * @author zhangzongshuo
 */
public class Solution0022 {
    public static void main(String []args) {
        List<String> results = new Solution().generateParenthesis(3);
        System.out.println(results);
    }

    public static class Solution {
        public List<String> generateParenthesis(int n) {
            List<String> list = new ArrayList<>();
            if (0 == n) {
                list.add("");
            } else {
                for (int i = 0; i < n; i++) {
                    for (String left : generateParenthesis(i)) {
                        for (String right : generateParenthesis(n - i - 1)) {
                            list.add("(" + left + ")" + right);
                        }
                    }
                }
            }
            return list;
        }
    }
}
