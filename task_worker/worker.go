package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jasonjoo2010/goschedule/core/definition"
)

type baseTask struct {
	Name       string
	counterMap map[string]int64
}

func (demo baseTask) Select(parameter, ownSign string, items []definition.TaskItem, eachFetchNum int) []interface{} {
	if demo.counterMap == nil {
		demo.counterMap = make(map[string]int64, 1)
	}
	result := make([]interface{}, 0, len(items)*eachFetchNum)
	for _, item := range items {
		cnt, ok := demo.counterMap[item.Id]
		if !ok {
			cnt = 1
		}
		for i := 0; i < eachFetchNum; i++ {
			result = append(result, fmt.Sprint(item.Id, ":", cnt))
			cnt++
		}
		demo.counterMap[item.Id] = cnt
	}
	time.Sleep(time.Second)
	return result
}

type demoSingleTask struct {
	baseTask
}

func (demo demoSingleTask) Execute(task interface{}, ownSign string) bool {
	if demo.Name != "" {
		fmt.Print("Task(", demo.Name, ") ", task, "\n")
	} else {
		fmt.Println("Task", task)
	}
	time.Sleep(100 * time.Millisecond)
	return true
}

type demoBatchTask struct {
	baseTask
}

func (demo demoBatchTask) Execute(tasks []interface{}, ownSign string) bool {
	builder := strings.Builder{}
	builder.WriteString("Got ")
	builder.WriteString(strconv.Itoa(len(tasks)))
	builder.WriteString(" tasks:\n")
	for _, task := range tasks {
		str, ok := task.(string)
		if !ok {
			continue
		}
		builder.WriteString("\t")
		builder.WriteString(str)
		builder.WriteString("\n")
	}
	fmt.Println(builder.String())
	time.Sleep(100 * time.Millisecond)
	return true
}
