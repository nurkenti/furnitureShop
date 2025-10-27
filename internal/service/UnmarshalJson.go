package service

import (
	"encoding/json"
	"fmt"

	"github.com/nurkenti/furnitureShop/internal/domain"
)

func (s *Storage) UnmarshalJSON(data []byte) error {

	type productJson struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	var temp map[int]productJson
	if err := json.Unmarshal(data, &temp); err != nil {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.products = make(map[int]domain.Product)

	for id, p := range temp {
		var product domain.Product
		var err error

		switch p.Type {
		case "chair":
			var chair domain.Chair
			err = json.Unmarshal(p.Data, &chair)
			product = &chair
		case "wardrobe":
			var wardrobe domain.Wardrobe
			err = json.Unmarshal(p.Data, &wardrobe)
			product = &wardrobe
		case "conditioner":
			var cond domain.Conditioner
			err = json.Unmarshal(p.Data, &cond)
			product = &cond
		default:
			return fmt.Errorf("Продукт не известен %s", p.Type)
		}

		if err != nil {
			return err
		}
		s.products[id] = product
	}
	return nil
}
