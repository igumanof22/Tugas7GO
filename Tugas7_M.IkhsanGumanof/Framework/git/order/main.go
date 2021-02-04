package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	cm "Tugas7_M.IkhsanGumanof/Framework/git/order/common"
	"Tugas7_M.IkhsanGumanof/Framework/git/order/middleware"
	"Tugas7_M.IkhsanGumanof/Framework/git/order/services"
	"Tugas7_M.IkhsanGumanof/Framework/git/order/transport"

	log "github.com/Sirupsen/logrus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func initHandlers() {

	var svc services.PaymentServices

	svc = services.PaymentService{}
	svc = middleware.BasicMiddleware()(svc)

	root := cm.Config.RootURL

	http.Handle(fmt.Sprintf("%s/orders", root), httptransport.NewServer(
		transport.OrderEndpoint(svc), transport.DecodeRequest, transport.EncodeResponse,
	))

	//Handler baru customer
	http.Handle(fmt.Sprintf("%s/customers", root), httptransport.NewServer(
		transport.CustomerEndpoint(svc), transport.DecodeCustomerRequest, transport.EncodeResponse,
	))
	//Handler baru product
	http.Handle(fmt.Sprintf("%s/products", root), httptransport.NewServer(
		transport.ProductEndpoint(svc), transport.DecodeProductRequest, transport.EncodeResponse,
	))

	http.Handle(fmt.Sprintf("%s/fastpay", root), httptransport.NewServer(
		transport.FastPayEndpoint(svc), transport.DecodeFastPayRequest, transport.EncodeResponse,
	))

	http.Handle(fmt.Sprintf("%s/trips", root), httptransport.NewServer(
		transport.TripsEndpoint(svc), transport.DecodeTripRequest, transport.EncodeResponse,
	))

	http.Handle(fmt.Sprintf("%s/payment", root), httptransport.NewServer(
		transport.InquiryEndpoint(svc), transport.DecodeInquiryPaymentRequest, transport.EncodeResponse,
	))

	http.Handle(fmt.Sprintf("%s/call", root), httptransport.NewServer(
		transport.CallEndpoint(svc), transport.DecodeCallRequest, transport.EncodeResponse,
	))
}

var logger *log.Entry

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	//log.SetReportCaller(true)
}

func main() {

	configFile := flag.String("conf", "conf-dev.yml", "main configuration file")
	flag.Parse()
	initLogger()
	log.WithField("file", *configFile).Info("Loading configuration file")
	cm.LoadConfigFromFile(configFile)
	initHandlers()

	var err error
	if cm.Config.RootURL != "" || cm.Config.ListenPort != "" {
		err = http.ListenAndServe(cm.Config.ListenPort, nil)
	}

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}

}
