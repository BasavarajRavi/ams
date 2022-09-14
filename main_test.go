package main

import (
	"ams/controllers"
	"ams/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

const expected = "Expected status: %d, but got: %d"

func TestGetUser(t *testing.T) {

	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	router := mux.NewRouter().StrictSlash(true)
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/getusers", nil)

	controllers.GetUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf(expected, http.StatusOK, w.Code)
	}
}

func TestCreateEmployee(t *testing.T) {

	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	router := mux.NewRouter().StrictSlash(true)
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/createemployee", nil)

	controllers.GetUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf(expected, http.StatusOK, w.Code)
	}
}

func TestGetAssetById(t *testing.T) {

	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	router := mux.NewRouter().StrictSlash(true)
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/asset/{id}", nil)

	controllers.GetUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf(expected, http.StatusOK, w.Code)
	}
}

func TestUpdateAsset(t *testing.T) {

	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	router := mux.NewRouter().StrictSlash(true)
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/update/{id}", nil)

	controllers.GetUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf(expected, http.StatusOK, w.Code)
	}
}

func TestUserCreate(t *testing.T) {

	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	router := mux.NewRouter().StrictSlash(true)
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/usercreate", nil)

	controllers.GetUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf(expected, http.StatusOK, w.Code)
	}
}

func TestCreateAsset(t *testing.T) {

	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	router := mux.NewRouter().StrictSlash(true)
	RegisterUserRoutes(router)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/createassets", nil)

	controllers.GetUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf(expected, http.StatusOK, w.Code)
	}
}
