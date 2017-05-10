package trie // import "kkn.fi/trie"

const r = 256 // extended ascii

type (
	// r-way trie node
	node struct {
		next     []*node
		isString bool // isWord
	}
	// Trie represents an ordered set of strings over
	// the extended ASCII alphabet.
	// It supports the usual Add, Contains, and Delete
	// functions. It also provides character-based functions for
	// finding the string in the set that is the longest prefix
	// of a given prefix, finding all strings in the set that
	// start with a given prefix, and finding all strings in the set
	// that match a given pattern.
	//
	// This implementation uses a 256-way trie.
	// The Add, Contains, Delete, and
	// LongestPrefixOf functions take time proportional to the length
	// of the key (in the worst case). Construction takes constant time.
	Trie struct {
		root   *node
		length int
	}
	stringQueue []string
)

// New returns an empty trie.
func New() *Trie {
	return &Trie{}
}

// Contains returns true if the set contains key and false otherwise.
func (t *Trie) Contains(key string) bool {
	x := t.get(t.root, []rune(key), 0)
	if x == nil {
		return false
	}
	return x.isString
}

func (t *Trie) get(x *node, key []rune, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		return x
	}
	c := key[d]
	return t.get(x.next[c], key, d+1)
}

// Add adds a key to the set if not present.
// If key is empty function will silently return
func (t *Trie) Add(key string) {
	if key == "" {
		return
	}
	t.root = t.add(t.root, []rune(key), 0)
}

func (t *Trie) add(x *node, key []rune, d int) *node {
	if x == nil {
		x = &node{
			next: make([]*node, r),
		}
	}
	if d == len(key) {
		if !x.isString {
			t.length++
		}
		x.isString = true
	} else {
		c := key[d]
		x.next[c] = t.add(x.next[c], key, d+1)
	}
	return x
}

// Len returns the number of strings in the set.
func (t *Trie) Len() int {
	return t.length
}

// IsEmpty returns true if set is empty.
func (t *Trie) IsEmpty() bool {
	return t.length == 0
}

// KeysWithPrefix returns all the keys in the set that match prefix.
func (t *Trie) KeysWithPrefix(prefix string) []string {
	results := &stringQueue{}
	x := t.get(t.root, []rune(prefix), 0)
	t.collect(x, []rune(prefix), results)
	return results.slice()
}

func (t *Trie) collect(x *node, prefix []rune, results *stringQueue) {
	if x == nil {
		return
	}
	if x.isString {
		results.enqueue(string(prefix))
	}
	for c := 0; c < r; c++ {
		prefix = append(prefix, rune(c))
		t.collect(x.next[c], prefix, results)
		prefix = prefix[0 : len(prefix)-1]
	}
}

// KeysThatMatch all of the keys in the set that match pattern,
// where '.' symbol is treated as a wildcard character.
func (t *Trie) KeysThatMatch(pattern string) []string {
	results := new(stringQueue)
	t.collectWildcard(t.root, []rune(""), []rune(pattern), results)
	return results.slice()
}

func (t *Trie) collectWildcard(x *node, prefix, pattern []rune, results *stringQueue) {
	if x == nil {
		return
	}
	d := len(prefix)
	if d == len(pattern) && x.isString {
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
		prefix = append(prefix, rune(c))
		t.collectWildcard(x.next[c], prefix, pattern, results)
	}
}

// LongestPrefixOf Returns the string in the set that is the
// longest prefix of query, or an empty string, if no such string.
func (t *Trie) LongestPrefixOf(query string) string {
	length := t.longestPrefixOf(t.root, []rune(query), 0, -1)
	if length == -1 {
		return ""
	}
	return string([]rune(query)[:length])
}

// returns the length of the longest string key in the subtrie
// rooted at x that is a prefix of the query string,
// assuming the first d character match and we have already
// found a prefix match of length length
func (t *Trie) longestPrefixOf(x *node, query []rune, d, length int) int {
	if x == nil {
		return length
	}
	if x.isString {
		length = d
	}
	if d == len(query) {
		return length
	}
	c := query[d]
	return t.longestPrefixOf(x.next[c], query, d+1, length)
}

// Delete deletes the key from the set if it is present.
func (t *Trie) Delete(key string) {
	t.root = t.delete(t.root, []rune(key), 0)
}

func (t *Trie) delete(x *node, key []rune, d int) *node {
	if x == nil {
		return nil
	}
	if d == len(key) {
		if x.isString {
			t.length--
		}
		x.isString = false
	} else {
		c := key[d]
		x.next[c] = t.delete(x.next[c], key, d+1)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.isString {
		return x
	}
	for c := 0; c < r; c++ {
		if x.next[c] != nil {
			return x
		}
	}
	return nil
}

// Keys returns all the keys in the set.
func (t *Trie) Keys() []string {
	return t.KeysWithPrefix("")
}

func (q *stringQueue) enqueue(x string) {
	*q = append(*q, x)
}

func (q *stringQueue) slice() []string {
	return *q
}
