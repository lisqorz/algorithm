package algorithm

import (
	"fmt"
	"testing"
)

func TestNewLruCache(t *testing.T) {
	cache := NewLruCache(5)
	cache.Add(3, 123)
	cache.Add("3", 333)
	cache.Add("23453", 12312)
	cache.Add("1233", 123)
	cache.Add("234afds3", "asdf")
	cache.Add("31234123ad", "asfg234")
	node := cache.end
	for {
		if nil == node {
			break
		}
		fmt.Println("last ", node.data)
		node = node.prev
	}
}
