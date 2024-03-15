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

		superAdmin := entities.User{
			Email:     UserAdminEmail,
			Password:  string(hashedPassword),
			FirstName: "Super",
			LastName:  "Admin",
			Role:      "superadmin",
		}
		if err := db.Create(&superAdmin).Error; err != nil {
			log.Fatalf("Failed to create super admin: %v", err)
		}
	}
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
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = db

	log.Println("🔌 Connected to the database successfully.")
	var models = []interface{}{&entities.User{}, &entities.Account{}, &entities.Building{}, &entities.Area{}, &entities.Equipment{}, &entities.System{}, &entities.Parameter{}}

	db.AutoMigrate(models...)
	createSuperAdmin(db)
}
