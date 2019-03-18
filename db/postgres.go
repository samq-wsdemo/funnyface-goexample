// filename postgres.go
// author Yoshiya Ito <myon53@gmail.com>
// date 2019-03-13
package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
	gorm.NowFunc = func() time.Time {
		return time.Now().Truncate(time.Second)
	}
}

// Database Setup
// !!! You have to call this function after viper setting file
func SetupDB() {
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")
	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	database := viper.GetString("postgres.database")
	sslmode := viper.GetString("postgres.sslmode")
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, database, password, sslmode,
	)
	fmt.Println(connection)
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	DB = db
}
