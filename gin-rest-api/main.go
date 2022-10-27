package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/svm-code/go/rest-api/docs"
	swaggofiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Employee struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	DeptId  int    `json:"deptId"`
	Phone   int64  `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

var employees = []Employee{
	Employee{
		Id:      "1",
		Name:    "Suraj",
		DeptId:  1,
		Phone:   9527957483,
		Email:   "suraj.more@github.com",
		Address: "washim",
	},
	Employee{
		Id:      "2",
		Name:    "Sanket",
		DeptId:  2,
		Phone:   9899799697,
		Email:   "sanket.kasar@github.com",
		Address: "Sangamner",
	},
}

// @BasePath /
// Get all employess
// @Summary It returns the list of employee
// @Schemes http
// @Description list of all employee
// @Tags List Of All Employee
// @Accept json
// @Produce json
// @Success 200 {object} []Employee
// @Router /emp [get]
func getEmoployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

// @BasePath /
// Get Employee by id
// @Summary It returns the list of employee
// @Schemes http
// @Description It will get employee detais for ID
// @Tags Get Employee By Id
// @Accept json
// @Produce json
// @Success 200 {object} Employee
// @Router /emp/{id} [get]
func getEmoployee(c *gin.Context) {
	id := c.Param("id")
	employee := Employee{}
	for _, emp := range employees {
		if emp.Id == id {
			employee = emp
		}
	}
	c.IndentedJSON(http.StatusOK, employee)
}

func main() {
	router := gin.Default()
	router.GET("/", HealthCheck)
	router.GET("/emp", getEmoployees)
	router.GET("/emp/:id", getEmoployee)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggofiles.Handler))

	if err := router.Run("localhost:8080"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(docs.SwaggerInfo)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "Wel-come to Employee rest api(gin), swagger and mongodb example")
}
