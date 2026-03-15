package scenarios

import (
	"fmt"
	"k6-game-test/pkg/net"
	"k6-game-test/proto/pb"
	"math/rand"
	stdnet "net"
	"time"
)

// AddRandomItemScenario 发送一次随机道具增加请求
func AddRandomItemScenario() {
	// 1. 建立 TCP 连接 (修改为 8999 端口)
	conn, err := stdnet.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		return
	}
	defer conn.Close()

	// 2. 构造随机请求数据
	// 根据 config/items.json，ID 范围是 1-10
	itemID := int32(rand.Intn(10) + 1)

	// 简单的类型映射逻辑
	var itemType int32 = 1
	switch {
	case itemID <= 2:
		itemType = 1
	case itemID <= 4:
		itemType = 2
	case itemID <= 6:
		itemType = 3
	case itemID == 7:
		itemType = 4
	case itemID == 8:
		itemType = 5
	case itemID == 9:
		itemType = 6
	case itemID == 10:
		itemType = 7
	}

	req := &pb.GetItemReq{
		ItemId:   itemID,
		ItemType: itemType,
		Count:    1,
	}

	// 3. 发送消息 (假设 MsgID_GET_ITEM_REQ = 3001)
	err = net.SendProtoMsg(conn, 3001, req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
	} else {
		fmt.Println("成功触发随机道具增加场景")
	}

	// 等待3秒
	time.Sleep(3 * time.Second)
}
