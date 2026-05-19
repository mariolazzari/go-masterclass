package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mtx     sync.Mutex
}

func (b *BankAccount) Deposit(val int) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	b.balance += val
}

func (b *BankAccount) WithDraw(val int) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	b.balance -= val
}

func (ba *BankAccount) Balance() int {
	ba.mtx.Lock()
	defer ba.mtx.Unlock()

	return ba.balance
}

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mtx sync.Mutex

	for range 10 {
		wg.Go(func() {
			mtx.Lock()
			counter++
			mtx.Unlock()
		})
	}

	wg.Wait()
	fmt.Println(counter)

	ba := BankAccount{
		balance: 10,
	}

	for i := range 10 {
		wg.Add(1)

		go func(val int) {
			defer wg.Done()
			if val%2 == 1 {
				ba.Deposit(val)
			} else {
				ba.WithDraw(val)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(ba.Balance())
}
