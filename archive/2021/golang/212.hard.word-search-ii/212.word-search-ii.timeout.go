/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 *
 * https://leetcode.com/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (37.54%)
 * Likes:    3769
 * Dislikes: 147
 * Total Accepted:    305.5K
 * Total Submissions: 810.4K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
  '["oath","pea","eat","rain"]'
 *
 * Given an m x n board of characters and a list of strings words, return all
 * words on the board.
 *
 * Each word must be constructed from letters of sequentially adjacent cells,
 * where adjacent cells are horizontally or vertically neighboring. The same
 * letter cell may not be used more than once in a word.
 *
 *
 * Example 1:
 *
 *
 * Input: board =
 * [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
 * words = ["oath","pea","eat","rain"]
 * Output: ["eat","oath"]
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["a","b"],["c","d"]], words = ["abcb"]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n == board[i].length
 * 1 <= m, n <= 12
 * board[i][j] is a lowercase English letter.
 * 1 <= words.length <= 3 * 10^4
 * 1 <= words[i].length <= 10
 * words[i] consists of lowercase English letters.
 * All the strings of words are unique.
 *
 *
*/

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	// 核心还是 word-search 的逻辑。问题在于怎么能一次行判断多个 words 是否存在，快速终止失败条件。
	trie := getTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	ans := make(map[string]bool)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if len(ans) == len(words) {
				break
			}
			backtrack(board, &trie, row, col, visited, make([]byte, 0), ans)
		}
	}
	result := make([]string, 0, len(ans))
	for k := range ans {
		result = append(result, k)
	}
	return result
}

// word-search
func backtrack(board [][]byte, trie *Trie, x, y int, visited [][]bool, prefix []byte, ans map[string]bool) {
	// 如果 prefix 匹配，就写到结果里
	if trie.Search(prefix) {
		ans[string(prefix)] = true
	}
	// 越界了
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}
	// 访问过了
	if visited[x][y] {
		return
	}

	// choose
	prefix = append(prefix, board[x][y])

	visited[x][y] = true                     // choose
	defer func() { visited[x][y] = false }() // unchoose

	// 当前访问路径不是任何word的前缀
	if !trie.StartsWith(prefix) {
		return
	}
	backtrack(board, trie, x-1, y, visited, prefix, ans)
	backtrack(board, trie, x+1, y, visited, prefix, ans)
	backtrack(board, trie, x, y-1, visited, prefix, ans)
	backtrack(board, trie, x, y+1, visited, prefix, ans)

	prefix = prefix[:len(prefix)-1] // unchoose
}

type TrieNode struct {
	Value    byte
	Children map[byte]*TrieNode // value -> Node
	isEnd    bool
}

type Trie struct {
	Root *TrieNode
}

/** Initialize your data structure here. */
func getTrie() Trie {
	return Trie{
		Root: &TrieNode{
			Children: make(map[byte]*TrieNode),
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this.Root
	for i := range word {
		r := word[i]
		node, ok := curr.Children[r]
		if !ok {
			node = &TrieNode{
				Value:    r,
				Children: make(map[byte]*TrieNode),
			}
			curr.Children[r] = node
		}
		curr = node
	}
	curr.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word []byte) bool {
	curr := this.Root
	for _, r := range word {
		node, ok := curr.Children[r]
		if !ok {
			return false
		}
		curr = node
	}
	return curr.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix []byte) bool {
	curr := this.Root
	for _, r := range prefix {
		node, ok := curr.Children[r]
		if !ok {
			return false
		}
		curr = node
	}
	return true
}

// @lc code=end

