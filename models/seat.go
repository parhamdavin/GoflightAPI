package models

type SeatDetail struct {
    ID            uint        `gorm:"primaryKey"`
    SeatNumber    string
    FlightID      uint
    Flight        FlightDetail
    TravelClassID uint
    TravelClass   TravelClass
    FlightCostID  uint
    FlightCost    FlightCost
    IsReserved    bool
    ReservationID uint       `gorm:"unique"`
     // Nullable to allow unreserved seats
}
