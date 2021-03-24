package algorithm

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type skipNode struct {
	key int
	data interface{}
	next [] *skipNode
}


type SkipList struct {
	head   *skipNode
	tail   *skipNode
	length int
	level  int
	mut    *sync.RWMutex
	rand   *rand.Rand
}

func NewSkipList(level int) *SkipList {
	list := &SkipList{}
	if level <= 0 {
		level = 32
	}
	list.level = level
	list.head = &skipNode{next: make([]*skipNode, level, level)}
	list.tail = &skipNode{}
	list.mut = &sync.RWMutex{}
	list.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := range list.head.next {
		list.head.next[index] = list.tail
	}
	return list
}
func (list *SkipList) Add(key int, data interface{}) {

	list.mut.Lock()

	defer list.mut.Unlock()

	//1.确定插入深度

	level := list.randomLevel()
	fmt.Printf("insert level:%d\n", level)

	//2.查找插入部位

	update := make([]*skipNode, level, level)

	node := list.head

	for index := level - 1; index >= 0; index-- {
		for {

			node1 := node.next[index]

			if node1 == list.tail || node1.key > key {

				update[index] = node //找到一个插入部位

				break

			} else if node1.key == key {

				node1.data = data
				return
			} else {

				node = node1

			}

		}

	}

	//3.执行插入

	newNode := &skipNode{key, data, make([]*skipNode, level, level)}

	for index, node := range update {

		node.next[index], newNode.next[index] = newNode, node.next[index]

	}
	list.length++
}

func (list *SkipList) randomLevel() int {

	level := 1

	for ; level < list.level && list.rand.Uint32()&0x1 == 1; level++ {
	}
	return level
}

func (list *SkipList) Find(key int) interface{} {

	list.mut.RLock()

	defer list.mut.RUnlock()

	node := list.head

	for index := len(node.next) - 1; index >= 0; index-- {

		for {

			node1 := node.next[index]

			if node1 == list.tail || node1.key > key {

				break

			} else if node1.key == key {

				return node1.data

			} else {

				node = node1

			}

		}

	}

	return nil
}
