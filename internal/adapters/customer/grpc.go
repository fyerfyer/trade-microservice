package customer

import (
	"context"

	"google.golang.org/grpc"
	"trade-microservice.fyerfyer.net/internal/application/domain"
	"trade-microservice.fyerfyer.net/internal/ports"
	pb "trade-microservice.fyerfyer.net/proto/customer"
)

type Adapter struct {
	api    ports.CustomerPort
	port   int
	server *grpc.Server
	pb.UnimplementedCustomerServer
}

func NewAdapter(api ports.CustomerPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a *Adapter) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	customer, err := a.api.CreateCustomer(req.GetCustomerName())
	if err != nil {
		return nil, err
	}

	return &pb.CreateCustomerResponse{
		CustomerId: customer.ID,
		Success:    true,
	}, nil
}

func (a *Adapter) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customer, err := a.api.GetCustomer(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.GetCustomerResponse{Customer: &pb.CustomerEntity{
		CustomerId: customer.ID,
		Name:       customer.Name,
		Status:     customer.Status,
		Balance:    customer.Balance,
	}}, nil
}

func (a *Adapter) DeactiveCustomer(ctx context.Context, req *pb.DeactivateCustomerRequest) (*pb.DeactivateCustomerResponse, error) {
	err := a.api.DeactiveCustomer(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.DeactivateCustomerResponse{Status: "inactive"}, nil
}

func (a *Adapter) ActivateCustomer(ctx context.Context, req *pb.ActivateCustomerRequest) (*pb.ActivateCustomerResponse, error) {
	err := a.api.ReactiveCustomer(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.ActivateCustomerResponse{Status: "active"}, nil
}

func (a *Adapter) SubmitOrder(ctx context.Context, req *pb.SubmitOrderRequest) (*pb.SubmitOrderResponse, error) {
	success, err := a.api.SubmitOrder(req.GetCustomerId(), convertPBIntoDomainItems(req.GetOrderItems()))
	if err != nil {
		return &pb.SubmitOrderResponse{
			Success: success,
			Message: err.Error(),
		}, err
	}

	return &pb.SubmitOrderResponse{
		Success: success,
		Message: "order processed successfully",
	}, nil
}

func (a *Adapter) GetUnpaidOrders(ctx context.Context, req *pb.GetUnpaidOrdersRequest) (*pb.GetUnpaidOrdersResponse, error) {
	orders, err := a.api.GetUnpaidOrders(req.GetCustomerId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUnpaidOrdersResponse{
		UnpaidOrders: convertDomainOrderIntoPB(orders),
	}, nil
}

func convertPBIntoDomainItems(items []*pb.OrderItem) []domain.OrderItem {
	var domainItems []domain.OrderItem
	for _, item := range items {
		domainItems = append(domainItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	return domainItems
}

func convertDomainItemsIntoPB(items []domain.OrderItem) []*pb.OrderItem {
	var pbItems []*pb.OrderItem
	for _, item := range items {
		pbItems = append(pbItems, &pb.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}

	return pbItems
}

func convertDomainOrderIntoPB(orders []domain.Order) []*pb.Order {
	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			OrderId: order.ID,
			Items:   convertDomainItemsIntoPB(order.Items),
			Status:  order.Status,
		})
	}

	return pbOrders
}
