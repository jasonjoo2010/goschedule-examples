package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jasonjoo2010/goschedule/core"
	"github.com/jasonjoo2010/goschedule/core/worker"
	"github.com/jasonjoo2010/goschedule/store/database"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("Create db instance error:", err.Error())
		return
	}
	store := database.New("/schedule/demo/simple", db)
	if store == nil {
		fmt.Println("Create db store failed")
		return
	}
	manager, err := core.New(store)
	if err != nil {
		fmt.Println(err)
		return
	}
	worker.Register(&HotSellingRefresher{})
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
