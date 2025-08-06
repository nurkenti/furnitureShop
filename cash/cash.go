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

// func ActionSel(a int) error {

// 	if err := Mybank.Sell(a); err != nil {
// 		fmt.Errorf("Ошибка при продаже", err)
// 		return err
// 	}

// 	return nil

// }

// // func (b *Bank) Sell(prise int) error {
// // 	Total := Mybank.money
// // 	sell, err := BayCash(Total, prise)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	Total = Bank.money

// 	fmt.Printf("Банк : %d, Сумма товара %d, Что осталось: %d\n", Total, prise, sell)
// 	return nil
// }

// func (b *Bank) Bay(prise int) error {
// 	Total := Mybank.money
// 	sell, err := SellCash(Total, prise)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("Банк : %d, Сумма товара %d, Что осталось: %d\n", Total, prise, sell)
// 	return nil
// }
// func SellCash(a, b int) (int, error) {
// 	return a + b, nil
// }
// func BayCash(a, b int) (int, error) {
// 	return a - b, nil
// }

// // func (b *Bank) save() error {
// // 	type TotalPriseSave struct {
// // 		MyBank int `json:"myBank"`
// // 	}
// // 	file, err := os.Create(b.file)
// // 	if err != nil {
// // 		return fmt.Errorf("ошибка создания файла: %w", err)
// // 	}

// // }
