package controllers

import (
	"net/http"

	"Tugas7_M.IkhsanGumanof/echo-rest/models"

	"github.com/labstack/echo"
)

//FetchAllCustomers ...
func FetchAllEmployees(c echo.Context) (err error) {

	result, err := models.FetchEmployees()

	return c.JSON(http.StatusOK, result)
}

//AddEmployees
func AddEmployees(c echo.Context) (err error) {

	result, err := models.AddEmployee(c)

	return c.JSON(http.StatusOK, result)
}

//UpdateEmployees
func UpdateEmployees(c echo.Context) (err error) {

	result, err := models.UpdateEmployee(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteEmployees ...
func DeleteEmployees(c echo.Context) (err error) {

	result, err := models.DeleteEmployee(c)

	return c.JSON(http.StatusOK, result)
}
