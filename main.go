package main

import (
	"fmt"
	"k6-game-test/engine"
	"k6-game-test/scenarios"
	"time"
)

func main() {
	// 配置参数拆分
	vus := 1 // 并发数
	qps := 1 // 吞吐量 (设置较大以避免阻塞，由场景内 Sleep 控制频率)
	duration := 1 * time.Minute

	fmt.Printf("开始压测: VUs=%d, Target QPS=%d, 持续时间=%v\n", vus, qps, duration)

	// 运行场景
	engine.Run(vus, qps, duration, scenarios.AddRandomItemScenario)

	fmt.Println("压测结束")
}
