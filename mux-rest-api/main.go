package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/svm-code/go/mux-rest-api/docs"
	services "github.com/svm-code/go/mux-rest-api/services"
	httpSwagger "github.com/swaggo/http-swagger"
)

func hi(w http.ResponseWriter, r *http.Request) {
	dataInfo := docs.SwaggerInfo
	w.Write([]byte("Hi Suraj!" + fmt.Sprint(dataInfo)))

}

// @title          E.M.S. API's
// @version        1.0
// @description    This is a simple crud opration for employee
// @termsOfService http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  suraj.more@quickheal.com
// @license.name   Quick Heal
// @license.url    https://www.quickheal.com/useful-tools
// @host           localhost:8080
// @BasePath       /

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hi)
	router.HandleFunc("/emp/{id}", services.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/emp/{id}", services.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/emp", services.SaveEmpoyee).Methods("POST")
	router.HandleFunc("/emp", services.GetAllEmpoyee)
	router.HandleFunc("/emp/{name}", services.GetEmpoyee)

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
