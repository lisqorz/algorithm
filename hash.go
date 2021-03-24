package algorithm

import "fmt"

type hashNode struct {
	hashcode int
	key      string
	data     interface{}
	next     *hashNode
}

type Hash struct {
	length int
	entry  [32]*hashNode
}

func NewHashTable() *Hash {
	return &Hash{}
}
func (o *hashNode) insert(node *hashNode) {
	for true {
		cur := o.next
		if cur == nil {
			o.next = node
			return
		}
		o = cur
	}
}
func (o *hashNode) find(key string) *hashNode {
	node := o
	for node != nil {
		if node.key == key {
			return node
		}
		node = node.next
	}
	return nil
}

func newHashNode(key string, value interface{}) *hashNode {
	return &hashNode{
		hashcode: hashCode(key),
		data:     value,
		key:      key,
		next:     nil,
	}
}
func (h *Hash) Put(key string, value interface{}) {
	node := newHashNode(key, value)
	fmt.Println("node", node.hashcode)
	if h.entry[node.hashcode] == nil {
		h.entry[node.hashcode] = newHashNode(key, value)
	} else {
		h.entry[node.hashcode].insert(node)
	}
}

func (h *Hash) Get(key string) *hashNode {
	code := hashCode(key)
	return h.entry[code].find(key)

}

func (h *Hash) Del(key string) {

}
func hash(key interface{}) int {
	if key == nil {
		return 0
	}
	c := hashCode(key)
	return c ^ (c >> 16)
}

func hashCode(value interface{}) int {
	a := value.(string)
	var b int
	for c := range a {
		b = 31*b + c
	}
	return b % 7
}
