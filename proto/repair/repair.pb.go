// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repair/repair.proto

package repair

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/emptypb"
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

type Status int32

const (
	Status_PROCESSING Status = 0
	Status_SUCCESS    Status = 1
	Status_FAILED     Status = 2
	Status_ABANDON    Status = 3
)

var Status_name = map[int32]string{
	0: "PROCESSING",
	1: "SUCCESS",
	2: "FAILED",
	3: "ABANDON",
}

var Status_value = map[string]int32{
	"PROCESSING": 0,
	"SUCCESS":    1,
	"FAILED":     2,
	"ABANDON":    3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{0}
}

type FixKV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	Val                  string   `protobuf:"bytes,2,opt,name=val,proto3" json:"val"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FixKV) Reset()         { *m = FixKV{} }
func (m *FixKV) String() string { return proto.CompactTextString(m) }
func (*FixKV) ProtoMessage()    {}
func (*FixKV) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{0}
}

func (m *FixKV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixKV.Unmarshal(m, b)
}
func (m *FixKV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixKV.Marshal(b, m, deterministic)
}
func (m *FixKV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixKV.Merge(m, src)
}
func (m *FixKV) XXX_Size() int {
	return xxx_messageInfo_FixKV.Size(m)
}
func (m *FixKV) XXX_DiscardUnknown() {
	xxx_messageInfo_FixKV.DiscardUnknown(m)
}

var xxx_messageInfo_FixKV proto.InternalMessageInfo

func (m *FixKV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FixKV) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type RepairRequest struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source"`
	Market               string   `protobuf:"bytes,2,opt,name=market,proto3" json:"market"`
	CounterIds           []string `protobuf:"bytes,3,rep,name=counter_ids,json=counterIds,proto3" json:"counter_ids"`
	Start                int64    `protobuf:"varint,4,opt,name=start,proto3" json:"start"`
	End                  int64    `protobuf:"varint,5,opt,name=end,proto3" json:"end"`
	LineType             int64    `protobuf:"varint,6,opt,name=line_type,json=lineType,proto3" json:"line_type"`
	Workers              int64    `protobuf:"varint,7,opt,name=workers,proto3" json:"workers"`
	FixKvs               []*FixKV `protobuf:"bytes,8,rep,name=fix_kvs,json=fixKvs,proto3" json:"fix_kvs"`
	RepairAll            bool     `protobuf:"varint,9,opt,name=repair_all,json=repairAll,proto3" json:"repair_all"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepairRequest) Reset()         { *m = RepairRequest{} }
func (m *RepairRequest) String() string { return proto.CompactTextString(m) }
func (*RepairRequest) ProtoMessage()    {}
func (*RepairRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{1}
}

func (m *RepairRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairRequest.Unmarshal(m, b)
}
func (m *RepairRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairRequest.Marshal(b, m, deterministic)
}
func (m *RepairRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairRequest.Merge(m, src)
}
func (m *RepairRequest) XXX_Size() int {
	return xxx_messageInfo_RepairRequest.Size(m)
}
func (m *RepairRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepairRequest proto.InternalMessageInfo

func (m *RepairRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RepairRequest) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *RepairRequest) GetCounterIds() []string {
	if m != nil {
		return m.CounterIds
	}
	return nil
}

func (m *RepairRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *RepairRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *RepairRequest) GetLineType() int64 {
	if m != nil {
		return m.LineType
	}
	return 0
}

func (m *RepairRequest) GetWorkers() int64 {
	if m != nil {
		return m.Workers
	}
	return 0
}

func (m *RepairRequest) GetFixKvs() []*FixKV {
	if m != nil {
		return m.FixKvs
	}
	return nil
}

func (m *RepairRequest) GetRepairAll() bool {
	if m != nil {
		return m.RepairAll
	}
	return false
}

type RepairResponse struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	ProgressId           int64    `protobuf:"varint,3,opt,name=progress_id,json=progressId,proto3" json:"progress_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepairResponse) Reset()         { *m = RepairResponse{} }
func (m *RepairResponse) String() string { return proto.CompactTextString(m) }
func (*RepairResponse) ProtoMessage()    {}
func (*RepairResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{2}
}

func (m *RepairResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairResponse.Unmarshal(m, b)
}
func (m *RepairResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairResponse.Marshal(b, m, deterministic)
}
func (m *RepairResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairResponse.Merge(m, src)
}
func (m *RepairResponse) XXX_Size() int {
	return xxx_messageInfo_RepairResponse.Size(m)
}
func (m *RepairResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepairResponse proto.InternalMessageInfo

func (m *RepairResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RepairResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *RepairResponse) GetProgressId() int64 {
	if m != nil {
		return m.ProgressId
	}
	return 0
}

type ProgressRequest struct {
	ProgressId           int64    `protobuf:"varint,1,opt,name=progress_id,json=progressId,proto3" json:"progress_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProgressRequest) Reset()         { *m = ProgressRequest{} }
func (m *ProgressRequest) String() string { return proto.CompactTextString(m) }
func (*ProgressRequest) ProtoMessage()    {}
func (*ProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{3}
}

func (m *ProgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProgressRequest.Unmarshal(m, b)
}
func (m *ProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProgressRequest.Marshal(b, m, deterministic)
}
func (m *ProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProgressRequest.Merge(m, src)
}
func (m *ProgressRequest) XXX_Size() int {
	return xxx_messageInfo_ProgressRequest.Size(m)
}
func (m *ProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProgressRequest proto.InternalMessageInfo

func (m *ProgressRequest) GetProgressId() int64 {
	if m != nil {
		return m.ProgressId
	}
	return 0
}

type ProgressResponse struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=repair.Status" json:"status"`
	Success              int64    `protobuf:"varint,4,opt,name=success,proto3" json:"success"`
	Failed               int64    `protobuf:"varint,5,opt,name=failed,proto3" json:"failed"`
	Total                int64    `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
	FailedStocks         []string `protobuf:"bytes,7,rep,name=failed_stocks,json=failedStocks,proto3" json:"failed_stocks"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProgressResponse) Reset()         { *m = ProgressResponse{} }
func (m *ProgressResponse) String() string { return proto.CompactTextString(m) }
func (*ProgressResponse) ProtoMessage()    {}
func (*ProgressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{4}
}

func (m *ProgressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProgressResponse.Unmarshal(m, b)
}
func (m *ProgressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProgressResponse.Marshal(b, m, deterministic)
}
func (m *ProgressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProgressResponse.Merge(m, src)
}
func (m *ProgressResponse) XXX_Size() int {
	return xxx_messageInfo_ProgressResponse.Size(m)
}
func (m *ProgressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProgressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProgressResponse proto.InternalMessageInfo

func (m *ProgressResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ProgressResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *ProgressResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PROCESSING
}

func (m *ProgressResponse) GetSuccess() int64 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *ProgressResponse) GetFailed() int64 {
	if m != nil {
		return m.Failed
	}
	return 0
}

func (m *ProgressResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ProgressResponse) GetFailedStocks() []string {
	if m != nil {
		return m.FailedStocks
	}
	return nil
}

type SyncFutuStockRequest struct {
	Market               string   `protobuf:"bytes,1,opt,name=market,proto3" json:"market"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncFutuStockRequest) Reset()         { *m = SyncFutuStockRequest{} }
func (m *SyncFutuStockRequest) String() string { return proto.CompactTextString(m) }
func (*SyncFutuStockRequest) ProtoMessage()    {}
func (*SyncFutuStockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{5}
}

func (m *SyncFutuStockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncFutuStockRequest.Unmarshal(m, b)
}
func (m *SyncFutuStockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncFutuStockRequest.Marshal(b, m, deterministic)
}
func (m *SyncFutuStockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncFutuStockRequest.Merge(m, src)
}
func (m *SyncFutuStockRequest) XXX_Size() int {
	return xxx_messageInfo_SyncFutuStockRequest.Size(m)
}
func (m *SyncFutuStockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncFutuStockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncFutuStockRequest proto.InternalMessageInfo

func (m *SyncFutuStockRequest) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

type SyncFutuStockResponse struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=repair.Status" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncFutuStockResponse) Reset()         { *m = SyncFutuStockResponse{} }
func (m *SyncFutuStockResponse) String() string { return proto.CompactTextString(m) }
func (*SyncFutuStockResponse) ProtoMessage()    {}
func (*SyncFutuStockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{6}
}

func (m *SyncFutuStockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncFutuStockResponse.Unmarshal(m, b)
}
func (m *SyncFutuStockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncFutuStockResponse.Marshal(b, m, deterministic)
}
func (m *SyncFutuStockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncFutuStockResponse.Merge(m, src)
}
func (m *SyncFutuStockResponse) XXX_Size() int {
	return xxx_messageInfo_SyncFutuStockResponse.Size(m)
}
func (m *SyncFutuStockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncFutuStockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncFutuStockResponse proto.InternalMessageInfo

func (m *SyncFutuStockResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SyncFutuStockResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *SyncFutuStockResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PROCESSING
}

type PlateIndexRequest struct {
	CounterId            string   `protobuf:"bytes,1,opt,name=counter_id,json=counterId,proto3" json:"counter_id"`
	IpoDate              string   `protobuf:"bytes,2,opt,name=ipo_date,json=ipoDate,proto3" json:"ipo_date"`
	IsAuto               bool     `protobuf:"varint,3,opt,name=is_auto,json=isAuto,proto3" json:"is_auto"`
	IsForce              bool     `protobuf:"varint,4,opt,name=is_force,json=isForce,proto3" json:"is_force"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlateIndexRequest) Reset()         { *m = PlateIndexRequest{} }
func (m *PlateIndexRequest) String() string { return proto.CompactTextString(m) }
func (*PlateIndexRequest) ProtoMessage()    {}
func (*PlateIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{7}
}

func (m *PlateIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlateIndexRequest.Unmarshal(m, b)
}
func (m *PlateIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlateIndexRequest.Marshal(b, m, deterministic)
}
func (m *PlateIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlateIndexRequest.Merge(m, src)
}
func (m *PlateIndexRequest) XXX_Size() int {
	return xxx_messageInfo_PlateIndexRequest.Size(m)
}
func (m *PlateIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlateIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlateIndexRequest proto.InternalMessageInfo

func (m *PlateIndexRequest) GetCounterId() string {
	if m != nil {
		return m.CounterId
	}
	return ""
}

func (m *PlateIndexRequest) GetIpoDate() string {
	if m != nil {
		return m.IpoDate
	}
	return ""
}

func (m *PlateIndexRequest) GetIsAuto() bool {
	if m != nil {
		return m.IsAuto
	}
	return false
}

func (m *PlateIndexRequest) GetIsForce() bool {
	if m != nil {
		return m.IsForce
	}
	return false
}

type PlateIndexResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=repair.Status" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlateIndexResponse) Reset()         { *m = PlateIndexResponse{} }
func (m *PlateIndexResponse) String() string { return proto.CompactTextString(m) }
func (*PlateIndexResponse) ProtoMessage()    {}
func (*PlateIndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{8}
}

func (m *PlateIndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlateIndexResponse.Unmarshal(m, b)
}
func (m *PlateIndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlateIndexResponse.Marshal(b, m, deterministic)
}
func (m *PlateIndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlateIndexResponse.Merge(m, src)
}
func (m *PlateIndexResponse) XXX_Size() int {
	return xxx_messageInfo_PlateIndexResponse.Size(m)
}
func (m *PlateIndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PlateIndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PlateIndexResponse proto.InternalMessageInfo

func (m *PlateIndexResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PlateIndexResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *PlateIndexResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PROCESSING
}

type StockModificationRequest struct {
	OldCounterId         string   `protobuf:"bytes,1,opt,name=old_counter_id,json=oldCounterId,proto3" json:"old_counter_id"`
	NewCounterId         string   `protobuf:"bytes,2,opt,name=new_counter_id,json=newCounterId,proto3" json:"new_counter_id"`
	ModifyTimestamp      int64    `protobuf:"varint,3,opt,name=modify_timestamp,json=modifyTimestamp,proto3" json:"modify_timestamp"`
	Market               string   `protobuf:"bytes,4,opt,name=market,proto3" json:"market"`
	IsForce              bool     `protobuf:"varint,5,opt,name=is_force,json=isForce,proto3" json:"is_force"`
	ClearDest            bool     `protobuf:"varint,6,opt,name=clearDest,proto3" json:"clearDest"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StockModificationRequest) Reset()         { *m = StockModificationRequest{} }
func (m *StockModificationRequest) String() string { return proto.CompactTextString(m) }
func (*StockModificationRequest) ProtoMessage()    {}
func (*StockModificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{9}
}

func (m *StockModificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockModificationRequest.Unmarshal(m, b)
}
func (m *StockModificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockModificationRequest.Marshal(b, m, deterministic)
}
func (m *StockModificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockModificationRequest.Merge(m, src)
}
func (m *StockModificationRequest) XXX_Size() int {
	return xxx_messageInfo_StockModificationRequest.Size(m)
}
func (m *StockModificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StockModificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StockModificationRequest proto.InternalMessageInfo

func (m *StockModificationRequest) GetOldCounterId() string {
	if m != nil {
		return m.OldCounterId
	}
	return ""
}

func (m *StockModificationRequest) GetNewCounterId() string {
	if m != nil {
		return m.NewCounterId
	}
	return ""
}

func (m *StockModificationRequest) GetModifyTimestamp() int64 {
	if m != nil {
		return m.ModifyTimestamp
	}
	return 0
}

func (m *StockModificationRequest) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *StockModificationRequest) GetIsForce() bool {
	if m != nil {
		return m.IsForce
	}
	return false
}

func (m *StockModificationRequest) GetClearDest() bool {
	if m != nil {
		return m.ClearDest
	}
	return false
}

type StockModificationResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=repair.Status" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StockModificationResponse) Reset()         { *m = StockModificationResponse{} }
func (m *StockModificationResponse) String() string { return proto.CompactTextString(m) }
func (*StockModificationResponse) ProtoMessage()    {}
func (*StockModificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{10}
}

func (m *StockModificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockModificationResponse.Unmarshal(m, b)
}
func (m *StockModificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockModificationResponse.Marshal(b, m, deterministic)
}
func (m *StockModificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockModificationResponse.Merge(m, src)
}
func (m *StockModificationResponse) XXX_Size() int {
	return xxx_messageInfo_StockModificationResponse.Size(m)
}
func (m *StockModificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StockModificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StockModificationResponse proto.InternalMessageInfo

func (m *StockModificationResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StockModificationResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *StockModificationResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PROCESSING
}

type SetShortDataRequest struct {
	StartDate            string   `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetShortDataRequest) Reset()         { *m = SetShortDataRequest{} }
func (m *SetShortDataRequest) String() string { return proto.CompactTextString(m) }
func (*SetShortDataRequest) ProtoMessage()    {}
func (*SetShortDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{11}
}

func (m *SetShortDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetShortDataRequest.Unmarshal(m, b)
}
func (m *SetShortDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetShortDataRequest.Marshal(b, m, deterministic)
}
func (m *SetShortDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetShortDataRequest.Merge(m, src)
}
func (m *SetShortDataRequest) XXX_Size() int {
	return xxx_messageInfo_SetShortDataRequest.Size(m)
}
func (m *SetShortDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetShortDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetShortDataRequest proto.InternalMessageInfo

func (m *SetShortDataRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

type SetShortDataResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=repair.Status" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetShortDataResponse) Reset()         { *m = SetShortDataResponse{} }
func (m *SetShortDataResponse) String() string { return proto.CompactTextString(m) }
func (*SetShortDataResponse) ProtoMessage()    {}
func (*SetShortDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a10c3a264744c38, []int{12}
}

func (m *SetShortDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetShortDataResponse.Unmarshal(m, b)
}
func (m *SetShortDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetShortDataResponse.Marshal(b, m, deterministic)
}
func (m *SetShortDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetShortDataResponse.Merge(m, src)
}
func (m *SetShortDataResponse) XXX_Size() int {
	return xxx_messageInfo_SetShortDataResponse.Size(m)
}
func (m *SetShortDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetShortDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetShortDataResponse proto.InternalMessageInfo

func (m *SetShortDataResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SetShortDataResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *SetShortDataResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PROCESSING
}

func init() {
	proto.RegisterEnum("repair.Status", Status_name, Status_value)
	proto.RegisterType((*FixKV)(nil), "repair.FixKV")
	proto.RegisterType((*RepairRequest)(nil), "repair.RepairRequest")
	proto.RegisterType((*RepairResponse)(nil), "repair.RepairResponse")
	proto.RegisterType((*ProgressRequest)(nil), "repair.ProgressRequest")
	proto.RegisterType((*ProgressResponse)(nil), "repair.ProgressResponse")
	proto.RegisterType((*SyncFutuStockRequest)(nil), "repair.SyncFutuStockRequest")
	proto.RegisterType((*SyncFutuStockResponse)(nil), "repair.SyncFutuStockResponse")
	proto.RegisterType((*PlateIndexRequest)(nil), "repair.PlateIndexRequest")
	proto.RegisterType((*PlateIndexResponse)(nil), "repair.PlateIndexResponse")
	proto.RegisterType((*StockModificationRequest)(nil), "repair.StockModificationRequest")
	proto.RegisterType((*StockModificationResponse)(nil), "repair.StockModificationResponse")
	proto.RegisterType((*SetShortDataRequest)(nil), "repair.SetShortDataRequest")
	proto.RegisterType((*SetShortDataResponse)(nil), "repair.SetShortDataResponse")
}

func init() { proto.RegisterFile("repair/repair.proto", fileDescriptor_2a10c3a264744c38) }

var fileDescriptor_2a10c3a264744c38 = []byte{
	// 1005 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5f, 0x53, 0x1b, 0x37,
	0x10, 0x8f, 0x63, 0xf0, 0x9f, 0x35, 0x18, 0x47, 0x10, 0x38, 0x0c, 0x4c, 0xdd, 0x6b, 0x26, 0x43,
	0xdb, 0x19, 0x33, 0x43, 0xfb, 0x9c, 0x19, 0x83, 0x21, 0x71, 0x29, 0x7f, 0x7a, 0x97, 0xe4, 0xa1,
	0x0f, 0xbd, 0x51, 0x7c, 0x32, 0xa8, 0x3e, 0x5b, 0x57, 0x49, 0x07, 0xf8, 0xb1, 0x2f, 0xfd, 0x76,
	0x7d, 0xeb, 0x27, 0xe8, 0x27, 0xe9, 0x48, 0x3a, 0x9d, 0x6d, 0xec, 0xb6, 0x13, 0x27, 0x3c, 0x59,
	0xfb, 0x5b, 0xed, 0x6a, 0xf7, 0xa7, 0x9f, 0xf7, 0x04, 0xeb, 0x9c, 0xc4, 0x98, 0xf2, 0x03, 0xf3,
	0xd3, 0x8c, 0x39, 0x93, 0x0c, 0x15, 0x8c, 0x55, 0xdf, 0xb9, 0x66, 0xec, 0x3a, 0x22, 0x07, 0x1a,
	0xfd, 0x90, 0xf4, 0x0e, 0xc8, 0x20, 0x96, 0x23, 0xb3, 0xc9, 0xfd, 0x16, 0x96, 0x4f, 0xe9, 0xfd,
	0xd9, 0x7b, 0x54, 0x83, 0x7c, 0x9f, 0x8c, 0x9c, 0x5c, 0x23, 0xb7, 0x5f, 0xf6, 0xd4, 0x52, 0x21,
	0xb7, 0x38, 0x72, 0x9e, 0x1a, 0xe4, 0x16, 0x47, 0xee, 0x1f, 0x4f, 0x61, 0xd5, 0xd3, 0x49, 0x3d,
	0xf2, 0x5b, 0x42, 0x84, 0x44, 0x9b, 0x50, 0x10, 0x2c, 0xe1, 0x5d, 0x92, 0x06, 0xa6, 0x96, 0xc2,
	0x07, 0x98, 0xf7, 0x89, 0x4c, 0xc3, 0x53, 0x0b, 0x7d, 0x01, 0x95, 0x2e, 0x4b, 0x86, 0x92, 0xf0,
	0x80, 0x86, 0xc2, 0xc9, 0x37, 0xf2, 0xfb, 0x65, 0x0f, 0x52, 0xa8, 0x13, 0x0a, 0xb4, 0x01, 0xcb,
	0x42, 0x62, 0x2e, 0x9d, 0xa5, 0x46, 0x6e, 0x3f, 0xef, 0x19, 0x43, 0x95, 0x42, 0x86, 0xa1, 0xb3,
	0xac, 0x31, 0xb5, 0x44, 0x3b, 0x50, 0x8e, 0xe8, 0x90, 0x04, 0x72, 0x14, 0x13, 0xa7, 0xa0, 0xf1,
	0x92, 0x02, 0xde, 0x8e, 0x62, 0x82, 0x1c, 0x28, 0xde, 0x31, 0xde, 0x27, 0x5c, 0x38, 0x45, 0xed,
	0xb2, 0x26, 0x7a, 0x09, 0xc5, 0x1e, 0xbd, 0x0f, 0xfa, 0xb7, 0xc2, 0x29, 0x35, 0xf2, 0xfb, 0x95,
	0xc3, 0xd5, 0x66, 0xca, 0x99, 0x66, 0xc1, 0x2b, 0xf4, 0xe8, 0xfd, 0xd9, 0xad, 0x40, 0x7b, 0x00,
	0x06, 0x0f, 0x70, 0x14, 0x39, 0xe5, 0x46, 0x6e, 0xbf, 0xe4, 0x95, 0x0d, 0xd2, 0x8a, 0x22, 0xf7,
	0x17, 0xa8, 0x5a, 0x1e, 0x44, 0xcc, 0x86, 0x82, 0x20, 0x04, 0x4b, 0x5d, 0x16, 0x1a, 0x1a, 0xf2,
	0x9e, 0x5e, 0xa3, 0x2d, 0x28, 0x12, 0xce, 0x83, 0x81, 0xb8, 0xb6, 0x2c, 0x10, 0xce, 0xcf, 0xc5,
	0xb5, 0x62, 0x21, 0xe6, 0xec, 0x9a, 0x13, 0x21, 0x02, 0x1a, 0x3a, 0x79, 0x1d, 0x03, 0x16, 0xea,
	0x84, 0xee, 0x21, 0xac, 0x5d, 0xa5, 0x96, 0x65, 0xfa, 0x41, 0x4c, 0x6e, 0x26, 0xe6, 0xaf, 0x1c,
	0xd4, 0xc6, 0x41, 0x8b, 0x94, 0xf5, 0x12, 0x0a, 0x42, 0x62, 0x99, 0x08, 0x5d, 0x51, 0xf5, 0xb0,
	0x6a, 0xb9, 0xf1, 0x35, 0xea, 0xa5, 0x5e, 0x45, 0xaf, 0x48, 0xba, 0x5d, 0x22, 0x44, 0x7a, 0x4b,
	0xd6, 0x54, 0xd7, 0xde, 0xc3, 0x34, 0x22, 0xf6, 0xaa, 0x52, 0x4b, 0xdd, 0xaa, 0x64, 0x12, 0x47,
	0xe9, 0x4d, 0x19, 0x03, 0x7d, 0x05, 0xab, 0xc6, 0x1f, 0x08, 0xc9, 0xba, 0x7d, 0x75, 0x59, 0x4a,
	0x0e, 0x2b, 0x06, 0xf4, 0x35, 0xe6, 0x36, 0x61, 0xc3, 0x1f, 0x0d, 0xbb, 0xa7, 0x89, 0x4c, 0x34,
	0x32, 0xa1, 0xbc, 0x54, 0x61, 0xb9, 0x49, 0x85, 0xb9, 0x11, 0x3c, 0x7f, 0xb0, 0xff, 0x11, 0xa9,
	0x70, 0x7f, 0xcf, 0xc1, 0xb3, 0xab, 0x08, 0x4b, 0xd2, 0x19, 0x86, 0xe4, 0xde, 0xd6, 0xb6, 0x07,
	0x30, 0x56, 0x79, 0x5a, 0x5f, 0x39, 0x13, 0x39, 0xda, 0x86, 0x12, 0x8d, 0x59, 0x10, 0x62, 0x49,
	0xd2, 0x63, 0x8b, 0x34, 0x66, 0x6d, 0x2c, 0x75, 0x41, 0x54, 0x04, 0x38, 0x91, 0x4c, 0x1f, 0x5c,
	0xf2, 0x0a, 0x54, 0xb4, 0x12, 0xc9, 0x74, 0x8c, 0x08, 0x7a, 0x4c, 0xfd, 0xd5, 0x96, 0xb4, 0xa7,
	0x48, 0xc5, 0xa9, 0x32, 0x5d, 0x0a, 0x68, 0xb2, 0x84, 0x39, 0xed, 0xae, 0x7e, 0xae, 0x76, 0xff,
	0xce, 0x81, 0xa3, 0x59, 0x3d, 0x67, 0x21, 0xed, 0xd1, 0x2e, 0x96, 0x94, 0x0d, 0x6d, 0xd7, 0x2f,
	0xa0, 0xca, 0xa2, 0x30, 0x98, 0xe9, 0x7c, 0x85, 0x45, 0xe1, 0x71, 0xd6, 0xfc, 0x0b, 0xa8, 0x0e,
	0xc9, 0xdd, 0xe4, 0x2e, 0x53, 0xca, 0xca, 0x90, 0xdc, 0x8d, 0x77, 0x7d, 0x0d, 0xb5, 0x81, 0x3a,
	0x62, 0x14, 0x48, 0x3a, 0x20, 0x42, 0xe2, 0x41, 0x9c, 0xfe, 0x4d, 0xd6, 0x0c, 0xfe, 0xd6, 0xc2,
	0x13, 0x42, 0x58, 0x9a, 0x1a, 0x35, 0x93, 0x8c, 0x2d, 0x4f, 0x31, 0x86, 0x76, 0xa1, 0xdc, 0x8d,
	0x08, 0xe6, 0x6d, 0x22, 0xa4, 0x96, 0x64, 0xc9, 0x1b, 0x03, 0x6e, 0x0c, 0xdb, 0x73, 0x7a, 0x7c,
	0x4c, 0x5a, 0xbf, 0x87, 0x75, 0x9f, 0x48, 0xff, 0x86, 0x71, 0xd9, 0xc6, 0x12, 0x4f, 0xc8, 0x48,
	0x8f, 0x3f, 0xa3, 0x94, 0x54, 0x46, 0x1a, 0x51, 0x5a, 0x71, 0xfb, 0xb0, 0x31, 0x1d, 0xf5, 0x88,
	0x25, 0x7e, 0xf3, 0x0a, 0x0a, 0x06, 0x41, 0x55, 0x80, 0x2b, 0xef, 0xf2, 0xf8, 0xc4, 0xf7, 0x3b,
	0x17, 0xaf, 0x6b, 0x4f, 0x50, 0x05, 0x8a, 0xfe, 0xbb, 0x63, 0x65, 0xd7, 0x72, 0x08, 0xa0, 0x70,
	0xda, 0xea, 0xfc, 0x78, 0xd2, 0xae, 0x3d, 0x55, 0x8e, 0xd6, 0x51, 0xeb, 0xa2, 0x7d, 0x79, 0x51,
	0xcb, 0x1f, 0xfe, 0x09, 0x50, 0xf9, 0x29, 0x61, 0x92, 0x98, 0xb9, 0x89, 0x5e, 0x41, 0xc5, 0xac,
	0xce, 0xd4, 0xd4, 0x46, 0xcf, 0xed, 0xb1, 0x53, 0x9f, 0x97, 0xfa, 0xe6, 0x43, 0xd8, 0xb4, 0xe8,
	0x3e, 0x41, 0xc7, 0x50, 0x33, 0xd8, 0x39, 0x1b, 0xca, 0x9b, 0x4f, 0x4d, 0xa2, 0xd4, 0xe4, 0xdf,
	0x60, 0x4e, 0xc4, 0xc7, 0x27, 0xb9, 0x80, 0xd5, 0xa9, 0x81, 0x83, 0x76, 0x33, 0x0a, 0xe7, 0xcc,
	0xad, 0xfa, 0xde, 0xbf, 0x78, 0xb3, 0x7c, 0x47, 0x50, 0x79, 0x4d, 0xa4, 0x9d, 0xe4, 0x68, 0xcb,
	0xee, 0x7f, 0xf0, 0x41, 0xa8, 0x3b, 0xb3, 0x8e, 0x89, 0xc6, 0x56, 0x7c, 0xc9, 0xe2, 0x4f, 0x4b,
	0x72, 0x04, 0x6b, 0x3e, 0x91, 0x9a, 0xda, 0x37, 0x54, 0x48, 0xc6, 0x47, 0x1f, 0x4f, 0xce, 0xa5,
	0xd6, 0xe8, 0x78, 0x3c, 0xd9, 0x44, 0xdb, 0xd9, 0xb9, 0x0f, 0x87, 0x67, 0xbd, 0x3e, 0xcf, 0x95,
	0x25, 0x3c, 0x07, 0x34, 0x95, 0xd0, 0xdc, 0xfc, 0xc2, 0xe9, 0x32, 0x19, 0x5e, 0x71, 0xda, 0x5d,
	0x40, 0x41, 0x2d, 0xfb, 0x10, 0x68, 0xbd, 0xd1, 0x95, 0x2c, 0xa0, 0x9f, 0x13, 0x40, 0xa9, 0x08,
	0x39, 0x0e, 0x49, 0x9b, 0x48, 0x4c, 0xa3, 0x05, 0xd2, 0x74, 0x60, 0xd3, 0x60, 0xa7, 0x8c, 0xdf,
	0x61, 0x1e, 0xb6, 0xc2, 0x5f, 0x17, 0xad, 0xe8, 0x67, 0x78, 0x36, 0x33, 0x00, 0x51, 0x63, 0x3c,
	0x18, 0xe6, 0xcf, 0xff, 0xfa, 0x97, 0xff, 0xb1, 0x23, 0xcb, 0xed, 0x8f, 0x87, 0x96, 0xee, 0xd7,
	0x0a, 0x62, 0x27, 0x0b, 0x9e, 0x1d, 0x84, 0xf5, 0xdd, 0xf9, 0xce, 0x2c, 0xe9, 0x7b, 0xd8, 0xb2,
	0x9e, 0x2b, 0x26, 0xa8, 0x3a, 0xf2, 0xb3, 0xe4, 0xfd, 0x01, 0x6a, 0xef, 0x62, 0x35, 0x7c, 0xc7,
	0xf5, 0xa2, 0xcd, 0xa6, 0x79, 0x4e, 0x37, 0xed, 0x73, 0xba, 0x79, 0xa2, 0x9e, 0xd3, 0xff, 0x9b,
	0xeb, 0x1c, 0xd6, 0x27, 0x72, 0xd9, 0x32, 0x17, 0x4d, 0xf7, 0xa1, 0xa0, 0xf7, 0x7f, 0xf7, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x50, 0xf4, 0x03, 0xfa, 0x0b, 0x00, 0x00,
}
