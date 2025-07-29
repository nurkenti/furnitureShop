package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github/kaiiiman/chairStore/warehouse"
	"os"
	"sync"
)

// А это база данных которые мы сохроняем все товары

var (
	ErrNotFound = errors.New("товар не найден")
)

type Storage struct {
	chairs map[int]warehouse.Chair
	mu     sync.Mutex
	file   string
}

func NewStorage(file string) *Storage {
	return &Storage{
		chairs: make(map[int]warehouse.Chair),
		file:   file,
	}
}

func (s *Storage) AddChair(chair warehouse.Chair) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.chairs[chair.Id]; exist {
		return errors.New("cтул с таким id уже существует")
	}

	s.chairs[chair.Id] = chair
	if err := s.save(); err != nil {
		return fmt.Errorf("Не удалось добавить данные %v", err)
	}
	fmt.Printf("✅ Стул добавлен: %+v\n", chair)
	return nil

}

func (s *Storage) DelChair(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exist := s.chairs[id]; !exist {
		return fmt.Errorf("Товар %d не существует", id)
	}

	delete(s.chairs, id)

	if err := s.save(); err != nil {
		return fmt.Errorf("Ошибка при сохранений данных: %v", err)
	}
	fmt.Printf("🗑️ Стул с ID %d успешно удален\n", id)
	return nil
}

func (s *Storage) save() error {
	data, err := json.Marshal(s.chairs)
	if err != nil {
		return nil
	}
	return os.WriteFile(s.file, data, 0644)
}

func (s *Storage) Load() error {
	data, err := os.ReadFile(s.file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.chairs)
}
