package engine

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Run 压测执行，拆分并发度(VUs)和吞吐量(QPS)
// vus: 并发 Goroutine 数量
// qps: 每秒总请求数限制
// duration: 持续时间
// scenario: 执行的业务场景
func Run(vus int, qps int, duration time.Duration, scenario func()) {
	var wg sync.WaitGroup
	// 创建一个限流器，允许每秒 qps 次突发请求
	limiter := rate.NewLimiter(rate.Limit(qps), qps)

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	for i := 0; i < vus; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// 检查是否结束
				select {
				case <-ctx.Done():
					return
				default:
					// 获取令牌，确保吞吐量限制
					if err := limiter.Wait(ctx); err != nil {
						return
					}
					// 执行业务场景
					scenario()
				}
			}
		}()
	}
	wg.Wait()
}
