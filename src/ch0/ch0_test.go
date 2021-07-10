package ch0

import (
	"fmt"
	"sync"
	"testing"
)

func TestSimpleFor(t *testing.T) {
	for i := 0; i < 10*10000; i++ {
		fmt.Println(i)
	}
}

func TestMultiFor(t *testing.T) {
	var tickets chan byte
	var tokenWg sync.WaitGroup
	max_tickets := 10 * 10000
	tx_size := 10 * 10000

	tickets = make(chan byte, max_tickets)
	for i := 0; i < max_tickets; i++ {
		tickets <- 1
	}

	tokenWg.Add(tx_size)
	i := 0
	for i < tx_size {
		i = i + 1
		<-tickets
		go func() {
			defer func() {
				tokenWg.Done()
				tickets <- 1
			}()
			fmt.Println(i)
		}()
	}

	tokenWg.Wait()
}
