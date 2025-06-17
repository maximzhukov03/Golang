package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type cache struct{
	client *redis.Client
}

type Cacher interface {
    Set(key string, value interface{}) error
    Get(key string) (interface{}, error)
}

func NewCache(client *redis.Client) Cacher {
    return &cache{
        client: client,
    }
}

func (c *cache) Set(key string, value interface{}) error{
	err := c.client.Set(key, value, 5 * time.Minute)
	if err != nil{
		log.Println("Ошибка сохранения данных в кэше:", err)
		return fmt.Errorf("%s: %w", key, err)
	}
	return nil
}

func (c *cache) Get(key string) (interface{}, error){
	result, err := c.client.Get(key).Result()
	if err == redis.Nil{
		fmt.Printf("not found by key %key%")
		return nil, fmt.Errorf("not found by key %s", key)
	}
	return result, nil
}

type User struct {
	ID int
    Name string
    Age int
}

func main() {
	// Создание клиента Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
	})

	cache := NewCache(client)
	
	// Установка значения по ключу
	err := cache.Set("some:key", "value")
	if err != nil {
        panic(err)
    }
	
	// Получение значения по ключу
	value, err := cache.Get("some:key")
	if err != nil {
        panic(err)
    }
	
	fmt.Println(value)
	
	user := &User{
        ID: 1,
        Name: "John",
        Age: 30,
    }
	// Установка значения по ключу
	err = cache.Set(fmt.Sprintf("user:%v", user.ID), user)
	if err != nil {
        panic(err)
    }
	
	// Получение значения по ключу
	value, err = cache.Get("user:1")
	if err != nil {
        panic(err)
    }
	
	fmt.Println(value)
}