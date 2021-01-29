package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/naumyegor/taxi-service/internal/config"
	"log"
)

func Migrate(env *config.Environment) {
	driver, err := postgres.WithInstance(env.DB, &postgres.Config{})
	if err != nil {
		env.Logger.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations/", "postgres", driver)
	if err != nil {
		env.Logger.Fatalln(err)
	}

	err = m.Steps(1)
	if err != nil {
		env.Logger.Fatalln(err)
	}

	log.Println("Migrations completed successfully.")

}
