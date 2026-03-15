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
	buf := make([]byte, 8+len(data))
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(data)+4))
	binary.BigEndian.PutUint32(buf[4:8], msgID)
	copy(buf[8:], data)

	_, err = conn.Write(buf)
	return err
}
