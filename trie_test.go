package trie

import "testing"

func TestStringQueue(t *testing.T) {
	q := new(stringQueue)
	q.enqueue("foo")
	q.enqueue("bar")
	q.enqueue("foobar")
	expected := []string{"foo", "bar", "foobar"}
	slice := q.stringSlice()
	for i, s := range expected {
		if slice[i] != s {
			t.Errorf("expected '%v', but got '%v'", s, slice[i])
		}
	}
	if len(slice) != 3 {
		t.Errorf("expected slice len 3, but got %d", len(slice))
	}
}
