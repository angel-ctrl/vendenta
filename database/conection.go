package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	//Postgres Driver imported
	"github.com/vendenta/models"
	_ "github.com/lib/pq"
	config "github.com/spf13/viper"
)

var DB = ConnectDB()

func ConnectDB() *gorm.DB {
	config.SetConfigFile("postgres.env")
	config.ReadInConfig()

	//Connect to DB
	var DB *gorm.DB
	DB, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Get("host"),
		config.GetInt("port"),
		config.Get("user"),
		config.Get("password"),
		config.Get("database")))

	fmt.Println("Connected to DB")

	if err != nil {
		log.Fatalf("Error in connect the DB %v", err)
		return nil
	}
	if err := DB.DB().Ping(); err != nil {
		log.Fatalln("Error in make ping the DB " + err.Error())
		return nil
	}
	if DB.Error != nil {
		log.Fatalln("Any Error in connect the DB " + err.Error())
		return nil
	}
	log.Println("DB connected")

	return DB
}

func PingDataBase() int {
	if err := DB.DB().Ping(); err != nil {
		return 0
	}
	return 1
}

func AutoMigrateDB() error {

	err := DB.AutoMigrate(&models.Account{}, 
		&models.ProfileAccount{},
		&models.Quiz{},
		&models.Questions{},
		&models.Answerds{},
		&models.UserExample{},
		&models.EmailExample{},
	)

	DB.Model(&models.Questions{}).AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Answerds{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")

	return err.Error
}
