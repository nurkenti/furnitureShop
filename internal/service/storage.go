package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/nurkenti/furnitureShop/internal/domain"
)

// А это база данных которые мы сохроняем все товары

var (
	ErrNotFound = errors.New("товар не найден")
)

type Storage struct {
	products map[int]domain.Product
	mu       sync.RWMutex
	file     string
}

func NewStorage(file string) *Storage {
	return &Storage{
		products: make(map[int]domain.Product),
		file:     file,
	}
}

func (s *Storage) AddProduct(product domain.Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.products[product.GetID()]; exist {
		return errors.New("Товар с таким id уже существует")
	} // exist мы проверям наличие ключа в map

	s.products[product.GetID()] = product // вызывает метод save и он возвращает ошибку. Если ошибка не нил  то хана
	fmt.Printf("✅ Товар добавлен: %+v\n", product)
	return s.save()

}
func (s *Storage) DelProduct(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exist := s.products[id]; !exist {
		return fmt.Errorf("Товар %d не существует", id)
	}

	delete(s.products, id)

	if err := s.save(); err != nil {
		return fmt.Errorf("Ошибка при сохранений данных: %v", err)
	}
	fmt.Printf("🗑️ Товар с ID %d успешно удален\n", id)
	return nil
}

func (s *Storage) save() error {
	type productSave struct {
		Type string         `json:"type"`
		Data domain.Product `json:"data"`
	}

	toSave := make(map[int]productSave)
	for id, p := range s.products {
		// Определяем тип товара
		var typeStr string
		switch p.(type) {
		case *domain.Chair:
			typeStr = "chair"
		case *domain.Wardrobe:
			typeStr = "wardrobe"
		case *domain.Conditioner:
			typeStr = "conditioner"
		default:
			return fmt.Errorf("неизвестный тип товара при сохранении: %T", p)
		}
		toSave[id] = productSave{
			Type: typeStr,
			Data: p,
		}
	}

	file, err := os.Create(s.file)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(toSave)
}
func (s *Storage) Load() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := os.ReadFile(s.file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Файла нет - это не ошибка
		}
		return err
	}

	if len(data) == 0 {
		return nil // Пустой файл - не ошибка
	}

	type productLoad struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	var loaded map[int]productLoad
	if err := json.Unmarshal(data, &loaded); err != nil {
		return fmt.Errorf("ошибка с разборам файла %w", err)
	}
	s.products = make(map[int]domain.Product)
	for id, item := range loaded {
		var p domain.Product

		switch item.Type {
		case "chair":
			p = &domain.Chair{}
		case "wardrobe":
			p = &domain.Wardrobe{}
		case "conditioner":
			p = &domain.Conditioner{}
		default:
			return fmt.Errorf("неизвестный тип товара при загрузке: %s", item.Type)
		}
		if err := json.Unmarshal(item.Data, p); err != nil {
			return fmt.Errorf("ошибка разбора товара %d: %w", id, err)
		}
		s.products[id] = p
	}

	return nil
}
