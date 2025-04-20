package main

import "fmt"

// Структура пользователя
type User struct {
    ID   int
    Name string
    Age  int
}

// Функция слияния двух отсортированных массивов пользователей
func Merge(arr1 []User, arr2 []User) []User {
    return merge(arr1, arr2)
}

func merge(left, right []User) []User {
    result := make([]User, 0, len(left)+len(right))
    for len(left) > 0 && len(right) > 0 {
        if left[0].ID <= right[0].ID {
            result = append(result, left[0])
            left = left[1:]
        } else {
            result = append(result, right[0])
            right = right[1:]
        }
    }
    result = append(result, left...)
    result = append(result, right...)
    return result
}

func main(){
	users1 := []User{
        {ID: 1, Name: "Алексей", Age: 30},
        {ID: 2, Name: "Иван", Age: 25},
        {ID: 3, Name: "Сергей", Age: 35},
        {ID: 4, Name: "Дмитрий", Age: 28},
        {ID: 5, Name: "Анна", Age: 22},
    }
    users2 := []User{
        {ID: 6, Name: "Елена", Age: 40},
        {ID: 7, Name: "Олег", Age: 32},
        {ID: 8, Name: "Мария", Age: 27},
        {ID: 9, Name: "Николай", Age: 29},
        {ID: 10, Name: "Татьяна", Age: 31},
    }
    mergedUsers := Merge(users1, users2)
    fmt.Println(mergedUsers)
}