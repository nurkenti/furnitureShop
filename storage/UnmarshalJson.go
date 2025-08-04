package storage

import (
	"encoding/json"
	"fmt"
	"github/kaiiiman/chairStore/warehouse"
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
	s.products = make(map[int]warehouse.Product)

	for id, p := range temp {
		var product warehouse.Product
		var err error

		switch p.Type {
		case "chair":
			var chair warehouse.Chair
			err = json.Unmarshal(p.Data, &chair)
			product = &chair
		case "wardrobe":
			var wardrobe warehouse.Wardrobe
			err = json.Unmarshal(p.Data, &wardrobe)
			product = &wardrobe
		case "conditioner":
			var cond warehouse.Conditioner
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
