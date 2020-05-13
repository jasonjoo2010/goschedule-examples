package main

import (
	"fmt"
	"time"

	"github.com/jasonjoo2010/goschedule/core"
	"github.com/jasonjoo2010/goschedule/store/redis"
)

func main() {
	manager, err := core.New(redis.New("/schedule/project/demo", "127.0.0.1", 6379))
	if err != nil {
		fmt.Println(err)
		return
	}
	manager.Start()
	for {
		time.Sleep(time.Second)
	}
}
