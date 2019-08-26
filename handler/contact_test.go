package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/checkaayush/sampark/handler"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

var contactID string

func getMockDB() *mgo.Session {
	db, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestCreateContact(t *testing.T) {
	// Setup database
	db := getMockDB()
	if _, err := db.DB("sampark").C("contacts").RemoveAll(nil); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup Echo
	e := echo.New()
	payload := []byte(`{"name":"Jon Snow", "email":"jon@gmail.com"}`)

	req := httptest.NewRequest(http.MethodPost, "/v1/contacts", bytes.NewBuffer(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := &handler.Handler{db}

	// Assertions
	if assert.NoError(t, h.CreateContact(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var m map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &m)

		assert.Equal(t, "Jon Snow", m["name"])
		assert.Equal(t, "jon@gmail.com", m["email"])

		contactID = m["id"].(string)
	}
}

func TestGetContactByID(t *testing.T) {
	// Cleanup database
	db := getMockDB()
	defer db.Close()

	// Setup Echo
	e := echo.New()

	fmt.Println("/v1/contacts/" + contactID)
	req := httptest.NewRequest(http.MethodGet, "/v1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/contacts/:id")
	c.SetParamNames("id")
	c.SetParamValues(contactID)
	h := &handler.Handler{db}

	// Assertions
	if assert.NoError(t, h.GetContactByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var m map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &m)

		assert.Equal(t, "Jon Snow", m["name"])
		assert.Equal(t, "jon@gmail.com", m["email"])
	}
}

func TestUpdateContactByID(t *testing.T) {
	// Cleanup database
	db := getMockDB()
	defer db.Close()

	// Setup Echo
	e := echo.New()

	fmt.Println("/v1/contacts/" + contactID)
	payload := []byte(`{"name":"Robert Stark"}`)
	req := httptest.NewRequest(http.MethodPatch, "/v1", bytes.NewBuffer(payload))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/contacts/:id")
	c.SetParamNames("id")
	c.SetParamValues(contactID)
	h := &handler.Handler{db}

	// Assertions
	if assert.NoError(t, h.UpdateContactByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var m map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &m)

		assert.Equal(t, "Contact updated", m["message"])
	}
}

func TestFetchContacts(t *testing.T) {
	// Cleanup database
	db := getMockDB()
	defer db.Close()

	// Setup Echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/contacts", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := &handler.Handler{db}

	// Assertions
	if assert.NoError(t, h.FetchContacts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestFetchContactsByName(t *testing.T) {
	// Cleanup database
	db := getMockDB()
	defer db.Close()

	// Setup Echo
	e := echo.New()
	q := make(url.Values)
	q.Set("name", "Robert Stark")
	req := httptest.NewRequest(http.MethodPost, "/v1/contacts?"+q.Encode(), nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := &handler.Handler{db}

	// Assertions
	if assert.NoError(t, h.FetchContacts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var r []map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &r)

		assert.Equal(t, "Robert Stark", r[0]["name"])
	}
}

func TestFetchContactsByEmail(t *testing.T) {
	// Cleanup database
	db := getMockDB()
	defer db.Close()

	// Setup Echo
	e := echo.New()
	q := make(url.Values)
	q.Set("email", "jon@gmail.com")
	req := httptest.NewRequest(http.MethodPost, "/v1/contacts?"+q.Encode(), nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := &handler.Handler{db}

	// Assertions
	if assert.NoError(t, h.FetchContacts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var r []map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &r)

		assert.Equal(t, "jon@gmail.com", r[0]["email"])
	}
}

func TestDeleteContactByID(t *testing.T) {
	// Cleanup database
	db := getMockDB()
	defer db.Close()

	// Setup Echo
	e := echo.New()

	fmt.Println("/v1/contacts/" + contactID)
	req := httptest.NewRequest(http.MethodDelete, "/v1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/contacts/:id")
	c.SetParamNames("id")
	c.SetParamValues(contactID)
	h := &handler.Handler{db}

	// Assertions
	if assert.NoError(t, h.DeleteContactByID(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
