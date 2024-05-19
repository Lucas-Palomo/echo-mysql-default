package util

import (
	"database/sql"
	"echo-mysql-default/internal/domain"
	"github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func OpenDatabase(config *domain.Config) (*sql.DB, error) {
	mysqlConfig := mysql.Config{
		User:                 config.DatabaseUser,
		Passwd:               config.DatabasePass,
		DBName:               config.DatabaseName,
		Addr:                 config.DatabaseAddr,
		ParseTime:            true,
		MultiStatements:      true,
		InterpolateParams:    true,
		AllowNativePasswords: true,
	}

	return sql.Open("mysql", mysqlConfig.FormatDSN())

}

func InitDB(db *sql.DB, migrationsDir string) {

	migrationSrc := migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	applied, err := migrate.Exec(db, "mysql", migrationSrc, migrate.Up)

	if err != nil {
		log.Fatalf("Failed to execute migrations:\n%s", GetStackTrace(err))
	}

	log.Printf("Applied %d migrations\n", applied)

}
