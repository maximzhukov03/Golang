package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var res string
	res = contextWithTimeout(context.Background(), 1*time.Second, 2*time.Second)
	fmt.Println(res)
	res = contextWithTimeout(context.Background(), 2*time.Second, 1*time.Second)
	fmt.Println(res)
}

func contextWithTimeout(ctx context.Context, contextTimeout time.Duration, timeAfter time.Duration) string {
	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()
	timer := time.NewTimer(timeAfter)
	defer timer.Stop()
	select{
	case <-ctx.Done():
		return "превышено время ожидания"
	case <-timer.C:
		return "превышено время ожидания контекста"
	}
}

// func contextWithDeadline(ctx context.Context, contextDeadline time.Duration, timeAfter time.Duration) string {
// 	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(contextDeadline))
// 	defer cancel()
// 	timer := time.NewTimer(timeAfter)
// 	defer timer.Stop()
//     select {
//     case <-ctx.Done():
//         return "context deadline exceeded"
//     case <-timer.C:
//         return "time after exceeded"
//     }
// }