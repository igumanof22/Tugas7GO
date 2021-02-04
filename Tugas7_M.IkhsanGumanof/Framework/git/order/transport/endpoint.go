package transport

import (
	"context"

	cm "Tugas7_M.IkhsanGumanof/Framework/git/order/common"
	"Tugas7_M.IkhsanGumanof/Framework/git/order/services"

	log "github.com/Sirupsen/logrus"

	"github.com/go-kit/kit/endpoint"
)

func invalidRequest() cm.Message {
	return cm.Message{
		Result: &cm.Result{
			Code:   99,
			Remark: "Invalid Request",
		},
	}
}

func OrderEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if req, ok := request.(cm.Message); ok {
			return svc.OrderHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}

func CustomerEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if req, ok := request.(cm.Customers); ok {
			return svc.CustomerHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}
func ProductEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if req, ok := request.(cm.Products); ok {
			return svc.ProductHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}

func FastPayEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(cm.FastPayRequest); ok {
			return svc.FastPayHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}

func CallEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if req, ok := request.(cm.FastPayRequest); ok {
			return svc.CallHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}

func InquiryEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if req, ok := request.(cm.InquiryRequest); ok {
			return svc.InquiryHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}

func TripsEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if req, ok := request.(cm.MyTrips); ok {
			return svc.TripsHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}
