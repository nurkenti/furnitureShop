package storage

import (
	"fmt"
	"github/kaiiiman/chairStore/warehouse"
	"strings"
)

// Тут мы делаем поиск товаров
func (s *Storage) Find(filters map[string]interface{}) ([]warehouse.Chair, error) {
	s.mu.Lock() // Блокировка для чтения
	defer s.mu.Unlock()

	// Вспомогательные функции проверки типов
	getStringFilter := func(key string) (string, bool) {
		val, exists := filters[key]
		if !exists {
			return "", false
		}
		s, ok := val.(string)
		return s, ok
	}

	getIntFilter := func(key string) (int, bool) {
		val, exists := filters[key]
		if !exists {
			return 0, false
		}
		// Обработка разных числовых типов
		switch v := val.(type) {
		case int:
			return v, true
		case float64: // Для JSON-чисел
			return int(v), true
		default:
			return 0, false
		}
	}

	var result []warehouse.Chair
	for _, chair := range s.chairs {
		match := true

		// Фильтр по имени
		if name, ok := getStringFilter("name"); ok {
			if !strings.Contains(strings.ToLower(chair.Name), strings.ToLower(name)) {
				match = false
			}
		}

		// Фильтр по материалу
		if material, ok := getStringFilter("material"); ok {
			if !strings.EqualFold(chair.Material, material) {
				match = false
			}
		}

		// Фильтр по цене
		if minPrice, ok := getIntFilter("min_price"); ok {
			if chair.Price < minPrice {
				match = false
			}
		}

		if match {
			result = append(result, chair)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Такого товара нету в базе: %+v", filters)
	}

	return result, nil
}
