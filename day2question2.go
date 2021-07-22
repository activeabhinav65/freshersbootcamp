package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var value=int(0)
func calc(wg *sync.WaitGroup){
	defer wg.Done()
	value+=(rand.Intn(10))

	s:=rand.Intn(5)
	time.Sleep(time.Duration(s)*time.Millisecond)
}
func main() {
	var wg sync.WaitGroup
	for i:=0;i<200;i++{
		wg.Add(1)
		go calc(&wg)
	}
	wg.Wait()
	value=value/200
	fmt.Println(value)


}

