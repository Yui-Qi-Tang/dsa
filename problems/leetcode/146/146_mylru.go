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

func constructorv30(size int) *lruv30 {
	lru := &lruv30{
		maxsize: size,
		cache:   make(map[int]*nodev30),
		left:    new(nodev30),
		right:   new(nodev30),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev30 struct {
	key, value int
	prev, next *nodev30
}

type lruv30 struct {
	cache       map[int]*nodev30
	maxsize     int
	left, right *nodev30
}

func (lru *lruv30) remove(n *nodev30) {
	n.next.prev = n.prev
	n.prev.next = n.next
}
func (lru *lruv30) insert(n *nodev30) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}
func (lru *lruv30) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}
func (lru *lruv30) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev30{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv29(size int) *lruv29 {

	lru := &lruv29{
		maxSize: size,
		cache:   make(map[int]*nodev29),
		left:    new(nodev29),
		right:   new(nodev29),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

type lruv29 struct {
	cache       map[int]*nodev29
	maxSize     int
	left, right *nodev29
}

type nodev29 struct {
	key, value int
	prev, next *nodev29
}

func (lru *lruv29) remove(n *nodev29) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv29) insert(n *nodev29) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv29) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv29) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev29{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv28(size int) *lruv28 {
	lru := &lruv28{
		cache:   make(map[int]*nodev28),
		left:    new(nodev28),
		right:   new(nodev28),
		maxsize: size,
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev28 struct {
	key, value int
	prev, next *nodev28
}

type lruv28 struct {
	cache       map[int]*nodev28
	maxsize     int
	left, right *nodev28
}

func (lru *lruv28) remove(n *nodev28) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv28) insert(n *nodev28) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv28) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv28) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev28{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv27(size int) *lruv27 {
	lru := &lruv27{
		maxSize: size,
		cache:   make(map[int]*nodev27),
		left:    new(nodev27),
		right:   new(nodev27),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev27 struct {
	key, value int
	prev, next *nodev27
}

type lruv27 struct {
	cache       map[int]*nodev27
	maxSize     int
	left, right *nodev27
}

func (lru *lruv27) remove(n *nodev27) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv27) insert(n *nodev27) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv27) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv27) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev27{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn
	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv26(size int) *lruv26 {
	lru := &lruv26{
		maxSize: size,
		cache:   make(map[int]*nodev26),
		left:    new(nodev26),
		right:   new(nodev26),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type lruv26 struct {
	cache       map[int]*nodev26
	maxSize     int
	left, right *nodev26
}

type nodev26 struct {
	key, value int
	prev, next *nodev26
}

func (lru *lruv26) remove(n *nodev26) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv26) insert(n *nodev26) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv26) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv26) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev26{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv25(size int) *lruv25 {
	lru := &lruv25{
		maxSize: size,
		cache:   make(map[int]*nodev25),
		left:    new(nodev25),
		right:   new(nodev25),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev25 struct {
	key, value int
	prev, next *nodev25
}

type lruv25 struct {
	cache       map[int]*nodev25
	maxSize     int
	left, right *nodev25
}

func (lru *lruv25) remove(n *nodev25) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv25) insert(n *nodev25) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv25) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv25) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev25{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv24(size int) *lruv24 {

	lru := &lruv24{
		maxSize: size,
		cache:   make(map[int]*nodev24),
		left:    new(nodev24),
		right:   new(nodev24),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev24 struct {
	key, value int
	prev, next *nodev24
}

type lruv24 struct {
	maxSize     int
	cache       map[int]*nodev24
	left, right *nodev24
}

func (lru *lruv24) remove(n *nodev24) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv24) insert(n *nodev24) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv24) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv24) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev24{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn
	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv23(size int) *lruv23 {
	lru := &lruv23{
		maxSize: size,
		cache:   make(map[int]*nodev23),
		right:   new(nodev23),
		left:    new(nodev23),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev23 struct {
	key, value int
	prev, next *nodev23
}

type lruv23 struct {
	maxSize     int
	cache       map[int]*nodev23
	left, right *nodev23
}

func (lru *lruv23) remove(n *nodev23) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv23) insert(n *nodev23) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv23) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv23) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev23{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv22(size int) *lruv22 {
	lru := &lruv22{
		maxSize: size,
		cache:   make(map[int]*nodev22),
		left:    new(nodev22),
		right:   new(nodev22),
	}
	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

type nodev22 struct {
	key, value int
	prev, next *nodev22
}

type lruv22 struct {
	maxSize     int
	cache       map[int]*nodev22
	left, right *nodev22
}

func (lru *lruv22) remove(n *nodev22) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv22) insert(n *nodev22) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv22) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv22) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev22{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv21(size int) *lruv21 {
	lru := &lruv21{
		maxSize: size,
		cache:   map[int]*nodev21{},
		left:    new(nodev21),
		right:   new(nodev21),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev21 struct {
	key, value int
	prev, next *nodev21
}

type lruv21 struct {
	maxSize     int
	cache       map[int]*nodev21
	left, right *nodev21
}

func (lru *lruv21) remove(n *nodev21) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv21) insert(n *nodev21) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv21) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv21) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev21{key: key, value: value}
	lru.cache[key] = nn
	lru.insert(nn)
	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv20(size int) *lruv20 {
	lru := &lruv20{
		maxSize: size,
		cache:   make(map[int]*nodev20),
		left:    new(nodev20),
		right:   new(nodev20),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev20 struct {
	key, value int
	prev, next *nodev20
}

type lruv20 struct {
	maxSize     int
	cache       map[int]*nodev20
	left, right *nodev20
}

func (lru *lruv20) remove(n *nodev20) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv20) insert(n *nodev20) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv20) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv20) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev20{key: key, value: value}
	lru.cache[key] = nn
	lru.insert(nn)

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv19(size int) *lruv19 {
	lru := &lruv19{
		maxSize: size,
		cache:   map[int]*nodev19{},
		left:    &nodev19{},
		right:   &nodev19{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev19 struct {
	key, value int
	prev, next *nodev19
}

type lruv19 struct {
	maxSize     int
	cache       map[int]*nodev19
	left, right *nodev19
}

func (lru *lruv19) remove(n *nodev19) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv19) insert(n *nodev19) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv19) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv19) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev19{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv18(size int) *lruv18 {
	lru := &lruv18{
		cache:   make(map[int]*nodev18),
		left:    new(nodev18),
		right:   new(nodev18),
		maxsize: size,
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev18 struct {
	key, value int
	prev, next *nodev18
}

type lruv18 struct {
	cache       map[int]*nodev18
	left, right *nodev18
	maxsize     int
}

func (lru *lruv18) remove(n *nodev18) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv18) insert(n *nodev18) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv18) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}
func (lru *lruv18) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev18{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv17(size int) *lruv17 {
	lru := &lruv17{
		maxsize: size,
		cache:   map[int]*nodev17{},
		left:    &nodev17{},
		right:   &nodev17{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev17 struct {
	key, value int
	prev, next *nodev17
}

type lruv17 struct {
	maxsize     int
	cache       map[int]*nodev17
	left, right *nodev17
}

func (lru *lruv17) remove(n *nodev17) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv17) insert(n *nodev17) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv17) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}
func (lru *lruv17) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev17{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	for len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv16(size int) *lruv16 {
	lru := &lruv16{
		maxsize: size,
		cache:   map[int]*nodev16{},
		left:    new(nodev16),
		right:   new(nodev16),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev16 struct {
	key, value int
	prev, next *nodev16
}

type lruv16 struct {
	maxsize     int
	cache       map[int]*nodev16
	left, right *nodev16
}

func (lru *lruv16) remove(n *nodev16) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv16) insert(n *nodev16) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv16) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv16) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev16{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn
	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv15(size int) *lruv15 {
	lru := &lruv15{
		maxsize: size,
		left:    new(nodev15),
		right:   new(nodev15),
		cache:   make(map[int]*nodev15),
	}

	lru.right.prev, lru.left.next = lru.left, lru.right

	return lru
}

type nodev15 struct {
	key, value int
	prev, next *nodev15
}

type lruv15 struct {
	maxsize     int
	left, right *nodev15
	cache       map[int]*nodev15
}

func (lru *lruv15) remove(n *nodev15) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv15) insert(n *nodev15) {
	n.prev = lru.right.prev
	n.next = lru.right

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv15) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv15) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		delete(lru.cache, n.key)
	}

	nn := &nodev15{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn
	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv14(size int) *lruv14 {
	lru := &lruv14{
		maxsize: size,
		cache:   map[int]*nodev14{},
		left:    &nodev14{},
		right:   &nodev14{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev14 struct {
	key, value int
	prev, next *nodev14
}

type lruv14 struct {
	maxsize     int
	cache       map[int]*nodev14
	left, right *nodev14
}

func (lru *lruv14) remove(n *nodev14) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv14) insert(n *nodev14) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv14) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv14) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev14{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv13(size int) *lruv13 {
	lru := &lruv13{
		maxsize: size,
		left:    &nodev13{},
		right:   &nodev13{},
		cache:   make(map[int]*nodev13),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type lruv13 struct {
	cache       map[int]*nodev13
	left, right *nodev13
	maxsize     int
}

type nodev13 struct {
	key, value int
	prev, next *nodev13
}

func (lru *lruv13) remove(n *nodev13) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv13) insert(n *nodev13) {
	n.prev = lru.right.prev
	n.next = lru.right

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv13) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruv13) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev13{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv12(size int) *lruv12 {
	lru := &lruv12{
		maxSize: size,
		cache:   make(map[int]*nodev12),
		left:    new(nodev12),
		right:   new(nodev12),
	}

	lru.left.next, lru.right.prev = lru.right, lru.left
	return lru
}

type nodev12 struct {
	key, value int
	prev, next *nodev12
}

type lruv12 struct {
	cache       map[int]*nodev12 // key: node
	left, right *nodev12
	maxSize     int
}

func (lru *lruv12) remove(n *nodev12) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (lru *lruv12) insert(n *nodev12) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv12) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv12) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev12{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn
	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv11(size int) *lruv11 {
	lru := &lruv11{
		maxSize: size,
		cache:   make(map[int]*nodev11),
		left:    new(nodev11),
		right:   new(nodev11),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev11 struct {
	key, value int
	prev, next *nodev11
}

type lruv11 struct {
	maxSize     int
	cache       map[int]*nodev11
	left, right *nodev11
}

func (lru *lruv11) remove(n *nodev11) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv11) insert(n *nodev11) {
	n.prev = lru.right.prev
	n.next = lru.right

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv11) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv11) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev11{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv10(size int) *lruv10 {
	lru := &lruv10{
		maxSize: size,
		cache:   make(map[int]*nodev10),
		left:    &nodev10{},
		right:   &nodev10{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev10 struct {
	key, value int
	prev, next *nodev10
}

type lruv10 struct {
	maxSize     int
	cache       map[int]*nodev10
	left, right *nodev10
}

func (lru *lruv10) remove(n *nodev10) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv10) insert(n *nodev10) {
	n.prev = lru.right.prev
	n.next = lru.right

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv10) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv10) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev10{key: key, value: value}
	lru.cache[key] = nn
	lru.insert(nn)

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv9(size int) *lruv9 {
	lru := &lruv9{
		maxsize: size,
		left:    &nodev9{},
		right:   &nodev9{},
		cache:   make(map[int]*nodev9),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev9 struct {
	key, value int
	prev, next *nodev9
}

type lruv9 struct {
	maxsize     int
	left, right *nodev9
	cache       map[int]*nodev9
}

func (lru *lruv9) remove(n *nodev9) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv9) insert(n *nodev9) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv9) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv9) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev9{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv8(size int) *lruv8 {
	lru := &lruv8{
		maxsize: size,
		left:    &nodev8{},
		right:   &nodev8{},
		cache:   make(map[int]*nodev8),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev8 struct {
	key, value int
	prev, next *nodev8
}

type lruv8 struct {
	// size
	maxsize int
	// limitation of left, right; lru => left.Next, most recently used right.prev
	left, right *nodev8
	// cache for get: int:node
	cache map[int]*nodev8
}

func (lru *lruv8) remove(n *nodev8) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv8) insert(n *nodev8) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv8) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv8) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev8{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv7(size int) *lruv7 {
	lru := &lruv7{
		maxsize: size,
		left:    &nodev7{},
		right:   &nodev7{},
		cache:   make(map[int]*nodev7),
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left

	return lru
}

type nodev7 struct {
	key, value int
	prev, next *nodev7
}

type lruv7 struct {
	maxsize     int
	cache       map[int]*nodev7
	left, right *nodev7
}

func (lru *lruv7) remove(n *nodev7) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv7) insert(n *nodev7) {
	n.next = lru.right
	n.prev = lru.right.prev

	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv7) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}

func (lru *lruv7) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev7{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn

	if len(lru.cache) > lru.maxsize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv6(size int) *lruv6 {
	lru := &lruv6{
		maxSize: size,
		cache:   make(map[int]*nodev6),
		left:    &nodev6{},
		right:   &nodev6{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

type nodev6 struct {
	key, value int
	prev, next *nodev6
}

type lruv6 struct {
	maxSize     int
	cache       map[int]*nodev6
	left, right *nodev6 // limitation of the lru list
}

func (lru *lruv6) remove(n *nodev6) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv6) insert(n *nodev6) {
	n.prev = lru.right.prev
	n.next = lru.right
	lru.right.prev.next = n
	lru.right.prev = n
}

func (lru *lruv6) Get(key int) int {
	if v := lru.cache[key]; v != nil {
		lru.remove(v)
		lru.insert(v)
		return v.value
	}
	return -1
}

func (lru *lruv6) Put(key, value int) {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	nn := &nodev6{key: key, value: value}
	lru.insert(nn)
	lru.cache[key] = nn
	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv5(size int) *lruv5 {
	lru := &lruv5{
		maxSize: size,
		left:    &nodev5{},
		right:   &nodev5{},
		cache:   make(map[int]*nodev5),
	}

	// init the list
	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

type nodev5 struct {
	key, value int
	prev, next *nodev5
}

type lruv5 struct {
	// maxsize
	maxSize int
	// limitation for the linked list
	left, right *nodev5
	// cache: key: node
	cache map[int]*nodev5
}

func (l *lruv5) insert(n *nodev5) {
	n.prev = l.right.prev
	n.next = l.right
	l.right.prev.next = n
	l.right.prev = n
}

func (l *lruv5) remove(n *nodev5) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *lruv5) Get(key int) int {
	if n := l.cache[key]; n != nil {
		l.remove(n)
		l.insert(n)
		return n.value
	}
	return -1
}

func (l *lruv5) Put(key, value int) {
	if n := l.cache[key]; n != nil {
		l.remove(n)
	}

	nn := &nodev5{key: key, value: value} // nn: newNode
	l.insert(nn)
	l.cache[key] = nn

	if len(l.cache) > l.maxSize {
		delete(l.cache, l.left.next.key)
		l.remove(l.left.next)
	}
}

func constructorv4(size int) *lruv4 {
	lru := &lruv4{
		maxSize: size,
		cache:   make(map[int]*nodev4),
		left:    &nodev4{},
		right:   &nodev4{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

type nodev4 struct {
	key, val   int
	prev, next *nodev4
}

type lruv4 struct {
	cache       map[int]*nodev4 // key: node
	maxSize     int
	left, right *nodev4 // left.Next -> last used; right.prev -> most recently used
}

func (v4 *lruv4) insert(n *nodev4) {
	n.prev = v4.right.prev
	n.next = v4.right
	v4.right.prev.next = n
	v4.right.prev = n
}

func (v4 *lruv4) remove(n *nodev4) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (v4 *lruv4) Get(key int) int {
	if n := v4.cache[key]; n != nil {
		v4.remove(n)
		v4.insert(n)
		return n.val
	}

	return -1
}

func (v4 *lruv4) Put(key, value int) {
	if n := v4.cache[key]; n != nil {
		v4.remove(n)
	}
	newNode := &nodev4{key: key, val: value}
	v4.insert(newNode)
	v4.cache[key] = newNode

	if len(v4.cache) > v4.maxSize {
		delete(v4.cache, v4.left.next.key)
		v4.remove(v4.left.next)
	}
}

func constructorv3(size int) *lruv3 {
	v3 := &lruv3{
		maxSize: size,
		cache:   make(map[int]*nodev3),
		left:    &nodev3{},
		right:   &nodev3{},
	}

	v3.left.next = v3.right
	v3.right.prev = v3.left
	return v3
}

type nodev3 struct {
	key, value int
	prev, next *nodev3
}

type lruv3 struct {
	cache       map[int]*nodev3
	maxSize     int
	left, right *nodev3
}

func (lru *lruv3) remove(n *nodev3) {

	if lru.right.prev == lru.left && lru.left.next == lru.right {
		return
	}

	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruv3) insert(n *nodev3) {
	n.prev = lru.right.prev
	lru.right.prev.next = n
	n.next = lru.right
	lru.right.prev = n
}

func (lru *lruv3) Get(key int) int {
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}
	return -1
}
func (lru *lruv3) Put(key, value int) {
	if v := lru.cache[key]; v != nil {
		lru.remove(v)
	}

	n := &nodev3{key: key, value: value}
	lru.insert(n)
	lru.cache[key] = n

	if len(lru.cache) > lru.maxSize {
		delete(lru.cache, lru.left.next.key)
		lru.remove(lru.left.next)
	}
}

func constructorv2(size int) *lruv2 {
	cache := &lruv2{
		maxSize: size,
		cache:   make(map[int]*nodev2),
		l:       &nodev2{},
		r:       &nodev2{},
	}

	cache.l.next = cache.r
	cache.r.prev = cache.l
	return cache
}

type nodev2 struct {
	key, value int
	next, prev *nodev2
}

type lruv2 struct {
	maxSize int
	cache   map[int]*nodev2
	l, r    *nodev2 // l.next: last used, r.prev: recently used
}

func (l *lruv2) remove(n *nodev2) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *lruv2) insert(n *nodev2) {
	n.next = l.r
	n.prev = l.r.prev
	l.r.prev.next = n
	l.r.prev = n
}

func (l *lruv2) Get(key int) int {
	if v := l.cache[key]; v != nil {
		l.remove(v)
		l.insert(v)
		return v.value
	}
	return -1
}

func (l *lruv2) Put(key, value int) {

	if v := l.cache[key]; v != nil {
		l.remove(v)
	}

	n := &nodev2{key: key, value: value}
	l.insert(n)
	l.cache[key] = n
	if len(l.cache) > l.maxSize {
		delete(l.cache, l.l.next.key)
		l.remove(l.l.next)
	}
}

func constructorv1(size int) *lruCachev1 {
	cache := &lruCachev1{
		size:  size,
		cache: make(map[int]*nodev1),
		l:     &nodev1{},
		r:     &nodev1{},
	}

	cache.l.next = cache.r
	cache.r.prev = cache.l
	return cache
}

type nodev1 struct {
	key, value int
	prev, next *nodev1
}

type lruCachev1 struct {
	cache map[int]*nodev1 // key: node
	l, r  *nodev1         // l: last used; r: recently used
	size  int             // buffer size
}

func (lru *lruCachev1) empty() bool {
	return lru.l.next == lru.r && lru.r.prev == lru.l
}

func (lru *lruCachev1) remove(n *nodev1) {

	if lru.empty() {
		return
	}

	n.prev.next = n.next
	n.next.prev = n.prev
}

func (lru *lruCachev1) insert(n *nodev1) {
	n.next = lru.r
	n.prev = lru.r.prev
	lru.r.prev.next = n
	lru.r.prev = n
	lru.cache[n.key] = n
	if len(lru.cache) > lru.size {
		delete(lru.cache, lru.l.next.key)
		lru.l.next.next.prev = lru.l
		lru.l.next = lru.l.next.next
	}
}

func (lru *lruCachev1) Get(key int) int {

	if n := lru.cache[key]; n != nil {
		lru.remove(n)
		lru.insert(n)
		return n.value
	}

	return -1
}

func (lru *lruCachev1) Put(key, value int) {

	n := &nodev1{key: key, value: value}
	if exist := lru.cache[key]; exist != nil {
		lru.remove(exist)
		lru.insert(n)
		return
	}
	lru.insert(n)
}
