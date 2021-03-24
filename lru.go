package algorithm

type lruNode struct {
	key  interface{}
	data interface{}
	next *lruNode
	prev *lruNode
}
type LruCache struct {
	limit int
	head  *lruNode
	end   *lruNode
	value map[interface{}]*lruNode
}

func (l *LruCache) Get(key interface{}) (interface{},bool) {
	a,b := l.value[key]
	return a,b
}

func (l *LruCache) Add(key interface{}, value interface{}) {
	if len(l.value) >= l.limit {
		oldKey := l.removeNode(l.head)
		delete(l.value, oldKey)
	}
	node := &lruNode{data: value, key: key}
	if l.end != nil {
		l.end.next = node
		node.prev = l.end
		node.next = nil
	}
	l.end = node
	if l.head == nil {
		l.head = node
	}
	l.value[key] = node
}
func (l *LruCache) removeNode(node *lruNode) interface{} {
	if node == l.head && node == l.end {

	} else if node == l.end {
		l.end = l.end.prev
		l.end.next = nil
	} else if node == l.head {
		l.head = l.head.next
		l.head.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	return node.key
}

func NewLruCache(cap int) *LruCache {
	return &LruCache{limit: cap, value: make(map[interface{}]*lruNode)}
}
