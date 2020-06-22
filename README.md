# leetcode
leetcode link and hint

#### 3. longest-substring-without-repeating-characters
- link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
- hint: if _, ok := maxMap[s[i]]; ok { if i-sidx > maxLen { maxLen = i - sidx } }

#### 5. longest-palindromic-substring
- link: https://leetcode.com/problems/longest-palindromic-substring/
- hint: if s[i] == s[i-1] {...}  if i+1 < len(s) && s[i+1] == s[i-1] {...}

#### 6. zigzag-conversion
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

#### 24.swap-nodes-in-pairs
- link: https://leetcode.com/problems/swap-nodes-in-pairs/
- hint: p1.Next = p2.Next; p2.Next = p1; p.Next = p2

#### 26.remove-duplicates-from-sorted-array
- link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
- hint: if v != c { i++  nums[i] = v   c = v }

#### 27.remove-element
- link: https://leetcode.com/problems/remove-element/
- hint: if v != val { num[i] = v  i++ }

#### 28.implement-strstr
- link: https://leetcode.com/problems/implement-strstr/
- hint: if haystack[i+j] != needle[j] { isHave = false  break }

#### 29.divide-two-integers
- link: https://leetcode.com/problems/divide-two-integers/
- hint: for x + x <= a { c = c + c; x = x + x; }

#### 31.next-permutation
- link: https://leetcode.com/problems/next-permutation/
- hint: for ; i >= 0; i-- { if nums[i] < nums[i+1] { break } }

#### 33.search-in-rotated-sorted-array
- link: https://leetcode.com/problems/search-in-rotated-sorted-array/
- hint: for ; i >= 0; i-- { if nums[i] < nums[i+1] { break } }

#### 34.find-first-and-last-position-of-element-in-sorted-array
- link: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
- hint: if target < nums[mid] && target >= nums[i] {  return binarySearch(nums, i, mid-1, target)  }

#### 35.search-insert-position
- link: https://leetcode.com/problems/search-insert-position/
- hint: if nums[idx] > target { return idx }

### 36.valid-sudoku
- link: https://leetcode.com/problems/valid-sudoku/
- hint: b[3*i+j] = board[3*(k/3)+i][3*(k%3)+j]

### 38.count-and-say
- link: https://leetcode.com/problems/count-and-say/
- hint: str := countAndSay(n - 1)

#### 39.combination-sum
- link: https://leetcode.com/problems/combination-sum/
- hint: dfs(candidates, list, i, target-candidates[i], sol)

#### 40.combination-sum-ii
- link: https://leetcode.com/problems/combination-sum-ii/
- hint: dfs(candidates, ret, temp, i+1, target-(*candidates)[i])

#### 43.multiply-strings
- link: https://leetcode.com/problems/multiply-strings/
- hint: res[p2] = sum % 10; res[p1] += sum / 10

#### 46.permutations
- link: https://leetcode.com/problems/permutations/
- hint: nums = nums[1:]; perms := permute(nums)

#### 47.permutations-ii
- link: https://leetcode.com/problems/permutations-ii/
- hint: if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] { continue }

#### 48.rotate-image
- link: https://leetcode.com/problems/ratate-image/
- hint: matrix[i][n-1-j+k], matrix[i+k][j], matrix[n-1-i][j-k], matrix[n-1-i-k][n-1-j] = matrix[n-1-i-k][n-1-j], matrix[i][n-1-j+k], matrix[i+k][j], matrix[n-1-i][j-k]

#### 49.group-anagrams
- link: https://leetcode.com/problems/group-anagrams/
- hint: sort.Sort(StrObjs(objs))

#### 50.powx-n
- link: https://leetcode.com/problems/powx-n/
- hint: if n%2 == 1 { res *= x }

#### 54.spiral-matrix
- link: https://leetcode.com/problems/spiral-matrix/
- hint: for left <= right && top <= bottom { }

#### 55.jump-game
- link: https://leetcode.com/problems/jump-game/
- hint: if reach >= i && nums[i]+i > reach { reach = nums[i] + i }

#### 56.merge-intervals
- link: https://leetcode.com/problems/merge-intervals/
- hint: sort.Sort(Intervals(intervals))

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