package database

import (
	"core-api/internal/domain/entities"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// createSuperAdmin creates a super admin user if it doesn't exist.
func createSuperAdmin(db *gorm.DB) {
	var count int64
	userAdminEmail := os.Getenv("USER_ADMIN_EMAIL")
	userAdminPwd := os.Getenv("USER_ADMIN_PWD")

	// Ensure environment variables are loaded.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check if the account with ID = 1 exists.
	accountID := uint(1)
	var account entities.Account
	if err := db.First(&account, accountID).Error; err != nil {
		account = entities.Account{
			ID:           accountID,
			Name:         "Admin Account",
			ContactEmail: userAdminEmail,
			ContactPhone: "1234567890",
			Plan:         "Premium",
		}
		if err := db.Create(&account).Error; err != nil {
			log.Fatalf("Failed to create admin account: %v", err)
		}
		log.Println("Admin account created successfully.")
	}

	// Check if the super admin user already exists.
	db.Model(&entities.User{}).Where("email = ?", userAdminEmail).Count(&count)
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userAdminPwd), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		superAdmin := entities.User{
			Email:     userAdminEmail,
			Password:  string(hashedPassword),
			FirstName: "Super",
			LastName:  "Admin",
			Role:      "superadmin",
			AccountID: &accountID,
		}
		if err := db.Create(&superAdmin).Error; err != nil {
			log.Fatalf("Failed to create super admin: %v", err)
		}
		log.Println("Super admin user created successfully.")
	}
}

// seedCategories inserts pre-registered categories along with their subcategories into the database if none exist.
func seedCategories(db *gorm.DB) {
	var count int64
	if err := db.Model(&entities.Category{}).Count(&count).Error; err != nil {
		log.Printf("Failed to count categories: %v", err)
		return
	}
	if count > 0 {
		log.Println("Categories already seeded.")
		return
	}

	// Predefined categories with their subcategories.
	predefinedCategories := []struct {
		Name          string
		Subcategories []string
	}{
		{"Residential Buildings", []string{"Single-family home", "Multi-family dwelling", "Apartment building", "Condominium", "Townhouse"}},
		{"Commercial Buildings", []string{"Office building", "Retail store", "Shopping mall", "Hotel", "Restaurant"}},
		{"Institutional Buildings", []string{"School", "University", "Hospital", "Government building", "Library"}},
		{"Industrial Buildings", []string{"Factory", "Warehouse", "Power plant", "Research facility", "Distribution center"}},
		{"Religious Buildings", []string{"Church", "Temple", "Mosque", "Synagogue", "Monastery"}},
		{"Cultural/Entertainment Buildings", []string{"Museum", "Theater", "Concert hall", "Cinema", "Art gallery"}},
		{"Sports/Recreational Buildings", []string{"Stadium", "Gymnasium", "Swimming pool", "Sports center", "Recreation center"}},
		{"Transportation Buildings", []string{"Airport", "Train station", "Bus terminal", "Parking structure", "Ferry terminal"}},
		{"Agricultural Buildings", []string{"Barn", "Greenhouse", "Silo", "Storage shed", "Processing facility"}},
		{"Military/Defense Buildings", []string{"Barrack", "Armory", "Base", "Training facility", "Command center"}},
	}

	// Insert main categories and their subcategories.
	for _, cat := range predefinedCategories {
		mainCat := entities.Category{
			Name: cat.Name,
		}
		if err := db.Create(&mainCat).Error; err != nil {
			log.Printf("Failed to create category '%s': %v", cat.Name, err)
			continue
		}
		log.Printf("Category '%s' created successfully.", cat.Name)
		// Create subcategories.
		for _, subName := range cat.Subcategories {
			subCat := entities.Category{
				Name:     subName,
				ParentID: &mainCat.ID,
			}
			if err := db.Create(&subCat).Error; err != nil {
				log.Printf("Failed to create subcategory '%s' for '%s': %v", subName, cat.Name, err)
			} else {
				log.Printf("Subcategory '%s' for '%s' created successfully.", subName, cat.Name)
			}
		}
	}
}

// InitDB initializes the database connection, performs migrations, and seeds initial data.
func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	fmt.Println(dbName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	fmt.Println("Connecting with DSN:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	log.Println("🔌 Connected to the database successfully.")

	// Auto-migrate only the entities that are currently used.
	err = db.AutoMigrate(
		&entities.User{},
		&entities.Account{},
		&entities.Building{},
		&entities.Project{},
		&entities.EfficiencyMeasure{},
		&entities.Contractor{},
		&entities.Coefficient{},
		&entities.Bms{},
		&entities.WeatherStation{},
		&entities.Regression{},
		&entities.Meter{},
		&entities.Target{},
		&entities.Subsidy{},
		&entities.Bill{},
		&entities.Category{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database schemas: %v", err)
	}

	// Create super admin user.
	createSuperAdmin(db)
	// Seed pre-registered categories and subcategories.
	seedCategories(db)
}
