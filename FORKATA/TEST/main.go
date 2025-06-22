package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

func monitorGoroutines(prevGoroutines int) {
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C{
		countGorout := runtime.NumGoroutine()

		percent := float64(countGorout - prevGoroutines) / float64(prevGoroutines)

		if percent < -0.2{
			log.Println("⚠️ Предупреждение: Количество горутин уменьшилось более чем на 20%!")
		}else if percent > 0.2{
			log.Println("⚠️ Предупреждение: Количество горутин увеличилось более чем на 20%!")
		}else{
			log.Printf("Текущее количество горутин: %d\n", countGorout)
		}
	}
}

func main() {	
	g, _ := errgroup.WithContext(context.Background())

	// Мониторинг горутин
	go func() {
		monitorGoroutines(runtime.NumGoroutine())
	}()

	// Имитация активной работы приложения с созданием горутин
	for i := 0; i < 64; i++ {
		g.Go(func() error {
			time.Sleep(5 * time.Second)
			return nil
		})
		time.Sleep(80 * time.Millisecond)
	}

	// Ожидание завершения всех горутин
	if err := g.Wait(); err != nil {
		fmt.Println("Ошибка:", err)
	}
}