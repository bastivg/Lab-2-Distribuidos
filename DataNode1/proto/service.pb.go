// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: service.proto

package proto

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

type MensajeContinente struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=Nombre,proto3" json:"Nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=Apellido,proto3" json:"Apellido,omitempty"`
	Estado   string `protobuf:"bytes,3,opt,name=Estado,proto3" json:"Estado,omitempty"`
}

func (x *MensajeContinente) Reset() {
	*x = MensajeContinente{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MensajeContinente) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MensajeContinente) ProtoMessage() {}

func (x *MensajeContinente) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MensajeContinente.ProtoReflect.Descriptor instead.
func (*MensajeContinente) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *MensajeContinente) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *MensajeContinente) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

func (x *MensajeContinente) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type AlmacenarEnDN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=Nombre,proto3" json:"Nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=Apellido,proto3" json:"Apellido,omitempty"`
	Id       int32  `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *AlmacenarEnDN) Reset() {
	*x = AlmacenarEnDN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlmacenarEnDN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlmacenarEnDN) ProtoMessage() {}

func (x *AlmacenarEnDN) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlmacenarEnDN.ProtoReflect.Descriptor instead.
func (*AlmacenarEnDN) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *AlmacenarEnDN) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *AlmacenarEnDN) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

func (x *AlmacenarEnDN) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PedirDN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *PedirDN) Reset() {
	*x = PedirDN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PedirDN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedirDN) ProtoMessage() {}

func (x *PedirDN) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedirDN.ProtoReflect.Descriptor instead.
func (*PedirDN) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *PedirDN) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConsultaPoblacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado string `protobuf:"bytes,1,opt,name=Estado,proto3" json:"Estado,omitempty"`
}

func (x *ConsultaPoblacion) Reset() {
	*x = ConsultaPoblacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultaPoblacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultaPoblacion) ProtoMessage() {}

func (x *ConsultaPoblacion) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultaPoblacion.ProtoReflect.Descriptor instead.
func (*ConsultaPoblacion) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *ConsultaPoblacion) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type NombreCompleto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=Nombre,proto3" json:"Nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=Apellido,proto3" json:"Apellido,omitempty"`
}

func (x *NombreCompleto) Reset() {
	*x = NombreCompleto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NombreCompleto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NombreCompleto) ProtoMessage() {}

func (x *NombreCompleto) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NombreCompleto.ProtoReflect.Descriptor instead.
func (*NombreCompleto) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *NombreCompleto) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *NombreCompleto) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type ListaNombres struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombres []*NombreCompleto `protobuf:"bytes,1,rep,name=nombres,proto3" json:"nombres,omitempty"`
}

func (x *ListaNombres) Reset() {
	*x = ListaNombres{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListaNombres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListaNombres) ProtoMessage() {}

func (x *ListaNombres) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListaNombres.ProtoReflect.Descriptor instead.
func (*ListaNombres) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListaNombres) GetNombres() []*NombreCompleto {
	if x != nil {
		return x.Nombres
	}
	return nil
}

type RespuestaGenerica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *RespuestaGenerica) Reset() {
	*x = RespuestaGenerica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaGenerica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaGenerica) ProtoMessage() {}

func (x *RespuestaGenerica) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaGenerica.ProtoReflect.Descriptor instead.
func (*RespuestaGenerica) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *RespuestaGenerica) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x5f, 0x0a, 0x11, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x53, 0x0a, 0x0d, 0x41, 0x6c, 0x6d, 0x61, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x45, 0x6e, 0x44, 0x4e, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x41, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x41, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x07, 0x50,
	0x65, 0x64, 0x69, 0x72, 0x44, 0x4e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x74, 0x61, 0x50, 0x6f, 0x62, 0x6c, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x45,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x22, 0x44, 0x0a, 0x0e, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x41, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x41, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x22, 0x3e, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x61, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x6f,
	0x52, 0x07, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x11, 0x52, 0x65, 0x73,
	0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x32, 0x89, 0x02, 0x0a, 0x09, 0x4d, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65,
	0x6e, 0x74, 0x65, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75,
	0x65, 0x73, 0x74, 0x61, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x61, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x4d, 0x53, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x61, 0x72, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x6c, 0x6d, 0x61, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x45, 0x6e, 0x44, 0x4e, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x61,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x4d, 0x53, 0x61, 0x73, 0x6b,
	0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x72, 0x44, 0x4e, 0x1a,
	0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4f,
	0x4e, 0x55, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x74, 0x61, 0x50, 0x6f, 0x62, 0x6c, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x1a, 0x12,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x4e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x73, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_proto_goTypes = []interface{}{
	(*MensajeContinente)(nil), // 0: main.MensajeContinente
	(*AlmacenarEnDN)(nil),     // 1: main.AlmacenarEnDN
	(*PedirDN)(nil),           // 2: main.PedirDN
	(*ConsultaPoblacion)(nil), // 3: main.ConsultaPoblacion
	(*NombreCompleto)(nil),    // 4: main.NombreCompleto
	(*ListaNombres)(nil),      // 5: main.ListaNombres
	(*RespuestaGenerica)(nil), // 6: main.RespuestaGenerica
}
var file_service_proto_depIdxs = []int32{
	4, // 0: main.ListaNombres.nombres:type_name -> main.NombreCompleto
	0, // 1: main.MyService.SendContinentMsg:input_type -> main.MensajeContinente
	1, // 2: main.MyService.SendOMSdepositar:input_type -> main.AlmacenarEnDN
	2, // 3: main.MyService.SendOMSask:input_type -> main.PedirDN
	3, // 4: main.MyService.SendONUMsg:input_type -> main.ConsultaPoblacion
	6, // 5: main.MyService.SendContinentMsg:output_type -> main.RespuestaGenerica
	6, // 6: main.MyService.SendOMSdepositar:output_type -> main.RespuestaGenerica
	4, // 7: main.MyService.SendOMSask:output_type -> main.NombreCompleto
	5, // 8: main.MyService.SendONUMsg:output_type -> main.ListaNombres
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MensajeContinente); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlmacenarEnDN); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PedirDN); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultaPoblacion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NombreCompleto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListaNombres); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaGenerica); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
