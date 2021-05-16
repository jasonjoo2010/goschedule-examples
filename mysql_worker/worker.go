package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	REFRESH_PERIOD = 30 * time.Second
	DELAY_UNIT     = 10 * time.Millisecond
)

type HotSellingRefresher struct {
	needStop    bool
	notifier    chan int
	lastRefresh time.Time
}

func (w *HotSellingRefresher) delay(d time.Duration) {
	// actually we change the cycle into 30 seconds
	tmr := time.NewTimer(d)
LOOP:
	for false == w.needStop {
		select {
		case <-tmr.C:
			break LOOP
		default:
			time.Sleep(DELAY_UNIT)
		}
	}
}

func (w *HotSellingRefresher) refresh() {
	// simulate the cost refreshing
	time.Sleep(time.Duration(rand.Intn(500)+1) * time.Millisecond)
	fmt.Println(time.Now().Format(time.RFC3339), "refreshed")
	w.lastRefresh = time.Now()
}

func (w *HotSellingRefresher) refreshOrWait() {
	// actually we change the cycle into 30 seconds
	now := time.Now()
	expire := w.lastRefresh.Add(REFRESH_PERIOD)
	if now.Before(expire) {
		w.delay(expire.Sub(now))
	}
	w.refresh()
}

func (w *HotSellingRefresher) loop() {
	fmt.Println("enter the loop")
	defer func() {
		w.notifier <- 1
		fmt.Println("exit the loop")
	}()
	for false == w.needStop {
		w.refreshOrWait()
	}
}

func (w *HotSellingRefresher) Start(strategyId, parameter string) error {
	if w.notifier == nil {
		w.notifier = make(chan int)
	}
	go w.loop()
	fmt.Println("worker started")

	return nil
}

func (w *HotSellingRefresher) Stop(strategyId, parameter string) error {
	fmt.Println("prepare to stop")
	w.needStop = true
	<-w.notifier
	fmt.Println("stopped")

	return nil
}
