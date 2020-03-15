# leetcode
leetcode link and hint

### 3. longest-substring-without-repeating-characters
- link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
- hint: if _, ok := maxMap[s[i]]; ok { if i-sidx > maxLen { maxLen = i - sidx } }

### 5. longest-palindromic-substring
- link: https://leetcode.com/problems/longest-palindromic-substring/
- hint: if s[i] == s[i-1] {...}  if i+1 < len(s) && s[i+1] == s[i-1] {...}

### 6. zigzag-conversion
- link: https://leetcode-cn.com/problems/zigzag-conversion/
- hint: if i == 0 {...} else if i == numRows-1 {...}  } else {...}

#### 7.reverse-integer
- link: https://leetcode.com/problems/reverse-integer/
- hint: r[i], r[j] = r[j], r[i]

#### 8.string-to-integer(atoi)
- link: https://leetcode-cn.com/problems/string-to-integer-atoi/
- hint: num = num*10 + int64(ch)

#### 12.integer-to-roman
- link: https://leetcode.com/problems/integer-to-roman/submissions/
- hint: romanArr := [15]string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

#### 9.palindrome-number
- link: https://leetcode.com/problems/palindrome-number/
- hint: s[i] != s[j]

#### 11.container-with-most-water
- link: https://leetcode.com/problems/container-with-most-water/
- hint: if height[i] < height[j] { i++ } else { j-- }

#### 13.roman-to-integer
- link: https://leetcode.com/problems/roman-to-integer/
- hint: romanMap := map[string]int {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400,
   "CM": 900,}

#### 14.longest-common-prefix
- link: https://leetcode.com/problems/longest-common-prefix/
- hint: strs[0][i] != strs[j][i]

#### 15.3sum
- link: https://leetcode.com/problems/3sum/
- hint: negSum := -1*(nums[i]+nums[j])

#### 16.3sum-closest
- link: https://leetcode.com/problems/3sum-closest/
- hint: sort.Ints(nums); num := nums[i]+nums[j]+nums[k]

#### 17.letter-combinations-of-a-phonee-number
- link: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
- hint: sub := letterCombinations(digits[1:]); newArr[i*len(sub)+k] = n

#### 18.4sum
- link: https://leetcode.com/problems/4sum/
- hint: negSum := target - nums[i] - nums[j] - nums[k]

#### 19.remove-nth-node-from-end-of-list
- link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/submissions/
- hint: if i == n { pre = p1  p1 = p1.Next }

#### 20.valid-parentheses
- link: https://leetcode.com/problems/valid-parentheses/
- hint: stack := make([]byte, len(s)/2);   index := 0

#### 21.merge-two-sorted-lists
- link: https://leetcode.com/problems/merge-two-sorted-lists/
- hint: if cl2.Val < cl1.Val { val = cl2.Val  cl2 = cl2.Next } else { cl1 = cl1.Next }

#### 22.generate-parentheses
- link: https://leetcode.com/problems/generate-parentheses/
- hint: res = append(res, "("+left+")"+right)

#### 26.remove-duplicates-from-sorted-array
- link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
- hint: if v != c { i++  nums[i] = v   c = v }

#### 27.remove-element
- link: https://leetcode.com/problems/remove-element/
- hint: if v != val { num[i] = v  i++ }

#### 28.implement-strstr
- link: https://leetcode.com/problems/implement-strstr/
- hint: if haystack[i+j] != needle[j] { isHave = false  break }

#### 64.minimum-path-sum
- link: https://leetcode.com/problems/minimum-path-sum/
- hint: path[i][j] = min(path[i-1][j], path[i][j-1]) + grid[i][j]

#### 79.word-search
- link: https://leetcode.com/problems/word-search/
- hint: search(board, word, i-1, j, idx+1) || search(board, word, i, j+1, idx+1) || search(board, word, i+1, j, idx+1) || search(board, word, i, j-1, idx+1)

#### 105.construct-binary-tree-from-preorder-and-inorder-traversal
- link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
- hint: tree.Left = buildTree(preorder[1:inIdx+1], inorder[0:inIdx]); tree.Right = buildTree(preorder[inIdx+1:len(preorder)], inorder[inIdx+1:len(inorder)])

#### 121.best-time-to-buy-and-sell-stock
- link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
- hint: prices[i] - min > profit

#### 136.single-number
- link: https://leetcode.com/problems/single-number/
- hint: ret = ret ^ num

#### 155.min-stack
- link: https://leetcode.com/problems/min-stack/
- hint: this.Stack = append(this.Stack, x - min)

#### 167.two-sum-ii-input-array-is-sorted
- link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
- hint: if idx, val := mapValue[target - val]; ok

#### 172.factorial-trailing-zeroes
- link: https://leetcode.com/problems/factorial-trailing-zeroes/
- hint: count = n / 5 + trailingZeroes(n / 5)

#### 198.house-robber
- link: https://leetcode.com/problems/house-robber/
- hint: bp[i] = max(bp[i-2] + val, bp[i-1])

#### 417.pacific-atlantic-water-flow
- link: https://leetcode.com/problems/pacific-atlantic-water-flow/
- hint: search(matrix, i, j, val, dpt)   i, i-1, i+1; j, j-1, j+1

#### 593. valid-square
- link: https://leetcode.com/problems/valid-square/
- hint: v[0] != 0 && v[0] == v[2] && v[1] == v[3] && v[4] == v[5] && 2 * v[0] == v[4]

#### 771.jewels-and-stones
- link: https://leetcode.com/problems/jewels-and-stones/
- hint: if _, ok := jMap[string(S[i])]; ok

#### 1123.lowest-common-ancestor-of-deepest-leaves
- link: https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/solution/
- hint: if leftDepth == rightDepth return tree