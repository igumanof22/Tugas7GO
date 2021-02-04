package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "Tugas7_M.IkhsanGumanof/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) InquiryHandler(ctx context.Context, req cm.InquiryRequest) (res cm.InquiryPayment) {
	var db *sql.DB
	var err error
	defer panicRecovery()

	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)

	db, err = sql.Open("mysql", mySQL)

	if err != nil {
		panic(err.Error())
	}

	res.TrxID = req.TrxID

	var payment cm.InquiryPayment
	var payments []cm.InquiryPayment

	sql := `SELECT
				trx_id,
				IFNULL(merchant_id,'') merchant_id,
				IFNULL(merchant,'') merchant,
				IFNULL(bill_no,'') bill_no,
				IFNULL(payment_reff,'') payment_reff,
				IFNULL(payment_date,'') payment_date ,
				IFNULL(payment_status_code,'') payment_status_code ,
				IFNULL(payment_status_desc,'') payment_status_desc ,
				IFNULL(response_code,'') response_code ,
				IFNULL(response_desc,'') response_desc
			FROM payment WHERE trx_id = ?`

	result, err := db.Query(sql, req.TrxID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&payment.TrxID, &payment.MerchantID, &payment.Merchant, &payment.BillNo, &payment.PaymentReff,
			&payment.PaymentDate, &payment.PaymentStatusCode, &payment.PaymentStatusDesc, &payment.ResponseCode,
			&payment.ResponseDesc)

		if err != nil {
			panic(err.Error())
		}
		payments = append(payments, payment)
	}
	payment.Response = req.Request
	payment.TrxID = req.TrxID
	payment.MerchantID = req.MerchantID
	payment.BillNo = req.BillNo

	res = payment

	return
}
