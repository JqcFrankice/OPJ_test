package scenarios

import (
	"fmt"
	"k6-game-test/pkg/net"
	"math/rand"
	"net"
	// "k6-game-test/proto/pb" // 生成 pb 代码后请取消注释
)

// AddRandomItemScenario 发送一次随机道具增加请求
func AddRandomItemScenario() {
	// 1. 建立 TCP 连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	defer conn.Close()

	// 2. 构造随机请求数据
	// 注意：pb 包需要先运行 protoc 生成！
	// req := &pb.GetItemReq{
	// 	ItemId:   rand.Int31n(100),
	// 	ItemType: rand.Int31n(5),
	// 	Count:    rand.Int31n(10) + 1,
	// }

	// 3. 发送消息 (假设 MsgID_GET_ITEM_REQ = 3001)
	// err = net.SendProtoMsg(conn, 3001, req)

	fmt.Println("成功触发随机道具增加场景 (需先生成 pb 代码)")
}
