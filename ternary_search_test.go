package trie_test

import (
	"testing"

	"kkn.fi/trie"
)

func TestTernarySearchGet(t *testing.T) {
	ts := trie.NewTernarySearch()
	for i, w := range data {
		ts.Put(w, i)
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
		if ts.Get(td.key) != td.value {
			t.Errorf("expected key '%v' to return %d, but got %d", td.key, td.value, ts.Get(td.key))
		}
	}
}

func TestTernarySearchIsEmpty(t *testing.T) {
	ts := trie.NewTernarySearch()
	if !ts.IsEmpty() {
		t.Error("expected empty trie")
	}
	for i, w := range data {
		ts.Put(w, i)
	}
	if ts.IsEmpty() {
		t.Error("expected non empty trie")
	}
}

func TestTernarySearchLen(t *testing.T) {
	ts := trie.NewTernarySearch()
	for i, w := range data {
		ts.Put(w, i)
	}
	if ts.Len() != 7 {
		t.Errorf("expected trie len 7, but got %d", ts.Len())
	}
}

func TestTernarySearch(t *testing.T) {
	ts := trie.NewTernarySearch()
	for i, w := range data {
		ts.Put(w, i)
	}
	if ts.Len() != 7 {
		t.Errorf("expected trie len 7, but got %d", ts.Len())
	}
}

func TestTernarySearchDisallowsEmptyKey(t *testing.T) {
	ts := trie.NewTernarySearch()
	if ts.Get("") != nil {
		t.Error("expected trie to contain key '' with value <nil>")
	}
	ts.Put("", "value")
	if ts.Get("") != nil {
		t.Error("expected trie to contain key '' with value <nil>")
	}
}
