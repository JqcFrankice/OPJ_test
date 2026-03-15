package scenarios

import (
	"fmt"
	"time"
)

// ExampleScenario 是一个示例压测场景
func ExampleScenario() {
	// 在此处实现具体的协议交互逻辑
	// 例如:
	// conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	// ... 发送 pb 消息 ...
	// ... 接收响应 ...

	fmt.Println("执行一次虚拟用户操作")
	time.Sleep(100 * time.Millisecond) // 模拟思考时间
}
