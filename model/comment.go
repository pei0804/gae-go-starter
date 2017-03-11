package model

import "time"

type Comments struct {
	ID        int        `gorm:"AUTO_INCREMENT" json:"id"`
	Name      string     `gorm:"type:varchar(30)" json:"name"`
	Comment   string     `gorm:"type:text" json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
