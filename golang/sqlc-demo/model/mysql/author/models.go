// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package author

import (
	dto "sqlc-demo/model/dto"
)

type Author struct {
	ID     int64             `json:"id"`
	Name   string            `json:"name"`
	Bio    *string           `json:"bio"`
	Config *dto.AuthorConfig `json:"config"`
}