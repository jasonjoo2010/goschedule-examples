package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/jasonjoo2010/goschedule/utils"
)

const (
	REFRESH_PERIOD = 30 * time.Second
)

type HotSellingRefresher struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	wg        sync.WaitGroup

	lastRefresh time.Time
}

func (w *HotSellingRefresher) refresh() {
	// simulate the cost refreshing
	time.Sleep(time.Duration(rand.Intn(500)+1) * time.Millisecond)
	fmt.Println(time.Now().Format(time.RFC3339), "refreshed")
	w.lastRefresh = time.Now()
}

func (w *HotSellingRefresher) loop(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("enter the loop")
	defer func() {
		fmt.Println("exit the loop")
		wg.Done()
	}()

	ticker := time.NewTicker(REFRESH_PERIOD)
	defer ticker.Stop()

LOOP:
	for {
		w.refresh()
		if !utils.DelayContext(ctx, REFRESH_PERIOD) {
			break LOOP
		}
	}
}

func (w *HotSellingRefresher) Start(strategyId, parameter string) error {
	w.ctx, w.ctxCancel = context.WithCancel(context.Background())
	w.wg.Add(1)
	go w.loop(w.ctx, &w.wg)
	fmt.Printf("worker started: %p\n", w)
	return nil
}

func (w *HotSellingRefresher) Stop(strategyId, parameter string) error {
	fmt.Printf("prepare to stop: %p\n", w)
	w.ctxCancel()
	w.wg.Wait()

	fmt.Println("stopped")
	return nil
}
