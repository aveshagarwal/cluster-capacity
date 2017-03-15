// Code generated by protoc-gen-gogo.
// source: internal/internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal/internal.proto

It has these top-level messages:
	Point
	Aux
	IteratorOptions
	Measurements
	Measurement
	Interval
	IteratorStats
	Series
	SeriesList
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Point struct {
	Name             *string        `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Tags             *string        `protobuf:"bytes,2,req,name=Tags" json:"Tags,omitempty"`
	Time             *int64         `protobuf:"varint,3,req,name=Time" json:"Time,omitempty"`
	Nil              *bool          `protobuf:"varint,4,req,name=Nil" json:"Nil,omitempty"`
	Aux              []*Aux         `protobuf:"bytes,5,rep,name=Aux" json:"Aux,omitempty"`
	Aggregated       *uint32        `protobuf:"varint,6,opt,name=Aggregated" json:"Aggregated,omitempty"`
	FloatValue       *float64       `protobuf:"fixed64,7,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64         `protobuf:"varint,8,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string        `protobuf:"bytes,9,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool          `protobuf:"varint,10,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	Stats            *IteratorStats `protobuf:"bytes,11,opt,name=Stats" json:"Stats,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}

func (m *Point) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Point) GetTags() string {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return ""
}

func (m *Point) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Point) GetNil() bool {
	if m != nil && m.Nil != nil {
		return *m.Nil
	}
	return false
}

func (m *Point) GetAux() []*Aux {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *Point) GetAggregated() uint32 {
	if m != nil && m.Aggregated != nil {
		return *m.Aggregated
	}
	return 0
}

func (m *Point) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Point) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Point) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Point) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *Point) GetStats() *IteratorStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Aux struct {
	DataType         *int32   `protobuf:"varint,1,req,name=DataType" json:"DataType,omitempty"`
	FloatValue       *float64 `protobuf:"fixed64,2,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64   `protobuf:"varint,3,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string  `protobuf:"bytes,4,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool    `protobuf:"varint,5,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Aux) Reset()         { *m = Aux{} }
func (m *Aux) String() string { return proto.CompactTextString(m) }
func (*Aux) ProtoMessage()    {}

func (m *Aux) GetDataType() int32 {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return 0
}

func (m *Aux) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Aux) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Aux) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Aux) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

type IteratorOptions struct {
	Expr             *string        `protobuf:"bytes,1,opt,name=Expr" json:"Expr,omitempty"`
	Aux              []string       `protobuf:"bytes,2,rep,name=Aux" json:"Aux,omitempty"`
	Sources          []*Measurement `protobuf:"bytes,3,rep,name=Sources" json:"Sources,omitempty"`
	Interval         *Interval      `protobuf:"bytes,4,opt,name=Interval" json:"Interval,omitempty"`
	Dimensions       []string       `protobuf:"bytes,5,rep,name=Dimensions" json:"Dimensions,omitempty"`
	Fill             *int32         `protobuf:"varint,6,opt,name=Fill" json:"Fill,omitempty"`
	FillValue        *float64       `protobuf:"fixed64,7,opt,name=FillValue" json:"FillValue,omitempty"`
	Condition        *string        `protobuf:"bytes,8,opt,name=Condition" json:"Condition,omitempty"`
	StartTime        *int64         `protobuf:"varint,9,opt,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64         `protobuf:"varint,10,opt,name=EndTime" json:"EndTime,omitempty"`
	Ascending        *bool          `protobuf:"varint,11,opt,name=Ascending" json:"Ascending,omitempty"`
	Limit            *int64         `protobuf:"varint,12,opt,name=Limit" json:"Limit,omitempty"`
	Offset           *int64         `protobuf:"varint,13,opt,name=Offset" json:"Offset,omitempty"`
	SLimit           *int64         `protobuf:"varint,14,opt,name=SLimit" json:"SLimit,omitempty"`
	SOffset          *int64         `protobuf:"varint,15,opt,name=SOffset" json:"SOffset,omitempty"`
	Dedupe           *bool          `protobuf:"varint,16,opt,name=Dedupe" json:"Dedupe,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *IteratorOptions) Reset()         { *m = IteratorOptions{} }
func (m *IteratorOptions) String() string { return proto.CompactTextString(m) }
func (*IteratorOptions) ProtoMessage()    {}

func (m *IteratorOptions) GetExpr() string {
	if m != nil && m.Expr != nil {
		return *m.Expr
	}
	return ""
}

func (m *IteratorOptions) GetAux() []string {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *IteratorOptions) GetSources() []*Measurement {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *IteratorOptions) GetInterval() *Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *IteratorOptions) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *IteratorOptions) GetFill() int32 {
	if m != nil && m.Fill != nil {
		return *m.Fill
	}
	return 0
}

func (m *IteratorOptions) GetFillValue() float64 {
	if m != nil && m.FillValue != nil {
		return *m.FillValue
	}
	return 0
}

func (m *IteratorOptions) GetCondition() string {
	if m != nil && m.Condition != nil {
		return *m.Condition
	}
	return ""
}

func (m *IteratorOptions) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *IteratorOptions) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *IteratorOptions) GetAscending() bool {
	if m != nil && m.Ascending != nil {
		return *m.Ascending
	}
	return false
}

func (m *IteratorOptions) GetLimit() int64 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *IteratorOptions) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *IteratorOptions) GetSLimit() int64 {
	if m != nil && m.SLimit != nil {
		return *m.SLimit
	}
	return 0
}

func (m *IteratorOptions) GetSOffset() int64 {
	if m != nil && m.SOffset != nil {
		return *m.SOffset
	}
	return 0
}

func (m *IteratorOptions) GetDedupe() bool {
	if m != nil && m.Dedupe != nil {
		return *m.Dedupe
	}
	return false
}

type Measurements struct {
	Items            []*Measurement `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Measurements) Reset()         { *m = Measurements{} }
func (m *Measurements) String() string { return proto.CompactTextString(m) }
func (*Measurements) ProtoMessage()    {}

func (m *Measurements) GetItems() []*Measurement {
	if m != nil {
		return m.Items
	}
	return nil
}

type Measurement struct {
	Database         *string `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	RetentionPolicy  *string `protobuf:"bytes,2,opt,name=RetentionPolicy" json:"RetentionPolicy,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex            *string `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget         *bool   `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Measurement) Reset()         { *m = Measurement{} }
func (m *Measurement) String() string { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()    {}

func (m *Measurement) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *Measurement) GetRetentionPolicy() string {
	if m != nil && m.RetentionPolicy != nil {
		return *m.RetentionPolicy
	}
	return ""
}

func (m *Measurement) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Measurement) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *Measurement) GetIsTarget() bool {
	if m != nil && m.IsTarget != nil {
		return *m.IsTarget
	}
	return false
}

type Interval struct {
	Duration         *int64 `protobuf:"varint,1,opt,name=Duration" json:"Duration,omitempty"`
	Offset           *int64 `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Interval) Reset()         { *m = Interval{} }
func (m *Interval) String() string { return proto.CompactTextString(m) }
func (*Interval) ProtoMessage()    {}

func (m *Interval) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Interval) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

type IteratorStats struct {
	SeriesN          *int64 `protobuf:"varint,1,opt,name=SeriesN" json:"SeriesN,omitempty"`
	PointN           *int64 `protobuf:"varint,2,opt,name=PointN" json:"PointN,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IteratorStats) Reset()         { *m = IteratorStats{} }
func (m *IteratorStats) String() string { return proto.CompactTextString(m) }
func (*IteratorStats) ProtoMessage()    {}

func (m *IteratorStats) GetSeriesN() int64 {
	if m != nil && m.SeriesN != nil {
		return *m.SeriesN
	}
	return 0
}

func (m *IteratorStats) GetPointN() int64 {
	if m != nil && m.PointN != nil {
		return *m.PointN
	}
	return 0
}

type Series struct {
	Name             *string  `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Tags             []byte   `protobuf:"bytes,2,opt,name=Tags" json:"Tags,omitempty"`
	Aux              []uint32 `protobuf:"varint,3,rep,name=Aux" json:"Aux,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Series) Reset()         { *m = Series{} }
func (m *Series) String() string { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()    {}

func (m *Series) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Series) GetTags() []byte {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Series) GetAux() []uint32 {
	if m != nil {
		return m.Aux
	}
	return nil
}

type SeriesList struct {
	Items            []*Series `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SeriesList) Reset()         { *m = SeriesList{} }
func (m *SeriesList) String() string { return proto.CompactTextString(m) }
func (*SeriesList) ProtoMessage()    {}

func (m *SeriesList) GetItems() []*Series {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "internal.Point")
	proto.RegisterType((*Aux)(nil), "internal.Aux")
	proto.RegisterType((*IteratorOptions)(nil), "internal.IteratorOptions")
	proto.RegisterType((*Measurements)(nil), "internal.Measurements")
	proto.RegisterType((*Measurement)(nil), "internal.Measurement")
	proto.RegisterType((*Interval)(nil), "internal.Interval")
	proto.RegisterType((*IteratorStats)(nil), "internal.IteratorStats")
	proto.RegisterType((*Series)(nil), "internal.Series")
	proto.RegisterType((*SeriesList)(nil), "internal.SeriesList")
}
