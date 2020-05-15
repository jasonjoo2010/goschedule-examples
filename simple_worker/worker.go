package main

import (
	"fmt"
	"time"

	"github.com/jasonjoo2010/goschedule/core/worker"
)

type DemoStrategy struct {
	worker.Worker
	notifier chan int
}

func (demo *DemoStrategy) loop() {
	i := 1
LOOP:
	for {
		fmt.Println("working: ", i)
		select {
		case <-demo.notifier:
			break LOOP
		default:
			i++
			time.Sleep(time.Second)
		}
	}
}

func (demo *DemoStrategy) Start(strategyId, parameter string) {
	if demo.notifier == nil {
		demo.notifier = make(chan int)
	}
	go demo.loop()
	fmt.Println("worker started")
}

func (demo *DemoStrategy) Stop(strategyId string) {
	fmt.Println("prepare to stop")
	demo.notifier <- 1
	fmt.Println("worker exited")
}
