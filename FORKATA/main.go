package main

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type HashMap struct{
	data map[uint32]interface{}
	hash func(string)uint32
}

func (hm *HashMap) Set(key string, value interface{}){
	hashKey := hm.hash(key)
	hm.data[hashKey] = value
}

func (hm *HashMap) Get(key string) (interface{}, bool){
	hashKey := hm.hash(key)
	value, ok := hm.data[hashKey]
	return value, ok
}

func crc32Hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func crc64Hash(key string) uint32 {
	table := crc64.MakeTable(crc64.ECMA)
	return uint32(crc64.Checksum([]byte(key), table))
}

func crc16Hash(key string) uint32 { //Мясо из gpt по переводу crc16
	var crc uint16 = 0xFFFF
	for _, b := range []byte(key) {
		crc ^= uint16(b)
		for i := 0; i < 8; i++ {
			if (crc & 0x0001) != 0 {
				crc >>= 1
				crc ^= 0xA001
			} else {
				crc >>= 1
			}
		}
	}
	return uint32(crc)
}

func crc8Hash(key string) uint32 { //Мясо из gpt по переводу crc8
	var crc uint8 = 0
	for _, b := range []byte(key) {
		crc ^= b
		for i := 0; i < 8; i++ {
			if (crc & 0x80) != 0 {
				crc = (crc << 1) ^ 0x07
			} else {
				crc <<= 1
			}
		}
	}
	return uint32(crc)
}


func NewHashMap(size int, crcs ...func(*HashMap)) *HashMap{
	hm := &HashMap{
		data: make(map[uint32]interface{}),
		hash: nil,
	}
	for _, crc := range crcs{
		crc(hm)
	}
	return hm
}

func WithHashCRC64() func(*HashMap){
	return func(hm *HashMap){
		hm.hash = crc64Hash
	}
}

func WithHashCRC32() func(*HashMap){
	return func(hm *HashMap){
		hm.hash = crc32Hash
	}
}


func WithHashCRC16() func(*HashMap){
	return func(hm *HashMap){
		hm.hash = crc16Hash
	}
}

func WithHashCRC8() func(*HashMap){
	return func(hm *HashMap){
		hm.hash = crc8Hash
	}
}

func MeassureTime(f func()) time.Duration{
	s := time.Now()
	f()
	time := time.Since(s)
	return time
}


func main() {
	m := NewHashMap(16, WithHashCRC64())
	since := MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC32())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC16())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC8())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)
}