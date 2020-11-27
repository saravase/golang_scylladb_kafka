package reporting

import (
	"golang_scylladb_kafka/productstore"
	"golang_scylladb_kafka/uuid"
)

type ProductManager struct {
	uuidGenerator uuid.Generator
	store         productstore.Store
}

func NewProductManager(gen uuid.Generator, store productstore.Store) *ProductManager {
	return &ProductManager{
		uuidGenerator: gen,
		store:         store,
	}
}

func (m *ProductManager) CreateProduct(request CreateProductRequest) (response CreateProductResponse, err error) {
	productID := m.uuidGenerator.Generate()

	r := productstore.CreateProductRequest{
		ProductID: productID,
		UserID:    request.UserID,
		Status:    productstore.ProductStatusPending.String(),
		Title:     request.Title,
	}

	err = m.store.CreateProduct(r)
	response = CreateProductResponse{
		ProductID: productID,
	}
	return
}
