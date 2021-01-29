package config

import (
	"database/sql"
	"log"
	"os"
)

type Environment struct {
	DB     *sql.DB
	Logger *log.Logger
}

func NewEnvironment() *Environment {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	db, err := sql.Open("postgres", "postgres://taxi:taxi@localhost/taxi?sslmode=disable")
	if err != nil {
		logger.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		logger.Fatalln(err)
	}

	logger.Println("Connection with DB established.")
	return &Environment{
		DB:     db,
		Logger: logger,
	}
}

func (env *Environment) GetDB() *sql.DB {
	return env.DB
}

func (env *Environment) GetLogger() *log.Logger {
	return env.Logger
}
