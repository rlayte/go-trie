package trie

type Trie struct {
	Value      interface{}
	terminated bool
	children   map[string]*Trie
}

func NewTrie() *Trie {
	return &Trie{children: map[string]*Trie{}}
}

func (t *Trie) Add(k string) *Trie {
	if len(k) == 0 {
		t.terminated = true
		return t
	}

	head := string(k[0])
	tail := string(k[1:])

	child, ok := t.children[head]

	if !ok {
		child = NewTrie()
		t.children[head] = child
	}

	return child.Add(tail)
}

func (t *Trie) removeChild(k string, ancestor *Trie, c string) {
	if len(k) == 0 {
		t.terminated = false

		if len(t.children) == 0 {
			delete(ancestor.children, c)
		}

		return
	}

	head := string(k[0])
	tail := string(k[1:])

	if child, ok := t.children[head]; ok {
		if len(t.children) > 1 {
			ancestor = t
			c = head
		}

		child.removeChild(tail, ancestor, c)
	}
}

func (t *Trie) Remove(k string) {
	head := string(k[0])
	t.removeChild(k, t, head)
}

func (t *Trie) Get(k string) (*Trie, bool) {
	if len(k) == 0 {
		return t, t.terminated
	}

	head := string(k[0])
	tail := string(k[1:])

	if child, ok := t.children[head]; ok {
		return child.Get(tail)
	}

	return t, false
}

func (t *Trie) Children(prefix string) []string {
	children := []string{}

	for key, child := range t.children {
		if child.terminated {
			children = append(children, prefix+key)
		}

		children = append(children, child.Children(prefix+key)...)
	}

	return children
}
