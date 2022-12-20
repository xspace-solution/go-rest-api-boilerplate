package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jsweet314/go-api-boilerplate/models"
	_ "github.com/lib/pq"
)

// databaseConfig info for connecting to the database
type databaseConfig struct {
	client   string
	user     string
	password string
	dbName   string
	host     string
	port     string
}

// dbConfig creates a DB config struct
func dbConfig() *databaseConfig {
	return &databaseConfig{
		client:   getEnv("DB_CLIENT", "postgres"),
		user:     getEnv("DB_USER", "postgres"),
		password: getEnv("DB_PASSWORD", ""),
		dbName:   getEnv("DB_NAME", "postgres"),
		host:     getEnv("DB_HOST", "localhost"),
		port:     getEnv("DB_PORT", "5432"),
	}
}

func (db *databaseConfig) getClient() string {
	return db.client
}

func (db *databaseConfig) getConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", db.user, db.password, db.dbName, db.host, db.port)
}

// ConnectToDatabase returns a db connection
func ConnectToDatabase() *sql.DB {
	db, err := sql.Open(dbConfig().getClient(), dbConfig().getConnectionString())

	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: '%v'", err))
	}

	models.SetDatabase(db)

	return db
}
