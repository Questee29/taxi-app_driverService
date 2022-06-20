package authorization

import (
	"context"
	"database/sql"
	"log"

	model "github.com/Questee29/taxi-app_driverService/models/order"
)

type orderRepository struct {
	db *sql.DB
}

func NewTaxiRepository(db *sql.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}
func (repository *orderRepository) FindFreeDriver(ctx context.Context, user model.UserRequest) (int32, error) {
	var DriverID int32

	query := `SELECT id 
	FROM drivers 
	WHERE taxi_type=$1 and status=$2 LIMIT 1
	
	`
	row := repository.db.QueryRow(query, user.TaxiType, "free")
	if err := row.Scan(&DriverID); err != nil {
		return 0, err
	}
	if err := repository.SetStatusBusy(DriverID); err != nil {
		return 0, err
	}
	log.Println(DriverID)
	return DriverID, nil
}
func (repository *orderRepository) SetStatusBusy(driverID int32) error {
	statusBusy := "busy"
	query := `UPDATE drivers
	SET status = $1
	WHERE id=$2`
	_, err := repository.db.Exec(query, statusBusy, driverID)
	if err != nil {
		return err
	}
	return nil
}
