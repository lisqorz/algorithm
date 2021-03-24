package algorithm

import (
	"testing"
)

func TestNewHashTable(t *testing.T) {
	hash :=NewHashTable()
	hash.Put("a", "a")
	hash.Put("ab", "ab")
	hash.Put("abc", "abc")
	hash.Put("abcd", "abcd")
	hash.Put("abcde", "abcde")
	hash.Put("abcdef", "abcdef")
	hash.Put("abcdefg", "abcdefg")
	hash.Put("abcdefgh", "abcdefgh")
	t.Logf("%#v",hash.Get("a"))
	t.Logf("%#v",hash.Get("ab"))
	t.Logf("%#v",hash.Get("abc"))
	t.Logf("%#v",hash.Get("abcd"))
	t.Logf("%#v",hash.Get("abcde"))
	t.Logf("%#v",hash.Get("abcdef"))
	t.Logf("%#v",hash.Get("abcdefg"))
	t.Logf("%#v",hash.Get("abcdefgh"))
}
