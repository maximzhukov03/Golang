package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"
)

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n int) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

// LoadData loads data from a JSON file at the given path into the list.
func (d *DoubleLinkedList) LoadData(path string) error {
    bytes, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    var commits []Commit
    if err := json.Unmarshal(bytes, &commits); err != nil {
        return err
    }
    QuickSort(commits)
    d.Init(commits)
    return nil
}

func (d *DoubleLinkedList) Init(c []Commit){
	d.head = nil
	d.tail = nil
	d.curr = nil
	d.len = 0

	for _, comm := range c{
		node := &Node{data: &comm}
		
		if d.tail == nil{
			d.head = node
			d.tail = node
		} else {
			d.tail.next = node
			node.prev = d.tail
			d.tail = node
		}
		d.len = d.len + 1
	}	
}


// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	count := 1
	current := d.head
	for current.next != nil {
		count += 1
		current = current.next
	}
	return count
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
    if d.curr == nil {
        if d.head == nil {
            return nil
        }

        d.curr = d.head
    } else {

        d.curr = d.curr.next
    }
    return d.curr
}

// Prev получение предыдущего элемента
func (d *DoubleLinkedList) Prev() *Node {
    if d.curr == nil {

        if d.tail == nil {
            return nil
        }
        d.curr = d.tail
    } else {
        d.curr = d.curr.prev
    }
    return d.curr
}

// Insert вставка элемента после n элемента

// Insert inserts a new node with commit c at position n.
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 0 || n > d.len {
		return errors.New("index out of bounds")
	}
	newNode := &Node{data: &c}
	if n == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
		}
	} else if n == d.len {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		newNode.next = current
		newNode.prev = current.prev
		current.prev.next = newNode
		current.prev = newNode
	}
	d.len++
	return nil
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {

    if n < 0 || n >= d.len {
        return errors.New("INDEX ERROR")
    }

    if n == 0 { // Если элемент первый
        d.head = d.head.next
        if d.head != nil {
            d.head.prev = nil
        } else {
            d.tail = nil
        }
        d.len--
        return nil
    }

    if n == d.len-1 { // Если элемент последний
        d.tail = d.tail.prev
        d.tail.next = nil
        d.len--
        return nil
    }

    current := d.head // Если элемент где-то в списке
    for i := 0; i < n; i++ {
        current = current.next
    }

    current.prev.next = current.next
    current.next.prev = current.prev
    d.len--

    return nil
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.curr == d.head{
		d.head = d.head.next
        if d.head != nil {
            d.head.prev = nil
        } else {
            d.tail = nil
        }
        d.len--
        return nil
	}
	
	if d.curr == d.tail{
		d.tail = d.tail.prev
        d.tail.next = nil
        d.len--
        return nil
	}

	d.curr.prev.next = d.curr.next
    d.curr.next.prev = d.curr.prev
    d.len--
	
	return nil
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	current := d.head
	index := 0
	for current != d.curr{
		index += 1
		current = current.next
	}
	return index, nil
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	if d.tail == nil {
		return nil
	}
	node := d.tail
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.len--
	return node
}


// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
    if d.head == nil {
        return nil
    }
    node := d.head
    d.head = d.head.next
    if d.head != nil {
        d.head.prev = nil
    } else {
        d.tail = nil
    }
    d.len--
    return node
}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	current := d.head
	for current != nil{
		if current.data.UUID == uuID{
			return current
		}
	}
	return nil
}

// Search поиск коммита по message
func (d *DoubleLinkedList) Search(message string) *Node {
	current := d.head
	for current != nil{
		if current.data.Message == message{
			return current
		}
	}
	return nil
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
    current := d.head
    var prev *Node
    var next *Node
    d.tail = d.head
    for current != nil {
        next = current.next
        current.next = prev
        current.prev = next
        prev = current
        current = next
    }
    d.head = prev
    return d
}

type Commit struct {
    Message string `json:"message"`
    UUID    string `json:"uuid"`
    Date    string `json:"date"` 
}

func compareCommitsEssential(comm1, comm2 Commit) int{
	date1, _ := time.Parse("2006-01-02", comm1.Date)
	date2, _ := time.Parse("2006-01-02", comm2.Date)

	if date1.Before(date2){
		return -1
	} else if date1.After(date2){
		return 1
	}
	return 0
}

// QuickSort sorts an array of Commits in ascending order by their Date.
func QuickSort(commits []Commit) {
	if len(commits) < 2 {
		return
	}

	left, right := 0, len(commits)-1

	// Pick a pivot.
	pivotIndex := len(commits) / 2

	// Move the pivot to the right.
	commits[pivotIndex], commits[right] = commits[right], commits[pivotIndex]

	// Pile elements smaller than the pivot on the left.
	for i := range commits {
		if compareCommitsEssential(commits[i], commits[right]) < 0 {
			commits[i], commits[left] = commits[left], commits[i]
			left++
		}
	}

	// Place the pivot after the last smaller element.
	commits[left], commits[right] = commits[right], commits[left]

	// Go down the rabbit hole.
	QuickSort(commits[:left])
	QuickSort(commits[left+1:])
}	

func GenerateData(numCommits int) []Commit {
	var commits []Commit
	gofakeit.Seed(0) // Initialize the random seed

	// Define how many commits you want to generate
	for i := 0; i < numCommits; i++ {
		commit := Commit{
			Message: gofakeit.Sentence(5),                                                                                                                // Generate a random sentence with 5 words
			UUID:    gofakeit.UUID(),                                                                                                                     // Generate a random UUID
			Date:    gofakeit.DateRange(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)).Format("2006-01-02"), // Generate a random date between 2020 and 2022
		}
		commits = append(commits, commit)
	}

	return commits
}

func main(){
	list := &DoubleLinkedList{}
	err := list.LoadData("test.json")
	if err != nil{
		fmt.Errorf("Ошибка в обработке")
	}
	fmt.Println(list.head.data)
}