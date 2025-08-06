package storage

import (
	"fmt"
	"strconv"
)

func ToInt(val interface{}) (int, error) {
	switch v := val.(type) {
	case int:
		return v, nil
	case float32:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("cannot convert %T to int", val)
	}
}
