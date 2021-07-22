package main

import (
	"fmt"
	"sync"
)
var alphabets = make(map[string]int)
func count(s string,wg *sync.WaitGroup){
	a:=len(s)
	defer wg.Done()
	for i:=0;i<a;i++ {
		x:=string(s[i])
		alphabets[x]++
	}
}
func main(){
	s1:="quick"
	s2:="brown"
	s3:="fox"
	s4:="lazy"
	s5:="dog"
	var wg sync.WaitGroup
	wg.Add(1)
	 go count(s1,&wg)
	wg.Add(1)
	 go count(s2,&wg)
	wg.Add(1)
	 go count(s3,&wg)
	wg.Add(1)
	 go count(s4,&wg)
	wg.Add(1)
	 go count(s5,&wg)
     wg.Wait()
	for k, _ := range alphabets {
		fmt.Println(k, alphabets[k])
	}




}
