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

#### 58.length-of-last-word
- link: https://leetcode.com/problems/length-of-last-word/
- hint: for i := count - 1; i >= 0; i-- { }

#### 59.spiral-matrix-ii
- link: https://leetcode.com/problems/spiral-matrix-ii/
- hint: for k := 1; k <= n*n && left < right && top < bottom; { }

#### 60.permutation-sequence
- link: https://leetcode.com/problems/permutation-sequence/
- hint: cur := factorial(n - depth - 1)

#### 61.rotate-list
- link: https://leetcode.com/problems/rotate-list/
- hint: i, j := 0, 0; p1, p2 := head, sentry

#### 62.unique-paths
- link: https://leetcode.com/problems/unique-paths/
- hint: dp[i][j] = dp[i-1][j] + dp[i][j-1]

#### 63.unique-paths-ii
- link: https://leetcode.com/problems/unique-paths-ii/
- hint: dp[i][j] = dp[i-1][j] + dp[i][j-1]

#### 64.minimum-path-sum
- link: https://leetcode.com/problems/minimum-path-sum/
- hint: path[i][j] = min(path[i-1][j], path[i][j-1]) + grid[i][j]

#### 66.plus-one
- link: https://leetcode.com/problems/plus-one/
- hint: n := digits[i] + carry

#### 67.add-binary
- link: https://leetcode.com/problems/add-binary/
- hint: if carry == 1 { res[0] = '1' }

#### 69.sqrtx
- link: https://leetcode.com/problems/sqrtx/
- hint: k := (i + j) / 2

#### 70.climbing-stairs
- link: https://leetcode.com/problems/climbing-stairs/
- hint: dp[i] = dp[i-1] + dp[i-2]

#### 71.simplify-path
- link: https://leetcode.com/problems/simplify-path/
- hint: stack.Push(str)

#### 73.set-matrix-zeroes
- link: https://leetcode.com/problems/set-matrix-zeroes/
- hint: rz, cz := false, false

#### 74.search-a-2d-matrix
- link: https://leetcode.com/problems/search-a-2d-matrix/
- hint: if target == matrix[row][k] { }

#### 75.sort-colors
- link: https://leetcode.com/problems/sort-colors/
- hint: if nums[idx] == 0 { } else if nums[idx] == 2 { }

#### 77.combinations
- link: https://leetcode.com/problems/combinations/
- hint: dfs(res, temp, n, k, i+1)

#### 78.subsets
- link: https://leetcode.com/problems/subsets/
- hint: temp := combine(count, i); idxs = append(idxs, temp...)

#### 79.word-search
- link: https://leetcode.com/problems/word-search/
- hint: search(board, word, i-1, j, idx+1) || search(board, word, i, j+1, idx+1) || search(board, word, i+1, j, idx+1) || search(board, word, i, j-1, idx+1)

#### 80.remove-duplicates-from-sorted-array-ii
- link: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
- hint: nums[rc-1] = nums[i]

#### 81.search-in-rotated-sorted-array-ii
- link: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
- hint: if nums[i] == nums[k] { i++ }

#### 82.remove-duplicates-from-sorted-list-ii
- link: https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
- hint: ll, last, cur := dumy, dumy, head

#### 83.remove-duplicates-from-sorted-list
- link: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
- hint: if cur == head || cur.Val != last.Val { last.Next = cur; last = cur }

#### 86.partition-list
- link: https://leetcode.com/problems/partition-list/
- hint: cur, leftCur, rightCur := head, left, right

#### 88.merge-sorted-array
- link: https://leetcode.com/problems/merge-sorted-array/
- hint: copy(temp, nums1[0:m])

#### 89.gray-code
- link: https://leetcode.com/problems/gray-code/
- hint: res[int(j)] = res[int(end-j-1)] | (1 << (i - 1))

#### 90.subsets-ii
- link: https://leetcode.com/problems/subsets-ii/
- hint: dfs(nums, res, temp, n, k, i+1)

#### 91.decode-ways
- link: https://leetcode.com/problems/decode-ways/
- hint: cur, pre := 1, 1

#### 92.reverse-linked-list-ii
- link: https://leetcode.com/problems/reverse-linked-list-ii/
- hint: pre, last, cur := dummy, head, head

#### 93.restore-ip-address
- link: https://leetcode.com/problems/restore-ip-address/
- hint: dfs(s, res, temp, depth+1, j+1)

#### 94.binary-tree-inorder-traversal
- link: https://leetcode.com/problems/binary-tree-inorder-traversal/
- hint: *res = append(*res, root.Val)

#### 95.unique-binary-search-trees-ii
- link: https://leetcode.com/problems/unique-binary-search-trees-ii/
- hint: lefts := preTrees(i, k-1); rights := preTrees(k+1, j)

#### 96.unique-binary-search-trees
- link: https://leetcode.com/problems/unique-binary-search-trees/
- hint: num += dp[j-1] * dp[i-j]

#### 98.validate-binary-search-tree
- link: https://leetcode.com/problems/validate-binary-search-tree/
- hint: if len(*temp) > 0 && root.Val <= (*temp)[len(*temp)-1] { return false }

#### 100.same-tree
- link: https://leetcode.com/problems/same-tree/
- hint: if isSameTree(p.Left, q.Left) == false { return false }

#### 101.symmetric-tree
- link: https://leetcode.com/problems/symmetric-tree/
- hint: temp = append(temp, node.Left); temp = append(temp, node.Right)

#### 102.binary-tree-level-order-traversal
- link: https://leetcode.com/problems/binary-tree-level-order-traversal/
- hint: if node.Left != nil { temp = append(temp, node.Left) }

#### 103.binary-tree-zigzag-level-order-traversal
- link: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
- hint: dir = ^dir

#### 104.maximum-depth-of-binary-tree
- link: https://leetcode.com/problems/maximum-depth-of-binary-tree/
- hint: left := maxDepth(root.Left); right := maxDepth(root.Right)

#### 105.construct-binary-tree-from-preorder-and-inorder-traversal
- link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
- hint: tree.Left = buildTree(preorder[1:inIdx+1], inorder[0:inIdx]); tree.Right = buildTree(preorder[inIdx+1:len(preorder)], inorder[inIdx+1:len(inorder)])

#### 106.construct-binary-tree-from-inorder-and-postorder-traversal
- link: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
- hint: root.Left = buildTree(inorder[0:inIdx], postorder[0:inIdx])

#### 107.binary-tree-level-order-traversal-ii
- link: https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
- hint: res[i], res[j] = res[j], res[i]

#### 108.convert-sorted-array-to-binary-search-tree
- link: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
- hint: root.Left = sortedArrayToBST(nums[0 : count/2])

#### 109.convert-sorted-list-to-binary-search-tree
- link: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
- hint: nums = append(nums, head.Val)

#### 110.balance-binary-tree
- link: https://leetcode.com/problems/balance-binary-tree/
- hint: if left == right || left == right+1 || left == right-1 { return isBalanced(root.Left) && isBalanced(root.Right) }

#### 111.minimum-depth-of-binary-tree
- link: https://leetcode.com/problems/minimum-depth-of-binary-tree/
- hint: dfs(root.Left, depth+1, minDepth)

#### 112.path-sum
- link: https://leetcode.com/problems/path-sum/
- hint: if root.Left == nil && root.Right == nil && root.Val == sum { return true }

#### 113.path-sum-ii
- link: https://leetcode.com/problems/path-sum-ii/
- hint: preTree(root.Left, sum-root.Val, temp, res)

#### 114.flatten-binary-tree-to-linked-list
- link: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
- hint: flatten(l)

#### 116.populating-next-right-pointers-in-each-node
- link: https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
- hint: head.Right.Next = head.Next.Left

#### 114.populating-next-right-pointers-in-each-node-ii
- link: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
- hint: pre, leftNode = cur.Left, cur.Left

#### 118.pascals-triangle
- link: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
- hint: temp[j] = res[i-2][j] + res[i-2][j-1]

#### 119.pascals-triangle-ii
- link: https://leetcode.com/problems/pascals-triangle-ii/
- hint: temp := res[j] + pre

#### 120.triangle
- link: https://leetcode.com/problems/triangle/
- hint: dp[j] = pre + triangle[i][j]

#### 121.best-time-to-buy-and-sell-stock
- link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
- hint: prices[i] - min > profit

#### 122.best-time-to-buy-and-sell-stock-ii
- link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
- hint: profit += prices[i] - prices[i-1]

#### 125.valid-palindrome
- link: https://leetcode.com/problems/valid-palindrome/
- hint: if i < len(s) && j >= 0 && s[i] != s[j] { }

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