package menu

import "fmt"

func MenuClient() {
	prompt := "Выберите товар: "
	menuItems := []string{"Chair", "Wardrobe", "Conditioner"}
	u := &UserInput{}
	CreatMenu(prompt, menuItems, u)
	fmt.Println("Вы выбрали:", u.option.Text)
}
func ClientChairModel() { // Модель стула
	prompt2 := "Какая модель вас устраевает: "
	menuItems2 := []string{"Sonyx", "Kurumi"}
	u2 := &UserInput{}
	CreatMenu(prompt2, menuItems2, u2)
	fmt.Println(u2.option.Text, " Отличный выбор!")
}
func CliChSum() { //Покупатель пишет сколько надо
	fmt.Println("Сколько вы хотите приобретить: ")
	fmt.Print("Number: ")
	var userAnswere int
	_, err := fmt.Scan(&userAnswere)

	if err != nil {
		fmt.Println("Надо написать количество!")
		var discard int
		fmt.Scan(&discard)
		fmt.Printf("%d Отлично", userAnswere)
	} else {
		fmt.Printf("Вам нужно %d количество, Отлично!", userAnswere)
		fmt.Println("")
	}
	fmt.Println("Мы сейчас проверим база данных")
}
