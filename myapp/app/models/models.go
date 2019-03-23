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
	Title                     string `gorm:"type:varchar(128)" sql:"not null"`
	Number                    int    `gorm:"type:INT UNSIGNED not null;"`
	Comment                   string `gorm:"type:varchar(256)" `
}

type Table2 struct {
	CommonModel
	Title                     string `gorm:"type:varchar(128)" sql:"not null"`
	Number                    int    `gorm:"type:INT UNSIGNED not null;"`
  Table1_id                 int `gorm:"type:INT UNSIGNED not null;"`
	Comment                   string `gorm:"type:varchar(256)" `
}
