package main

import (
	"fmt"
	"k6-game-test/engine"
	"k6-game-test/scenarios"
	"time"
)

func main() {
	// 配置参数拆分
	vus := 10 // 并发数 (Concurrency: 10个用户同时在线)
	qps := 10 // 吞吐量 (Throughput: 每秒处理10个请求)
	duration := 1 * time.Minute

	fmt.Printf("开始压测: VUs=%d, Target QPS=%d, 持续时间=%v\n", vus, qps, duration)

	// 运行场景
	engine.Run(vus, qps, duration, scenarios.ExampleScenario)

	fmt.Println("压测结束")
}
