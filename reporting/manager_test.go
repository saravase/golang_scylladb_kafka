package reporting

import (
	"errors"
	"testing"

	mockproductstore "github.com/saravase/golang_scylladb_kafka/productstore/mock"
	mockuuid "github.com/saravase/golang_scylladb_kafka/uuid/mock"

	"github.com/saravase/golang_scylladb_kafka/productstore"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ProductManagerSuite struct {

	// Define the suite, and absorb the built-in basic suite
	// functionality from testify - including a T() method which
	// returns the current testing context
	suite.Suite

	// Assertions provides assertion methods around the TestingT interface.
	*require.Assertions

	// A Controller represents the top-level control of a mock ecosystem.
	// It defines the scope and lifetime of mock objects, as well as their expectations.
	// It is safe to call Controller's methods from multiple goroutines.
	// Each test should create a new Controller and invoke Finish via defer
	ctrl *gomock.Controller

	// NewProductManager(gen uuid.Generator, store productstore.Store)
	mockProductStore  *mockproductstore.MockStore
	mockUUIDGenerator *mockuuid.MockGenerator

	manager *ProductManager
}

func TestProductManagerSuite(t *testing.T) {
	suite.Run(t, new(ProductManagerSuite))
}

// Before run test
func (s *ProductManagerSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.mockProductStore = mockproductstore.NewMockStore(s.ctrl)
	s.mockUUIDGenerator = mockuuid.NewMockGenerator(s.ctrl)
	s.manager = NewProductManager(s.mockUUIDGenerator, s.mockProductStore)
}

// After run test
func (s *ProductManagerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *ProductManagerSuite) TestCreateProduct() {
	productID := "product-01"
	userID := "user-01"
	title := "plant"

	s.mockUUIDGenerator.EXPECT().Generate().Return(productID).Times(1)

	s.mockProductStore.EXPECT().CreateProduct(gomock.Eq(productstore.CreateProductRequest{
		ProductID: productID,
		UserID:    userID,
		Status:    productstore.ProductStatusPending.String(),
		Title:     title,
	})).Return(nil).Times(1)

	actualResponse, err := s.manager.CreateProduct(CreateProductRequest{
		UserID: userID,
		Title:  title,
	})
	s.NoError(err)

	expectedResponse := CreateProductResponse{
		ProductID: productID,
	}

	s.Equal(expectedResponse, actualResponse)
}

func (s *ProductManagerSuite) TestCreateProductError() {

	s.mockUUIDGenerator.EXPECT().Generate().Return("productID").Times(1)

	createError := errors.New("Create product error")
	s.mockProductStore.EXPECT().CreateProduct(gomock.Any()).Return(createError).Times(1)

	_, err := s.manager.CreateProduct(CreateProductRequest{})
	s.Equal(createError, err)

}
