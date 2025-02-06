package main

// Implement Trie (Prefix Tree)

// Trie ...
type Trie struct {
	Children map[rune]*Trie
	IsEnd    bool
}

// Constructor ...
func Constructor() Trie {
	return Trie{Children: make(map[rune]*Trie)}
}

// Insert ...
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		if node.Children[ch] == nil {
			node.Children[ch] = &Trie{Children: make(map[rune]*Trie)}
		}
		node = node.Children[ch]
	}
	node.IsEnd = true
}

// Search ...
func (t *Trie) Search(word string) bool {
	node := t.findNode(word)
	return node != nil && node.IsEnd
}

// StartsWith ...
func (t *Trie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

func (t *Trie) findNode(word string) *Trie {
	node := t
	for _, ch := range word {
		if node.Children[ch] == nil {
			return nil
		}
		node = node.Children[ch]
	}
	return node
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
