// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: proto/order/order.proto

package order

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

type ProcessOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint64       `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	OrderItems []*OrderItem `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
}

func (x *ProcessOrderRequest) Reset() {
	*x = ProcessOrderRequest{}
	mi := &file_proto_order_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessOrderRequest) ProtoMessage() {}

func (x *ProcessOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessOrderRequest.ProtoReflect.Descriptor instead.
func (*ProcessOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessOrderRequest) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ProcessOrderRequest) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type ProcessOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProcessOrderResponse) Reset() {
	*x = ProcessOrderResponse{}
	mi := &file_proto_order_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessOrderResponse) ProtoMessage() {}

func (x *ProcessOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessOrderResponse.ProtoReflect.Descriptor instead.
func (*ProcessOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
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
	mi := &file_proto_order_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_order_proto_msgTypes[2]
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
	return file_proto_order_order_proto_rawDescGZIP(), []int{2}
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

var File_proto_order_order_proto protoreflect.FileDescriptor

var file_proto_order_order_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0x69, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2e, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x69, 0x0a, 0x09, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0x50, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x47, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_order_order_proto_rawDescOnce sync.Once
	file_proto_order_order_proto_rawDescData = file_proto_order_order_proto_rawDesc
)

func file_proto_order_order_proto_rawDescGZIP() []byte {
	file_proto_order_order_proto_rawDescOnce.Do(func() {
		file_proto_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_order_proto_rawDescData)
	})
	return file_proto_order_order_proto_rawDescData
}

var file_proto_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_order_order_proto_goTypes = []any{
	(*ProcessOrderRequest)(nil),  // 0: order.ProcessOrderRequest
	(*ProcessOrderResponse)(nil), // 1: order.ProcessOrderResponse
	(*OrderItem)(nil),            // 2: order.OrderItem
}
var file_proto_order_order_proto_depIdxs = []int32{
	2, // 0: order.ProcessOrderRequest.order_items:type_name -> order.OrderItem
	0, // 1: order.Order.ProcessOrder:input_type -> order.ProcessOrderRequest
	1, // 2: order.Order.ProcessOrder:output_type -> order.ProcessOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_order_order_proto_init() }
func file_proto_order_order_proto_init() {
	if File_proto_order_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_order_proto_goTypes,
		DependencyIndexes: file_proto_order_order_proto_depIdxs,
		MessageInfos:      file_proto_order_order_proto_msgTypes,
	}.Build()
	File_proto_order_order_proto = out.File
	file_proto_order_order_proto_rawDesc = nil
	file_proto_order_order_proto_goTypes = nil
	file_proto_order_order_proto_depIdxs = nil
}
