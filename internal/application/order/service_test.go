package order

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	redis "trade-microservice.fyerfyer.net/internal/application/cache"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/ports"
)

type Application struct {
	repo    Repository
	cache   ports.Cache
	payment ports.PaymentPort
}

func NewApplication(repo Repository, cache ports.Cache, payment ports.PaymentPort) *Application {
	return &Application{
		repo:    repo,
		cache:   cache,
		payment: payment,
	}
}

func (a *Application) ProcessOrder(order *domain.Order) (*domain.Order, error) {
	if err := a.repo.Save(order); err != nil {
		return nil, err
	}

	if err := a.cache.Set(redis.GetOrderKey(order.ID, order.Status), order, 0); err != nil {
		return nil, err
	}

	totalPrice := order.TotalPrice()
	_, err := a.payment.Charge(order.CustomerID, order.ID, totalPrice)
	if err != nil {
		return nil, err
	}

	return order, nil
}

type mockedPayment struct {
	mock.Mock
}

func (p *mockedPayment) Charge(customerID, orderID uint64, totalPrice float32) (*domain.Payment, error) {
	args := p.Called(customerID, orderID, totalPrice)
	return args.Get(0).(*domain.Payment), args.Error(1)
}

type mockedRepo struct {
	mock.Mock
}

func (r *mockedRepo) Save(order *domain.Order) error {
	args := r.Called(order)
	return args.Error(0)
}

func (r *mockedRepo) Update(order *domain.Order) error {
	args := r.Called(order)
	return args.Error(0)
}

func (r *mockedRepo) Delete(orderID uint64) error {
	args := r.Called(orderID)
	return args.Error(0)
}

type mockedCache struct {
	mock.Mock
}

func (c *mockedCache) Set(key string, data interface{}, expire int) error {
	args := c.Called(key, data, expire)
	return args.Error(0)
}

func (c *mockedCache) Exists(key string) bool {
	args := c.Called(key)
	return args.Bool(0)
}

func (c *mockedCache) Get(key string) ([]byte, error) {
	args := c.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *mockedCache) Delete(key string) (bool, error) {
	args := c.Called(key)
	return args.Bool(0), args.Error(1)
}

func Test_Should_Process_Order(t *testing.T) {
	payment := new(mockedPayment)
	repo := new(mockedRepo)
	cache := new(mockedCache)
	repo.On("Save", mock.Anything).Return(nil)
	cache.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	payment.On("Charge", mock.Anything, mock.Anything, mock.Anything).Return(&domain.Payment{
		Status:  "success",
		Message: "Payment processed successfully",
	}, nil)

	application := NewApplication(repo, cache, payment)
	_, err := application.ProcessOrder(&domain.Order{
		CustomerID: 1,
		Items: []domain.OrderItem{
			{
				ProductCode: "camera",
				UnitPrice:   12.3,
				Quantity:    3,
			},
		},
		CreatedAt: time.Time{},
	})

	assert.Nil(t, err)
	repo.AssertNumberOfCalls(t, "Save", 1)
	payment.AssertNumberOfCalls(t, "Charge", 1)
	cache.AssertNumberOfCalls(t, "Set", 1)
}

// go install github.com/vektra/mockery/v2@v2.50.0
