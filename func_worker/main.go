package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/jasonjoo2010/goschedule/core"
	"github.com/jasonjoo2010/goschedule/core/worker"
	"github.com/jasonjoo2010/goschedule/store/redis"
	"github.com/jasonjoo2010/goschedule/types"
)

type HotSellingRefresher struct {
}

func (w *HotSellingRefresher) refresh() {
	// simulate the cost refreshing
	time.Sleep(time.Duration(rand.Intn(500)+1) * time.Millisecond)
	fmt.Println(time.Now().Format(time.RFC3339), "refreshed")
}

func main() {
	store := redis.New("/schedule/demo/func", "127.0.0.1", 6379)
	defer store.Close()

	manager, err := core.New(types.ScheduleConfig{}, store)
	if err != nil {
		fmt.Println(err)
		return
	}
	refresher := HotSellingRefresher{}
	worker.RegisterFunc("HotSellingRefresher", func(strategyId, parameter string) {
		refresher.refresh()
	})
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
			time.Sleep(time.Second)
		}
	}
}
