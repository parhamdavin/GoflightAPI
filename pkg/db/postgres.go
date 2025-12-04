package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"example.com/airline-reservation/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func SeedData(db *gorm.DB) {
// 	airports := []models.Airport{
// 		{Name: "Imam Khomeini", AirportCity: "Tehran", AirportCountry: "Iran"},
// 		{Name: "Istanbul Airport", AirportCity: "Istanbul", AirportCountry: "Turkey"},
// 	}
// 	for i := range airports {
// 		db.FirstOrCreate(&airports[i], models.Airport{Name: airports[i].Name})
// 	}

// 	travelClasses := []models.TravelClass{
// 		{Name: "Economy"},
// 		{Name: "Business"},
// 		{Name: "First Class"},
// 	}
// 	for i := range travelClasses {
// 		db.FirstOrCreate(&travelClasses[i], models.TravelClass{Name: travelClasses[i].Name})
// 	}

// 	companies := []models.Company{
// 		{Name: "IranAir", LogoURL: "http://example.com/iranair.png", Country: "Iran"},
// 		{Name: "Turkish Airlines", LogoURL: "http://example.com/turkish.png", Country: "Turkey"},
// 	}
// 	for i := range companies {
// 		db.FirstOrCreate(&companies[i], models.Company{Name: companies[i].Name})
// 	}

// 	calendar := models.Calendar{Date: time.Now().AddDate(0, 0, 7)}
// 	db.FirstOrCreate(&calendar, models.Calendar{Date: calendar.Date})

// 	flightCost := models.FlightCost{Cost: 120.50, CalendarID: calendar.ID}
// 	db.FirstOrCreate(&flightCost, models.FlightCost{CalendarID: calendar.ID})

// 	services := []models.FlightService{
// 		{Name: "Meal"},
// 		{Name: "WiFi"},
// 	}
// 	for i := range services {
// 		db.FirstOrCreate(&services[i], models.FlightService{Name: services[i].Name})
// 	}

// 	flight := models.FlightDetail{
// 		FlightNumber:  "IR-456",
// 		FlightType:    "International",
// 		SourceID:      airports[0].ID,
// 		DestinationID: airports[1].ID,
// 		CalendarID:    calendar.ID,
// 		CompanyID:     companies[0].ID,
// 		FlightCostID:  flightCost.ID,
// 		Services:      services,
// 	}
// 	db.FirstOrCreate(&flight, models.FlightDetail{FlightNumber: flight.FlightNumber})

// 	seats := []models.SeatDetail{
// 		{SeatNumber: "1A", FlightID: flight.ID, TravelClassID: travelClasses[0].ID, FlightCostID: flightCost.ID},
// 		{SeatNumber: "1B", FlightID: flight.ID, TravelClassID: travelClasses[0].ID, FlightCostID: flightCost.ID},
// 	}
// 	for i := range seats {
// 		db.FirstOrCreate(&seats[i], models.SeatDetail{SeatNumber: seats[i].SeatNumber, FlightID: flight.ID})
// 	}

// 	user := models.User{
// 		Name:        "Ali Rezaei",
// 		SSID:        "1234567890",
// 		PhoneNumber: "09121234567",
// 		PassportId:  "IR123456",
// 		Email:       "ali@example.com",
// 		Password:    "hashedpassword",
// 		Agency:      false,
// 		Active:      true,
// 		Role:        "Customer",
// 	}
// 	db.FirstOrCreate(&user, models.User{Email: user.Email})

// 	wallet := models.Wallet{
// 		UserId:  user.ID,
// 		Balance: 1000.0,
// 	}
// 	db.FirstOrCreate(&wallet, models.Wallet{UserId: user.ID})

// 	transaction := models.WalletTransaction{
// 		WalletID:    wallet.ID,
// 		Amount:      -120.5,
// 		Type:        "Payment",
// 		Status:      "Completed",
// 		Description: "Flight booking",
// 		CreatedAt:   time.Now(),
// 	}
// 	db.FirstOrCreate(&transaction, models.WalletTransaction{WalletID: wallet.ID, Description: transaction.Description})

// 	reservation := models.Reservation{
// 		UserID: user.ID,
// 		SeatID: seats[0].ID,
// 	}
// 	db.FirstOrCreate(&reservation, models.Reservation{UserID: user.ID, SeatID: seats[0].ID})

// 	payment := models.PaymentStatus{
// 		ReservationID: reservation.ID,
// 		Amount:        120.5,
// 		PaymentStatus: "Completed",
// 		PaymentMethod: "Credit Card",
// 		PaymentDate:   time.Now(),
// 	}
// 	db.FirstOrCreate(&payment, models.PaymentStatus{ReservationID: reservation.ID})

// 	offering := models.ServiceOffering{
// 		FlightID:  flight.ID,
// 		ServiceID: services[0].ID,
// 		OfferedYN: true,
// 		FromDate:  time.Now(),
// 		ToDate:    time.Now().AddDate(0, 0, 30),
// 	}
// 	db.FirstOrCreate(&offering, models.ServiceOffering{FlightID: offering.FlightID, ServiceID: offering.ServiceID})

// 	log.Println("✅ Seed data inserted successfully")
// }



func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set in environment variables")
	}

	var db *gorm.DB
	var err error

	// Retry connecting to DB up to 5 times
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ Database connected!")
			break
		}
		log.Printf("🔁 Attempt %d: Failed to connect to DB, retrying in 2 seconds...\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("❌ Failed to connect to database after retries: ", err)
	}

	// Run migrations
	err = db.AutoMigrate(
		&models.Airport{},
		&models.FlightDetail{},
		&models.FlightService{},
		&models.ServiceOffering{},
		&models.TravelClass{},
		&models.SeatDetail{},
		&models.User{},
		&models.Reservation{},
		&models.PaymentStatus{},
		&models.Company{},
		&models.Calendar{},
		&models.FlightCost{},
		&models.Wallet{},
		&models.WalletTransaction{},
	)
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	// SeedData(db)
	fmt.Println("✅ Migration and seed completed successfully!")

	return db
}




