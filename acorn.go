package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func jiggling(ctx context.Context, interval time.Duration) {
	relPos := 10
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit")
			return
		case <-ticker.C:
			robotgo.MoveRelative(relPos, relPos)
			time.Sleep(time.Millisecond * 200)
			relPos *= -1
		}
	}
}
