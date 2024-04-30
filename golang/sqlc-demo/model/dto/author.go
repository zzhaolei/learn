package dto

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type AuthorConfig struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (a *AuthorConfig) Scan(src any) error {
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, a)
	}
	return fmt.Errorf("invalid type %T", src)
}

func (a *AuthorConfig) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}
