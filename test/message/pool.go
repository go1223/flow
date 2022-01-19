package message

import (
	"fmt"
	"sync"
)

var locker sync.Mutex
var chan_pool = make(map[string]*chan string, 10)

func GetRChan(id string) (c *chan string) {
	defer func() {
		fmt.Printf("getRchan %p\n", c)
	}()
	locker.Lock()
	defer locker.Unlock()
	var ok bool
	c, ok = chan_pool[id]
	if !ok {
		var newC = make(chan string)
		chan_pool[id] = &newC
		return &newC
	}
	return
}

func GetWChan(id string) (c *chan string) {
	defer func() {
		fmt.Printf("getWchan %p\n", c)
	}()
	locker.Lock()
	defer locker.Unlock()
	var ok bool
	c, ok = chan_pool[id]
	if !ok {
		var newC = make(chan string)
		chan_pool[id] = &newC
		return &newC
	}
	return
}
