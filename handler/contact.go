package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/checkaayush/sampark/model"
)

// CreateContact creates a contact in Sampark
func (h *Handler) CreateContact(c echo.Context) (err error) {
	contact := &model.Contact{
		ID: bson.NewObjectId(),
	}
	if err = c.Bind(contact); err != nil {
		return
	}

	// Name and email are mandatory
	if contact.Name == "" || contact.Email == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid / missing fields"}
	}

	// Save post in database
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("sampark").C("contacts").Insert(contact); err != nil {
		if mgo.IsDup(err) {
			msg := "Contact with given email already exists"
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: msg}
		}
		return
	}
	return c.JSON(http.StatusCreated, contact)
}

// FetchContacts fetches a list of contacts from the database
func (h *Handler) FetchContacts(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	// Retrieve contacts from database
	contacts := []*model.Contact{}
	db := h.DB.Clone()
	if err = db.DB("sampark").C("contacts").
		Find(bson.M{}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&contacts); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, contacts)
}

// GetContactByID fetches a single contact by id from database
func (h *Handler) GetContactByID(c echo.Context) (err error) {
	// Check if provided id is a valid object id hex
	if ok := bson.IsObjectIdHex(c.Param("id")); !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid id",
		})
	}

	id := bson.ObjectIdHex(c.Param("id"))
	contact := &model.Contact{}

	db := h.DB.Clone()
	if err = db.DB("sampark").C("contacts").FindId(id).One(contact); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, contact)
}

// DeleteContactByID fetches a single contact by id from database
func (h *Handler) DeleteContactByID(c echo.Context) (err error) {
	// Check if provided id is a valid object id hex
	if ok := bson.IsObjectIdHex(c.Param("id")); !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid id",
		})
	}

	id := bson.ObjectIdHex(c.Param("id"))

	db := h.DB.Clone()
	if err = db.DB("sampark").C("contacts").Remove(bson.M{"_id": id}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}
	defer db.Close()

	return c.NoContent(http.StatusNoContent)
}

// UpdateContactByID fetches a single contact by id from database
func (h *Handler) UpdateContactByID(c echo.Context) (err error) {
	// Check if provided id is a valid object id hex
	if ok := bson.IsObjectIdHex(c.Param("id")); !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid id",
		})
	}
	id := bson.ObjectIdHex(c.Param("id"))

	payload := make(map[string]string)
	err = json.NewDecoder(c.Request().Body).Decode(&payload)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid / missing fields"}
	}

	// New name
	name, ok := payload["name"]
	if !ok {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid / missing fields"}
	}

	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": name}}

	// Update contact in database
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("sampark").C("contacts").Update(query, update); err != nil {
		return
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Contact updated",
	})
}
