package migrations

import (
	"database/sql"
	"log"
	"path/filepath"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func NewPgMigrator(db *sql.DB) (*migrate.Migrate, error) {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("failed to get path")
	}

	sourceUrl := "file://" + filepath.Dir(path)

	log.Println(sourceUrl)

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatalf("failed to create migrator driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(sourceUrl, "mysql", driver)

	return m, err
}
