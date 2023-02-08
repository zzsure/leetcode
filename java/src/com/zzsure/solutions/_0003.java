package com.zzsure.solutions;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Set;

public class _0003 {
    public static class Solution {
        public int lengthOfLongestSubstring(String s) {
            Set<Character> set = new HashSet<>();
            int rk = 0, maxLen = 0, n = s.length();
            for (int i = 0; i < n; i++) {
                if (i != 0) {
                    set.remove(s.charAt(i-1));
                }
                while(rk < n && !set.contains(s.charAt(rk))) {
                    set.add(s.charAt(rk));
                    ++rk;
                }
                maxLen = Math.max(maxLen, rk-i);
            }
            return maxLen;
        }
    }
}
