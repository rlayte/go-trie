package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()

	keys := []string{"Hello", "World", "Help", "Hell"}
	missing := []string{"Hel", "Nope"}

	for _, key := range keys {
		trie.Add(key)
	}

	for _, key := range keys {
		_, ok := trie.Get(key)

		if !ok {
			t.Error("Missing key:", key)
		}
	}

	for _, key := range missing {
		_, ok := trie.Get(key)

		if ok {
			t.Error("Should not have key:", key)
		}
	}

	child, _ := trie.Get("Hello")
	child.Value = "Updated"

	if child, _ := trie.Get("Hello"); child.Value != "Updated" {
		t.Error("Hello should now equal 100 but is", child.Value)
	}

	children := trie.Children("")

	if len(children) != 4 {
		t.Error("There should be 4 children. Actually:", len(children))
	}

	h, _ := trie.Get("H")
	nested := h.Children("")

	if len(nested) != 3 {
		t.Error("There should be 2 children from 'H'. Actually:", len(nested))
	}

	trie.Remove("Hell")

	if _, ok := trie.Get("Hell"); ok {
		t.Error("Hell should have been removed")
	}

	if _, ok := trie.Get("Hello"); !ok {
		t.Error("Hello should not have been removed")
	}
}
