// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: proto/customer/customer.proto

package customer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerName string `protobuf:"bytes,1,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_proto_customer_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_proto_customer_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreateUserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type SubmitOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint64       `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	OrderItems []*OrderItem `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
}

func (x *SubmitOrderRequest) Reset() {
	*x = SubmitOrderRequest{}
	mi := &file_proto_customer_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitOrderRequest) ProtoMessage() {}

func (x *SubmitOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitOrderRequest.ProtoReflect.Descriptor instead.
func (*SubmitOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitOrderRequest) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *SubmitOrderRequest) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type SubmitOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SubmitOrderResponse) Reset() {
	*x = SubmitOrderResponse{}
	mi := &file_proto_customer_customer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitOrderResponse) ProtoMessage() {}

func (x *SubmitOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitOrderResponse.ProtoReflect.Descriptor instead.
func (*SubmitOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitOrderResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SubmitOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCode string  `protobuf:"bytes,1,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	UnitPrice   float32 `protobuf:"fixed32,2,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	Quantity    int32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_proto_customer_customer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{4}
}

func (x *OrderItem) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *OrderItem) GetUnitPrice() float32 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetUnpaidOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetUnpaidOrdersRequest) Reset() {
	*x = GetUnpaidOrdersRequest{}
	mi := &file_proto_customer_customer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUnpaidOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnpaidOrdersRequest) ProtoMessage() {}

func (x *GetUnpaidOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnpaidOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetUnpaidOrdersRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{5}
}

func (x *GetUnpaidOrdersRequest) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetUnpaidOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnpaidOrders []*Order `protobuf:"bytes,1,rep,name=unpaid_orders,json=unpaidOrders,proto3" json:"unpaid_orders,omitempty"`
}

func (x *GetUnpaidOrdersResponse) Reset() {
	*x = GetUnpaidOrdersResponse{}
	mi := &file_proto_customer_customer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUnpaidOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnpaidOrdersResponse) ProtoMessage() {}

func (x *GetUnpaidOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnpaidOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetUnpaidOrdersResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{6}
}

func (x *GetUnpaidOrdersResponse) GetUnpaidOrders() []*Order {
	if x != nil {
		return x.UnpaidOrders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64       `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Items   []*OrderItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Status  string       `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_proto_customer_customer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{7}
}

func (x *Order) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeactivateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *DeactivateUserRequest) Reset() {
	*x = DeactivateUserRequest{}
	mi := &file_proto_customer_customer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeactivateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateUserRequest) ProtoMessage() {}

func (x *DeactivateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateUserRequest.ProtoReflect.Descriptor instead.
func (*DeactivateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{8}
}

func (x *DeactivateUserRequest) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type DeactivateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeactivateUserResponse) Reset() {
	*x = DeactivateUserResponse{}
	mi := &file_proto_customer_customer_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeactivateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateUserResponse) ProtoMessage() {}

func (x *DeactivateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateUserResponse.ProtoReflect.Descriptor instead.
func (*DeactivateUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{9}
}

func (x *DeactivateUserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ActivateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ActivateUserRequest) Reset() {
	*x = ActivateUserRequest{}
	mi := &file_proto_customer_customer_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserRequest) ProtoMessage() {}

func (x *ActivateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserRequest.ProtoReflect.Descriptor instead.
func (*ActivateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{10}
}

func (x *ActivateUserRequest) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type ActivateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ActivateUserResponse) Reset() {
	*x = ActivateUserResponse{}
	mi := &file_proto_customer_customer_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserResponse) ProtoMessage() {}

func (x *ActivateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_customer_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserResponse.ProtoReflect.Descriptor instead.
func (*ActivateUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_customer_proto_rawDescGZIP(), []int{11}
}

func (x *ActivateUserResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_customer_customer_proto protoreflect.FileDescriptor

var file_proto_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x6b, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x49, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x09, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e,
	0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x70, 0x61,
	0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x75,
	0x6e, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x0c, 0x75, 0x6e, 0x70, 0x61, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x22, 0x65, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x30, 0x0a, 0x16, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x36, 0x0a, 0x13, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x14,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa3, 0x03, 0x0a,
	0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x44, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x12, 0x1f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_customer_customer_proto_rawDescOnce sync.Once
	file_proto_customer_customer_proto_rawDescData = file_proto_customer_customer_proto_rawDesc
)

func file_proto_customer_customer_proto_rawDescGZIP() []byte {
	file_proto_customer_customer_proto_rawDescOnce.Do(func() {
		file_proto_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_customer_customer_proto_rawDescData)
	})
	return file_proto_customer_customer_proto_rawDescData
}

var file_proto_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_customer_customer_proto_goTypes = []any{
	(*CreateUserRequest)(nil),       // 0: customer.CreateUserRequest
	(*CreateUserResponse)(nil),      // 1: customer.CreateUserResponse
	(*SubmitOrderRequest)(nil),      // 2: customer.SubmitOrderRequest
	(*SubmitOrderResponse)(nil),     // 3: customer.SubmitOrderResponse
	(*OrderItem)(nil),               // 4: customer.OrderItem
	(*GetUnpaidOrdersRequest)(nil),  // 5: customer.GetUnpaidOrdersRequest
	(*GetUnpaidOrdersResponse)(nil), // 6: customer.GetUnpaidOrdersResponse
	(*Order)(nil),                   // 7: customer.Order
	(*DeactivateUserRequest)(nil),   // 8: customer.DeactivateUserRequest
	(*DeactivateUserResponse)(nil),  // 9: customer.DeactivateUserResponse
	(*ActivateUserRequest)(nil),     // 10: customer.ActivateUserRequest
	(*ActivateUserResponse)(nil),    // 11: customer.ActivateUserResponse
}
var file_proto_customer_customer_proto_depIdxs = []int32{
	4,  // 0: customer.SubmitOrderRequest.order_items:type_name -> customer.OrderItem
	7,  // 1: customer.GetUnpaidOrdersResponse.unpaid_orders:type_name -> customer.Order
	4,  // 2: customer.Order.items:type_name -> customer.OrderItem
	0,  // 3: customer.Customer.CreateCustomer:input_type -> customer.CreateUserRequest
	2,  // 4: customer.Customer.SubmitOrder:input_type -> customer.SubmitOrderRequest
	5,  // 5: customer.Customer.GetUnpaidOrders:input_type -> customer.GetUnpaidOrdersRequest
	8,  // 6: customer.Customer.DeactivateCustomer:input_type -> customer.DeactivateUserRequest
	10, // 7: customer.Customer.ActivateUser:input_type -> customer.ActivateUserRequest
	1,  // 8: customer.Customer.CreateCustomer:output_type -> customer.CreateUserResponse
	3,  // 9: customer.Customer.SubmitOrder:output_type -> customer.SubmitOrderResponse
	6,  // 10: customer.Customer.GetUnpaidOrders:output_type -> customer.GetUnpaidOrdersResponse
	9,  // 11: customer.Customer.DeactivateCustomer:output_type -> customer.DeactivateUserResponse
	11, // 12: customer.Customer.ActivateUser:output_type -> customer.ActivateUserResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_customer_customer_proto_init() }
func file_proto_customer_customer_proto_init() {
	if File_proto_customer_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_customer_customer_proto_goTypes,
		DependencyIndexes: file_proto_customer_customer_proto_depIdxs,
		MessageInfos:      file_proto_customer_customer_proto_msgTypes,
	}.Build()
	File_proto_customer_customer_proto = out.File
	file_proto_customer_customer_proto_rawDesc = nil
	file_proto_customer_customer_proto_goTypes = nil
	file_proto_customer_customer_proto_depIdxs = nil
}
