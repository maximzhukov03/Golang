package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Node struct {
	index int
	data  *User
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(user *User) *BinaryTree {
    if t.root == nil{
        t.root = &Node{
            index: user.ID,
            data: user,
        }
    } else {
        t.root.insert(user)
    }

    return t
}

func (n *Node) insert(user *User) {
    if user.ID < n.index{
        if n.left == nil{
            n.left = &Node{
                index: user.ID,
                data: user,
            }
        } else {
            n.left.insert(user)
        }
    } else {
        if n.right == nil{
            n.right = &Node{
                index: user.ID,
                data: user,
            }
        } else {
            n.right.insert(user)
        }
    }
}

func (t *BinaryTree) search(key int) *User {
    if t.root.index == key{
        return t.root.data
    } else {
        user := t.root.search(key)
        if user != nil{
            return user
        }
        return nil
    }
}

func (n *Node) search(key int) *User {
    if n == nil{
        return nil
    }
    if key < n.index {
        if n.left != nil{
            return n.left.search(key)
        }
    } else if key > n.index{
        if n.right != nil{
            return n.right.search(key)
        }
    } else {
        return n.data
    }
    return nil
}

func generateData(n int) *BinaryTree {
	rand.Seed(time.Now().UnixNano())
	bt := &BinaryTree{}
	for i := 0; i < n; i++ {
		val := rand.Intn(100)
		bt.insert(&User{
			ID:   val,
			Name: fmt.Sprintf("User%d", val),
			Age:  rand.Intn(50) + 20,
		})
	}
	return bt
}

func main() {
	bt := generateData(50)
	user := bt.search(30)
	if user != nil {
		fmt.Printf("Найден пользователь: %+v\n", user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}