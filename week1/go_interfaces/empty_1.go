package main

import (
	"fmt"
	"strconv"
)

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Недостаточно денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

type Pay interface {
	Pay(int) error
}

func (w *Wallet) String() string {
	return "Кошелек в котором " + strconv.Itoa(w.Cash) + " денег."
}

//------------

func main() {
	myWallet := &Wallet{Cash: 100}
	fmt.Printf("Raw payment  : %#v\n", myWallet)
	fmt.Printf("Способ оплаты: %s\n", myWallet)
}
