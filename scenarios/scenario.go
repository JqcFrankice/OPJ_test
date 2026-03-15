package scenarios

import (
	"fmt"
	"time"
)

// ExampleScenario 是一个示例压测场景
// 为了实现 QPS=10，我们可以通过控制 sleep 时间来控制单用户的请求频率
// 如果有 10 个 VU，每个 VU 每秒发起 1 次请求，总 QPS 就是 10
func ExampleScenario() {
	// 在此处实现具体的协议交互逻辑

	// 模拟请求逻辑
	// conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	// ... 发送 pb 消息 ...

	fmt.Println("执行一次虚拟用户操作")

	// 模拟思考时间：1秒钟发起1次请求，10个用户总QPS即为10
	time.Sleep(1 * time.Second)
}
