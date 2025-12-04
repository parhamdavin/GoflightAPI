package main

import (
	"fmt"

	"example.com/airline-reservation/config"
	"example.com/airline-reservation/internal/airport"
	"example.com/airline-reservation/internal/flights"
	"example.com/airline-reservation/internal/payment"
	"example.com/airline-reservation/internal/reservations"
	"example.com/airline-reservation/internal/wallet"

	// "example.com/airline-reservation/internal/auth"
	"example.com/airline-reservation/internal/users"
	"github.com/labstack/echo/v4"

	"example.com/airline-reservation/pkg/container"
	"example.com/airline-reservation/pkg/db"
)

func main() {
	database := db.InitDB()
	e := echo.New()
	config.LoadEnv()
	

	modules := []container.Module{
		&users.UserModule{E: e},
		&airport.AirportModule{E:e},
		&flights.FlightModule{E:e},
		&reservations.ReservationModule{E: e},
		&payment.PaymentModule{E: e},
		&wallet.WalletModule{E: e},
		// &auth.AutModule{E:e},
	}

	app := container.NewAppContainer(database, modules)
	fmt.Println("App container initialized:", app)

	e.Logger.Fatal(e.Start(":8080"))
	

	
}