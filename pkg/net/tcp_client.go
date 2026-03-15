package net

import (
	"encoding/binary"
	"google.golang.org/protobuf/proto"
	"net"
)

// SendProtoMsg 封装发送 Protobuf 消息 (带长度前缀)
func SendProtoMsg(conn net.Conn, msgID uint32, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	// 封包协议: 4字节总长度 + 4字节MsgID + 数据体
	// 注意：大多数游戏服务器使用 LittleEndian，请务必与服务端确认！
	dataLen := uint32(len(data) + 4)

	buf := make([]byte, 8+len(data))

	// 使用 LittleEndian 写入 (如果是 Zinx 默认协议)
	binary.LittleEndian.PutUint32(buf[0:4], dataLen)
	binary.LittleEndian.PutUint32(buf[4:8], msgID)
	copy(buf[8:], data)

	_, err = conn.Write(buf)
	return err
}
