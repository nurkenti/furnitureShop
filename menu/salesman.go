package menu

import (
	"bufio"
	"fmt"
	"github/kaiiiman/chairStore/storage"
	"github/kaiiiman/chairStore/warehouse"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Salesman1() {
	fmt.Println("Вы выбрали роль Продовца), Ваша задача прибрести товар и продать их покупателям")
	fmt.Println("Хороший игры!")
}
func waitEnter(reader *bufio.Reader) {
	fmt.Print("")
	_, _ = reader.ReadString('\n') // Ожидание подтверждения } //
}

func Doing() {
	reader := bufio.NewReader(os.Stdin)
	for {
		prompt := "Выберите действие: "
		menuItems := []string{"Купить товар", "Поиск", "Удалить товар", "Продать товар", "Выход"}
		u := &UserInput{}

		CreatMenu(prompt, menuItems, u)
		fmt.Printf("\nВы выбрали: %s\n", u.option.Text)

		switch u.option.ID {
		case 0:
			BuyChair()
		case 1:
			LoadStorage("Чтобы найти товар напишите модель: ")
		case 2:
			DelChairMan()
		case 3:
			SellChair()
		case 4:
			Timeloading(1, "Завершить работу...")
			return
		}
		Timeloading(1, "")
		fmt.Println("")
		waitEnter(reader)
	}

}

func LoadStorage(s string) {
	db := storage.NewStorage("data, json")
	if err := db.Load(); err != nil {
		log.Println("Не удалось загрузить данные", err)
	}
	fmt.Println(s)
	var nameUserAns string
	fmt.Print("Данные о товаре: ")
	_, err := fmt.Scan(&nameUserAns)
	if err != nil {
		fmt.Println("Ошибка! Надо написать", err)
	}
	Timeloading(2, "Поиск...")
	// Поиск товара
	chairs, err := db.Find(map[string]interface{}{
		"name": nameUserAns,
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(chairs)
}

func BuyChair() {
	// Тут у меня загрузил склад
	db := storage.NewStorage("data, json")
	if err := db.Load(); err != nil {
		log.Println("Не удалось загрузить данные", err)
	}

	fmt.Println("Что вы хотите купить ?")
	prompt := "Выберите товар: "
	menuItems := []string{"Chair", "Wardrobe", "Conditioner"}
	u := &UserInput{}
	CreatMenu(prompt, menuItems, u)
	fmt.Println("Вы выбрали:", u.option.Text)
	// Model
	promptModel := "Выберите товар: "
	menuItemsModel := []string{"Sonyx", "Kurumi"}
	uModel := &UserInput{}
	CreatMenu(promptModel, menuItemsModel, uModel)
	fmt.Println("Вы выбрали:", uModel.option.Text)
	// Material
	promptMaterial := "Выберите материал: "
	menuItemsMaterial := []string{"wood", "metal", "plastic"}
	uMaterial := &UserInput{}
	CreatMenu(promptMaterial, menuItemsMaterial, uMaterial)
	fmt.Println("Вы выбрали:", uMaterial.option.Text)
	// Price
	promptPrice := "Прайс: "
	menuItemsPrice := []string{"5000", "15000", "30000"}
	uPrice := &UserInput{}
	CreatMenu(promptPrice, menuItemsPrice, uPrice)
	fmt.Println("Вы выбрали:", uPrice.option.Text)
	num, err := strconv.Atoi(uPrice.option.Text)
	if err != nil {
		fmt.Println("Ошибка из string в int", err)
	}

	//InStock
	fmt.Print("Количество стульев: ")
	var uInStock int
	fmt.Scan(&uInStock)

	// ID
	ids := rand.Intn(1000)

	Timeloading(3, "Идет процесс покупки...")
	// Загружаем товар
	db.AddChair(warehouse.Chair{
		Id:       ids,
		Name:     uModel.option.Text,
		Material: uMaterial.option.Text,
		Price:    num,
		InStock:  uInStock,
	})
	fmt.Println("Вы купили chair")
}

func DelChairMan() {
	db := storage.NewStorage("data, json")
	if err := db.Load(); err != nil {
		log.Println("Не удалось загрузить данные", err)
	}
	var i int
	LoadStorage("Чтобы удалить товар из базы нужно найти id c помощю модели")

	fmt.Print("ID: ")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Ошибка при вводе цифр", err)
	}
	Timeloading(2, "Процесс удаление товара...")
	if err := db.DelChair(i); err != nil {
		log.Fatal()
	}
}

func SellChair() {
	db := storage.NewStorage("data, json")
	if err := db.Load(); err != nil {
		log.Println("Не удалось загрузить данные", err)
	}
	fmt.Println("Вы хотите продать товар")
	LoadStorage("Чтобы найти Id и количество на складе введите модель: ")
	fmt.Print("Пожалуйста ведите id и количество : ")
	var idr int
	var instock int
	_, err := fmt.Scan(&idr, &instock)
	if err != nil {
		fmt.Println("Ошибка при вводе цифр", err)
	}
	Timeloading(4, "Обработка покупки...")

	soldChair, err := db.Sell(idr, instock)
	if err != nil {
		log.Fatal("Ошибка продажи ", err)
	}
	fmt.Printf("💰 Продано %d стульев модели '%s'\n", 3, soldChair.Name)
	fmt.Printf("📊 Остаток на складе: %d\n", soldChair.InStock)
}

func Timeloading(n time.Duration, s string) {
	fmt.Println(s)
	time.Sleep(n * time.Second)
}
