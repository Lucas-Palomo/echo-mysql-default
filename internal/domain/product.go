package domain

import "time"

type Product struct {
	Id          int        `json:"id" param:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       int64      `json:"price"`
	Online      bool       `json:"online"`
	CreatedAt   *time.Time `json:"createdAt" db:"created_at"`
	ModifiedAt  *time.Time `json:"modifiedAt" db:"modified_at"`
}
