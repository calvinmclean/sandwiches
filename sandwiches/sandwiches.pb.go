// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sandwiches.proto

package sandwiches

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type IngredientRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngredientRequest) Reset()         { *m = IngredientRequest{} }
func (m *IngredientRequest) String() string { return proto.CompactTextString(m) }
func (*IngredientRequest) ProtoMessage()    {}
func (*IngredientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{1}
}

func (m *IngredientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngredientRequest.Unmarshal(m, b)
}
func (m *IngredientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngredientRequest.Marshal(b, m, deterministic)
}
func (m *IngredientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngredientRequest.Merge(m, src)
}
func (m *IngredientRequest) XXX_Size() int {
	return xxx_messageInfo_IngredientRequest.Size(m)
}
func (m *IngredientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IngredientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IngredientRequest proto.InternalMessageInfo

func (m *IngredientRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Ingredient struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ingredient) Reset()         { *m = Ingredient{} }
func (m *Ingredient) String() string { return proto.CompactTextString(m) }
func (*Ingredient) ProtoMessage()    {}
func (*Ingredient) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{2}
}

func (m *Ingredient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ingredient.Unmarshal(m, b)
}
func (m *Ingredient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ingredient.Marshal(b, m, deterministic)
}
func (m *Ingredient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingredient.Merge(m, src)
}
func (m *Ingredient) XXX_Size() int {
	return xxx_messageInfo_Ingredient.Size(m)
}
func (m *Ingredient) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingredient.DiscardUnknown(m)
}

var xxx_messageInfo_Ingredient proto.InternalMessageInfo

func (m *Ingredient) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ingredient) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Ingredient) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Ingredient) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MultipleIngredient struct {
	Ingredients          []*Ingredient `protobuf:"bytes,1,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MultipleIngredient) Reset()         { *m = MultipleIngredient{} }
func (m *MultipleIngredient) String() string { return proto.CompactTextString(m) }
func (*MultipleIngredient) ProtoMessage()    {}
func (*MultipleIngredient) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{3}
}

func (m *MultipleIngredient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultipleIngredient.Unmarshal(m, b)
}
func (m *MultipleIngredient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultipleIngredient.Marshal(b, m, deterministic)
}
func (m *MultipleIngredient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultipleIngredient.Merge(m, src)
}
func (m *MultipleIngredient) XXX_Size() int {
	return xxx_messageInfo_MultipleIngredient.Size(m)
}
func (m *MultipleIngredient) XXX_DiscardUnknown() {
	xxx_messageInfo_MultipleIngredient.DiscardUnknown(m)
}

var xxx_messageInfo_MultipleIngredient proto.InternalMessageInfo

func (m *MultipleIngredient) GetIngredients() []*Ingredient {
	if m != nil {
		return m.Ingredients
	}
	return nil
}

type RecipeRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecipeRequest) Reset()         { *m = RecipeRequest{} }
func (m *RecipeRequest) String() string { return proto.CompactTextString(m) }
func (*RecipeRequest) ProtoMessage()    {}
func (*RecipeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{4}
}

func (m *RecipeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecipeRequest.Unmarshal(m, b)
}
func (m *RecipeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecipeRequest.Marshal(b, m, deterministic)
}
func (m *RecipeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecipeRequest.Merge(m, src)
}
func (m *RecipeRequest) XXX_Size() int {
	return xxx_messageInfo_RecipeRequest.Size(m)
}
func (m *RecipeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecipeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecipeRequest proto.InternalMessageInfo

func (m *RecipeRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Recipe struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Bread                int32    `protobuf:"varint,3,opt,name=bread,proto3" json:"bread,omitempty"`
	Meats                []int32  `protobuf:"varint,4,rep,packed,name=meats,proto3" json:"meats,omitempty"`
	Cheeses              []int32  `protobuf:"varint,5,rep,packed,name=cheeses,proto3" json:"cheeses,omitempty"`
	Toppings             []int32  `protobuf:"varint,6,rep,packed,name=toppings,proto3" json:"toppings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recipe) Reset()         { *m = Recipe{} }
func (m *Recipe) String() string { return proto.CompactTextString(m) }
func (*Recipe) ProtoMessage()    {}
func (*Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{5}
}

func (m *Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipe.Unmarshal(m, b)
}
func (m *Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipe.Marshal(b, m, deterministic)
}
func (m *Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipe.Merge(m, src)
}
func (m *Recipe) XXX_Size() int {
	return xxx_messageInfo_Recipe.Size(m)
}
func (m *Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_Recipe proto.InternalMessageInfo

func (m *Recipe) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Recipe) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Recipe) GetBread() int32 {
	if m != nil {
		return m.Bread
	}
	return 0
}

func (m *Recipe) GetMeats() []int32 {
	if m != nil {
		return m.Meats
	}
	return nil
}

func (m *Recipe) GetCheeses() []int32 {
	if m != nil {
		return m.Cheeses
	}
	return nil
}

func (m *Recipe) GetToppings() []int32 {
	if m != nil {
		return m.Toppings
	}
	return nil
}

type MultipleRecipe struct {
	Recipes              []*Recipe `protobuf:"bytes,1,rep,name=recipes,proto3" json:"recipes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MultipleRecipe) Reset()         { *m = MultipleRecipe{} }
func (m *MultipleRecipe) String() string { return proto.CompactTextString(m) }
func (*MultipleRecipe) ProtoMessage()    {}
func (*MultipleRecipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{6}
}

func (m *MultipleRecipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultipleRecipe.Unmarshal(m, b)
}
func (m *MultipleRecipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultipleRecipe.Marshal(b, m, deterministic)
}
func (m *MultipleRecipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultipleRecipe.Merge(m, src)
}
func (m *MultipleRecipe) XXX_Size() int {
	return xxx_messageInfo_MultipleRecipe.Size(m)
}
func (m *MultipleRecipe) XXX_DiscardUnknown() {
	xxx_messageInfo_MultipleRecipe.DiscardUnknown(m)
}

var xxx_messageInfo_MultipleRecipe proto.InternalMessageInfo

func (m *MultipleRecipe) GetRecipes() []*Recipe {
	if m != nil {
		return m.Recipes
	}
	return nil
}

type MenuItemRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MenuItemRequest) Reset()         { *m = MenuItemRequest{} }
func (m *MenuItemRequest) String() string { return proto.CompactTextString(m) }
func (*MenuItemRequest) ProtoMessage()    {}
func (*MenuItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{7}
}

func (m *MenuItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuItemRequest.Unmarshal(m, b)
}
func (m *MenuItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuItemRequest.Marshal(b, m, deterministic)
}
func (m *MenuItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuItemRequest.Merge(m, src)
}
func (m *MenuItemRequest) XXX_Size() int {
	return xxx_messageInfo_MenuItemRequest.Size(m)
}
func (m *MenuItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MenuItemRequest proto.InternalMessageInfo

func (m *MenuItemRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MenuItem struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MenuItem) Reset()         { *m = MenuItem{} }
func (m *MenuItem) String() string { return proto.CompactTextString(m) }
func (*MenuItem) ProtoMessage()    {}
func (*MenuItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{8}
}

func (m *MenuItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuItem.Unmarshal(m, b)
}
func (m *MenuItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuItem.Marshal(b, m, deterministic)
}
func (m *MenuItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuItem.Merge(m, src)
}
func (m *MenuItem) XXX_Size() int {
	return xxx_messageInfo_MenuItem.Size(m)
}
func (m *MenuItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuItem.DiscardUnknown(m)
}

var xxx_messageInfo_MenuItem proto.InternalMessageInfo

func (m *MenuItem) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MenuItem) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *MenuItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MultipleMenuItem struct {
	MenuItems            []*MenuItem `protobuf:"bytes,1,rep,name=menuItems,proto3" json:"menuItems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MultipleMenuItem) Reset()         { *m = MultipleMenuItem{} }
func (m *MultipleMenuItem) String() string { return proto.CompactTextString(m) }
func (*MultipleMenuItem) ProtoMessage()    {}
func (*MultipleMenuItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d271faa439776c6e, []int{9}
}

func (m *MultipleMenuItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultipleMenuItem.Unmarshal(m, b)
}
func (m *MultipleMenuItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultipleMenuItem.Marshal(b, m, deterministic)
}
func (m *MultipleMenuItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultipleMenuItem.Merge(m, src)
}
func (m *MultipleMenuItem) XXX_Size() int {
	return xxx_messageInfo_MultipleMenuItem.Size(m)
}
func (m *MultipleMenuItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MultipleMenuItem.DiscardUnknown(m)
}

var xxx_messageInfo_MultipleMenuItem proto.InternalMessageInfo

func (m *MultipleMenuItem) GetMenuItems() []*MenuItem {
	if m != nil {
		return m.MenuItems
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "sandwiches.Empty")
	proto.RegisterType((*IngredientRequest)(nil), "sandwiches.IngredientRequest")
	proto.RegisterType((*Ingredient)(nil), "sandwiches.Ingredient")
	proto.RegisterType((*MultipleIngredient)(nil), "sandwiches.MultipleIngredient")
	proto.RegisterType((*RecipeRequest)(nil), "sandwiches.RecipeRequest")
	proto.RegisterType((*Recipe)(nil), "sandwiches.Recipe")
	proto.RegisterType((*MultipleRecipe)(nil), "sandwiches.MultipleRecipe")
	proto.RegisterType((*MenuItemRequest)(nil), "sandwiches.MenuItemRequest")
	proto.RegisterType((*MenuItem)(nil), "sandwiches.MenuItem")
	proto.RegisterType((*MultipleMenuItem)(nil), "sandwiches.MultipleMenuItem")
}

func init() { proto.RegisterFile("sandwiches.proto", fileDescriptor_d271faa439776c6e) }

var fileDescriptor_d271faa439776c6e = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0x3a, 0x71, 0xd2, 0x4c, 0x68, 0x48, 0x47, 0x15, 0x32, 0xa1, 0x94, 0xb0, 0x5c, 0x72,
	0x40, 0x3d, 0x98, 0x0b, 0x07, 0x40, 0xa2, 0xa2, 0x84, 0x1e, 0xca, 0xc1, 0x47, 0x6e, 0x69, 0x3c,
	0x6a, 0x57, 0xaa, 0x9d, 0xc5, 0xbb, 0x11, 0xea, 0x1f, 0xe0, 0xc6, 0x1f, 0xe0, 0xc0, 0x6f, 0x45,
	0xf1, 0x7a, 0xd7, 0x5b, 0x58, 0xdf, 0xe6, 0xe3, 0xed, 0xf8, 0xbd, 0x37, 0x23, 0xc3, 0x4c, 0xad,
	0xcb, 0xfc, 0x87, 0xd8, 0xdc, 0x92, 0x3a, 0x93, 0xd5, 0x56, 0x6f, 0x11, 0xda, 0x0a, 0x1f, 0x41,
	0x7c, 0x51, 0x48, 0x7d, 0xcf, 0x5f, 0xc1, 0xd1, 0x65, 0x79, 0x53, 0x51, 0x2e, 0xa8, 0xd4, 0x19,
	0x7d, 0xdf, 0x91, 0xd2, 0x38, 0x85, 0x48, 0xe4, 0x09, 0x5b, 0xb0, 0x65, 0x9c, 0x45, 0x22, 0xe7,
	0xdf, 0x00, 0x5a, 0x10, 0x22, 0x0c, 0xca, 0x75, 0x41, 0x75, 0x7f, 0x9c, 0xd5, 0x31, 0x1e, 0x43,
	0x2c, 0x2b, 0xb1, 0xa1, 0x24, 0x5a, 0xb0, 0x25, 0xcb, 0x4c, 0xb2, 0x47, 0xea, 0x7b, 0x49, 0x49,
	0xdf, 0x20, 0xf7, 0x71, 0x33, 0x7b, 0xe0, 0x66, 0x7f, 0x05, 0xbc, 0xda, 0xdd, 0x69, 0x21, 0xef,
	0xc8, 0xfb, 0xc6, 0x5b, 0x98, 0x08, 0x97, 0xa9, 0x84, 0x2d, 0xfa, 0xcb, 0x49, 0xfa, 0xe4, 0xcc,
	0xd3, 0xe4, 0xb1, 0xf6, 0xa1, 0xfc, 0x05, 0x1c, 0x66, 0xb4, 0x11, 0x92, 0xba, 0xc4, 0xfc, 0x66,
	0x30, 0x34, 0x88, 0xa0, 0x12, 0x03, 0x8f, 0x2c, 0x7c, 0xaf, 0xec, 0xba, 0xa2, 0x75, 0x5e, 0x8b,
	0x88, 0x33, 0x93, 0x60, 0x02, 0x71, 0x41, 0x6b, 0xad, 0x92, 0xc1, 0xa2, 0xbf, 0x8c, 0xcf, 0xa3,
	0x19, 0xcb, 0x4c, 0x01, 0x4f, 0x60, 0xb4, 0xb9, 0x25, 0x52, 0xa4, 0x92, 0xd8, 0xf5, 0x6c, 0x09,
	0x4f, 0xe1, 0x40, 0x6f, 0xa5, 0x14, 0xe5, 0x8d, 0x4a, 0x86, 0xae, 0xed, 0x6a, 0xfc, 0x03, 0x4c,
	0xad, 0x1b, 0x0d, 0xc7, 0xd7, 0x30, 0xaa, 0xea, 0xc8, 0xba, 0x80, 0xbe, 0x0b, 0x8d, 0x54, 0x0b,
	0xe1, 0x2f, 0xe1, 0xf1, 0x15, 0x95, 0xbb, 0x4b, 0x4d, 0x45, 0x97, 0xfe, 0x4f, 0x70, 0x60, 0x21,
	0xff, 0xf6, 0xba, 0xd7, 0x58, 0xdb, 0xd4, 0x6f, 0x6d, 0xe2, 0x9f, 0x61, 0x66, 0x89, 0xba, 0x69,
	0x29, 0x8c, 0x8b, 0x26, 0xb6, 0x64, 0x8f, 0x7d, 0xb2, 0x8e, 0x59, 0x0b, 0x4b, 0xff, 0x30, 0x98,
	0xb4, 0xab, 0x54, 0xf8, 0x05, 0x0e, 0x57, 0xa4, 0xbd, 0x4b, 0x78, 0xde, 0xb1, 0x74, 0xa3, 0x6e,
	0xde, 0x71, 0x13, 0xbc, 0x87, 0x17, 0x30, 0x7d, 0x30, 0x49, 0xe1, 0x91, 0x8f, 0xad, 0xcf, 0x7f,
	0x7e, 0xfa, 0x80, 0xdf, 0x7f, 0x77, 0xc8, 0x7b, 0xe9, 0x4f, 0x06, 0x23, 0xe3, 0xb2, 0xc2, 0x77,
	0x30, 0x5e, 0x91, 0x6e, 0x16, 0xf3, 0x34, 0xb0, 0x87, 0x86, 0x54, 0x60, 0x45, 0xbc, 0x87, 0xef,
	0x01, 0xdc, 0xeb, 0x20, 0x99, 0x79, 0x88, 0x8c, 0x7d, 0x9e, 0xfe, 0x62, 0x30, 0xd8, 0x3b, 0x88,
	0xe7, 0x30, 0x59, 0x91, 0x76, 0xae, 0x3f, 0x0b, 0x5a, 0xdc, 0x30, 0x09, 0xfa, 0xcf, 0x7b, 0xf8,
	0x11, 0x1e, 0x79, 0x33, 0x82, 0x6c, 0x4e, 0x42, 0x6c, 0xda, 0x11, 0xd7, 0xc3, 0xfa, 0xaf, 0xf2,
	0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x33, 0x78, 0x8c, 0x69, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IngredientsClient is the client API for Ingredients service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IngredientsClient interface {
	GetIngredient(ctx context.Context, in *IngredientRequest, opts ...grpc.CallOption) (*Ingredient, error)
	GetIngredients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultipleIngredient, error)
}

type ingredientsClient struct {
	cc *grpc.ClientConn
}

func NewIngredientsClient(cc *grpc.ClientConn) IngredientsClient {
	return &ingredientsClient{cc}
}

func (c *ingredientsClient) GetIngredient(ctx context.Context, in *IngredientRequest, opts ...grpc.CallOption) (*Ingredient, error) {
	out := new(Ingredient)
	err := c.cc.Invoke(ctx, "/sandwiches.Ingredients/GetIngredient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingredientsClient) GetIngredients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultipleIngredient, error) {
	out := new(MultipleIngredient)
	err := c.cc.Invoke(ctx, "/sandwiches.Ingredients/GetIngredients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngredientsServer is the server API for Ingredients service.
type IngredientsServer interface {
	GetIngredient(context.Context, *IngredientRequest) (*Ingredient, error)
	GetIngredients(context.Context, *Empty) (*MultipleIngredient, error)
}

// UnimplementedIngredientsServer can be embedded to have forward compatible implementations.
type UnimplementedIngredientsServer struct {
}

func (*UnimplementedIngredientsServer) GetIngredient(ctx context.Context, req *IngredientRequest) (*Ingredient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIngredient not implemented")
}
func (*UnimplementedIngredientsServer) GetIngredients(ctx context.Context, req *Empty) (*MultipleIngredient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIngredients not implemented")
}

func RegisterIngredientsServer(s *grpc.Server, srv IngredientsServer) {
	s.RegisterService(&_Ingredients_serviceDesc, srv)
}

func _Ingredients_GetIngredient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngredientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngredientsServer).GetIngredient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwiches.Ingredients/GetIngredient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngredientsServer).GetIngredient(ctx, req.(*IngredientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ingredients_GetIngredients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngredientsServer).GetIngredients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwiches.Ingredients/GetIngredients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngredientsServer).GetIngredients(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ingredients_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sandwiches.Ingredients",
	HandlerType: (*IngredientsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIngredient",
			Handler:    _Ingredients_GetIngredient_Handler,
		},
		{
			MethodName: "GetIngredients",
			Handler:    _Ingredients_GetIngredients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sandwiches.proto",
}

// RecipesClient is the client API for Recipes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecipesClient interface {
	GetRecipe(ctx context.Context, in *RecipeRequest, opts ...grpc.CallOption) (*Recipe, error)
	GetRecipes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultipleRecipe, error)
}

type recipesClient struct {
	cc *grpc.ClientConn
}

func NewRecipesClient(cc *grpc.ClientConn) RecipesClient {
	return &recipesClient{cc}
}

func (c *recipesClient) GetRecipe(ctx context.Context, in *RecipeRequest, opts ...grpc.CallOption) (*Recipe, error) {
	out := new(Recipe)
	err := c.cc.Invoke(ctx, "/sandwiches.Recipes/GetRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesClient) GetRecipes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultipleRecipe, error) {
	out := new(MultipleRecipe)
	err := c.cc.Invoke(ctx, "/sandwiches.Recipes/GetRecipes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipesServer is the server API for Recipes service.
type RecipesServer interface {
	GetRecipe(context.Context, *RecipeRequest) (*Recipe, error)
	GetRecipes(context.Context, *Empty) (*MultipleRecipe, error)
}

// UnimplementedRecipesServer can be embedded to have forward compatible implementations.
type UnimplementedRecipesServer struct {
}

func (*UnimplementedRecipesServer) GetRecipe(ctx context.Context, req *RecipeRequest) (*Recipe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (*UnimplementedRecipesServer) GetRecipes(ctx context.Context, req *Empty) (*MultipleRecipe, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipes not implemented")
}

func RegisterRecipesServer(s *grpc.Server, srv RecipesServer) {
	s.RegisterService(&_Recipes_serviceDesc, srv)
}

func _Recipes_GetRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServer).GetRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwiches.Recipes/GetRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServer).GetRecipe(ctx, req.(*RecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recipes_GetRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServer).GetRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwiches.Recipes/GetRecipes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServer).GetRecipes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recipes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sandwiches.Recipes",
	HandlerType: (*RecipesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecipe",
			Handler:    _Recipes_GetRecipe_Handler,
		},
		{
			MethodName: "GetRecipes",
			Handler:    _Recipes_GetRecipes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sandwiches.proto",
}

// MenuClient is the client API for Menu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MenuClient interface {
	GetMenuItem(ctx context.Context, in *MenuItemRequest, opts ...grpc.CallOption) (*MenuItem, error)
	GetMenuItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultipleMenuItem, error)
}

type menuClient struct {
	cc *grpc.ClientConn
}

func NewMenuClient(cc *grpc.ClientConn) MenuClient {
	return &menuClient{cc}
}

func (c *menuClient) GetMenuItem(ctx context.Context, in *MenuItemRequest, opts ...grpc.CallOption) (*MenuItem, error) {
	out := new(MenuItem)
	err := c.cc.Invoke(ctx, "/sandwiches.Menu/GetMenuItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuClient) GetMenuItems(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MultipleMenuItem, error) {
	out := new(MultipleMenuItem)
	err := c.cc.Invoke(ctx, "/sandwiches.Menu/GetMenuItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServer is the server API for Menu service.
type MenuServer interface {
	GetMenuItem(context.Context, *MenuItemRequest) (*MenuItem, error)
	GetMenuItems(context.Context, *Empty) (*MultipleMenuItem, error)
}

// UnimplementedMenuServer can be embedded to have forward compatible implementations.
type UnimplementedMenuServer struct {
}

func (*UnimplementedMenuServer) GetMenuItem(ctx context.Context, req *MenuItemRequest) (*MenuItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuItem not implemented")
}
func (*UnimplementedMenuServer) GetMenuItems(ctx context.Context, req *Empty) (*MultipleMenuItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuItems not implemented")
}

func RegisterMenuServer(s *grpc.Server, srv MenuServer) {
	s.RegisterService(&_Menu_serviceDesc, srv)
}

func _Menu_GetMenuItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServer).GetMenuItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwiches.Menu/GetMenuItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServer).GetMenuItem(ctx, req.(*MenuItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Menu_GetMenuItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServer).GetMenuItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandwiches.Menu/GetMenuItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServer).GetMenuItems(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Menu_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sandwiches.Menu",
	HandlerType: (*MenuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMenuItem",
			Handler:    _Menu_GetMenuItem_Handler,
		},
		{
			MethodName: "GetMenuItems",
			Handler:    _Menu_GetMenuItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sandwiches.proto",
}
