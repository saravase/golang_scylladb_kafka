package reporting

import (
	mockproductstore "golang_scylladb_kafka/productstore/mock"

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

	mockProductStore *mockproductstore.MockStore
}
