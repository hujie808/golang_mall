//场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Bans struct {
	visitIp map[string]struct{}
}

func NewBans() *Bans {
	return &Bans{visitIp: make(map[string]struct{})}
}

//is access
func (b *Bans) is_access(access_ip string) bool {
	banMutex.Lock()
	defer banMutex.Unlock()
	if _, ok := b.visitIp[access_ip]; ok {
		return true
	}
	b.visitIp[access_ip] = struct{}{}
	go b.verification3minute(access_ip)
	return false
}

func (b *Bans) verification3minute(access_ip string) {
	time.Sleep(3 * time.Minute)
	banMutex.Lock()
	zIp := b.visitIp
	delete(zIp, access_ip)
	b.visitIp = zIp
	banMutex.Unlock()
}

var banMutex sync.Mutex

func main() {

	newBan := NewBans()
	wait := new(sync.WaitGroup)
	var succeed int64
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			ipEnd := j
			wait.Add(1)
			go func() {
				defer wait.Done()
				accessIp := fmt.Sprintf("192.168.0.%d", ipEnd)
				if !newBan.is_access(accessIp) {
					fmt.Println(accessIp)
					atomic.AddInt64(&succeed, 1)
				} else {

				}

			}()
		}
	}
	wait.Wait()
	fmt.Printf("success:%d", succeed)

}
