package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // used by gorm
)

// Config for the database.
type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// FKey is a foreign key descriptor.
type FKey struct {
	Model interface{}
	Args  [4]string
}

var db *gorm.DB

// DB returns the current database connection
func DB() *gorm.DB {
	return db
}

// Setup database (connection, migrations)
func Setup(config Config, tables []interface{}, fks ...FKey) error {
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=2s",
		config.Username, config.Password, config.Host, config.Port, config.Database,
	)
	var err error
	db, err = gorm.Open("mysql", uri)
	if err != nil {
		return err
	}
	db.AutoMigrate(tables...)
	for _, fk := range fks {
		a := fk.Args
		db.Model(fk.Model).AddForeignKey(a[0], a[1], a[2], a[3])
	}
	return nil
}

// Close current database connection
func Close() {
	if db != nil {
		db.Close()
	}
}
