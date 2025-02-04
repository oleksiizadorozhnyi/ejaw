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

func (s SellerService) CreateSeller(seller *models.Seller) error {
	return s.repo.Create(seller)
}

func (s SellerService) UpdateSeller(seller *models.Seller) error {
	return s.repo.Update(seller)
}

func (s SellerService) DeleteSeller(phonenumber string) error {
	return s.repo.DeleteByPhone(phonenumber)
}

func (s SellerService) GetSellers() ([]models.Seller, error) {
	return s.repo.GetSellers()
}
