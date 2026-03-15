package scenarios

import (
	"fmt"
)

// ExampleScenario 是一个示例压测场景
// 业务逻辑完全专注于执行一次用户操作
// 请求频率(QPS)由 engine 控制
func ExampleScenario() {
	// 在此处实现具体的协议交互逻辑
	// 例如:
	// conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	// ... 发送 pb 消息 ...

	fmt.Println("执行一次虚拟用户操作")
}
