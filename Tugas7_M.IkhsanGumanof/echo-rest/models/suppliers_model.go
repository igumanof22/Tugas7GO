package models

import (
	"fmt"
	"net/http"

	"Tugas7_M.IkhsanGumanof/echo-rest/common"
	"Tugas7_M.IkhsanGumanof/echo-rest/db"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

func FetchSupplier() (res Response, err error) {

	var errs *ex.AppError
	var cust common.Suppliers
	var custObj []common.Suppliers
	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	con := db.CreateCon()

	sqlQuery := `SELECT
					IFNULL(CompanyName,''),
					IFNULL(ContactName,'') ContactName,
					IFNULL(ContactTitle,'') ContactTitle,
					IFNULL(Address,'') Address,
					IFNULL(City,'') City,
					IFNULL(Country,'') Country,
					IFNULL(Phone,'') Phone ,
					IFNULL(PostalCode,'') PostalCode
				FROM suppliers `

	rows, err := con.Query(sqlQuery)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&cust.CompanyName, &cust.ContactName, &cust.ContactTitle, &cust.Address, &cust.City,
			&cust.Country, &cust.Phone, &cust.PostalCode)

		if err != nil {
			errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
			errMessage = err.Error()
			return res, err
		}

		custObj = append(custObj, cust)

	}

	res.Status = http.StatusOK
	res.Message = "succses"
	res.Data = custObj

	return res, nil
}

//StoreSupplier ...
func StoreSupplier(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `INSERT INTO suppliers (CompanyName,ContactName,ContactTitle,Address,City,Country,Phone,PostalCode)
					 VALUES (?,?,?,?,?,?,?,?)`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle, req.Address,
		req.City, req.Country, req.Phone, req.PostalCode)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"ContactName ADD": req.ContactName,
	}

	return res, nil
}

//UpdateSupplier ...
func UpdateSupplier(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `UPDATE suppliers SET CompanyName = ?, ContactTitle = ?, Address = ? WHERE  ContactName = '` + req.ContactName + `'`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName, req.ContactTitle, req.Address)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"row affected :": req.ContactName,
	}

	return res, nil
}

//DeleteSupplier ...
func DeleteSupplier(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `DELETE FROM suppliers WHERE  ContactName = ?`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.ContactName)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"row deleted :": req.ContactName,
	}

	return res, nil
}
