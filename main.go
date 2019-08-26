package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"

	"github.com/checkaayush/sampark/handler"
)

func getEnvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}

	return val
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.WARN)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database connection
	db, err := mgo.Dial(getEnvWithDefault("MONGO_URI", "mongodb:27017"))
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Create database indices
	dbname := getEnvWithDefault("MONGO_DBNAME", "sampark")
	if err = db.Copy().DB(dbname).C("contacts").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Routes
	v1 := e.Group("/v1")
	v1.GET("/health", h.Health)
	v1.POST("/contacts", h.CreateContact)
	v1.GET("/contacts", h.FetchContacts)
	v1.GET("/contacts/:id", h.GetContactByID)
	v1.DELETE("/contacts/:id", h.DeleteContactByID)
	v1.PATCH("/contacts/:id", h.UpdateContactByID)

	// Start server
	addr := getEnvWithDefault("SERVER_ADDR", ":5000")
	e.Logger.Fatal(e.Start(addr))
}
