package database

import (
	"fmt"
	"github.com/chiragsamtani/template-pattern/shared/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenMysqlConn(d config.RootConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		d.GetDatabaseUsername(),
		d.GetDatabasePassword(),
		d.GetDatabaseHost(),
		d.GetDatabaseName())

	db, err := gorm.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil

}
