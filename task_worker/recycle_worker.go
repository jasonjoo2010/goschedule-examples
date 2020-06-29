package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jasonjoo2010/goschedule/core/definition"
)

type orderRecycleTask struct {
}

func (demo *orderRecycleTask) Select(parameter, ownSign string, items []definition.TaskItem, eachFetchNum int) []interface{} {
	var result []interface{}
	if rand.Intn(10) == 0 {
		// 10% possibility no data returned, we are done this time
		fmt.Println(ownSign, time.Now().Format(time.RFC3339), "no more")
		return result
	}
	cnt := rand.Intn(eachFetchNum) + 1
	for cnt > 0 {
		result = append(result, rand.Intn(9999999999))
		cnt--
	}
	fmt.Println(ownSign, time.Now().Format(time.RFC3339), "select", len(result), "expired orders")
	return result
}

func (demo *orderRecycleTask) Execute(task interface{}, ownSign string) bool {
	// simulate execute cost
	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(200)))
	fmt.Println(ownSign, time.Now().Format(time.RFC3339), "close order", task)
	return true
}
