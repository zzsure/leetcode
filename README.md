# leetcode
leetcode link and hint

#### 7.reverse-integer
- link: https://leetcode-cn.com/problems/reverse-integer/submissions/
- hint: r[i], r[j] = r[j], r[i]

#### 13.roman-to-integer
- link: https://leetcode-cn.com/problems/roman-to-integer/submissions/
- hint: romanMap := map[string]int {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400,
   "CM": 900,}

#### 64.minimum-path-sum
- link: https://leetcode-cn.com/problems/minimum-path-sum/
- hint: path[i][j] = min(path[i-1][j], path[i][j-1]) + grid[i][j]

#### 79.word-search
- link: https://leetcode-cn.com/problems/word-search/
- hint: search(board, word, i-1, j, idx+1) || search(board, word, i, j+1, idx+1) || search(board, word, i+1, j, idx+1) || search(board, word, i, j-1, idx+1)

#### 105.construct-binary-tree-from-preorder-and-inorder-traversal
- link: https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
- hint: tree.Left = buildTree(preorder[1:inIdx+1], inorder[0:inIdx]); tree.Right = buildTree(preorder[inIdx+1:len(preorder)], inorder[inIdx+1:len(inorder)])

#### 121.best-time-to-buy-and-sell-stock
- link: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
- hint: prices[i] - min > profit

#### 136.single-number
- link: https://leetcode-cn.com/problems/single-number/
- hint: ret = ret ^ num

#### 155.min-stack
- link: https://leetcode-cn.com/problems/min-stack/
- hint: this.Stack = append(this.Stack, x - min)

#### 167.two-sum-ii-input-array-is-sorted
- link: https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
- hint: if idx, val := mapValue[target - val]; ok

#### 172.factorial-trailing-zeroes
- link: https://leetcode-cn.com/problems/factorial-trailing-zeroes/
- hint: count = n / 5 + trailingZeroes(n / 5)

#### 198.house-robber
- link: https://leetcode-cn.com/problems/house-robber/
- hint: bp[i] = max(bp[i-2] + val, bp[i-1])

#### 417.pacific-atlantic-water-flow
- link: https://leetcode-cn.com/problems/pacific-atlantic-water-flow/
- hint: search(matrix, i, j, val, dpt)   i, i-1, i+1; j, j-1, j+1

#### 593. valid-square
- link: https://leetcode-cn.com/problems/valid-square/
- hint: v[0] != 0 && v[0] == v[2] && v[1] == v[3] && v[4] == v[5] && 2 * v[0] == v[4]

#### 771.jewels-and-stones
- link: https://leetcode-cn.com/problems/jewels-and-stones/
- hint: if _, ok := jMap[string(S[i])]; ok

#### 1123.lowest-common-ancestor-of-deepest-leaves
- link: https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves/solution/
- hint: if leftDepth == rightDepth return tree