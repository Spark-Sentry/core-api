package database

import (
	"core-api/internal/domain/entities"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
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
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	awsRegion := os.Getenv("AWS_REGION")

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Generate an IAM auth token
	authToken, err := rdsutils.BuildAuthToken(fmt.Sprintf("%s:3306", dbHost), awsRegion, dbUser, sess.Config.Credentials)
	if err != nil {
		log.Fatalf("Failed to generate IAM auth token: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, authToken, dbHost, dbName)
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
