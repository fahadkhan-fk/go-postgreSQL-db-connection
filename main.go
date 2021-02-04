package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres" // required
	"gorm.io/driver/postgres"                    // postgres driver needed to open the connection
	"gorm.io/gorm"                               // this is gorm latest version [ gorm-v2 ]
)

// we declare the db and err variables global so that we can use them out of this package as well when needed.
var (
	db  *gorm.DB
	err error
)

func main() {

	// dbURI is a connection string that holds your database connection. You can directly pass your credentials in it but the best way is to use the
	// environment variables while working in a production.

	// host = a host is required here, currently for your local machine it will be eg: `localhost`
	// user = enter your postgres database user's name
	// dbname = enter your database name here
	// port = enter the database port for postgres here eg: `5432`
	// sslmode = it is a security feature by postgres you can enter -> `disable`
	// password = enter your postgres user's password here.

	// you can use exact same db connection string by making same credentials of yours as below or write your own in place of them.
	dbURI := "host=localhost user=postgres dbname=golang_database port=5432 sslmode=disable password=postgres"

	// gorm.Open()  will open the database connection in `db` variable, if in any case an error occurs it will be stored in `err` variable.
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect database !!", err)
	} else {
		fmt.Println("Connection was successfull")
	}

}
