package main

import (
	"fmt"

	//"jwt-authentication-golang/controllers"
	"ams/controllers"
	"ams/database"

	_ "ams/docs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Asset Management System
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
//@query.collection.format multi

func main() {

	LoadAppConfig()

	database.Connect(AppConfig.ConnectionString)

	router := mux.NewRouter().StrictSlash(true)

	RegisterUserRoutes(router)

	log.Printf(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}

func RegisterUserRoutes(router *mux.Router) {

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(

		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition

		httpSwagger.DeepLinking(true),

		httpSwagger.DocExpansion("none"),

		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(

		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition

		httpSwagger.DeepLinking(true),

		httpSwagger.DocExpansion("none"),

		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodPost)

	router.HandleFunc("/api/createemployee", controllers.CreateEmployee).Methods("POST")

	router.HandleFunc("/api/getasset/{id}", controllers.GetAssetById).Methods("GET")

	router.HandleFunc("/api/getusers", controllers.GetUser).Methods("GET")

	//router.HandleFunc("/api/update/{id}", controllers.UpdateAsset).Methods("PUT")

	router.HandleFunc("/api/updates/{id}/{asset_id}", controllers.UpdateUser).Methods("PUT")

	router.HandleFunc("/api/usercreate", controllers.Users).Methods("POST")

	router.HandleFunc("/api/createassets", controllers.Asset).Methods("POST")
}
