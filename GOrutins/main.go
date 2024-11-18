package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var action = []string{ "logged in", "logged out", "create record", "delete record", "update record",}

type logItem struct{ /*Структура действия и время действия*/
    action string
    timestamp time.Time
}

type User struct{ /*Информация о пользователе*/
    id int
    email string
    logs []logItem /*Массив из структур активностей и времени активности*/
}


/*принимает в себя структуру Инфы о пользователе*/
/*    |   */
/*    V   */
func (u User) getActivityInfo() string{  /*Функция получения информации о пользователе*/
    
    out := fmt.Sprintf("ID: %d | email: %s\nActivity log:\n", u.id, u.email)
    for index, item := range u.logs{
        out += fmt.Sprintf("%d. [%s] at %s\n", index + 1, item.action, item.timestamp)
    }
    return out
}

func main(){
    
    wg := &sync.WaitGroup{}

    rand.Seed(time.Now().Unix())
    users := generateUsers(1000)
    for _, user := range users{
        wg.Add(1)
        go saveUserInfo(user, wg)
    }
    wg.Wait()
}

func saveUserInfo(user User, wg *sync.WaitGroup){
    fmt.Printf("WRITING IN FILE: \n")

    filename := fmt.Sprintf("logs/userID_%d.txt", user.id)
    file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil{
        fmt.Println(err)
    }

    file.WriteString(user.getActivityInfo())
    wg.Done()
}

func generateUsers(count int) []User{
    users := make([]User, count)

    for i := 0; i < count; i++{
        users[i] = User{
            id: i+1,
            email: fmt.Sprintf("user%d@mail.ru", i+1),
            logs: generateLogs(rand.Intn(1000)),
        }
    }
    return users
}

func generateLogs(count int) []logItem{
    logs := make([]logItem, count)

    for i := 0; i < count; i++{
        logs[i] = logItem{
            timestamp: time.Now(),
            action: action[rand.Intn(len(action) - 1)],
        }
    } 
    return logs
}