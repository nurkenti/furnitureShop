package menu

import (
	"fmt"
	"log"
	"os"

	"github.com/dixonwille/wmenu"
)

type UserInput struct {
	option wmenu.Opt
}

func (u *UserInput) OptFunc(option wmenu.Opt) error {
	u.option = option
	return nil
}

func CreatMenu(p string, m []string, u *UserInput) {
	menu := wmenu.NewMenu(p)
	menu.ChangeReaderWriter(os.Stdin, os.Stdout, os.Stderr)
	for i, item := range m {
		menu.Option(item, i, false, u.OptFunc)
	}
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
func Role() int {
	prompt := "Выберите роль: "
	menuItems := []string{"Продовец", "Покупатель"}
	u := &UserInput{}
	CreatMenu(prompt, menuItems, u)
	fmt.Println("Вы выбрали:", u.option.Text)
	fmt.Println("")
	return u.option.ID
}
