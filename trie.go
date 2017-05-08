package trie // import "kkn.fi/trie"

const r = 256 // extended ascii

type (
	// Interface defines common functions for all trie types.
	Interface interface {
		Put(key string, value interface{})
		Get(key string) interface{}
		Delete(key string)
		Contains(key string) bool
		IsEmpty() bool
		LongestPrefixOf(query string) string
		KeysWithPrefix(prefix string) []string
		KeysThatMatch(pattern string) []string
		Keys() []string
		Len() int
	}
	stringQueue []string
)

func (q *stringQueue) enqueue(x string) {
	*q = append(*q, x)
}

func (q *stringQueue) slice() []string {
	r := make([]string, 0, len(*q))
	return append(r, []string(*q)...)
}
