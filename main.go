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

	// Add middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Add Basic Auth
	e.Use(middleware.BasicAuth(func(user, pass string, c echo.Context) (bool, error) {
		authUser := getEnvWithDefault("AUTH_USERNAME", "admin")
		authPass := getEnvWithDefault("AUTH_PASSWORD", "admin")
		if user == authUser && pass == authPass {
			return true, nil
		}
		return false, nil
	}))

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
	addr := ":" + getEnvWithDefault("PORT", "5000")
	e.Logger.Fatal(e.Start(addr))
}
