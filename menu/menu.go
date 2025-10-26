package menu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dixonwille/wmenu"
)

type UserInput struct {
	option wmenu.Opt
}

func (u *UserInput) OptFunc(option wmenu.Opt) error {
	u.option = option
	return nil
}

func CreateMenu(p string, m []string, u *UserInput) {
	menu := wmenu.NewMenu(p)
	menu.ChangeReaderWriter(os.Stdin, os.Stdout, os.Stderr)
	for i, item := range m {
		menu.Option(item, i, false, u.OptFunc)
	}
	err := menu.Run()
	if err != nil {
		fmt.Println("ЩЩС")
		log.Fatal(err)
	}
}
func NewMenuTemplate(p string, items []string) (string, int, error) {
	prompt := p
	menuItems := items
	u := &UserInput{}
	CreateMenu(prompt, menuItems, u)
	fmt.Printf("Вы выбрали: %s\n", u.option.Text)
	return u.option.Text, u.option.ID, nil
}

func ClearInputBuffer() {
	file := os.Stdin
	stat, _ := file.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		// Устанавливаем очень короткий дедлайн
		file.SetReadDeadline(time.Now().Add(5 * time.Millisecond))

		scanner := bufio.NewScanner(file)
		if scanner.Scan() {
			_ = scanner.Text()
		}
		// Сбрасываем дедлайн
		file.SetReadDeadline(time.Time{})
	}
}

func Role() int {
	prompt := "Выберите роль: "
	menuItems := []string{"Продовец", "Покупатель"}
	u := &UserInput{}
	CreateMenu(prompt, menuItems, u)
	fmt.Println("Вы выбрали:", u.option.Text)
	fmt.Println("")
	return u.option.ID
}
