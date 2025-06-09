package database

import (
	"context"
	"database/sql"
	"log"
	"golang/project_Swagger_Pet/internal/models"
)

type OrderRepositoryPostgres struct {
	db *sql.DB
}

func NewOrderDb(db *sql.DB) *OrderRepositoryPostgres{
	return &OrderRepositoryPostgres{
		db: db,
	}
}

type OrderRepository interface {
    Create(ctx context.Context, order models.Order) error
    GetByID(ctx context.Context, id int64) (models.Order, error)
    Delete(ctx context.Context, id int64) error
}


func (r *OrderRepositoryPostgres) Create(ctx context.Context, order models.Order) error {
	query := `INSERT INTO orders (id, pet_id, quantity, ship_date) 
	          VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query,
		order.ID,
		order.PetID,
		order.Quantity,
		order.ShipDate,
	)
	if err != nil {
		log.Printf("Ошибка при создании заказа: %v", err)
	}
	return err
}

func (r *OrderRepositoryPostgres) GetByID(ctx context.Context, id int64) (models.Order, error) {
	query := `SELECT id, pet_id, quantity, ship_date 
	          FROM orders 
	          WHERE id = $1 AND is_deleted = FALSE`

	var order models.Order
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&order.ID,
		&order.PetID,
		&order.Quantity,
		&order.ShipDate,
	)
	if err != nil {
		log.Printf("Ошибка при получении заказа по ID: %v", err)
	}
	return order, err
}

func (r *OrderRepositoryPostgres) Delete(ctx context.Context, id int64) error {
	query := `UPDATE orders SET is_deleted = TRUE WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Ошибка при удалении заказа: %v", err)
	}
	return err
}