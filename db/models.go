package db

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// obj should be a pointer
func Get(obj any) error {
	query := DB.Find(obj)
	return query.Error
}

// model should be a pointer
func Update(model any) error {
	result := DB.Save(model)
	return result.Error
}

// model should be a pointer
func Create(model any) error {
	result := DB.Create(model)
	return result.Error
}

// model should be a pointer
func Delete(model any) error {
	result := DB.Delete(model)
	return result.Error
}
