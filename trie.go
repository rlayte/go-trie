package trie

type Trie struct {
	value      string
	terminated bool
	children   map[string]Trie
}

func NewTrie() Trie {
	return Trie{children: map[string]Trie{}}
}

func getHeadTail(s string) (string, string) {
	head := string(s[0])
	tail := string(s[1:])

	return head, tail
}

func (t *Trie) add(s string) {
	if len(s) == 0 {
		t.terminated = true
		return
	}

	head, tail := getHeadTail(s)

	child := NewTrie()
	child.value = head
	child.add(tail)
	t.children[head] = child
}

func (t *Trie) contains(s string) bool {
	if len(s) == 0 {
		return t.terminated
	}

	head, tail := getHeadTail(s)

	if child, ok := t.children[head]; ok {
		return child.contains(tail)
	}

	return false
}
