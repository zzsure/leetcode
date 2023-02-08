package com.zzsure.solutions;

import java.util.ArrayList;
import java.util.List;

/**
 * The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
 *
 * countAndSay(1) = "1"
 * countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
 * To determine how you "say" a digit string, split it into the minimal number of substrings such that each substring contains exactly one unique digit. Then for each substring, say the number of digits, then say the digit. Finally, concatenate every said digit.
 *
 * For example, the saying and conversion for digit string "3322251":
 * Given a positive integer n, return the nth term of the count-and-say sequence.
 *
 *
 *
 * Example 1:
 *
 * Input: n = 1
 * Output: "1"
 * Explanation: This is the base case.
 * Example 2:
 *
 * Input: n = 4
 * Output: "1211"
 * Explanation:
 * countAndSay(1) = "1"
 * countAndSay(2) = say "1" = one 1 = "11"
 * countAndSay(3) = say "11" = two 1's = "21"
 * countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211
 *
 * 1 <= n <= 30
 *
 * @author zhangzongshuo
 */
public class Solution0038 {
    public static void main(String []args) {
        String result = new Solution().countAndSay(4);
        System.out.println(result);
    }

    public static class Solution {
        public String countAndSay(int n) {
            if (1 == n) {
                return "1";
            }
            String str = countAndSay(n-1);
            StringBuilder builder = new StringBuilder();
            char[] nums = str.toCharArray();
            int count = 1;
            for (int i = 0; i < nums.length; i++) {
                int nowNum = nums[i] - '0';
                if (i == nums.length-1) {
                    builder.append(count);
                    builder.append(nowNum);
                    break;
                }
                int aftNum = nums[i+1] - '0';
                if (nowNum != aftNum) {
                    builder.append(count);
                    builder.append(nowNum);
                    count = 1;
                } else {
                    count++;
                }
            }
            return builder.toString();
        }
    }
}
