package trie // import "kkn.fi/trie"

type (
	sTNode struct {
		next  []*sTNode
		value interface{}
	}
	// SymbolTable represents an symbol table of key-value pairs, with
	// string keys and interface{} values. It supports the usual Put, Get, Contains,
	// Delete, Len, and IsEmpty functions. It also provides rune-based functions
	// for finding the string in the symbol table that is the longest prefix of
	// a given prefix, finding all strings in the symbol table that start with
	// a given prefix, and finding all strings in the symbol table that match
	// a given pattern. A symbol table implements the associative array abstraction:
	// when associating a value with a key that is already in the symbol table, the
	// convention is to replace the old value with the new value.  This trie uses
	// the convention that values cannot be nil. Setting the value associated with
	// a key to nil is equivalent to deleting the key from the symbol table.
	//
	// This implementation uses a 256-way trie. The Put, Contains, Delete, and
	// longest prefix functions take time proportional to the length of the key (in
	// the worst case). Construction takes constant time. The Len, and IsEmpty
	// functions take constant time. Construction takes constant time.
	SymbolTable struct {
		root   *sTNode
		length int
	}
)

// NewSymbolTable returns a trie based on symbol table implementation.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{}
}

// Put inserts the key-value pair into the trie, overwriting the old
// value with the new value if the key is already in the symbol table.
// If the value is nil, this effectively deletes the key from the symbol table.
// If key is empty this function will silently return.
func (t *SymbolTable) Put(key string, value interface{}) {
	if key == "" {
		return
	}
	if value == nil {
		t.Delete(key)
	} else {
		t.root = t.put(t.root, []rune(key), value, 0)
	}
}

func (t *SymbolTable) put(x *sTNode, key []rune, value interface{}, d int) *sTNode {
	if x == nil {
		x = &sTNode{
			next: make([]*sTNode, r),
		}
	}
	if d == len(key) {
		if x.value == nil {
			t.length++
		}
		x.value = value
		return x
	}
	c := key[d]
	x.next[c] = t.put(x.next[c], key, value, d+1)
	return x
}

// Get returns the value associated with the given key.
func (t *SymbolTable) Get(key string) interface{} {
	x := t.get(t.root, []rune(key), 0)
	if x == nil {
		return nil
	}
	return x.value
}

func (t *SymbolTable) get(x *sTNode, key []rune, d int) *sTNode {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	c := key[d]
	return t.get(x.next[c], key, d+1)
}

// Delete removes the key from the symbol table if the key is present.
func (t *SymbolTable) Delete(key string) {
	t.root = t.delete(t.root, []rune(key), 0)
}

func (t *SymbolTable) delete(x *sTNode, key []rune, d int) *sTNode {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.value != nil {
			t.length--
		}
		x.value = nil
	} else {
		c := key[d]
		x.next[c] = t.delete(x.next[c], key, d+1)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.value != nil {
		return x
	}
	for c := 0; c < r; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}

// Contains returns true if the trie contains key and false otherwise.
func (t *SymbolTable) Contains(key string) bool {
	return t.Get(key) != nil
}

// IsEmpty returns true if trie is empty and false otherwise.
func (t *SymbolTable) IsEmpty() bool {
	return t.length == 0
}

// LongestPrefixOf returns the string in the symbol table that is the
// longest prefix of query, or empty string, if no such string is found
// in the trie.
func (t *SymbolTable) LongestPrefixOf(query string) string {
	q := []rune(query)
	length := t.longestPrefixOf(t.root, q, 0, 0)
	return string(q[0:length])
}

func (t *SymbolTable) longestPrefixOf(x *sTNode, query []rune, d, length int) int {
	if x == nil {
		return length
	}
	if x.value != nil {
		length = d
	}
	if d == len(query) {
		return length
	}
	c := query[d]
	return t.longestPrefixOf(x.next[c], query, d+1, length)
}

// KeysWithPrefix returns all the keys in the trie that match prefix.
func (t *SymbolTable) KeysWithPrefix(prefix string) []string {
	results := new(stringQueue)
	x := t.get(t.root, []rune(prefix), 0)
	t.collect(x, []rune(prefix), results)
	return results.slice()
}

func (t *SymbolTable) collect(x *sTNode, prefix []rune, results *stringQueue) {
	if x == nil {
		return
	}
	if x.value != nil {
		results.enqueue(string(prefix))
	}
	for c := 0; c < r; c++ {
		prefix = append(prefix, rune(c))
		t.collect(x.next[c], prefix, results)
		prefix = prefix[0 : len(prefix)-1]
	}
}

// KeysThatMatch all of the keys in the symbol table that match pattern,
// where '.' symbol is treated as a wildcard character.
func (t *SymbolTable) KeysThatMatch(pattern string) []string {
	results := new(stringQueue)
	t.collectWildcard(t.root, []rune(""), []rune(pattern), results)
	return results.slice()
}

func (t *SymbolTable) collectWildcard(x *sTNode, prefix, pattern []rune, results *stringQueue) {
	if x == nil {
		return
	}
	d := len(prefix)
	if d == len(pattern) && x.value != nil {
		results.enqueue(string(prefix))
	}
	if d == len(pattern) {
		return
	}
	c := pattern[d]
	if c == '.' {
		for ch := 0; ch < r; ch++ {
			prefix = append(prefix, rune(ch))
			t.collectWildcard(x.next[ch], prefix, pattern, results)
			prefix = prefix[0 : len(prefix)-1]
		}
	} else {
		prefix = append(prefix, c)
		t.collectWildcard(x.next[c], prefix, pattern, results)
	}
}

// Keys returns all the keys in the trie.
func (t *SymbolTable) Keys() []string {
	return t.KeysWithPrefix("")
}

// Len returns the number of strings in the trie.
func (t *SymbolTable) Len() int {
	return t.length
}
