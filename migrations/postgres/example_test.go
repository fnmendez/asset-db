package postgres_test

import (
	"fmt"

	"github.com/owasp-amass/asset-db/migrations/postgres"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ExampleMigrations() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, _ := db.DB()

	migrationsSource := migrate.EmbedFileSystemMigrationSource{
		FileSystem: postgres.Migrations(),
		Root:       "/",
	}

	_, err = migrate.Exec(sqlDb, "postgres", migrationsSource, migrate.Up)
	if err != nil {
		panic(err)
	}

	tables := []string{"assets", "relations"}
	for _, table := range tables {
		fmt.Println(db.Migrator().HasTable(table))
	}

	// Output:
	// true
	// true
}
