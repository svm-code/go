package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/svm-code/go/mux-rest-api/daos"
)

type response struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

var success = "SUCCESS"
var failed = "FAILED"

// var FAILED
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	body := response{
		Result: success,
		Data:   data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func ErrorResponse(w http.ResponseWriter, data interface{}) {
	ErrorResponseWithStatus(w, data, http.StatusConflict)
}

func ErrorResponseWithStatus(w http.ResponseWriter, data interface{}, statusCode int) {
	body := response{
		Result: failed,
		Data:   fmt.Sprint(data),
	}
	log.Println("Error Logs: ", data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func GetReauiestBody(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&v)
	return err
}

func GetRequiestBody(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&v)
	return err
}

func GetEmpoyeeFromRequiestModel(r *http.Request, e *daos.Employee) error {
	var emp daos.EmployeeModel
	err := GetRequiestBody(r, &emp)
	if err == nil {
		*e = daos.Employee{
			Name:    emp.Name,
			Email:   emp.Email,
			DeptId:  emp.DeptId,
			Address: emp.Address,
			Phone:   emp.Phone,
		}
	}
	return err
}
