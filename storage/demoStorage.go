package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github/kaiiiman/chairStore/warehouse"
	"os"
	"sync"
)

type DemoStore struct {
	products map[int]warehouse.Store
	mu       sync.Mutex
	file     string
}

func NewDemoStorage(file string) *DemoStore {
	return &DemoStore{
		products: make(map[int]warehouse.Store),
		file:     file,
	}
}

func (d *DemoStore) AddDemo(product warehouse.Store) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	ch1 := product.Chairs.Id
	wd := product.Wardrobes.Id
	cd := product.Conditioners.Id

	if _, exist := d.products[ch1]; exist {
		return errors.New("товар id уже существует")
	}
	if _, exist := d.products[wd]; exist {
		return errors.New("товар id уже существует")
	}
	if _, exist := d.products[cd]; exist {
		return errors.New("товар id уже существует")
	}
	d.products[ch1] = product
	d.products[wd] = product
	d.products[cd] = product
	if err := d.save(); err != nil {
		return fmt.Errorf("Не удалось сохранить данные %v")
	}
	fmt.Printf("✅Товар добавлен: %+v\n", product)
	return nil
}

func (d *DemoStore) save() error {
	data, err := json.Marshal(d.products)
	if err != nil {
		return nil
	}
	return os.WriteFile(d.file, data, 0644)
}
