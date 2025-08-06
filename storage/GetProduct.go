package storage

import (
	"errors"
	"github/kaiiiman/chairStore/warehouse"
	"strings"
)

// Тут мы делаем поиск товаров
func (s *Storage) Find(filters map[string]interface{}) ([]warehouse.Product, error) {
	s.mu.Lock() // Блокировка для чтения
	defer s.mu.Unlock()

	var result []warehouse.Product

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
func matchesFilter(product warehouse.Product, filters map[string]interface{}) bool {
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
	case *warehouse.Chair:
		if material, ok := filters["material"].(string); ok {
			if !strings.EqualFold(p.Material, material) {
				return false
			}
		}
	case *warehouse.Conditioner:
		if version, ok := filters["version"].(string); ok {
			if !strings.EqualFold(p.Version, version) {
				return false
			}
		}
	}
	return true
}

// 	for _, product := range s.products {
// 		match := true
// 		// смотрим только значение map продукта

// 		if modelFilter, ok := filters["model"]; ok && match {
// 			if model, ok := modelFilter.(string); ok {
// 				if !strings.Contains(
// 					strings.ToLower(product.GetModel()),
// 					strings.ToLower(model),
// 				) {
// 					match = false
// 				}
// 			}
// 		}
// 		if minPriceFilter, ok := filters["price"]; ok && match {
// 			if minPrice, err := toInt(minPriceFilter); err != nil {
// 				if product.GetPrice() < minPrice {
// 					match = false
// 				}
// 			}
// 		}
// 		if minStockFilter, ok := filters["in_stock"]; ok && match {
// 			if minStock, err := toInt(minStockFilter); err != nil {
// 				if product.GetInStock() < minStock {
// 					match = false
// 				}
// 			}
// 		}
// 		if _, ok := filters["type"]; ok && match {
// 			switch p := product.(type) {
// 			case *warehouse.Chair:
// 				if material, ok := filters["material"]; ok {
// 					if m, ok := material.(string); ok {
// 						if !strings.EqualFold(p.Material, m) {
// 							match = false
// 						}
// 					}

//					}
//				case *warehouse.Wardrobe:
//					if material, ok := filters["material"]; ok {
//						if m, ok := material.(string); ok {
//							if !strings.EqualFold(p.Material, m) {
//								match = false
//							}
//						}
//					}
//				case *warehouse.Conditioner:
//					if version, ok := filters["version"]; ok {
//						if v, ok := version.(string); ok {
//							if !strings.EqualFold(p.Version, v) {
//								match = false
//							}
//						}
//					}
//				}
//			}
//			if match {
//				result = append(result, product)
//			}
//		}
//		return result, nil
//	}
