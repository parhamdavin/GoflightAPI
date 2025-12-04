package models

type TravelClass struct {
    ID       uint         `gorm:"primaryKey"`
    Name     string
    Seats    []SeatDetail `gorm:"foreignKey:TravelClassID"`
}