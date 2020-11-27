//go:generate mockgen -source store.go -destination mock/store_mock.go -package mock
package productstore

type Store interface {
	CreateProduct(r CreateProductRequest) error
}

type CreateProductRequest struct {
	ProductID string
	UserID    string
	Status    string
	Title     string
}
