package database

import (
	"context"
	"database/sql"
	"log"
	"golang/project_Swagger_Pet/internal/models"
)

type PetRepositoryPostgres struct {
	db *sql.DB
}

func NewPetDb(db *sql.DB) *PetRepositoryPostgres{
	return &PetRepositoryPostgres{
		db: db,
	}
}

type PetRepository interface {
    Create(ctx context.Context, pet models.Pet) error
    GetByID(ctx context.Context, id int64) (models.Pet, error)
    Update(ctx context.Context, pet models.Pet) error
    Delete(ctx context.Context, id int64) error
    FindByStatus(ctx context.Context, status string) ([]models.Pet, error)
}


func (d *PetRepositoryPostgres) Create(ctx context.Context, pet models.Pet) error{
	query := `INSERT INTO pets (id, name, status) VALUES ($1, $2, $3)`
	_, err := d.db.ExecContext(ctx, query, pet.ID, pet.Name, pet.Status)
	if err != nil{
		log.Println("Ошибка в создании Pet")
		return err
	}
	return nil
}

func (d *PetRepositoryPostgres) GetByID(ctx context.Context, id int64) (models.Pet, error) {
	query := `SELECT id, name, status FROM pets WHERE id = $1 AND is_deleted = FALSE`
	var p models.Pet
	err := d.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Status)
	return p, err
}

func (d *PetRepositoryPostgres) Delete(ctx context.Context, id int64) error {
	query := `UPDATE pets SET is_deleted = TRUE WHERE id = $1`
	_, err := d.db.ExecContext(ctx, query, id)
	return err
}

func (d *PetRepositoryPostgres) Update(ctx context.Context,  pet models.Pet) error{
	query := `UPDATE pets SET name = $1, status = $2 WHERE id = $3;`
	_, err := d.db.ExecContext(ctx, query, pet.Name, pet.Status, pet.ID)
	if err != nil{
		log.Println("Ошибка в Update Pet")
		return err
	}
	return nil
}

func (d *PetRepositoryPostgres) FindByStatus(ctx context.Context, status string) ([]models.Pet, error) {
    query := `SELECT id, name, status FROM pets WHERE is_deleted = FALSE AND status = $1`

    rows, err := d.db.QueryContext(ctx, query, status)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var pets []models.Pet
    for rows.Next() {
        var p models.Pet
        if err := rows.Scan(&p.ID, &p.Name, &p.Status); err != nil {
            return nil, err
        }
        pets = append(pets, p)
    }
    return pets, nil
}