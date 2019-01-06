package main

import (
	"Structure/src/system/app"
	"flag"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var port string
var dbhost string
var dbport string
var dbdatabase string
var dbuser string
var dbpass string
var dboptions string

func init() {
	flag.StringVar(&port, "port", "4000", "Assigning the port that the server should listen to.")
	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	if envPort := os.Getenv("PORT"); len(envPort) > 0 {
		port = envPort
	}

	if db_host := os.Getenv("DB_HOST"); len(db_host) > 0 {
		dbhost = db_host
	}
	if db_port := os.Getenv("DB_PORT"); len(db_port) > 0 {
		dbport = db_port
	}
	if db_database := os.Getenv("DB_DATABASE"); len(db_database) > 0 {
		dbdatabase = db_database
	}
	if db_user := os.Getenv("DB_USER"); len(db_user) > 0 {
		dbuser = db_user
	}
	if db_password := os.Getenv("DB_PASSWORD"); len(db_password) > 0 {
		dbpass = db_password
	}
}

func main() {
	// initialize database
	// db, err := DB.ConnectSql(dbdatabase, dbuser, dbpass)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db, err := DB.ConnectXorm(dbhost, dbport, dbdatabase, dbuser, dbpass, dboptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// initialize server
	s := app.NewServer()

	s.Init(port)
	s.Start()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
