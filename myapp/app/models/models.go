package models

import (
	"time"
)

type CommonModel struct {
	ID         int        `gorm:"type:int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY" json:"id"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at" sql:"not null;type:datetime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at" sql:"not null;type:datetime"`
}

type Table1 struct {
	CommonModel
	Title                     string `json:"shop_id" gorm:"type:varchar(128)" sql:"not null"`
	Number                    int    `json:"branch_id" gorm:"type:INT UNSIGNED not null;"`
	Comment                   string `json:"shop_name" gorm:"type:varchar(256)" `
}
