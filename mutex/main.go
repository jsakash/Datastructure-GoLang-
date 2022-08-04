package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance %d\n", amount, balance)
	balance += amount
	mutex.Unlock()
	wg.Done()
}
func withdraw(amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance %d\n", amount, balance)
	balance -= amount
	mutex.Unlock()
	wg.Done()
}
func main() {

	balance = 1000

	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	fmt.Printf("Current balance %d\n", balance)
}
