package service

import (
	"ejaw/internal/models"
	"ejaw/internal/repository"
)

type SellerService struct {
	repo *repository.SellerRepository
}

func NewSellerService(repo *repository.SellerRepository) (*SellerService, error) {
	return &SellerService{repo: repo}, nil
}

func (s SellerService) CreateOrUpdateSeller(seller *models.Seller) error {
	return s.repo.Create(seller)
}

func (s SellerService) GetSellers() ([]models.Seller, error) {
	return s.repo.GetSellers()
}
