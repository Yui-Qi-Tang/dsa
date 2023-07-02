package b75

/*
146. LRU Cache
Medium
16.9K
755
Companies
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.



Example 1:

Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4


Constraints:

1 <= capacity <= 3000
0 <= key <= 104
0 <= value <= 105
At most 2 * 105 calls will be made to get and put.
Accepted
1.3M
Submissions
3.2M
Acceptance Rate
40.7%
*/

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type LRUCache struct {
	cache   map[int]*Node
	maxSize int
	l, r    *Node // l: last recently used, r: most recently used
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		maxSize: capacity,
		cache:   make(map[int]*Node, capacity),
		l:       &Node{},
		r:       &Node{},
	}

	c.l.Next, c.r.Prev = c.r, c.l
	return c
}

func (this *LRUCache) empty() bool {
	return this.r.Prev == this.l && this.l.Next == this.r
}

func (this *LRUCache) evict() int {
	if this.empty() {
		return -1
	}
	evictKey := this.l.Next.Key
	this.l.Next.Next.Prev = this.l
	this.l.Next = this.l.Next.Next
	return evictKey
}

func (this *LRUCache) insert(n *Node) {
	n.Prev, n.Next = this.r.Prev, this.r
	this.r.Prev.Next, this.r.Prev = n, n
}

func (this *LRUCache) remove(n *Node) {
	if this.empty() {
		return
	}

	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func (this *LRUCache) Get(key int) int {
	n := this.cache[key]
	if n == nil {
		return -1
	}

	if n.Next == this.l {
		return n.Val
	}

	this.remove(n)
	this.insert(n)
	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
	v := this.cache[key]
	if v != nil {
		v.Val = value
		this.remove(v)
		this.insert(v)
		return
	}

	n := &Node{Key: key, Val: value}
	this.insert(n)
	this.cache[key] = n
	if len(this.cache) > this.maxSize {
		delete(this.cache, this.evict())
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
