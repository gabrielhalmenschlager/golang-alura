package models

type Pokemon struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"unique;not null"`
	Type     string `json:"type" gorm:"not null"`
	ImageURL string `json:"image_url"`
}
