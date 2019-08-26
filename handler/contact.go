package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"github.com/checkaayush/sampark/model"
)

func (h *Handler) CreateContact(c echo.Context) (err error) {
	contact := &model.Contact{
		ID: bson.NewObjectId(),
	}
	if err = c.Bind(contact); err != nil {
		return
	}

	// Validation
	if contact.Name == "" || contact.Email == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid / missing fields"}
	}

	// Find user from database
	db := h.DB.Clone()
	defer db.Close()
	// if err = db.DB("twitter").C("users").FindId(u.ID).One(u); err != nil {
	// 	if err == mgo.ErrNotFound {
	// 		return echo.ErrNotFound
	// 	}
	// 	return
	// }

	// Save post in database
	if err = db.DB("sampark").C("contacts").Insert(contact); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, contact)
}

func (h *Handler) FetchContacts(c echo.Context) (err error) {
	// userID := userIDFromToken(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	// Retrieve posts from database
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
