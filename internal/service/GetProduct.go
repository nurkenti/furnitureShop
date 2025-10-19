package service

import (
	"errors"
	"strings"

	"github.com/nurkenti/furnitureShop/internal/domain"
)

// Тут мы делаем поиск товаров
func (s *Storage) Find(filters map[string]interface{}) ([]domain.Product, error) {
	s.mu.Lock() // Блокировка для чтения
	defer s.mu.Unlock()

	var result []domain.Product

	for _, product := range s.products {
		if matchesFilter(product, filters) {
			result = append(result, product)
		}
	}
	if len(result) == 0 {
		return nil, errors.New("товар не найден")

	}
	return result, nil
}
func matchesFilter(product domain.Product, filters map[string]interface{}) bool {
	if model, ok := filters["model"].(string); ok { // тут мы проверяем мап на ключ. Если есть то идет дальше ok
		if !strings.Contains(
			strings.ToLower(product.GetModel()),
			strings.ToLower(model),
		) {
			return false
		}
	}
	if minPrice, ok := filters["min_price"].(int); ok {
		if product.GetPrice() < minPrice {
			return false
		}

	}
	if minStock, ok := filters["min_stock"].(int); ok {
		if product.GetInStock() < minStock {
			return false
		}
	}

	switch p := product.(type) {
	case *domain.Chair:
		if material, ok := filters["material"].(string); ok {
			if !strings.EqualFold(p.Material, material) {
				return false
			}
		}
	case *domain.Conditioner:
		if version, ok := filters["version"].(string); ok {
			if !strings.EqualFold(p.Version, version) {
				return false
			}
		}
	}
	return true
}
