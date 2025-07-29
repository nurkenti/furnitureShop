package main

import (
	"github/kaiiiman/chairStore/menu"
)

func main() {
	Salesman()

	// 	switch menu.Role() {
	// 	case 0:
	// 		Salesman()
	// 		menu.ClearScreen()

	// 	case 1:
	// 		Client()
	// 	}
	// }

}
func Salesman() {
	menu.Doing()
}
func Client() {
	menu.MenuClient()
	menu.ClientChairModel()
	menu.CliChSum()

}
