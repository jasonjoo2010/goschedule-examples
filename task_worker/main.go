package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jasonjoo2010/goschedule/core"
	"github.com/jasonjoo2010/goschedule/core/worker/task_worker"
	"github.com/jasonjoo2010/goschedule/store/redis"
	"github.com/jasonjoo2010/goschedule/types"
)

func main() {
	inst := demoSingleTask{
		baseTask: baseTask{
			Name: "single instanced",
		},
	}
	store := redis.New("/schedule/demo/task", "127.0.0.1", 6379)
	defer store.Close()

	manager, err := core.New(types.ScheduleConfig{}, store)
	if err != nil {
		fmt.Println(err)
		return
	}
	task_worker.RegisterTaskTypeName("singleExecutor", &demoSingleTask{})
	task_worker.RegisterTaskTypeName("batchExecutor", &demoBatchTask{})
	task_worker.RegisterTaskInstName("singleStreamExecutor", inst)
	task_worker.RegisterTaskTypeName("orderRecycleTask", &orderRecycleTask{})
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
