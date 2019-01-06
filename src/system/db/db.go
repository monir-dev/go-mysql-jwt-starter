package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// "github.com/go-xorm/xorm"
)

// func ConnectXorm(host string, port string, database string, user string, pass string, options string) (db *xorm.Engine, err error) {
// 	return xorm.NewEngine("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&"+options)
// }

func Connect() (db *gorm.DB) {

	dbuser := "root"
	dbpass := ""
	dbdatabase := "node-start"
	db, err := gorm.Open("mysql", dbuser+":"+dbpass+"@/"+dbdatabase)
	if err != nil {
		panic(err)
	}
	return db
}

// func ConnectSql(database string, user string, pass string) (db *gorm.DB, err error) {
// 	return gorm.Open("mysql", user+":"+pass+"@/"+database)
// }
