package repository

import (
	"database/sql"
	"ejaw/config"
	"ejaw/internal/models"
	"fmt"
	_ "github.com/lib/pq"
)

type SellerRepository struct {
	db *sql.DB
}

func NewSellerRepository(postgres *config.Postgres) (*SellerRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", postgres.Host, postgres.User, postgres.Password, postgres.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &SellerRepository{db: db}, nil
}

func (r *SellerRepository) Create(seller *models.Seller) error {
	return r.db.QueryRow("INSERT INTO sellers (name, phone) VALUES ($1, $2) RETURNING id", seller.Name, seller.Phone).Scan(&seller.ID)
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
