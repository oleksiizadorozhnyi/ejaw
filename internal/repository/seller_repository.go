package repository

import (
	"database/sql"
	"ejaw/config"
	"ejaw/internal/models"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var ErrPhoneExists = errors.New("phone number already exists")

type SellerRepository struct {
	db *sql.DB
}

func NewSellerRepository(postgres *config.Postgres) (*SellerRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", postgres.Host, postgres.User, postgres.Password, postgres.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to apply migrations to database: %v", err)
	}

	return &SellerRepository{db: db}, nil
}

func runMigrations(db *sql.DB) error {
	if err := goose.Up(db, "./migrations"); err != nil {
		return fmt.Errorf("goose migration failed: %w", err)
	}
	return nil
}

func (r *SellerRepository) Create(seller *models.Seller) error {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sellers WHERE phone = $1)", seller.Phone).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return ErrPhoneExists
	}

	query := "INSERT INTO sellers (id, name, phone) VALUES (gen_random_uuid(), $1, $2) RETURNING id"
	return r.db.QueryRow(query, seller.Name, seller.Phone).Scan(&seller.ID)
}

func (r *SellerRepository) DeleteByPhone(phone string) error {
	result, err := r.db.Exec("DELETE FROM sellers WHERE phone = $1", phone)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("seller not found")
	}

	return nil
}

func (r *SellerRepository) GetSellers() ([]models.Seller, error) {
	rows, err := r.db.Query("SELECT id, name, phone FROM sellers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sellers []models.Seller
	for rows.Next() {
		var s models.Seller
		if err := rows.Scan(&s.ID, &s.Name, &s.Phone); err != nil {
			return nil, err
		}
		sellers = append(sellers, s)
	}
	return sellers, nil
}
