package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jasonjoo2010/goschedule/core"
	"github.com/jasonjoo2010/goschedule/core/worker"
	"github.com/jasonjoo2010/goschedule/store/redis"
)

func main() {
	manager, err := core.New(redis.New("/schedule/demo/simple", "127.0.0.1", 6379))
	if err != nil {
		fmt.Println(err)
		return
	}
	worker.Register(DemoStrategy{})
	manager.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)
LOOP:
	for {
		select {
		case <-c:
			manager.Shutdown()
			break LOOP
		default:
		}
		time.Sleep(time.Second)
	}
}
