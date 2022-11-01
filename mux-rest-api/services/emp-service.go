package services

import (
	"net/http"

	"github.com/gorilla/mux"
	daos "github.com/svm-code/go/mux-rest-api/daos"
	utils "github.com/svm-code/go/mux-rest-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get All the Empoyees
// @Summary     Get details of all employees
// @Description Get details of all employees
// @Tags        all-employes
// @Accept      json
// @Produce     json
// @Success     200 {array} daos.Employee
// @Router      /emp [get]
func GetAllEmpoyee(w http.ResponseWriter, r *http.Request) {
	allEmp := daos.Employee{}.FindAll()
	utils.SuccessResponse(w, allEmp)
}

// Get one Empoyee of Name
// @Summary     Get details of one employee of name
// @Description Get details of one employee of name
// @Tags        employee-by-name
// @Accept      json
// @Produce     json
// @Param name path string true "Emplopyee name"
// @Success     200 {object} daos.Employee
// @Router      /emp/{name} [get]
func GetEmpoyee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		utils.ErrorResponse(w, "name is missing in parameters:"+name)
		return
	}
	emp, err := daos.Employee.FindByName(daos.Employee{}, name)
	if err != nil {
		utils.ErrorResponse(w, err)
	} else {
		utils.SuccessResponse(w, emp)
	}
}

// Create new Empoyee
// @Summary     Create new employee
// @Description Create new employee
// @Tags        new-employee
// @Accept      json
// @Produce     json
// @Param employee body daos.EmployeeModel true "Create Emplopyee"
// @Success     200 {object} utils.response
// @Router      /emp [post]
func SaveEmpoyee(w http.ResponseWriter, r *http.Request) {
	emp := daos.Employee{}
	err := utils.GetEmpoyeeFromRequiestModel(r, &emp)
	if err != nil {
		utils.ErrorResponse(w, err)
		return
	}
	err = daos.Employee.Save(emp)
	if err != nil {
		utils.ErrorResponse(w, err)
	} else {
		utils.SuccessResponse(w, "Saved success fully")
	}

}

// Update Empoyee by id
// @Summary     Update Empoyee by id
// @Description Update Empoyee by id
// @Tags        update-employee
// @Accept      json
// @Produce     json
// @Param employee body daos.EmployeeModel true "Update Emplopyee"
// @Param id path string true "Emplopyee object id"
// @Success     200 {object} utils.response{}
// @Router      /emp/{id} [put]
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.ErrorResponse(w, "id is missing in parameters:"+id)
		return
	}
	emp := daos.Employee{}
	err := utils.GetEmpoyeeFromRequiestModel(r, &emp)
	if err != nil {
		utils.ErrorResponse(w, err)
		return
	}
	_id, _ := primitive.ObjectIDFromHex(id)
	emp.ID = _id
	err = emp.Update()
	if err != nil {
		utils.ErrorResponse(w, err)
	} else {
		utils.SuccessResponse(w, "Update success fully for id:"+id)
	}
}

// Delete Empoyee by id
// @Summary     Delete Empoyee by id
// @Description Delete Empoyee by id
// @Tags        delete-employee
// @Accept      json
// @Produce     json
// @Param id path string true "Emplopyee object id"
// @Success     200 {object} utils.response{}
// @Router      /emp/{id} [delete]
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.ErrorResponse(w, "id is missing in parameters:"+id)
		return
	}
	err := daos.Employee.Delete(daos.Employee{}, id)
	if err != nil {
		utils.ErrorResponse(w, err)
	} else {
		utils.SuccessResponse(w, "Deleted success fully id:"+id)
	}
}
