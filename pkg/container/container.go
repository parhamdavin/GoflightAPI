package container

import (
	"gorm.io/gorm"
)

type AppContainer struct {
	UserService   interface{}
	FlightService interface{}
	AirportService interface{}
	JWTService interface{}
	ReservationService interface{}
	PaymentService interface{}
	WalletService interface{}

}

func (a *AppContainer) RegisterModules(modules []Module, db *gorm.DB) {
	for _, m := range modules {
		m.Register(db, a)
	}
}


func NewAppContainer(db *gorm.DB, modules []Module) *AppContainer {
	container := &AppContainer{}

	for _, m := range modules {
		m.Register(db, container)
	}

	return container
}
