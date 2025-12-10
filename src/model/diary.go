package model

import (
	"time"

	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Note string
	UserID uint
	QuestionID uint
	Question Question `gorm:"foreignKey:QuestionID"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}