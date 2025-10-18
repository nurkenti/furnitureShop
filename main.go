package main

import (
	"github.com/nurkenti/furnitureShop/menu"
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
