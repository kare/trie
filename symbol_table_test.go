package trie_test

import (
	"testing"

	"kkn.fi/trie"
)

var data = []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}

func TestSymbolTableLen(t *testing.T) {
	st := trie.NewSymbolTable()
	for i, w := range data {
		st.Put(w, i)
	}
	if st.Len() != 7 {
		t.Errorf("expected len 7, but got %d", st.Len())
	}
}

func TestSymbolTableIsEmpty(t *testing.T) {
	st := trie.NewSymbolTable()
	if !st.IsEmpty() {
		t.Errorf("expected empty symbol table, but was not")
	}
	for i, w := range data {
		st.Put(w, i)
	}
	if st.IsEmpty() {
		t.Error("expected non empty symbol table")
	}
}

func TestSymbolTableKeys(t *testing.T) {
	st := trie.NewSymbolTable()
	for i, w := range data {
		st.Put(w, i)
	}
	keys := st.Keys()
	if len(keys) != len(data)-1 {
		t.Errorf("expected keys to return %d elements, but got %d", len(data)-1, len(keys))
	}
}

func TestSymbolTableDelete(t *testing.T) {
	st := trie.NewSymbolTable()
	st.Delete("null")
	for i, w := range data {
		st.Put(w, i)
	}
	st.Delete("she")
	if st.Len() != len(data)-2 {
		t.Error("expected delete to remove element")
	}
}

func TestSymbolTableGet(t *testing.T) {
	st := trie.NewSymbolTable()
	for i, w := range data {
		st.Put(w, i)
	}
	testData := []struct {
		key   string
		value interface{}
	}{
		{"by", 4},
		{"sea", 6},
		{"sells", 1},
		{"she", 0},
		{"shells", 3},
		{"shore", 7},
		{"the", 5},
		{"null", nil},
	}
	for _, td := range testData {
		if st.Get(td.key) != td.value {
			t.Errorf("expected key '%v' to return %d, but got %d", td.key, td.value, st.Get(td.key))
		}
	}
}

func TestSymbolTableDisallowsEmptyKey(t *testing.T) {
	st := trie.NewSymbolTable()
	st.Put("", "value")
	if st.Get("") != nil {
		t.Error("expected st to contain key '' with value <nil>")
	}
}

func TestSymbolTablePutDoesntAcceptNilValue(t *testing.T) {
	st := trie.NewSymbolTable()
	st.Put("key", "value")
	if !st.Contains("key") {
		t.Error("expected st to contain key 'key'")
	}
	st.Put("key", nil)
	if st.Contains("key") {
		t.Error("expected put nil value to remove key")
	}
}

func TestSymbolTableKeysThatMatch(t *testing.T) {
	st := trie.NewSymbolTable()
	for i, w := range data {
		st.Put(w, i)
	}

	results := st.KeysThatMatch(".he.l.")
	if len(results) != 1 {
		t.Errorf("expected 1, but got %d results", len(results))
	}
	expected := "shells"
	if results[0] != expected {
		t.Errorf("expected '%v' but got '%v'", expected, results[0])
	}

}

func TestSymbolTableLongestPrefix(t *testing.T) {
	st := trie.NewSymbolTable()
	for i, w := range data {
		st.Put(w, i)
	}

	prefix := st.LongestPrefixOf("shellsort")
	expected := "shells"
	if prefix != expected {
		t.Errorf("expected '%v' but got '%v'", expected, prefix)
	}

	prefix = st.LongestPrefixOf("kare")
	if prefix != "" {
		t.Errorf("expected '' but got '%v'", prefix)
	}
}

func TestSymbolTableKeysWithPrefix(t *testing.T) {
	st := trie.NewSymbolTable()
	for i, w := range data {
		st.Put(w, i)
	}

	expected := "shore"
	result := st.KeysWithPrefix("shor")
	if len(result) != 1 {
		t.Errorf("expected 1, but got %d results", len(result))
	}
	if result[0] != expected {
		t.Errorf("expected '%v', but got '%v'", expected, result[0])
	}
	result = st.KeysWithPrefix("kare")
	if len(result) != 0 {
		t.Errorf("expected 0, but got %d results", len(result))
	}
}
