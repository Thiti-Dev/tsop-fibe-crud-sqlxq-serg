package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Config is for configuration the pgsql
type Config struct{
	User		string
	Password 	string
	Host		string
	Port		string
	Name		string
	DisableTLS	bool
}

// OpenPgSQL is the function that will established the connection and return the active connection pgclient for further usage
func OpenPgSQL() (*sqlx.DB, error){
	//set env
	var cfg Config
	cfg.Host = `localhost`
	cfg.Name = `testgorm`
	cfg.User = `postgres`
	cfg.Password = `root`
	cfg.Port = `5432`
	cfg.DisableTLS = true
	sslmode := `disable`

	var dataSource = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, sslmode)
	return sqlx.Connect("postgres", dataSource)
}