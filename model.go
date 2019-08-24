package main

import (
	"database/sql"
	"errors"
	"time"
)

type contact struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *contact) getContact(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (c *contact) updateContact(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (c *contact) deleteContact(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (c *contact) createContact(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getContacts(db *sql.DB, start, count int) ([]contact, error) {
	return nil, errors.New("Not implemented")
}
