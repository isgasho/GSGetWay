package main

//下面是一个启动的例子
import (
	gsgetway "buguang01/GSGetWay"
	"fmt"
	"time"
)

func main() {
	fmt.Println("game Service get wa1111y.")
	i := 0
	gsgetway.Instance.Init("", "127.0.0.1:8080")
	gsgetway.Instance.Start()
	for {
		select {
		case <-gsgetway.Instance.Ctx.Done():
			fmt.Println("end2.")
			return
		default:
			time.Sleep(6 * time.Second)
			fmt.Printf("sleep %d.\n", i)
			if i++; i >= 10 {
				gsgetway.Instance.Stop()
				fmt.Println("end.")
			}

		}
	}

}
