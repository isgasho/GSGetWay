package gsgetway

import (
	"context"
	"fmt"
	"log"
	"net"
)

var Instance = &TcpThread{}

type TcpThread struct {
	Name      string
	TcpIp     string
	UserList  map[int]TcpUser
	Tcplisten *net.Listener
	Ctx       context.Context
	cancel    context.CancelFunc
}

func (this *TcpThread) Init(name string, ip string) {
	this.Name = name
	this.TcpIp = ip
	this.UserList = map[int]TcpUser{}
	this.Ctx, this.cancel = context.WithCancel(context.Background())
}

//开始服务
func (this *TcpThread) Start() {
	go this.Process()
}

//关闭服务
func (this *TcpThread) Stop() {
	if this.cancel != nil {
		this.cancel()
		(*this.Tcplisten).Close()
	}
}

//方法体
func (this *TcpThread) Process() {
	fmt.Println("process.")
	listener, err := net.Listen("tcp", this.TcpIp)
	if err != nil {
		log.Fatal(err)
	}
	this.Tcplisten = &listener
	defer (*this.Tcplisten).Close() //关闭监听的端口

	for {

		conn, err := (*this.Tcplisten).Accept() //用conn接收链接
		fmt.Println("accept.")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("SUCCESS.")
			ctx1, cal := context.WithCancel(this.Ctx)
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
