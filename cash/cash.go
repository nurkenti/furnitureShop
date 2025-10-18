package cash

import (
	"fmt"
)

type Bank struct {
	Money int
}

func MyBank(money int) Bank {
	return Bank{
		Money: money,
	}

}

func (b *Bank) AddMoney(a int) error {
	b.Money += a
	if b.Money < 0 {
		return fmt.Errorf("У вас нету денег")
	}
	fmt.Printf("Вы продали товар на сумму: %d\nНа счету %d\n", a, b)
	return nil

	return nil
}

func (b *Bank) SellMoney(a int) error {
	b.Money -= a
	if b.Money < 0 {
		fmt.Printf("Вам не хвотает: %d", b.Money)
		b.Money += a
		return fmt.Errorf("У вас нету денег")
	} else {
		b.Money += a
		fmt.Printf("У вас на счету: %d\n", b.Money)
		b.Money -= a
	}
	fmt.Printf("Вы купили товар на сумму: %d, У вас осталось %d\n", a, b)
	return nil
}
