package trie

import "testing"

func TestContains(t *testing.T) {
	trie := NewTrie()
	trie.add("Hello")
	trie.add("World")

	if !trie.contains("Hello") {
		t.Error("Should have key 'Hello'")
	}

	if !trie.contains("World") {
		t.Error("Should have key 'World'")
	}

	if trie.contains("Hel") {
		t.Error("Should not contain key 'Hel'")
	}

	if trie.contains("Skepta") {
		t.Error("Should not contain 'Skepta'")
	}
}
