package message

import (
	"fmt"
	"sync"
	"time"
)

type cc struct {
	c      *chan string
	status bool //运行，结束
	exit   bool //是否需要退出
	id     string
}

func (c cc) Run() {
	for i := 0; ; i++ {
		fmt.Println("message", i)
		if c.exit {
			fmt.Println(c.id, "exit")
			return
		}
		*c.c <- fmt.Sprintf("%d", i)
		time.Sleep(time.Second * 1)
	}
}

var p_locker sync.Mutex
var products = make(map[string]*cc, 10)

//注册消息生产线
func Register(id string) {
	p_locker.Lock()
	defer p_locker.Unlock()
	p := cc{exit: false, c: GetWChan(id), id: id}
	go p.Run()
	products[id] = &p //cc{exit: false,c: GetWChan(id),id: id}
}
