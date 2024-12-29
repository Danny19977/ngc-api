package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Danny19977/ngc-api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBSQL *sql.DB

// Connect initializes the database connection
func Connect() error {
	p := utils.Env("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return fmt.Errorf("failed to parse database port: %v", err)
	}

	DNS := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		utils.Env("DB_HOST"), port, utils.Env("DB_USER"), utils.Env("DB_PASSWORD"), utils.Env("DB_NAME"))

	connection, err := gorm.Open(postgres.Open(DNS), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return fmt.Errorf("could not connect to the database with GORM: %v", err)
	}

	DB = connection
	fmt.Println("Database Connected (GORM) 🎉!")

	// Connection via SQL
	dbSQL, err := sql.Open("postgres", DNS)
	if err != nil {
		return fmt.Errorf("could not connect to the database with SQL driver: %v", err)
	}
	DBSQL = dbSQL
	fmt.Println("Database Connected (SQL) 🎉!")

	return nil
}
