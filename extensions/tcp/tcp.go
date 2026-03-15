package tcp

import (
	"go.k6.io/k6/js/modules"
	"net"
)

// TCPClient 模拟 TCP 客户端扩展
type TCPClient struct{}

// Connect 建立连接
func (t *TCPClient) Connect(addr string) (*net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	return &conn, err
}

// 注册插件模块
func init() {
	modules.Register("k6/x/tcp", new(TCPClient))
}

// 注册插件模块
func init() {
	modules.Register("k6/x/tcp", new(TCPClient))
}
