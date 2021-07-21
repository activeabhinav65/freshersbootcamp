package main
import (
	"fmt"
	"math/rand"
)
type Employee interface {
	calsal(timespan int) int
}
type Full struct {
	basic int
}
type Con struct {
	basic int
}
type Free struct {
	basic int
}
func (e Full) calsal(days int) int {
	return days * e.basic
}
func (e Con) calsal(days int) int {
	return days * e.basic
}
func (e Free) calsal(hours int) int {
	if hours < 20 {
		return 0
	}
	return hours * e.basic
}
func main(){
	f1 := Full{basic:500}
	c1 := Con{basic: 100}
	fr1 := Free{basic: 10}
	data := [3]Employee{f1, c1, fr1}
	for idx, elem := range data {
		if idx == 2{
			fmt.Println(elem.calsal(rand.Intn(200)))
		}else {
			fmt.Println(elem.calsal(rand.Intn(28)))
		}
	}
}