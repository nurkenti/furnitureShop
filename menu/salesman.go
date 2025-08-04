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
	"strings"
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
			DelProductM()
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
	db := storage.NewStorage("data.json")
	if err := db.Load(); err != nil {
		log.Fatal("Не удалось загрузить данные", err)
	}

	// fmt.Println(s)
	// var modelAns string
	// _, err := fmt.Scan(&modelAns)
	// if err != nil {
	// 	log.Fatal("Ошибка! Надо написать", err)
	// }

	var typeAns string
	fmt.Print("Напишите тип товара :")
	_, errs := fmt.Scan(&typeAns)
	if errs != nil {
		log.Fatal("Ошибка Надо написать ", errs)
	}

	searchProduct := make(map[string]interface{})

	if typeAns != "" {
		switch typeAns {
		case "chair", "wardrobe", "conditioner":
			searchProduct["type"] = typeAns
		default:
			fmt.Println("Ошибка: недопустимый тип товара")
			return
		}
	}
	Timeloading(2, "Поиск...")
	// Поиск товара
	productCheck, err := db.Find(searchProduct)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(searchProduct) == 0 {
		fmt.Println("Товар не найден")
		return
	}

	fmt.Println("\nНайденный товары: ")
	for i, p := range productCheck {
		{
			fmt.Printf("%d.ID: %d, model: %s, price: %d, instock: %d\n",
				i+1,
				p.GetID(),
				p.GetModel(),
				p.GetPrice(),
				p.GetInStock())
		}
	}
}

func BuyChair() {
	// Тут у меня загрузил склад
	db := storage.NewStorage("data.json")
	if err := db.Load(); err != nil {
		log.Fatal("Не удалось загрузить данные", err)
	}

	fmt.Println("Что вы хотите купить ?")
	prompt := "Выберите товар: "
	menuItems := []string{"Chair", "Wardrobe", "Conditioner"}
	u := &UserInput{}
	CreatMenu(prompt, menuItems, u)
	// numProd, err := strconv.Atoi(u.option.Text)
	// if err != nil {
	// 	fmt.Println("Ошибка из string в int", err)
	// }
	switch u.option.ID {
	case 0:
		addChair(db)
	case 1:
		addWardrobe(db)
	case 2:
		addConditioner(db)

	}
	fmt.Println("Товар успешно добавлен и сохранен!")
}
func addChair(db *storage.Storage) {
	i := ID()
	m := Model("Sonyx", "Kurumi")
	ma := Material("wood", "metal", "plastic")
	pr := Price("5000", "10000", "20000")
	in := Instock("стульев")
	chair := &warehouse.Chair{
		BaseProduct: warehouse.BaseProduct{
			ID:      i,
			Model:   m,
			Price:   pr,
			InStock: in,
		},
		Material: ma,
		Type:     "chair",
	}
	Timeloading(3, "Идет процесс покупки...")
	if err := db.AddProduct(chair); err != nil {
		log.Fatal("Ошибка добавление товара", err)
	}
}

func addWardrobe(db *storage.Storage) {
	i := ID()
	m := Model("Unibi", "Facito")
	mat := Material("wood", "metal", "bamboo")
	p := Price("20000", "50000", "100000")
	in := Instock("шкафа")
	wardrobe := &warehouse.Wardrobe{
		BaseProduct: warehouse.BaseProduct{
			ID:    i,
			Model: m,

			Price:   p,
			InStock: in,
		},
		Material: mat,
		Type:     "wardrobe",
	}
	Timeloading(3, "Идет процесс покупки...")
	if err := db.AddProduct(wardrobe); err != nil {
		log.Fatal("Ошибка добавление товара", err)
	}
}
func addConditioner(db *storage.Storage) {
	i := ID()
	m := Model("Xpx", "Faca")
	mat := Version()
	p := Price("40000", "150000", "620000")
	in := Instock("кондиционеров")
	conditioner := &warehouse.Conditioner{
		BaseProduct: warehouse.BaseProduct{
			ID:      i,
			Model:   m,
			Price:   p,
			InStock: in,
		},
		Version: mat,
		Type:    "conditioner",
	}
	Timeloading(3, "Идет процесс покупки...")
	if err := db.AddProduct(conditioner); err != nil {
		log.Fatal("Ошибка добавление товара", err)
	}
}

func ID() int {
	ids := rand.Intn(1000)
	return ids
}

func Model(a, b string) string {
	promptModel := "Выберите модель: "
	menuItemsModel := []string{a, b}
	uModel := &UserInput{}
	CreatMenu(promptModel, menuItemsModel, uModel)
	fmt.Println("Вы выбрали:", uModel.option.Text)
	return uModel.option.Text
}

func Price(a, b, c string) int {
	promptPrice := "Прайс: "
	menuItemsPrice := []string{a, b, c}
	uPrice := &UserInput{}
	CreatMenu(promptPrice, menuItemsPrice, uPrice)
	fmt.Println("Вы выбрали:", uPrice.option.Text)
	num, err := strconv.Atoi(uPrice.option.Text)
	if err != nil {
		fmt.Println("Ошибка из string в int", err)
	}
	return num
}
func Material(a, b, c string) string {

	// Material
	promptMaterial := "Выберите материал: "
	menuItemsMaterial := []string{a, b, c}
	uMaterial := &UserInput{}
	CreatMenu(promptMaterial, menuItemsMaterial, uMaterial)
	fmt.Println("Вы выбрали:", uMaterial.option.Text)
	return uMaterial.option.Text
}
func Version() string {
	promptVersion := "Выберите Версию: "
	menuItemsVersion := []string{"super01", "cco2", "312ultra"}
	uVersion := &UserInput{}
	CreatMenu(promptVersion, menuItemsVersion, uVersion)
	fmt.Println("Вы выбрали:", uVersion.option.Text)
	return uVersion.option.Text

}
func Instock(a string) int {
	//InStock
	fmt.Printf("Количество %s: ", a)
	var uInStock int
	fmt.Scan(&uInStock)
	return uInStock

}

func DelProductM() {
	db := storage.NewStorage("data.json")
	if err := db.Load(); err != nil {
		log.Println("Не удалось загрузить данные", err)
	}

	fmt.Println("Чтобы удалить товар надо сначала найти его id")
	LoadStorage("Напишите тип товара")

	fmt.Println("Введите id чтобы удалить товар: ")
	var id int
	if _, err := fmt.Scan(&id); err != nil {
		log.Fatal("Ошибка при вводе ID")
	}
	// 5. Подтверждение удаления
	fmt.Printf("Вы уверены, что хотите удалить товар с ID %d? (y/n): ", id)
	var confirm string
	fmt.Scan(&confirm)
	if strings.ToLower(confirm) != "y" {
		fmt.Println("Удаление отменено")
		return
	}

	if err := db.DelProduct(id); err != nil {
		log.Fatal("Ошибки при удаление товара", err)
	}
	fmt.Println("Товар успешно удален")
}

func SellChair() {
	db := storage.NewStorage("data.json")
	if err := db.Load(); err != nil {
		log.Println("Не удалось загрузить данные", err)
	}
	fmt.Println("Вы хотите продать товар")
	LoadStorage("Чтобы найти Id и количество на складе введите модель: ")
	fmt.Print("Пожалуйста ведите id: ")
	var idr int
	var instock int
	_, err := fmt.Scan(&idr)
	fmt.Print("Пожалуйста ведите количество: ")
	_, errs := fmt.Scan(&instock)
	if err != nil {
		fmt.Println("Ошибка при вводе цифр", err)
	}
	if errs != nil {
		fmt.Println("Ошибка при вводе цифр", err)
	}
	Timeloading(4, "Обработка покупки...")

	soldChair, err := db.Sell(idr, instock)
	if err != nil {
		log.Fatal("Ошибка продажи ", err)
	}
	fmt.Printf("💰 Продано %d стульев модели '%s'\n", 3, soldChair.GetModel())
	fmt.Printf("📊 Остаток на складе: %d\n", soldChair.GetInStock())
}

func Timeloading(n time.Duration, s string) {
	fmt.Println(s)
	time.Sleep(n * time.Second)
}
