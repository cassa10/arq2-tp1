package mock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

type InterfaceMocks struct {
	Logger       *MockLogger
	CustomerRepo *MockCustomerRepository
	SellerRepo   *MockSellerRepository
	ProductRepo  *MockProductRepository
	OrderRepo    *MockOrderRepository
}

// NewInterfaceMocks create an *InterfaceMocks with their mocked interfaces initialized
func NewInterfaceMocks(t *testing.T) *InterfaceMocks {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return &InterfaceMocks{
		Logger:       NewMockLogger(ctrl),
		CustomerRepo: NewMockCustomerRepository(ctrl),
		SellerRepo:   NewMockSellerRepository(ctrl),
		ProductRepo:  NewMockProductRepository(ctrl),
		OrderRepo:    NewMockOrderRepository(ctrl),
	}
}
