package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App encapsulates the basic application structure
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize establishes connection with the database
func (a *App) Initialize(username, password, dbname string) {

}

// Run starts the application
func (a *App) Run(addr string) {

}
