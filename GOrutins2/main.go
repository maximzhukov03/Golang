package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var action = []string{ // Срез активностей
    "logged in",
    "logged out",
    "create info",
    "delete info",
    "update info",
}

type logItem struct{ // Структура Активности и времени
    action string
    timestamp time.Time
}

type User struct{ // Стрктура юзера
    id int
    email string
    logged []logItem
}

func (u User) getActivityInfo() string{ // Получение активности
    out := fmt.Sprintf("ID: %d | EMAIL: %s\nACTYVITY:\n", u.id, u.email)
    for index, item := range u.logged{
        out += fmt.Sprintf("[%d] [%s] AT:[%s]\n", index + 1, item.action, item.timestamp)
    }
    return out
}

func READERSTR() string{ // Функция для чтения Строки
    reader := bufio.NewReader(os.Stdin)
    str, _ := reader.ReadString('\n')
    str = strings.TrimSpace(str)
    return str
}

func (u User) getActivityFilter() string{ // Получение активности по фильтру
    out := fmt.Sprintf("ID: %d | EMAIL: %s\nACTYVITY:\n", u.id, u.email)

    fmt.Print("Слово для фильтра: ")
    filter := READERSTR()    

    for index, item := range u.logged{
        if item.action == filter{
            out += fmt.Sprintf("[%d] [%s] AT:[%s]\n", index + 1, item.action, item.timestamp)
        }
    }
    return out
}

func (u User) getActivityTime() string{ // Получение активности по времени
    out := fmt.Sprintf("ID: %d | EMAIL: %s\n", u.id, u.email)
    var timeStart, timeStop time.Time 
    for _, item := range u.logged{
        if item.action == "logged in"{
            timeStart = item.timestamp
        }
        if item.action == "logged out"{
            timeStop = item.timestamp
            break
        }
    }
    timeActivity := timeStop.Sub(timeStart)
    timeHour := int(timeActivity.Hours())
    timeMinetes := int(timeActivity.Minutes()) % 60
    out += fmt.Sprintf("Total activity time: %d Часов %d Минут", timeHour, timeMinetes)
    return out
}

func main(){

    t := time.Now()

    users := generateUsers(50) // Генерация юзеров(срез) при помощи функции

    wg := &sync.WaitGroup{} // Созание группы для горутин

    for _, user := range users{ //Итерируемся по юзерам чтобы записать их в файл
        wg.Add(1)
        go WriteFromFile(user, wg)
    }

    wg.Wait() // Ждем чтобы все горутины закончились
    fmt.Println("TIME DERATION: ", time.Since(t).String())//Сколько по времени заняло выполнение
}

func generateUsers(count int) []User{//Функция для генерации юзеров (На выходе структура)
    user := make([]User, count)

    for i := 0; i < count; i++{
        user[i] = User{
            id: i + 1,
            email: fmt.Sprintf("maximzhukov%d@mail.ru", i + 1),
            logged: generateLogs(rand.Intn(1000)),//генерация на рандом ДЕЙСТВИЙ
        }
    }

    return user
}


func generateLogs(count int) []logItem{ //Фунция для генерации Действий
    logs := make([]logItem, count)

    for i := 0; i < count; i++{
        logs[i] = logItem{
            action: action[rand.Intn(len(action) -1 )],
            timestamp: time.Now(),
        }
    }

    return logs
}

func WriteFromFile(user User, wg *sync.WaitGroup){ // Фунция для записи в файл юзеров отдельно
    time.Sleep(time.Millisecond * 10)
    fmt.Printf("WRITING USER %d FROM FILE uid_%d.txt\n", user.id, user.id)
    filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
    file, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0644)
    if err != nil{
        fmt.Println(err)
    }
    file.WriteString(user.getActivityInfo())
    wg.Done() //Горутина заканчивает работу
}