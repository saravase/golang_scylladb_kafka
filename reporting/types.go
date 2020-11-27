package reporting

type CreateProductRequest struct {
	UserID string
	Title  string
}

type CreateProductResponse struct {
	ProductID string
}
