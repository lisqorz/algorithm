package algorithm

import (
	"testing"
)

func TestNewSkipList(t *testing.T) {
	list := NewSkipList(32)
	for i := 0; i < 10000; i++ {
		list.Add(i, i+1)
	}
	ret := list.Find(1)
	t.Logf("%#v", ret)
}
