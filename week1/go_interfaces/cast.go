package main

import (
	"fmt"
)

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

// ---------------------
type Card struct {
	Balance    int
	ValidUntil string
	CardHolder string
	CVV        int
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Не хватает денег на карте")
	}
	c.Balance -= amount
	return nil
}

// ---------------------

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Не хватает денег на аккаунте")
	}
	a.Money -= amount
	return nil
}

// switch типов
func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата наличкой?")
	case *Card:
		plasticCard, ok := p.(*Card)
		if !ok {
			fmt.Println("Не удалось преобразовать к типу *Card")
		}
		fmt.Println("Вставляйте карту, ", plasticCard.CardHolder)
	default:
		fmt.Println("Что-то новое!")
	}

	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 110, CardHolder: "George"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 120, AppleID: "101010"}
	Buy(myMoney)
}
