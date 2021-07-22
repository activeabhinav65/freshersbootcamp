package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	balance:=500
	var mutex = &sync.Mutex{}
	for r := 0; r < 100; r++ {
		go func() {
			for {
					money := rand.Intn(100)
				mutex.Lock()
				balance += money
				mutex.Unlock()
					time.Sleep(time.Millisecond)
			}
		}()
	}
	for w := 0; w < 100; w++ {
		go func() {
			for {
				money2 := rand.Intn(100)
				mutex.Lock()
				if money2>balance {

				}
				if money2<=balance {
					balance = balance - money2
				}

				mutex.Unlock()
				time.Sleep(time.Millisecond)
			}
		}()
	}
		time.Sleep(time.Second)
		mutex.Lock()
	fmt.Println("final balance is ",balance)
	mutex.Unlock()
}


