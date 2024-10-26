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

// createSuperAdmin  creates a super admin user if it doesn't exist
func createSuperAdmin(db *gorm.DB) {
	var count int64
	var UserAdminEmail string = os.Getenv("USER_ADMIN_EMAIL")
	var UserAdminPwd string = os.Getenv("USER_ADMIN_PWD")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Model(&entities.User{}).Where("email = ?", UserAdminEmail).Count(&count)
	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserAdminPwd), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		accountID := uint(1)

		superAdmin := entities.User{
			Email:     UserAdminEmail,
			Password:  string(hashedPassword),
			FirstName: "Super",
			LastName:  "Admin",
			Role:      "superadmin",
			AccountID: &accountID,
		}
		if err := db.Create(&superAdmin).Error; err != nil {
			log.Fatalf("Failed to create super admin: %v", err)
		}
	}
	log.Println("Create SuperUser successfully.")

}

func InitDB() {
	err := godotenv.Load()
	if err != nil {
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

	if err := db.AutoMigrate(&entities.User{}, &entities.Account{}, &entities.Building{}, &entities.Area{}, &entities.Equipment{}, &entities.System{}); err != nil {
		log.Fatalf("Failed to auto-migrate database schemas: %v", err)
	}

	createSuperAdmin(db)
}
