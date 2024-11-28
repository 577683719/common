package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"sync"
	"time"
)

func main() {
	seconds := CalculateDurationUntilTomorrow()
	fmt.Println(seconds)
}

// 获取明天0点到现在的时间Duration
func CalculateDurationUntilTomorrow() time.Duration {
	// 获取当前时间
	now := time.Now()
	return datetime.BeginOfDay(datetime.AddDay(time.Now(), 1)).Sub(now)
}

type MyStruct struct {
	num int
	wg  *sync.WaitGroup
}

func NewMyStruct(num int, wg *sync.WaitGroup) MyStruct {
	myStruct := MyStruct{
		num: num,
		wg:  wg,
	}
	wg.Add(1)
	return myStruct
}

func (s MyStruct) Run() {
	defer s.wg.Done()
	val := s.num
	fact := s.num
	for i := s.num - 1; i > 0; i-- {
		fact *= i
	}

	fmt.Printf("Factorial of %d: %d\n", val, fact)
}
