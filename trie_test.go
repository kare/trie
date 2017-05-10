package trie_test

import (
	"testing"

	"kkn.fi/trie"
)

func TestTrieContains(t *testing.T) {
	st := trie.New()
	for _, w := range data {
		st.Add(w)
	}
	td := []struct {
		key    string
		exists bool
	}{
		{"by", true},
		{"sea", true},
		{"sells", true},
		{"she", true},
		{"shells", true},
		{"shore", true},
		{"the", true},
		{"null", false},
	}
	for _, test := range td {
		if test.exists != st.Contains(test.key) {
			t.Errorf("contains failed key: '%v'", test.key)
		}
	}
}

func TestTrieKeys(t *testing.T) {
	st := trie.New()
	for _, w := range data {
		st.Add(w)
	}
	keys := st.Keys()
	if len(keys) != len(data)-1 {
		t.Errorf("expected %d keys, but got %d\n", len(data)-1, len(keys))
	}
}

func TestTrieIsEmpty(t *testing.T) {
	st := trie.New()
	if !st.IsEmpty() {
		t.Errorf("expected set to be empty")
	}
	for _, w := range data {
		st.Add(w)
	}
	if st.IsEmpty() {
		t.Errorf("expected set to be empty")
	}
}

func TestTrieLen(t *testing.T) {
	st := trie.New()
	for _, w := range data {
		st.Add(w)
	}
	if st.Len() != len(data)-1 {
		t.Errorf("expected %d, but got %d", len(data)-1, st.Len())
	}
}

func TestTrieLongestPrefixOf(t *testing.T) {
	st := trie.New()
	for _, w := range data {
		st.Add(w)
	}
	expected := "shells"
	result := st.LongestPrefixOf("shellsort")
	if expected != result {
		t.Errorf("expected '%v', but got '%v'", expected, result)
	}

	expected = ""
	result = st.LongestPrefixOf("xshellsort")
	if expected != result {
		t.Errorf("expected '%v', but got '%v'", expected, result)
	}
	result = st.LongestPrefixOf("")
	if expected != result {
		t.Errorf("expected '%v', but got '%v'", expected, result)
	}
}

func TestTrieKeysWithPrefix(t *testing.T) {
	st := trie.New()
	for _, w := range data {
		st.Add(w)
	}
	expected := "shore"
	result := st.KeysWithPrefix("shor")
	if 1 != len(result) {
		t.Errorf("expected len 1, but got len %d", len(result))
	}
	if expected != result[0] {
		t.Errorf("expected '%v', but got '%v'", expected, result[0])
	}

	result = st.KeysWithPrefix("shortening")
	if 0 != len(result) {
		t.Errorf("expected 0, but got %d", len(result))
	}
}

func TestTrieKeysThatMatch(t *testing.T) {
	st := trie.New()
	for _, w := range data {
		st.Add(w)
	}

	expected := "shells"
	results := st.KeysThatMatch(".he.l.")
	if len(results) != 1 {
		t.Errorf("expected to have len 1 results, but got %d", len(results))
	}
	if expected != results[0] {
		t.Errorf("expected '%v', but got '%v'", expected, results[0])
	}
}

func TestTrieDelete(t *testing.T) {
	st := trie.New()
	st.Delete("she")
	for _, w := range data {
		st.Add(w)
	}

	st.Delete("she")
	if st.Len() != 6 {
		t.Errorf("expected len 6, but got %d", st.Len())
	}
}

var tmp string

func BenchmarkTrieLongestPrefixOf(b *testing.B) {
	st := trie.New()
	for _, w := range data {
		st.Add(w)
	}
	for n := 0; n < b.N; n++ {
		tmp = st.LongestPrefixOf("shellsort")
	}
}

var tmp2 []string

func BenchmarkTrieKeysWithPrefix(b *testing.B) {
	st := trie.New()
	for _, w := range data {
		st.Add(w)
	}
	for n := 0; n < b.N; n++ {
		tmp2 = st.KeysWithPrefix("shor")
	}
}
