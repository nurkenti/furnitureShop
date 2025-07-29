package storage

import (
	"fmt"
	"github/kaiiiman/chairStore/warehouse"
	"log"
)

func (s *Storage) Sell(id int, quantity int) (warehouse.Chair, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if quantity <= 0 {
		return warehouse.Chair{}, fmt.Errorf("invalid quantity: must be positive, got %d", quantity)
	}

	chair, exists := s.chairs[id]
	if !exists {
		return warehouse.Chair{}, fmt.Errorf("chair with ID %d not found", id)
	}

	if chair.InStock < quantity {
		return warehouse.Chair{}, fmt.Errorf("not enough chairs in stock (requested: %d, available: %d)", quantity, chair.InStock)
	}

	chair.InStock -= quantity
	s.chairs[id] = chair

	if err := s.save(); err != nil {
		return warehouse.Chair{}, fmt.Errorf("failed to save inventory: %w", err)
	}

	log.Printf("[SALE] chair_id=%d quantity=%d new_stock=%d", id, quantity, chair.InStock)
	return chair, nil
}
