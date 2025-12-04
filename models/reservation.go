package models

type Reservation struct {
    ID              uint            `gorm:"primaryKey"`
    UserID     uint
    User       User
    SeatID          uint            
    Seat            SeatDetail      `gorm:"foreignKey:SeatID"`
    PaymentStatuses []PaymentStatus `gorm:"foreignKey:ReservationID"`
}

