package database

import (
	"fmt"
	"strconv"

	helper "github.com/horlabyc/iSub/helpers"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := helper.LoadEnv("DB_PORT")
	dbPort, err := strconv.ParseUint(p, 10, 32)
	configData := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		helper.LoadEnv("DB_HOST"),
		dbPort,
		helper.LoadEnv("DB_USER"),
		helper.LoadEnv("DB_PASSWORD"),
		helper.LoadEnv("DB_NAME"),
	)
	DB, err = gorm.Open("postgres", configData)
	if err != nil {
		fmt.Println(
			err.Error(),
		)
		panic("connection to database failed")
	}
	fmt.Println("Connection Opened to Database")
}
