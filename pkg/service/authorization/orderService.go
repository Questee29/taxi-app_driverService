package authorization

import (
	"context"

	model "github.com/Questee29/taxi-app_driverService/models/order"
	pb "github.com/Questee29/taxi-app_driverService/proto/protob"
)

type OrderRepository interface {
	FindFreeDriver(ctx context.Context, user model.UserRequest) (int32, error)
}
type orderService struct {
	pb.UnimplementedOrderGrpcServer
	repository OrderRepository
}

func NewOrderService(repository OrderRepository) *orderService {
	return &orderService{
		repository: repository,
	}
}
func (service *orderService) FindFreeDriver(ctx context.Context, user model.UserRequest) (model.DriverResponse, error) {
	driverID, err := service.repository.FindFreeDriver(ctx, user)
	if err != nil {
		return model.DriverResponse{}, err
	}
	response := model.DriverResponse{
		Driverid: driverID,
		TaxType:  user.TaxiType,
	}

	return response, nil
}
