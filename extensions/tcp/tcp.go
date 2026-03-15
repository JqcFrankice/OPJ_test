package tcp

import (
	"encoding/binary"
	"fmt"
	"net"

	"go.k6.io/k6/js/modules"
	"google.golang.org/protobuf/proto"
	"k6-game-test/proto/pb" // 假设通过 protoc 生成后存放于此
)

// TCPClient 定义扩展方法
type TCPClient struct{}

// SendProtoMsg 封装发送 Protobuf 消息 (带长度前缀)
func (t *TCPClient) SendProtoMsg(conn net.Conn, msgID uint32, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	// 简单的封包协议: 4字节长度 + 4字节MsgID + 数据体
	buf := make([]byte, 8+len(data))
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(data)+4))
	binary.BigEndian.PutUint32(buf[4:8], msgID)
	copy(buf[8:], data)

	_, err = conn.Write(buf)
	return err
}

func init() {
	modules.Register("k6/x/tcp", new(TCPClient))
}
