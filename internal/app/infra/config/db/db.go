package db

import (
	"compra/internal/app/api/model/product_model"
	"compra/internal/app/api/model/purchase_model"
	"compra/internal/app/infra/config/configEnv"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB
var err error

func InitDB(config *configEnv.Config) *gorm.DB {

	dsn := fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", config.Mysql.Url)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}
	defer db.Close()

	if err := checkAndCreateDatabase(config); err != nil {
		log.Fatal(err)
	}

	dsn = fmt.Sprintf("%spurchases_db?charset=utf8&parseTime=True&loc=Local", config.Mysql.Url)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	db.AutoMigrate(&product_model.Product{}, &purchase_model.Purchase{})
	fmt.Println("Database and tables created successfully!")

	return db
}

func checkAndCreateDatabase(config *configEnv.Config) error {
	conn, err := sql.Open("mysql", config.Mysql.Url)
	if err != nil {
		return fmt.Errorf("Error connecting to MySQL server: %v", err)
	}
	defer conn.Close()

	var exists string
	err = conn.QueryRow("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = 'purchases_db'").Scan(&exists)

	if err == sql.ErrNoRows {
		_, err := conn.Exec("CREATE DATABASE purchases_db")
		if err != nil {
			return fmt.Errorf("Error creating database: %v", err)
		}
		fmt.Println("Database 'purchases_db' created successfully.")
	} else if err != nil {
		return fmt.Errorf("Error checking database: %v", err)
	}

	return nil
}
