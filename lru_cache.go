package ctci

type LRUNode[K comparable] struct {
	next, prev *LRUNode[K]
	key        K
	value      any
}

type LRUCache[K comparable] struct {
	hash       map[K]*LRUNode[K]
	head, tail *LRUNode[K]
	capacity   int
	items      int
}

func NewLRUCache[K comparable](capacity int) *LRUCache[K] {
	return &LRUCache[K]{
		hash:     make(map[K]*LRUNode[K], capacity),
		capacity: capacity,
	}
}

func (cache *LRUCache[K]) detach(node *LRUNode[K]) {
	if node == nil {
		return
	}

	if cache.head == node {
		cache.head = node.next
	}
	if cache.tail == node {
		cache.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}

func (cache *LRUCache[K]) setHead(node *LRUNode[K]) {
	if node == nil {
		return
	}

	node.next = cache.head
	if cache.head != nil {
		cache.head.prev = node
	}
	cache.head = node
	if cache.tail == nil {
		cache.tail = node
	}
}

func (cache *LRUCache[K]) Get(key K) (any, bool) {
	if node, ok := cache.hash[key]; !ok {
		return nil, false
	} else {
		cache.detach(node)
		cache.setHead(node)
		return node.value, true
	}
}

func (cache *LRUCache[K]) Set(key K, value any) {
	if node, ok := cache.hash[key]; ok {
		cache.detach(node)
		node.value = value
		cache.setHead(node)
	} else {
		if cache.items == cache.capacity {
			delete(cache.hash, cache.tail.key)
			cache.detach(cache.tail)
			cache.items--
		}
		node = &LRUNode[K]{key: key, value: value}
		cache.setHead(node)
		cache.items++
		cache.hash[key] = node
	}
}
