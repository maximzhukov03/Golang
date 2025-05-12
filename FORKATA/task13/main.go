package main

import (
	"fmt"
	"hash/crc32"	
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type HashMap struct{
	hash func(string) uint32
	cache map[string]interface{}
}

type Node struct{
	key string
	value interface{}
	next *Node
}

type HashMapSlice struct{
	buckets [][]entryNode
	size int
	hash func(string) uint32
	cache map[string]interface{}
}

type entryNode struct{
	key string
	value interface{}
}

type HashMapList struct{
	buckets []*Node
	size int
	hash func(string) uint32
	cache map[string]interface{}
}

func NewHashMapSlice(count int, options ...func(*HashMap)) HashMaper {
	h := &HashMap{
		hash: func(s string) uint32{
			return crc32.ChecksumIEEE([]byte(s))
		},	
		cache: make(map[string]interface{}),
	}

	for _, opt := range options{
		opt(h)
	}

	return &HashMapSlice{
		buckets: make([][]entryNode, count),
		size: count,
		hash: h.hash,
		cache: h.cache,
	}
}

func NewHashMapList(count int, options ...func(*HashMap)) HashMaper {
	h := &HashMap{
		hash: func(s string) uint32{
			return crc32.ChecksumIEEE([]byte(s))
		},
		cache: make(map[string]interface{}),
	}

	for _, opt := range options{
		opt(h)
	}

    return &HashMapList{
		buckets: make([]*Node, count),
		size: count,
		hash: h.hash,
		cache: h.cache,
	}
}

func (h *HashMapList)  Set(key string, value interface{}){
	idBucket := h.hash(key) % uint32(h.size)
	node := h.buckets[idBucket]
	for node != nil{
		if node.key == key{
			node.value = value 
			h.cache[key] = value
			return
		}
		node = node.next
	}

	h.buckets[idBucket] = &Node{key: key, value: value, next: h.buckets[idBucket]}
	h.cache[key] = value
}

func (h *HashMapList)  Get(key string) (interface{}, bool){
	idBucket := h.hash(key) % uint32(h.size)
	valList, ok := h.cache[key]
	if ok{
		return valList, true
	}
	node := h.buckets[idBucket]
	for node != nil{
		if node.key == key{
			h.cache[key] = node.value
			return node.value, true
		}
		node = node.next
	}
	return nil, false
	
}

func (h *HashMapSlice)  Set(key string, value interface{}){
	idBucket := h.hash(key) % uint32(h.size)
	for index, elemSlice := range h.buckets[idBucket]{
		if elemSlice.key == key{
			h.buckets[idBucket][index].value = value 
			h.cache[key] = value
		}
	}
	h.buckets[idBucket] = append(h.buckets[idBucket], entryNode{key: key, value: value})
	h.cache[key] = value
}

func (h *HashMapSlice)  Get(key string) (interface{}, bool){
	idBucket := h.hash(key) % uint32(h.size)
	val, ok := h.cache[key]
	if ok{
		return val, true
	}
	for _, elemBucket := range h.buckets[idBucket]{
		if elemBucket.key == key{
			return elemBucket.value, true
		}
	}
	return nil, false
}

func WithHashFunc(f func(string) uint32) func(*HashMap){
	return func(h *HashMap){
		h.hash = f
	}
}

func MeassureTime(f func()) time.Duration {
	start := time.Now()
	f()
	since := time.Since(start)
	return since
}

func main() {
	time := MeassureTime(TestSlice16)
	fmt.Println(time)
	time = MeassureTime(TestSlice1000)
	fmt.Println(time)

	time = MeassureTime(TestList16)
	fmt.Println(time)
	time = MeassureTime(TestList1000)
	fmt.Println(time)
}

func TestList16() {
	m := NewHashMapList(16)
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestList1000() {
	m := NewHashMapList(1000)
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}



func TestSlice16() {
	m := NewHashMapSlice(16)
	for i := 0; i < 16; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	for i := 0; i < 16; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}

func TestSlice1000() {
	m := NewHashMapSlice(1000)
	for i := 0; i < 1000; i++ {
		m.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	for i := 0; i < 1000; i++ {
		value, ok := m.Get(fmt.Sprintf("key%d", i))
		if !ok {
			fmt.Printf("Expected key to exist in the HashMap")
		}
		if value != fmt.Sprintf("value%d", i) {
			fmt.Printf("Expected value to be 'value%d', got '%v'", i, value)
		}
	}
}