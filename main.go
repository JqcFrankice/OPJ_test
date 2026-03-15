package main

import (
	"fmt"
	"k6-game-test/engine"
	"k6-game-test/scenarios"
	"time"
)

func main() {
	// 根据您的要求进行配置
	// DAU=10 (每日活跃用户，在此处体现为并发虚拟用户数 VUs)
	// QPS=10 (每秒查询数，即所有用户每秒总共发起 10 次请求)
	// 持续 1 分钟
	vus := 10
	duration := 1 * time.Minute

	fmt.Printf("开始压测: VUs=%d, QPS=10, 持续时间=%v\n", vus, duration)

	// 运行场景
	engine.Run(vus, duration, scenarios.ExampleScenario)

	fmt.Println("压测结束")
}
