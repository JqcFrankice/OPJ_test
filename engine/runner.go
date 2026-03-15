// engine/runner.go - 核心并发引擎
package engine

import (
	"sync"
	"time"
)

// Runner 执行压测
func Run(vus int, duration time.Duration, scenario func()) {
	var wg sync.WaitGroup
	stop := time.After(duration)

	for i := 0; i < vus; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					scenario()
				}
			}
		}()
	}
	wg.Wait()
}
