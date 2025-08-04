package storage

import (
	"fmt"
	"github/kaiiiman/chairStore/warehouse"
	"log"
)

func (s *Storage) Sell(id int, quantity int) (warehouse.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if quantity <= 0 {
		return nil, fmt.Errorf("invalid quantity: must be positive, got %d", quantity)
	}

	product, exists := s.products[id]
	if !exists {
		return nil, fmt.Errorf("product with ID %d not found", id)
	}

	productCopy := product.Clone()

	if err := product.ReduceStock(quantity); err != nil {
		return nil, fmt.Errorf("sale failed: %w", err)
	}
	if err := s.save(); err != nil {
		product.ReduceStock(-quantity)
		return nil, fmt.Errorf("failed to save inventory: %w", err)
	}
	log.Printf("[УСПЕШНАЯ ПРОДАЖА] ID товара: %d, Модель: %s, Продано: %d, Остаток: %d",
		id,
		product.GetModel(),
		quantity,
		product.GetInStock())

	return productCopy, nil
}
