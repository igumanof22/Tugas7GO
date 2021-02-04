package controllers

import (
	"net/http"

	"Tugas7_M.IkhsanGumanof/echo-rest/models"

	"github.com/labstack/echo"
)

//FetchAllSuppliers ...
func FetchAllSuppliers(c echo.Context) (err error) {

	result, err := models.FetchSupplier()

	return c.JSON(http.StatusOK, result)
}

//addSuppliers ...
func StoreSuppliers(c echo.Context) (err error) {

	result, err := models.StoreSupplier(c)

	return c.JSON(http.StatusOK, result)
}

//UpdateSuppliers ...
func UpdateSuppliers(c echo.Context) (err error) {

	result, err := models.UpdateSupplier(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteSuppliers ...
func DeleteSuppliers(c echo.Context) (err error) {

	result, err := models.DeleteSupplier(c)

	return c.JSON(http.StatusOK, result)
}
