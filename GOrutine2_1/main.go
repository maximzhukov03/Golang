package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var action = []string{
    "logged in",
    "logged out",
    "create info",
    "delete info",
    "update info",
}

type logItem struct{
    action string
    timestamp time.Time
}

type User struct{
    id int
    email string
    logged []logItem
}

func (u User) getActivityInfo() string{
    out := fmt.Sprintf("ID: %d | EMAIL: %s\nACTYVITY:\n", u.id, u.email)
    for index, item := range u.logged{
        out += fmt.Sprintf("[%d] [%s] AT:[%s]\n", index + 1, item.action, item.timestamp)
    }
    return out
}

func READERSTR() string{
    reader := bufio.NewReader(os.Stdin)
    str, _ := reader.ReadString('\n')
    str = strings.TrimSpace(str)
    return str
}

func (u User) getActivityFilter() string{
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

func (u User) getActivityTime() string{
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
    currentTime := time.Now()
    newTime := currentTime.Add(2 * time.Hour)
    u := User{
        id: 213123,
        email: "maximzhukov03@mail.ru",
        logged: []logItem{
            {action[0], time.Now()},
            {action[2], time.Now()},
            {action[4], time.Now()},
            {action[2], time.Now()},
            {action[4], time.Now()},
            {action[3], time.Now()},
            {action[4], time.Now()},
            {action[3], time.Now()},
            {action[2], time.Now()},
            {action[4], time.Now()},
            {action[2], time.Now()},
            {action[4], time.Now()},
            {action[2], time.Now()},
            {action[1], newTime},
        },  
    }

    fmt.Println(u.getActivityInfo())
    fmt.Println(u.getActivityFilter())    
    fmt.Println(u.getActivityTime())
}