package gsgetway

import (
	"context"
	"fmt"
	"net"
)

type TcpUser struct {
	MemberID int
	Hash     string
	UserType int
	Socket   net.Conn
	Data     chan byte
	ctx      context.Context
	cancel   context.CancelFunc
}

func (this *TcpUser) Read() []byte {

	return nil
}

func (this *TcpUser) Write(dach chan byte) error {

	return nil
}

func (this *TcpUser) Process(conn net.Conn) {
	fmt.Println("tcpuser process")

}
