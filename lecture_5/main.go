package main

import "fmt"

const (
	applePrice = 5.99
	pearPrice  = 7.00
	balance    = 23.00
)

func main() {
	summaryPrice := applePrice*float64(9) + pearPrice*float64(8)
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?: %d грн.\n", int64(summaryPrice))
	count := balance / applePrice
	fmt.Printf("2. Скільки груш ми можемо купити?: %d од.\n", int64(count))
	count = balance / pearPrice
	fmt.Printf("3. Скільки яблук ми можемо купити?: %d од.\n", int64(count))
	summaryPrice = applePrice*float64(2) + pearPrice*float64(2)
	fmt.Printf("4. Чи ми можемо купити 2 груші та 2 яблука?: %v \n", balance >= summaryPrice)
}
