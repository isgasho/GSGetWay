package gsgetway

import (
	"context"
	"log"
	"net"
)

type TcpThread struct {
	Name     string
	TcpIp    string
	UserList map[int]TcpUser
	ctx      context.Context
	cancel   context.CancelFunc
}

func (this *TcpThread) Init(name string, ip string) {
	this.Name = name
	this.TcpIp = ip
	this.UserList = map[int]TcpUser{}
	this.ctx, this.cancel = context.WithCancel(context.Background())
}

//开始服务
func (this *TcpThread) Start() {
	go this.Process()
}

//关闭服务
func (this *TcpThread) Stop() {
	if this.cancel != nil {
		this.cancel()
	}
}

//方法体
func (this *TcpThread) Process() {
	listener, err := net.Listen("tcp", this.TcpIp)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() //关闭监听的端口
	for {
		select {
		case <-this.ctx.Done():
			return
		default:
			conn, err := listener.Accept() //用conn接收链接
			if err != nil {
				log.Fatal(err)
			}
			ctx1, cal := context.WithCancel(this.ctx)
			tcpuser := TcpUser{
				ctx:    ctx1,
				cancel: cal,
			}
			go tcpuser.Process(conn)
		}

		// conn.Write([]byte("Yinzhengjie\n")) //通过conn的wirte方法将这些数据返回给客户端。
		// conn.Write([]byte("hello Golang\n"))
		// time.Sleep(time.Minute) //在结束这个链接之前需要睡一分钟在结束当前循环。
		// conn.Close()            //与客户端断开连接。
	}

}

var Instance = TcpThread{}
