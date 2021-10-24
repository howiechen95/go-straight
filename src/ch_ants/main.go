package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	a := 0
	a++

	//n := i.(int32)
	//atomic.AddInt32(&sum, n)
	//fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	a := 0
	a++
	//time.Sleep(1 * time.Millisecond)
	//fmt.Println("Hello World!")
}

func main() {
	st := time.Now()
	defer func() {
		fmt.Println("cost:", time.Now().Sub(st).Milliseconds())
	}()
	defer ants.Release()

	runTimes := 10000000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	//p, _ := ants.NewPoolWithFunc(100000, func(i interface{}) {
	//	myFunc(i)
	//	wg.Done()
	//})
	//defer p.Release()
	//// Submit tasks one by one.
	//for i := 0; i < runTimes; i++ {
	//	wg.Add(1)
	//	_ = p.Invoke(int32(i))
	//}
	//wg.Wait()
	//fmt.Printf("running goroutines: %d\n", p.Running())
	//fmt.Printf("finish all tasks, result is %d\n", sum)
}
