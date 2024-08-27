package concurrency

import (
	"log"
	"sync"
	"time"

	"math/rand"
)

func condExample() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))

			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员%d 准备就绪", i)
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判被唤醒一次")
	}
	c.L.Unlock()

	log.Println("所有运动员准备就绪")
}
