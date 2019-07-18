// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dtssdk/model/message.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MsgID int32

const (
	MsgID_ConnectID               MsgID = 0
	MsgID_DisconnectID            MsgID = 1
	MsgID_GetDefenceZoneRequestID MsgID = 2
	MsgID_GetDefenceZoneReplyID   MsgID = 3
	MsgID_SetDeviceRequestID      MsgID = 4
	MsgID_SetDeviceReplyID        MsgID = 5
	MsgID_ZoneTempNotifyID        MsgID = 6
	MsgID_ZoneAlarmNotifyID       MsgID = 7
	MsgID_DeviceEventNotifyID     MsgID = 8
	MsgID_TempSignalNotifyID      MsgID = 9
	MsgID_GetDeviceIDRequestID    MsgID = 10
	MsgID_GetDeviceIDReplyID      MsgID = 11
	MsgID_HeartBeatID             MsgID = 250
)

var MsgID_name = map[int32]string{
	0:   "ConnectID",
	1:   "DisconnectID",
	2:   "GetDefenceZoneRequestID",
	3:   "GetDefenceZoneReplyID",
	4:   "SetDeviceRequestID",
	5:   "SetDeviceReplyID",
	6:   "ZoneTempNotifyID",
	7:   "ZoneAlarmNotifyID",
	8:   "DeviceEventNotifyID",
	9:   "TempSignalNotifyID",
	10:  "GetDeviceIDRequestID",
	11:  "GetDeviceIDReplyID",
	250: "HeartBeatID",
}

var MsgID_value = map[string]int32{
	"ConnectID":               0,
	"DisconnectID":            1,
	"GetDefenceZoneRequestID": 2,
	"GetDefenceZoneReplyID":   3,
	"SetDeviceRequestID":      4,
	"SetDeviceReplyID":        5,
	"ZoneTempNotifyID":        6,
	"ZoneAlarmNotifyID":       7,
	"DeviceEventNotifyID":     8,
	"TempSignalNotifyID":      9,
	"GetDeviceIDRequestID":    10,
	"GetDeviceIDReplyID":      11,
	"HeartBeatID":             250,
}

func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}

func (MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{0}
}

//防区状态
type DefenceAreaState int32

const (
	DefenceAreaState_Normal      DefenceAreaState = 0
	DefenceAreaState_WarnDiffer  DefenceAreaState = 1
	DefenceAreaState_WarnUp      DefenceAreaState = 2
	DefenceAreaState_WarnTemp    DefenceAreaState = 3
	DefenceAreaState_AlarmDiffer DefenceAreaState = 4
	DefenceAreaState_AlarmUp     DefenceAreaState = 5
	DefenceAreaState_AlarmTemp   DefenceAreaState = 6
)

var DefenceAreaState_name = map[int32]string{
	0: "Normal",
	1: "WarnDiffer",
	2: "WarnUp",
	3: "WarnTemp",
	4: "AlarmDiffer",
	5: "AlarmUp",
	6: "AlarmTemp",
}

var DefenceAreaState_value = map[string]int32{
	"Normal":      0,
	"WarnDiffer":  1,
	"WarnUp":      2,
	"WarnTemp":    3,
	"AlarmDiffer": 4,
	"AlarmUp":     5,
	"AlarmTemp":   6,
}

func (x DefenceAreaState) String() string {
	return proto.EnumName(DefenceAreaState_name, int32(x))
}

func (DefenceAreaState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{1}
}

//防区状态
type FiberState int32

const (
	FiberState_SSTATEISOK   FiberState = 0
	FiberState_SSTATUSUNFIN FiberState = 1
	FiberState_SSTATUSFIN   FiberState = 2
	FiberState_SSTATUSBRK   FiberState = 3
	FiberState_SSTATUSTLO   FiberState = 4
	FiberState_SSTATUSLTM   FiberState = 5
)

var FiberState_name = map[int32]string{
	0: "SSTATEISOK",
	1: "SSTATUSUNFIN",
	2: "SSTATUSFIN",
	3: "SSTATUSBRK",
	4: "SSTATUSTLO",
	5: "SSTATUSLTM",
}

var FiberState_value = map[string]int32{
	"SSTATEISOK":   0,
	"SSTATUSUNFIN": 1,
	"SSTATUSFIN":   2,
	"SSTATUSBRK":   3,
	"SSTATUSTLO":   4,
	"SSTATUSLTM":   5,
}

func (x FiberState) String() string {
	return proto.EnumName(FiberState_name, int32(x))
}

func (FiberState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{2}
}

//设备
type DeviceEvent struct {
	ChannelID            int32      `protobuf:"varint,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	DeviceID             string     `protobuf:"bytes,2,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Timestamp            int64      `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Line                 string     `protobuf:"bytes,4,opt,name=Line,proto3" json:"Line,omitempty"`
	EventType            FiberState `protobuf:"varint,5,opt,name=EventType,proto3,enum=model.FiberState" json:"EventType,omitempty"`
	ChannelLength        float32    `protobuf:"fixed32,6,opt,name=ChannelLength,proto3" json:"ChannelLength,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeviceEvent) Reset()         { *m = DeviceEvent{} }
func (m *DeviceEvent) String() string { return proto.CompactTextString(m) }
func (*DeviceEvent) ProtoMessage()    {}
func (*DeviceEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{0}
}

func (m *DeviceEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceEvent.Unmarshal(m, b)
}
func (m *DeviceEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceEvent.Marshal(b, m, deterministic)
}
func (m *DeviceEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceEvent.Merge(m, src)
}
func (m *DeviceEvent) XXX_Size() int {
	return xxx_messageInfo_DeviceEvent.Size(m)
}
func (m *DeviceEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceEvent proto.InternalMessageInfo

func (m *DeviceEvent) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *DeviceEvent) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *DeviceEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DeviceEvent) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *DeviceEvent) GetEventType() FiberState {
	if m != nil {
		return m.EventType
	}
	return FiberState_SSTATEISOK
}

func (m *DeviceEvent) GetChannelLength() float32 {
	if m != nil {
		return m.ChannelLength
	}
	return 0
}

//防区
type DefenceZone struct {
	ID                   int32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ChannelID            int32            `protobuf:"varint,2,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	ZoneName             string           `protobuf:"bytes,3,opt,name=ZoneName,proto3" json:"ZoneName,omitempty"`
	Start                float32          `protobuf:"fixed32,4,opt,name=Start,proto3" json:"Start,omitempty"`
	Finish               float32          `protobuf:"fixed32,5,opt,name=Finish,proto3" json:"Finish,omitempty"`
	Tag                  string           `protobuf:"bytes,6,opt,name=Tag,proto3" json:"Tag,omitempty"`
	AlarmType            DefenceAreaState `protobuf:"varint,7,opt,name=AlarmType,proto3,enum=model.DefenceAreaState" json:"AlarmType,omitempty"`
	AlarmLoc             float32          `protobuf:"fixed32,8,opt,name=AlarmLoc,proto3" json:"AlarmLoc,omitempty"`
	MaxTemperature       float32          `protobuf:"fixed32,9,opt,name=MaxTemperature,proto3" json:"MaxTemperature,omitempty"`
	MinTemperature       float32          `protobuf:"fixed32,10,opt,name=MinTemperature,proto3" json:"MinTemperature,omitempty"`
	AverageTemperature   float32          `protobuf:"fixed32,11,opt,name=AverageTemperature,proto3" json:"AverageTemperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DefenceZone) Reset()         { *m = DefenceZone{} }
func (m *DefenceZone) String() string { return proto.CompactTextString(m) }
func (*DefenceZone) ProtoMessage()    {}
func (*DefenceZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{1}
}

func (m *DefenceZone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefenceZone.Unmarshal(m, b)
}
func (m *DefenceZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefenceZone.Marshal(b, m, deterministic)
}
func (m *DefenceZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefenceZone.Merge(m, src)
}
func (m *DefenceZone) XXX_Size() int {
	return xxx_messageInfo_DefenceZone.Size(m)
}
func (m *DefenceZone) XXX_DiscardUnknown() {
	xxx_messageInfo_DefenceZone.DiscardUnknown(m)
}

var xxx_messageInfo_DefenceZone proto.InternalMessageInfo

func (m *DefenceZone) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *DefenceZone) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *DefenceZone) GetZoneName() string {
	if m != nil {
		return m.ZoneName
	}
	return ""
}

func (m *DefenceZone) GetStart() float32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *DefenceZone) GetFinish() float32 {
	if m != nil {
		return m.Finish
	}
	return 0
}

func (m *DefenceZone) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *DefenceZone) GetAlarmType() DefenceAreaState {
	if m != nil {
		return m.AlarmType
	}
	return DefenceAreaState_Normal
}

func (m *DefenceZone) GetAlarmLoc() float32 {
	if m != nil {
		return m.AlarmLoc
	}
	return 0
}

func (m *DefenceZone) GetMaxTemperature() float32 {
	if m != nil {
		return m.MaxTemperature
	}
	return 0
}

func (m *DefenceZone) GetMinTemperature() float32 {
	if m != nil {
		return m.MinTemperature
	}
	return 0
}

func (m *DefenceZone) GetAverageTemperature() float32 {
	if m != nil {
		return m.AverageTemperature
	}
	return 0
}

//设置设备请求
type SetDeviceRequest struct {
	ZoneTempNotifyEnable    bool     `protobuf:"varint,1,opt,name=ZoneTempNotifyEnable,proto3" json:"ZoneTempNotifyEnable,omitempty"`
	ZoneAlarmNotifyEnable   bool     `protobuf:"varint,2,opt,name=ZoneAlarmNotifyEnable,proto3" json:"ZoneAlarmNotifyEnable,omitempty"`
	FiberStatusNotifyEnable bool     `protobuf:"varint,3,opt,name=FiberStatusNotifyEnable,proto3" json:"FiberStatusNotifyEnable,omitempty"`
	TempSignalNotifyEnable  bool     `protobuf:"varint,4,opt,name=TempSignalNotifyEnable,proto3" json:"TempSignalNotifyEnable,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *SetDeviceRequest) Reset()         { *m = SetDeviceRequest{} }
func (m *SetDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*SetDeviceRequest) ProtoMessage()    {}
func (*SetDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{2}
}

func (m *SetDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeviceRequest.Unmarshal(m, b)
}
func (m *SetDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeviceRequest.Marshal(b, m, deterministic)
}
func (m *SetDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeviceRequest.Merge(m, src)
}
func (m *SetDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_SetDeviceRequest.Size(m)
}
func (m *SetDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeviceRequest proto.InternalMessageInfo

func (m *SetDeviceRequest) GetZoneTempNotifyEnable() bool {
	if m != nil {
		return m.ZoneTempNotifyEnable
	}
	return false
}

func (m *SetDeviceRequest) GetZoneAlarmNotifyEnable() bool {
	if m != nil {
		return m.ZoneAlarmNotifyEnable
	}
	return false
}

func (m *SetDeviceRequest) GetFiberStatusNotifyEnable() bool {
	if m != nil {
		return m.FiberStatusNotifyEnable
	}
	return false
}

func (m *SetDeviceRequest) GetTempSignalNotifyEnable() bool {
	if m != nil {
		return m.TempSignalNotifyEnable
	}
	return false
}

//设置设备回执
type SetDeviceReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDeviceReply) Reset()         { *m = SetDeviceReply{} }
func (m *SetDeviceReply) String() string { return proto.CompactTextString(m) }
func (*SetDeviceReply) ProtoMessage()    {}
func (*SetDeviceReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{3}
}

func (m *SetDeviceReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeviceReply.Unmarshal(m, b)
}
func (m *SetDeviceReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeviceReply.Marshal(b, m, deterministic)
}
func (m *SetDeviceReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeviceReply.Merge(m, src)
}
func (m *SetDeviceReply) XXX_Size() int {
	return xxx_messageInfo_SetDeviceReply.Size(m)
}
func (m *SetDeviceReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeviceReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeviceReply proto.InternalMessageInfo

func (m *SetDeviceReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SetDeviceReply) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

//获取所有防区
type GetDefenceZoneRequest struct {
	Channel              int32    `protobuf:"varint,1,opt,name=Channel,proto3" json:"Channel,omitempty"`
	Search               string   `protobuf:"bytes,2,opt,name=Search,proto3" json:"Search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDefenceZoneRequest) Reset()         { *m = GetDefenceZoneRequest{} }
func (m *GetDefenceZoneRequest) String() string { return proto.CompactTextString(m) }
func (*GetDefenceZoneRequest) ProtoMessage()    {}
func (*GetDefenceZoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{4}
}

func (m *GetDefenceZoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDefenceZoneRequest.Unmarshal(m, b)
}
func (m *GetDefenceZoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDefenceZoneRequest.Marshal(b, m, deterministic)
}
func (m *GetDefenceZoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDefenceZoneRequest.Merge(m, src)
}
func (m *GetDefenceZoneRequest) XXX_Size() int {
	return xxx_messageInfo_GetDefenceZoneRequest.Size(m)
}
func (m *GetDefenceZoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDefenceZoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDefenceZoneRequest proto.InternalMessageInfo

func (m *GetDefenceZoneRequest) GetChannel() int32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *GetDefenceZoneRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

//获取所有防区
type GetDefenceZoneReply struct {
	Success              bool           `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrMsg               string         `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Rows                 []*DefenceZone `protobuf:"bytes,3,rep,name=Rows,proto3" json:"Rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetDefenceZoneReply) Reset()         { *m = GetDefenceZoneReply{} }
func (m *GetDefenceZoneReply) String() string { return proto.CompactTextString(m) }
func (*GetDefenceZoneReply) ProtoMessage()    {}
func (*GetDefenceZoneReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{5}
}

func (m *GetDefenceZoneReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDefenceZoneReply.Unmarshal(m, b)
}
func (m *GetDefenceZoneReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDefenceZoneReply.Marshal(b, m, deterministic)
}
func (m *GetDefenceZoneReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDefenceZoneReply.Merge(m, src)
}
func (m *GetDefenceZoneReply) XXX_Size() int {
	return xxx_messageInfo_GetDefenceZoneReply.Size(m)
}
func (m *GetDefenceZoneReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDefenceZoneReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetDefenceZoneReply proto.InternalMessageInfo

func (m *GetDefenceZoneReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetDefenceZoneReply) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *GetDefenceZoneReply) GetRows() []*DefenceZone {
	if m != nil {
		return m.Rows
	}
	return nil
}

//获取设备id 请求
type GetDeviceIDRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceIDRequest) Reset()         { *m = GetDeviceIDRequest{} }
func (m *GetDeviceIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceIDRequest) ProtoMessage()    {}
func (*GetDeviceIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{6}
}

func (m *GetDeviceIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceIDRequest.Unmarshal(m, b)
}
func (m *GetDeviceIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceIDRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceIDRequest.Merge(m, src)
}
func (m *GetDeviceIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceIDRequest.Size(m)
}
func (m *GetDeviceIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceIDRequest proto.InternalMessageInfo

//获取设备id 回执
type GetDeviceIDReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	DeviceID             string   `protobuf:"bytes,3,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceIDReply) Reset()         { *m = GetDeviceIDReply{} }
func (m *GetDeviceIDReply) String() string { return proto.CompactTextString(m) }
func (*GetDeviceIDReply) ProtoMessage()    {}
func (*GetDeviceIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{7}
}

func (m *GetDeviceIDReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceIDReply.Unmarshal(m, b)
}
func (m *GetDeviceIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceIDReply.Marshal(b, m, deterministic)
}
func (m *GetDeviceIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceIDReply.Merge(m, src)
}
func (m *GetDeviceIDReply) XXX_Size() int {
	return xxx_messageInfo_GetDeviceIDReply.Size(m)
}
func (m *GetDeviceIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceIDReply proto.InternalMessageInfo

func (m *GetDeviceIDReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetDeviceIDReply) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *GetDeviceIDReply) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

//防区温度，广播
type ZoneTempNotify struct {
	DeviceID             string         `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Timestamp            int64          `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Zones                []*DefenceZone `protobuf:"bytes,3,rep,name=Zones,proto3" json:"Zones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ZoneTempNotify) Reset()         { *m = ZoneTempNotify{} }
func (m *ZoneTempNotify) String() string { return proto.CompactTextString(m) }
func (*ZoneTempNotify) ProtoMessage()    {}
func (*ZoneTempNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{8}
}

func (m *ZoneTempNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZoneTempNotify.Unmarshal(m, b)
}
func (m *ZoneTempNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZoneTempNotify.Marshal(b, m, deterministic)
}
func (m *ZoneTempNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZoneTempNotify.Merge(m, src)
}
func (m *ZoneTempNotify) XXX_Size() int {
	return xxx_messageInfo_ZoneTempNotify.Size(m)
}
func (m *ZoneTempNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_ZoneTempNotify.DiscardUnknown(m)
}

var xxx_messageInfo_ZoneTempNotify proto.InternalMessageInfo

func (m *ZoneTempNotify) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *ZoneTempNotify) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ZoneTempNotify) GetZones() []*DefenceZone {
	if m != nil {
		return m.Zones
	}
	return nil
}

//防区警报 广播
type ZoneAlarmNotify struct {
	DeviceID             string         `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Timestamp            int64          `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Zones                []*DefenceZone `protobuf:"bytes,3,rep,name=Zones,proto3" json:"Zones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ZoneAlarmNotify) Reset()         { *m = ZoneAlarmNotify{} }
func (m *ZoneAlarmNotify) String() string { return proto.CompactTextString(m) }
func (*ZoneAlarmNotify) ProtoMessage()    {}
func (*ZoneAlarmNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{9}
}

func (m *ZoneAlarmNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZoneAlarmNotify.Unmarshal(m, b)
}
func (m *ZoneAlarmNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZoneAlarmNotify.Marshal(b, m, deterministic)
}
func (m *ZoneAlarmNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZoneAlarmNotify.Merge(m, src)
}
func (m *ZoneAlarmNotify) XXX_Size() int {
	return xxx_messageInfo_ZoneAlarmNotify.Size(m)
}
func (m *ZoneAlarmNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_ZoneAlarmNotify.DiscardUnknown(m)
}

var xxx_messageInfo_ZoneAlarmNotify proto.InternalMessageInfo

func (m *ZoneAlarmNotify) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *ZoneAlarmNotify) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ZoneAlarmNotify) GetZones() []*DefenceZone {
	if m != nil {
		return m.Zones
	}
	return nil
}

//设备状态， 广播
type DeviceEventNotify struct {
	ChannelID            int32      `protobuf:"varint,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	DeviceID             string     `protobuf:"bytes,2,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Timestamp            int64      `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Line                 string     `protobuf:"bytes,4,opt,name=Line,proto3" json:"Line,omitempty"`
	EventType            FiberState `protobuf:"varint,5,opt,name=EventType,proto3,enum=model.FiberState" json:"EventType,omitempty"`
	ChannelLength        float32    `protobuf:"fixed32,6,opt,name=ChannelLength,proto3" json:"ChannelLength,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeviceEventNotify) Reset()         { *m = DeviceEventNotify{} }
func (m *DeviceEventNotify) String() string { return proto.CompactTextString(m) }
func (*DeviceEventNotify) ProtoMessage()    {}
func (*DeviceEventNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{10}
}

func (m *DeviceEventNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceEventNotify.Unmarshal(m, b)
}
func (m *DeviceEventNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceEventNotify.Marshal(b, m, deterministic)
}
func (m *DeviceEventNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceEventNotify.Merge(m, src)
}
func (m *DeviceEventNotify) XXX_Size() int {
	return xxx_messageInfo_DeviceEventNotify.Size(m)
}
func (m *DeviceEventNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceEventNotify.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceEventNotify proto.InternalMessageInfo

func (m *DeviceEventNotify) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *DeviceEventNotify) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *DeviceEventNotify) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DeviceEventNotify) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *DeviceEventNotify) GetEventType() FiberState {
	if m != nil {
		return m.EventType
	}
	return FiberState_SSTATEISOK
}

func (m *DeviceEventNotify) GetChannelLength() float32 {
	if m != nil {
		return m.ChannelLength
	}
	return 0
}

//设备状态， 广播
type TempSignalNotify struct {
	DeviceID             string    `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Line                 string    `protobuf:"bytes,2,opt,name=Line,proto3" json:"Line,omitempty"`
	ChannelID            int32     `protobuf:"varint,3,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	Signal               []float32 `protobuf:"fixed32,4,rep,packed,name=Signal,proto3" json:"Signal,omitempty"`
	Timestamp            int64     `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TempSignalNotify) Reset()         { *m = TempSignalNotify{} }
func (m *TempSignalNotify) String() string { return proto.CompactTextString(m) }
func (*TempSignalNotify) ProtoMessage()    {}
func (*TempSignalNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{11}
}

func (m *TempSignalNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TempSignalNotify.Unmarshal(m, b)
}
func (m *TempSignalNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TempSignalNotify.Marshal(b, m, deterministic)
}
func (m *TempSignalNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TempSignalNotify.Merge(m, src)
}
func (m *TempSignalNotify) XXX_Size() int {
	return xxx_messageInfo_TempSignalNotify.Size(m)
}
func (m *TempSignalNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_TempSignalNotify.DiscardUnknown(m)
}

var xxx_messageInfo_TempSignalNotify proto.InternalMessageInfo

func (m *TempSignalNotify) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *TempSignalNotify) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *TempSignalNotify) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *TempSignalNotify) GetSignal() []float32 {
	if m != nil {
		return m.Signal
	}
	return nil
}

func (m *TempSignalNotify) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

//心跳
type HeartBeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeat) Reset()         { *m = HeartBeat{} }
func (m *HeartBeat) String() string { return proto.CompactTextString(m) }
func (*HeartBeat) ProtoMessage()    {}
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9e331ea920348e8, []int{12}
}

func (m *HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeat.Unmarshal(m, b)
}
func (m *HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeat.Marshal(b, m, deterministic)
}
func (m *HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeat.Merge(m, src)
}
func (m *HeartBeat) XXX_Size() int {
	return xxx_messageInfo_HeartBeat.Size(m)
}
func (m *HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeat proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("model.MsgID", MsgID_name, MsgID_value)
	proto.RegisterEnum("model.DefenceAreaState", DefenceAreaState_name, DefenceAreaState_value)
	proto.RegisterEnum("model.FiberState", FiberState_name, FiberState_value)
	proto.RegisterType((*DeviceEvent)(nil), "model.DeviceEvent")
	proto.RegisterType((*DefenceZone)(nil), "model.DefenceZone")
	proto.RegisterType((*SetDeviceRequest)(nil), "model.SetDeviceRequest")
	proto.RegisterType((*SetDeviceReply)(nil), "model.SetDeviceReply")
	proto.RegisterType((*GetDefenceZoneRequest)(nil), "model.GetDefenceZoneRequest")
	proto.RegisterType((*GetDefenceZoneReply)(nil), "model.GetDefenceZoneReply")
	proto.RegisterType((*GetDeviceIDRequest)(nil), "model.GetDeviceIDRequest")
	proto.RegisterType((*GetDeviceIDReply)(nil), "model.GetDeviceIDReply")
	proto.RegisterType((*ZoneTempNotify)(nil), "model.ZoneTempNotify")
	proto.RegisterType((*ZoneAlarmNotify)(nil), "model.ZoneAlarmNotify")
	proto.RegisterType((*DeviceEventNotify)(nil), "model.DeviceEventNotify")
	proto.RegisterType((*TempSignalNotify)(nil), "model.TempSignalNotify")
	proto.RegisterType((*HeartBeat)(nil), "model.HeartBeat")
}

func init() { proto.RegisterFile("dtssdk/model/message.proto", fileDescriptor_b9e331ea920348e8) }

var fileDescriptor_b9e331ea920348e8 = []byte{
	// 890 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x49, 0x51, 0x12, 0x47, 0x89, 0xb2, 0x9e, 0xc8, 0x36, 0x9b, 0xf6, 0x20, 0x10, 0x45,
	0x20, 0xf8, 0x60, 0x03, 0xee, 0x0f, 0x7a, 0xb5, 0xb3, 0x76, 0xba, 0x88, 0xac, 0x00, 0xa4, 0x84,
	0x02, 0x3d, 0x75, 0x2d, 0xaf, 0x65, 0xa2, 0x12, 0xa9, 0x92, 0x2b, 0xb7, 0x7a, 0x91, 0x3e, 0x47,
	0x1f, 0xa5, 0xa7, 0x3e, 0x44, 0x0f, 0x7d, 0x80, 0x9e, 0x8a, 0x5d, 0x52, 0x14, 0x49, 0x2b, 0x3e,
	0xf8, 0xd0, 0x43, 0x6e, 0xfc, 0x66, 0xbf, 0x99, 0x9d, 0x6f, 0x76, 0x66, 0x24, 0x78, 0x7d, 0x23,
	0xd3, 0xf4, 0xe6, 0xe7, 0x93, 0x45, 0x7c, 0x23, 0xe6, 0x27, 0x0b, 0x91, 0xa6, 0x7c, 0x26, 0x8e,
	0x97, 0x49, 0x2c, 0x63, 0xb4, 0xb5, 0xd1, 0xfb, 0xd3, 0x80, 0x0e, 0x15, 0xf7, 0xe1, 0x54, 0x5c,
	0xdc, 0x8b, 0x48, 0xe2, 0x17, 0xe0, 0xbc, 0xbd, 0xe3, 0x51, 0x24, 0xe6, 0x8c, 0xba, 0x46, 0xdf,
	0x18, 0xd8, 0xfe, 0xd6, 0x80, 0xaf, 0xa1, 0x9d, 0x91, 0x19, 0x75, 0xcd, 0xbe, 0x31, 0x70, 0xfc,
	0x02, 0x2b, 0xcf, 0x71, 0xb8, 0x10, 0xa9, 0xe4, 0x8b, 0xa5, 0x6b, 0xf5, 0x8d, 0x81, 0xe5, 0x6f,
	0x0d, 0x88, 0xd0, 0x18, 0x86, 0x91, 0x70, 0x1b, 0xda, 0x4b, 0x7f, 0xe3, 0x09, 0x38, 0xfa, 0xd2,
	0xf1, 0x7a, 0x29, 0x5c, 0xbb, 0x6f, 0x0c, 0xba, 0xa7, 0x7b, 0xc7, 0x3a, 0xad, 0xe3, 0xcb, 0xf0,
	0x5a, 0x24, 0x81, 0xe4, 0x52, 0xf8, 0x5b, 0x0e, 0x7e, 0x09, 0x2f, 0xf2, 0x5c, 0x86, 0x22, 0x9a,
	0xc9, 0x3b, 0xb7, 0xd9, 0x37, 0x06, 0xa6, 0x5f, 0x35, 0x7a, 0x7f, 0x9b, 0x4a, 0xd2, 0xad, 0x88,
	0xa6, 0xe2, 0xc7, 0x38, 0x12, 0xd8, 0x05, 0xb3, 0xd0, 0x62, 0x66, 0x89, 0x6e, 0x25, 0x9a, 0x3b,
	0x24, 0x2a, 0xaf, 0x11, 0x5f, 0x08, 0xad, 0xc2, 0xf1, 0x0b, 0x8c, 0x3d, 0xb0, 0x03, 0xc9, 0x13,
	0xa9, 0x55, 0x98, 0x7e, 0x06, 0xf0, 0x00, 0x9a, 0x97, 0x61, 0x14, 0xa6, 0x77, 0x5a, 0x83, 0xe9,
	0xe7, 0x08, 0x09, 0x58, 0x63, 0x3e, 0xd3, 0x39, 0x3a, 0xbe, 0xfa, 0xc4, 0x6f, 0xc0, 0x39, 0x9b,
	0xf3, 0x64, 0xa1, 0x05, 0xb7, 0xb4, 0xe0, 0xc3, 0x5c, 0x70, 0x9e, 0xf0, 0x59, 0x22, 0x78, 0x2e,
	0xbb, 0x60, 0xaa, 0x94, 0x34, 0x18, 0xc6, 0x53, 0xb7, 0xad, 0xaf, 0x28, 0x30, 0xbe, 0x81, 0xee,
	0x15, 0xff, 0x6d, 0x2c, 0x16, 0x4b, 0x91, 0x70, 0xb9, 0x4a, 0x84, 0xeb, 0x68, 0x46, 0xcd, 0xaa,
	0x79, 0x61, 0x54, 0xe6, 0x41, 0xce, 0xab, 0x58, 0xf1, 0x18, 0xf0, 0xec, 0x5e, 0x24, 0x7c, 0x26,
	0xca, 0xdc, 0x8e, 0xe6, 0xee, 0x38, 0xf1, 0xfe, 0x31, 0x80, 0x04, 0x42, 0x66, 0x5d, 0xe0, 0x8b,
	0x5f, 0x56, 0x22, 0x95, 0x78, 0x0a, 0x3d, 0x55, 0x33, 0xc5, 0x1b, 0xc5, 0x32, 0xbc, 0x5d, 0x5f,
	0x44, 0xfc, 0x7a, 0x2e, 0xf4, 0x1b, 0xb4, 0xfd, 0x9d, 0x67, 0xf8, 0x35, 0xec, 0x2b, 0xbb, 0x16,
	0x56, 0x71, 0x32, 0xb5, 0xd3, 0xee, 0x43, 0xfc, 0x0e, 0x0e, 0x8b, 0x56, 0x59, 0xa5, 0x15, 0x3f,
	0x4b, 0xfb, 0x7d, 0xec, 0x18, 0xbf, 0x85, 0x03, 0x95, 0x43, 0x10, 0xce, 0x22, 0x3e, 0xaf, 0x38,
	0x36, 0xb4, 0xe3, 0x47, 0x4e, 0xbd, 0x73, 0xe8, 0x96, 0xf4, 0x2e, 0xe7, 0x6b, 0x74, 0xa1, 0x15,
	0xac, 0xa6, 0x53, 0x91, 0xa6, 0xb9, 0xc0, 0x0d, 0x54, 0x9d, 0x71, 0x91, 0x24, 0x57, 0xe9, 0x2c,
	0x1f, 0x96, 0x1c, 0x79, 0x0c, 0xf6, 0xdf, 0xa9, 0x18, 0x45, 0x8f, 0x6e, 0x0a, 0xe7, 0x42, 0x2b,
	0xef, 0xc4, 0xbc, 0x5f, 0x37, 0x50, 0x85, 0x0a, 0x04, 0x4f, 0xa6, 0x77, 0x9b, 0x50, 0x19, 0xf2,
	0x62, 0x78, 0x55, 0x0f, 0xf5, 0xa4, 0x9c, 0xf0, 0x0d, 0x34, 0xfc, 0xf8, 0xd7, 0xd4, 0xb5, 0xfa,
	0xd6, 0xa0, 0x73, 0x8a, 0xd5, 0xb6, 0xd4, 0x81, 0xf5, 0xb9, 0xd7, 0x03, 0x7c, 0xb7, 0xd1, 0xcf,
	0x68, 0x9e, 0xb8, 0xf7, 0x13, 0x90, 0x8a, 0xf5, 0x69, 0x39, 0x94, 0xd7, 0x8b, 0x55, 0x5d, 0x2f,
	0x9e, 0x84, 0x6e, 0xb5, 0x6f, 0x2a, 0x6c, 0xe3, 0xb1, 0x65, 0x64, 0xd6, 0x97, 0xd1, 0x00, 0x6c,
	0x15, 0xeb, 0x31, 0xb1, 0x19, 0xc1, 0x5b, 0xc1, 0xcb, 0x5a, 0xe3, 0xfd, 0x2f, 0xd7, 0xfe, 0x65,
	0xc0, 0x5e, 0x69, 0x2b, 0xe7, 0x37, 0x7f, 0x02, 0xbb, 0xf9, 0x77, 0x03, 0x48, 0x7d, 0xb0, 0x1e,
	0xad, 0xe8, 0x26, 0x37, 0xb3, 0x94, 0x5b, 0xa5, 0x0e, 0x56, 0xbd, 0x0e, 0x6a, 0x52, 0x74, 0x74,
	0xb7, 0xd1, 0xb7, 0xd4, 0x3a, 0xce, 0x50, 0xb5, 0x06, 0x76, 0xad, 0x06, 0x5e, 0x07, 0x9c, 0xef,
	0x05, 0x4f, 0xe4, 0xb9, 0xe0, 0xf2, 0xe8, 0x0f, 0x13, 0xec, 0xab, 0x74, 0xc6, 0x28, 0xbe, 0x00,
	0xe7, 0x6d, 0x1c, 0x45, 0x62, 0x2a, 0x19, 0x25, 0xcf, 0x90, 0xc0, 0x73, 0x1a, 0xa6, 0xd3, 0xc2,
	0x62, 0xe0, 0xe7, 0x70, 0xb8, 0x73, 0x94, 0x19, 0x25, 0x26, 0x7e, 0xf6, 0x70, 0xce, 0x97, 0xf3,
	0x35, 0xa3, 0xc4, 0xc2, 0x03, 0xc0, 0xfa, 0xda, 0x64, 0x94, 0x34, 0xb0, 0x57, 0x59, 0xa7, 0x19,
	0xdb, 0x56, 0xd6, 0x6a, 0xf3, 0x33, 0x4a, 0x9a, 0xb8, 0x0f, 0x7b, 0xb5, 0xe6, 0x64, 0x94, 0xb4,
	0xf0, 0x10, 0x5e, 0x3d, 0xe8, 0x1d, 0x46, 0x49, 0x5b, 0xdd, 0x59, 0xaf, 0x3d, 0xa3, 0xc4, 0x41,
	0x17, 0x7a, 0x0f, 0x47, 0x9a, 0x51, 0x02, 0xca, 0xa3, 0x3e, 0xd6, 0x8c, 0x92, 0x0e, 0x12, 0xe8,
	0x14, 0xd5, 0x62, 0x94, 0xfc, 0x6b, 0x1c, 0xad, 0x80, 0xd4, 0x7f, 0xc2, 0x10, 0xa0, 0x39, 0x8a,
	0x93, 0x05, 0x9f, 0x93, 0x67, 0xd8, 0x05, 0xf8, 0x81, 0x27, 0x11, 0x0d, 0x6f, 0x6f, 0x45, 0x42,
	0x0c, 0x75, 0xa6, 0xf0, 0x64, 0x49, 0x4c, 0x7c, 0x0e, 0x6d, 0xf5, 0xad, 0x72, 0x23, 0x16, 0xbe,
	0x84, 0x8e, 0x56, 0x94, 0x53, 0x1b, 0xd8, 0x81, 0x96, 0x36, 0x4c, 0x96, 0xc4, 0x56, 0x0f, 0x92,
	0xfd, 0x30, 0x2a, 0x72, 0xf3, 0x28, 0x02, 0xd8, 0xb6, 0xa3, 0xba, 0x24, 0x08, 0xc6, 0x67, 0xe3,
	0x0b, 0x16, 0x7c, 0x78, 0x9f, 0x3d, 0x97, 0xc6, 0x93, 0x60, 0x32, 0xba, 0x64, 0x23, 0x62, 0x14,
	0x8c, 0x49, 0xa0, 0xb0, 0x59, 0xc2, 0xe7, 0xfe, 0x7b, 0x62, 0x95, 0xf0, 0x78, 0xf8, 0x81, 0x34,
	0x4a, 0x78, 0x38, 0xbe, 0x22, 0xf6, 0x75, 0x53, 0xff, 0x79, 0xfa, 0xea, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2e, 0xa8, 0xfb, 0xa1, 0x5a, 0x09, 0x00, 0x00,
}
