/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 *
 * https://leetcode.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (52.73%)
 * Likes:    4598
 * Dislikes: 70
 * Total Accepted:    419.6K
 * Total Submissions: 789.7K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * A trie (pronounced as "try") or prefix tree is a tree data structure used to
 * efficiently store and retrieve keys in a dataset of strings. There are
 * various applications of this data structure, such as autocomplete and
 * spellchecker.
 *
 * Implement the Trie class:
 *
 *
 * Trie() Initializes the trie object.
 * void insert(String word) Inserts the string word into the trie.
 * boolean search(String word) Returns true if the string word is in the trie
 * (i.e., was inserted before), and false otherwise.
 * boolean startsWith(String prefix) Returns true if there is a previously
 * inserted string word that has the prefix prefix, and false otherwise.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
 * [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
 * Output
 * [null, null, true, false, true, null, true]
 *
 * Explanation
 * Trie trie = new Trie();
 * trie.insert("apple");
 * trie.search("apple");   // return True
 * trie.search("app");     // return False
 * trie.startsWith("app"); // return True
 * trie.insert("app");
 * trie.search("app");     // return True
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= word.length, prefix.length <= 2000
 * word and prefix consist only of lowercase English letters.
 * At most 3 * 10^4 calls in total will be made to insert, search, and
 * startsWith.
 *
 *
*/

// @lc code=start
type TrieNode struct {
	Value    rune
	Children map[rune]*TrieNode // value -> Node
	isEnd    bool
}

type Trie struct {
	Root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this.Root
	for _, r := range word {
		node, ok := curr.Children[r]
		if !ok {
			node = &TrieNode{
				Value:    r,
				Children: make(map[rune]*TrieNode),
			}
			curr.Children[r] = node
		}
		curr = node
	}
	curr.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
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
func (this *Trie) StartsWith(prefix string) bool {
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

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

