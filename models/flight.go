package models

import (
	
)

type FlightDetail struct {
    ID            uint           `gorm:"primaryKey"`
    FlightNumber  string
    FlightType    string
    SourceID      uint
    DestinationID uint
    SeatCount uint
    Source        Airport        `gorm:"foreignKey:SourceID"`
    Destination   Airport        `gorm:"foreignKey:DestinationID"`
    CalendarID    uint
    CompanyID     uint
    Company       Company `gorm:"foreignKey:CompanyID"`
    Calendar      Calendar
    Seats         []SeatDetail   `gorm:"foreignKey:FlightID"`
    Services      []FlightService `gorm:"many2many:service_offerings;"`
    FlightCostID  uint            `gorm:"foreignKey:FlightCostID"`  
    FlightCost    FlightCost     
}

type FlightService struct {
    ID       uint           `gorm:"primaryKey"`
    Name     string
    Flights  []FlightDetail `gorm:"many2many:service_offerings;"`
}

type FlightCost struct {
    ID         uint     `gorm:"primaryKey"`
    Cost       float64
    CalendarID uint
    Calendar   Calendar
    SeatID     uint
}

