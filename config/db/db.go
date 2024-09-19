package db

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func MysqlConnect(host, username, password, db_name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		username, password, host, db_name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // This instructs GORM to not pluralize table names
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Failed to connect to the database")
	} else {
		connectFigure := figure.NewColorFigure("Connect to Mysql", "", "yellow", true)
		connectFigure.Print()
	}

	return db
}
