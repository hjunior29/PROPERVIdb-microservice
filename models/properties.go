package models

type Properties struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Address string `json:"address"`
	Price   float64 `json:"price"`
}