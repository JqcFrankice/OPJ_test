package main

import (
	"fmt"
	"k6-game-test/engine"
	"k6-game-test/scenarios"
	"time"
)

func main() {
	// 配置压测参数
	vus := 50                    // 并发虚拟用户数
	duration := 10 * time.Second // 持续时间

	fmt.Printf("开始压测: VUs=%d, 持续时间=%v\n", vus, duration)

	// 运行场景
	// 注意：您需要在 scenarios 包中实现 LoginScenario 或其他具体测试函数
	engine.Run(vus, duration, scenarios.ExampleScenario)

	fmt.Println("压测结束")
}
