package orm

import (
	"context"
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnToDB is conn to db
var ConnToDB *sql.DB

// func execMigrations() error {
// 	driver, e := postgres.WithInstance(ConnToDB, &postgres.Config{})
// 	if e != nil {
// 		return e
// 	}

// 	m, e := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
// 	if e != nil {
// 		return e
// 	}

// 	if e = m.Up(); e != nil && e != migrate.ErrNoChange {
// 		return e
// 	}

// 	return nil
// }

// InitDB init db, settings and tables
func InitDB(log *log.Logger) error {
	// establish connection to db
	log.Println("accessing database")
	dsn := "host=localhost user=gorm password=gorm dbname=photographer port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		return e
	}
	ConnToDB, _ = db.DB()
	log.Println("accessed")

	// make some sql settings
	log.Println("set up database configs")
	if _, e := ConnToDB.ExecContext(context.Background(), "PRAGMA foreign_keys = ON;PRAGMA case_sensitive_like = true;PRAGMA auto_vacuum = FULL;"); e != nil {
		return e
	}
	log.Println("database configured")

	// do migrations
	log.Println("making migrations")
	// if e := execMigrations(); e != nil {
	// return e
	// }
	log.Println("migrations completed")
	return nil
}
