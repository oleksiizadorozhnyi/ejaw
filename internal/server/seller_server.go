package server

import (
	"ejaw/config"
	"ejaw/internal/models"
	"ejaw/internal/repository"
	"ejaw/internal/service"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
)

type SellerServer struct {
	service       *service.SellerService
	adminUser     string
	adminPassword string
}

func NewSellerServer(service *service.SellerService, admin *config.Admin) *SellerServer {
	return &SellerServer{
		service:       service,
		adminUser:     admin.User,
		adminPassword: admin.Password,
	}
}

func (s *SellerServer) Run(addr string) error {
	http.HandleFunc("/sellers", s.basicAuth(s.GetSellers))
	http.HandleFunc("/create_seller", s.basicAuth(s.CreateSeller))
	http.HandleFunc("/delete_seller", s.basicAuth(s.DeleteSeller))

	return http.ListenAndServe(addr, nil)
}

func (s *SellerServer) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != s.adminUser || pass != s.adminPassword {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"Restricted\"")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func (s *SellerServer) CreateSeller(w http.ResponseWriter, r *http.Request) {
	var seller models.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		zap.L().Error("Invalid JSON payload", zap.Error(err))
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	if err := s.service.CreateOrUpdateSeller(&seller); err != nil {
		if errors.Is(err, repository.ErrPhoneExists) {
			http.Error(w, "Phone number already exist", http.StatusBadRequest)
			return
		}
		zap.L().Error("Internal server error", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(seller)
}

func (s *SellerServer) DeleteSeller(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Phone string `json:"phone"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		zap.L().Error("Invalid JSON payload", zap.Error(err))
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := s.service.DeleteSeller(request.Phone); err != nil {
		zap.L().Error("Failed to delete seller", zap.Error(err))
		http.Error(w, "Seller not found or deletion failed", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Seller deleted successfully"})
}

func (s *SellerServer) GetSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := s.service.GetSellers()
	if err != nil {
		zap.L().Error("Internal server error", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sellers)
}
