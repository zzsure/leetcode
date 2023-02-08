package com.zzsure.solutions;

/**
 *
 * 5. Longest Palindromic Substring
 * Given a string s, return the longest palindromic substring in s.
 * 1 <= s.length <= 1000
 * s consist of only digits and English letters (lower-case and/or upper-case),
 *
 * @author zzsure
 * @date 2020/10/29
 *
 */

public class _0005 {
    public static class Solution {
        public String longestPalindrome(String s) {
            if (0 == s.length() || 1 == s.length()) {
                return s;
            }
            if (2 == s.length()) {
                if (s.charAt(0) == s.charAt(1)) {
                    return s;
                }
                return s.substring(0, 1);
            }
            int maxLen = 0;
            String maxStr = "";
            for (int i = 0; i < s.length(); i++) {
                int j = 0;
                for (j = 0; (i-j) >= 0 && (i+j) < s.length(); j++) {
                    if (s.charAt(i-j) != s.charAt(i+j)) {
                        if (2*j-1 > maxLen) {
                            maxLen = 2*j-1;
                            maxStr = s.substring(i-j+1, i+j);
                        }
                        break ;
                    }
                }
                if ((i-j) < 0 || (i+j) >= s.length()) {
                    if (2*j-1 > maxLen) {
                        maxLen = 2*j-1;
                        maxStr = s.substring(i-j+1, i+j);
                    }
                }
                j = 0;
                for (j = 0; (i-j) >=0 && (i+j+1) < s.length(); j++) {
                    if (s.charAt(i-j) != s.charAt(i+j+1)) {
                        if (2*j > maxLen) {
                            maxLen = 2*j;
                            maxStr = s.substring(i-j+1, i+j+1);
                        }
                        break ;
                    }
                }
                if ((i-j) < 0 || (i+j+1) >= s.length()) {
                    if (2*j > maxLen) {
                        maxLen = 2*j;
                        maxStr = s.substring(i-j+1, i+j+1);
                    }
                }
            }
            return maxStr;
    }
    }
}
