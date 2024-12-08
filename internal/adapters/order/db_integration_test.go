package order

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"trade-microservice.fyerfyer.net/internal/application/domain"
)

// OrderDataBaseTestSuite defines the test suite
type OrderDataBaseTestSuite struct {
	suite.Suite
	DatasourceURL string
}

// SetUpSuite sets up the test environment
func (o *OrderDataBaseTestSuite) SetupSuite() {
	ctx := context.Background()
	port := nat.Port("3306/tcp")

	// Define database URL function
	dbURL := func(host string, port nat.Port) string {
		return fmt.Sprintf(
			"root:110119abc@tcp(%s:%s)/orders?charset=utf8mb4&parseTime=True&loc=Local",
			host, port.Port(),
		)
	}

	// Define container request
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/mysql:8.0.30",
		ExposedPorts: []string{string(port)},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "110119abc",
			"MYSQL_DATABASE":      "orders",
		},
		WaitingFor: wait.ForSQL(port, "mysql", dbURL).
			WithStartupTimeout(30 * time.Second).
			WithPollInterval(500 * time.Millisecond),
	}

	// Start the MySQL container
	mysqlContainer, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	o.Require().NoError(err, "failed to start MySQL container")

	// Get host, port from the container
	host, err := mysqlContainer.Host(ctx)
	o.Require().NoError(err, "failed to get container host")

	mappedPort, err := mysqlContainer.MappedPort(ctx, port)
	o.Require().NoError(err, "failed to get container mapped port")

	o.DatasourceURL = dbURL(host, mappedPort)
}

// Test_Should_Save_Order tests saving an order
func (o *OrderDataBaseTestSuite) Test_Should_Save_Order() {
	log.Printf("*********%v", o.DatasourceURL)
	gormRepo, err := NewGormRepository(o.DatasourceURL)
	o.Require().NoError(err, "failed to initialize Gorm repository")
	o.Require().NotNil(gormRepo, "repository is nil")

	order := &domain.Order{}
	saveErr := gormRepo.Save(order)
	o.NoError(saveErr, "failed to save order")
}

// Run the test suite
func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDataBaseTestSuite))
}
