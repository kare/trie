package trie // import "kkn.fi/trie"

type (
	tSNode struct {
		c     rune
		left  *tSNode
		mid   *tSNode
		right *tSNode
		value interface{}
	}
	// TernarySearch is a symbol table with string keys and interface{} values.
	// It implements ternary search trie.
	TernarySearch struct {
		length int
		root   *tSNode
	}
)

// NewTernarySearch returns an empty ternary search trie.
func NewTernarySearch() *TernarySearch {
	return &TernarySearch{}
}

// Contains return true for an existing key.
func (t TernarySearch) Contains(key string) bool {
	return t.Get(key) != nil
}

// Get returns value for a key. Get will return nil for an empty key or when key
// is not found.
func (t TernarySearch) Get(key string) interface{} {
	if key == "" {
		return nil
	}
	x := t.get(t.root, []rune(key), 0)
	if x == nil {
		return nil
	}
	return x.value
}

// return subtrie corresponding to given key
func (t TernarySearch) get(x *tSNode, key []rune, d int) *tSNode {
	if len(key) == 0 || x == nil {
		return nil
	}
	c := key[d]
	if c < x.c {
		return t.get(x.left, key, d)
	} else if c > x.c {
		return t.get(x.right, key, d)
	} else if d < len(key)-1 {
		return t.get(x.mid, key, d+1)
	}
	return x
}

// Put inserts string key into trie
// If key is empty this function will silently return
func (t *TernarySearch) Put(key string, val interface{}) {
	if key == "" {
		return
	}
	if !t.Contains(key) {
		t.length++
	}
	t.root = t.put(t.root, []rune(key), val, 0)
}

func (t *TernarySearch) put(x *tSNode, key []rune, val interface{}, d int) *tSNode {
	c := key[d]
	if x == nil {
		x = new(tSNode)
		x.c = c
	}
	if c < x.c {
		x.left = t.put(x.left, key, val, d)
	} else if c > x.c {
		x.right = t.put(x.right, key, val, d)
	} else if d < len(key)-1 {
		x.mid = t.put(x.mid, key, val, d+1)
	} else {
		x.value = val
	}
	return x
}

// Delete removes the key from the trie if the key is present.
func (t TernarySearch) Delete(key string) {
	panic("not implemented!")
}

// LongestPrefixOf returns longest prefix of argument prefix in trie
func (t TernarySearch) LongestPrefixOf(query string) string {
	if len(query) == 0 {
		return ""
	}
	length := 0
	x := t.root
	q := []rune(query)
	for i := 0; x != nil && i < len(q); {
		c := q[i]
		if c < x.c {
			x = x.left
		} else if c > x.c {
			x = x.right
		} else {
			i++
			if x.value != nil {
				length = i
			}
			x = x.mid
		}
	}
	return string(q[0:length])
}

// Keys returns all the keys in the trie.
func (t TernarySearch) Keys() []string {
	queue := new(stringQueue)
	t.collect(t.root, []rune(""), queue)
	return queue.slice()
}

// KeysWithPrefix returns all keys starting with given prefix.
func (t TernarySearch) KeysWithPrefix(prefix string) []string {
	queue := new(stringQueue)
	x := t.get(t.root, []rune(prefix), 0)
	if x == nil {
		return queue.slice()
	}
	if x.value != nil {
		queue.enqueue(prefix)
	}
	t.collect(x.mid, []rune(prefix), queue)
	return queue.slice()
}

// all keys in subtrie rooted at x with given prefix
func (t TernarySearch) collect(x *tSNode, prefix []rune, queue *stringQueue) {
	if x == nil {
		return
	}
	t.collect(x.left, prefix, queue)
	if x.value != nil {
		queue.enqueue(string(append(prefix, x.c)))
	}
	t.collect(x.mid, append(prefix, x.c), queue)
	t.collect(x.right, prefix, queue)
}

// KeysThatMatch returns all keys matching given wildcard pattern
func (t TernarySearch) KeysThatMatch(pattern string) []string {
	queue := new(stringQueue)
	t.collectWildcard(t.root, []rune(""), []rune(pattern), 0, queue)
	return queue.slice()
}

func (t TernarySearch) collectWildcard(x *tSNode, prefix, pattern []rune, i int, q *stringQueue) {
	if x == nil {
		return
	}
	c := pattern[i]
	if c == '.' || c < x.c {
		t.collectWildcard(x.left, prefix, pattern, i, q)
	}
	if c == '.' || c == x.c {
		if i == len(pattern)-1 && x.value != nil {
			q.enqueue(string(append(prefix, x.c)))
		}
		if i < len(pattern)-1 {
			t.collectWildcard(x.mid, append(prefix, x.c), pattern, i+1, q)
		}
		if c == '.' || c > x.c {
			t.collectWildcard(x.right, prefix, pattern, i, q)
		}
	}
}

// Len returns length of trie.
func (t TernarySearch) Len() int {
	return t.length
}

// IsEmpty returns true when trie is empty.
func (t TernarySearch) IsEmpty() bool {
	return t.length == 0
}
