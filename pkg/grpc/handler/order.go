package handler

import (
	"context"
	"log"

	model "github.com/Questee29/taxi-app_driverService/models/order"
	pb "github.com/Questee29/taxi-app_driverService/proto/protob"
)

type OrderService interface {
	FindFreeDriver(ctx context.Context, user model.UserRequest) (model.DriverResponse, error)
}

type OrderHandler struct {
	pb.UnimplementedOrderGrpcServer
	service OrderService
}

func NewOrderHandler(service OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) FindDriver(ctx context.Context, req *pb.FindDriverRequest) (*pb.FindDriverResponse, error) {
	uReq := model.UserRequest{
		ID:       req.GetUserid(),
		TaxiType: req.GetType().String(),
	}
	log.Printf("New request from user with id '%d'.Car type '%s'", uReq.ID, uReq.TaxiType)
	//creating model
	dResp, err := h.service.FindFreeDriver(ctx, uReq)
	if err != nil {
		return nil, err
	}
	log.Printf("%s car was found", dResp.TaxType)
	return toPBModel(dResp), nil
}

func toPBModel(driver model.DriverResponse) *pb.FindDriverResponse {
	value, ok := pb.CarType_value[driver.TaxType]
	if !ok {
		return nil
	}

	return &pb.FindDriverResponse{
		Driverid: driver.Driverid,
		Type:     *pb.CarType(value).Enum(),
	}

}
