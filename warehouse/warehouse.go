package warehouse

type Chair struct {
	Id       int    `json: "id"`   //01
	Name     string `json: "name"` //Sonyx, Kurumi
	Material string `json: "material"`
	InStock  int    `json:"in_stoke"`
	Price    int    `price`
}

type Wardrobe struct { //шкаф
	Id       int    `json: "id"`   //02
	Name     string `json: "name"` //Unibi, Facito
	Material string `json: "material"`
	InStoke  int    `json: "in_stock"`
	Price    int    `price`
}

type Conditioner struct {
	Id      int `json: "id"` //03
	Price   int `price`
	InStoke int `json: "in_stock"`
}
type Store struct {
	Chairs       *Chair
	Wardrobes    *Wardrobe
	Conditioners *Conditioner
	TotalSales   int
}
