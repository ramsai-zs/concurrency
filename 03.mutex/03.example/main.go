package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var mutex sync.Mutex

	// variable for bank balance
	var bankbalance int

	// print out starting values
	fmt.Printf("Initial amount of bank balance: $%0.2d", bankbalance)

	// define weekly revenue
	incomes := []Income{
		{"main job", 500},
		{"part time", 100},
		{"gifts", 50},
		{"Investments", 100},
	}

	// loop through 52 weeks and print how much it is mad;keeping runnig in total.

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for i = 1; i <= 52; i++ {
				mutex.Lock()
				temp := bankbalance
				temp += income.Amount
				bankbalance = temp
				fmt.Printf("On Week %d,you earned $%d.00 from %s\n", i, bankbalance, income.Source)
				mutex.Unlock()
			}
		}(i, income)

	}

	wg.Wait()

	// print out the final balance.
	fmt.Printf("Final bank balance $%d.00", bankbalance)
}
