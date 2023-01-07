package main

import (
	"fmt"
)

type Wallet struct {
	HolderName string
	Balance    int64
}

/*
Method
*/
func (walt *Wallet) showBalance() {
	fmt.Printf("Balance For %s\n", walt.HolderName)
	fmt.Printf("Balance: %d\n", walt.Balance)
}

func (walt *Wallet) addAmount(amount int64) {
	walt.Balance += amount
}

func (walt *Wallet) withdraw(amount int64) {
	if amount > walt.Balance {
		fmt.Println("Insufficient Balance")
	} else {
		walt.Balance -= amount
		fmt.Printf("Amount %d is withdrawed", amount)
	}
}

/*
This is Method Acting on struct
*/
func (walt *Wallet) getTaxableAmount(percent float64) float64 {
	balance := float64(walt.Balance)
	balance = balance * percent / 100
	return balance
}

func main() {

	var kunalWallet = Wallet{
		HolderName: "Kunal Singh",
		Balance:    0,
	}

	kunalWallet.addAmount(100)
	//kunalWallet.withdraw(70)
	kunalWallet.showBalance()
	fmt.Println(kunalWallet.getTaxableAmount(10))

}
