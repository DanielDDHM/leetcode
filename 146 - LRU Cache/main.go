package main

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{key: -1, value: -1}
	tail := &Node{key: -1, value: -1}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (lru *LRUCache) Get(key int) int {
	node, exists := lru.cache[key]
	if !exists {
		return -1
	}

	lru.moveToFront(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	if node, exists := lru.cache[key]; exists {
		node.value = value
		lru.moveToFront(node)
		return
	}

	newNode := &Node{key: key, value: value}
	lru.addToFront(newNode)
	lru.cache[key] = newNode

	if len(lru.cache) > lru.capacity {
		lru.removeLRU()
	}
}

func (lru *LRUCache) moveToFront(node *Node) {
	lru.removeNode(node)
	lru.addToFront(node)
}

func (lru *LRUCache) addToFront(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) removeLRU() {
	lruNode := lru.tail.prev
	lru.removeNode(lruNode)
	delete(lru.cache, lruNode.key)
}
