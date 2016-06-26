package model

import (
	"github.com/jinzhu/gorm"
)

// common structs
type Communicator struct {
	gorm.Model
	Id     *int64  `json:"id"`
	Sector *string `json:"sector"`
}
