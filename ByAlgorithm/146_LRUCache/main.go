package main

// 定义节点
type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

// 定义 lru 缓存类型
type LRUCache struct {
	size       int                  // 当前存储多少
	capacity   int                  // 容量多少
	cache      map[int]*DLinkedNode // 存储结构
	head, tail *DLinkedNode         // 头尾指针
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key: 0, value: 0}
}

// 初始化缓存，容量，头尾指针相连
func Construckor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail // 头结点连着尾结点
	l.tail.next = l.head // 尾结点连接这头结点
	return l
}

// 获取元素值
func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

// 添加元素
func (l *LRUCache) Put(key, value int) {
	// 元素不存在，判断长度是否大于容量，大于删除
	if _, ok := l.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			l.removeTail()
			l.size--
		}
	} else {
		// 元素存在，更新结点的值，将节点移到头部
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}
func (l *LRUCache) addToHead(node *DLinkedNode) {
	l.head.next = node
	node.pre = l.head
	l.head.next.pre = node
	node.next = l.head.next
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.pre
	l.removeNode(node)
	return node
}
