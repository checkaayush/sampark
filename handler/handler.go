package handler

import (
	"gopkg.in/mgo.v2"
)

// Handler encapsulates database connection for API handlers
type Handler struct {
	DB *mgo.Session
}
