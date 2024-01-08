package transformhelpers

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func StringToPgtypeText(str string) pgtype.Text {
	val := pgtype.Text{}
	val.Scan(str)
	return val
}

func Int32ToPgtypeInt4(num int32) pgtype.Int4 {
	val := pgtype.Int4{
		Int32: num,
		Valid: true,
	}
	return val
}

func ToPgtypeText(m protoreflect.Enum) pgtype.Text {
	return pgtype.Text{}
}

func BoolToPgtypeBool(b bool) pgtype.Bool {
	val := pgtype.Bool{}
	val.Scan(b)
	return val
}

func PgtypeTextToString(text pgtype.Text) string {
	return text.String
}

func PgtypeTextTo(text pgtype.Text) protoreflect.Enum {
	return nil
}

func PgtypeBoolToBool(b pgtype.Bool) bool {
	return b.Bool
}

func TimePtrToPgtypeTimestamptz(t *timestamppb.Timestamp) pgtype.Timestamptz {
	val := pgtype.Timestamptz{}
	val.Scan(t.AsTime())
	return val
}

func PgtypeTimestamptzToTimePtr(t pgtype.Timestamptz) *timestamppb.Timestamp {
	return timestamppb.New(t.Time)
}

func PgtypeInt4ToInt32(num pgtype.Int4) int32 {
	return num.Int32
}

func Int32ToPgtypeInterval(t int32) pgtype.Interval {
	val := pgtype.Interval{}
	val.Scan(t)
	return val
}

func PgtypeIntervalToInt32(t pgtype.Interval) int32 {
	return int32(t.Microseconds)
}

func Int64ToPgtypeInt4(i int64) pgtype.Int4 {
	return pgtype.Int4{
		Int32: int32(i),
		Valid: true,
	}
}

func UInt64ToPgtypeInt8(i uint64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: int64(i),
		Valid: true,
	}
}

func PgtypeInt4ToInt64(val pgtype.Int4) int64 {
	return int64(val.Int32)
}

func PgtypeTextToStringValue(val pgtype.Text) *wrapperspb.StringValue {
	if !val.Valid {
		return nil
	}
	return &wrapperspb.StringValue{
		Value: val.String,
	}
}

func StringValueToPgtypeText(val *wrapperspb.StringValue) pgtype.Text {
	if val == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{String: val.Value, Valid: true}
}

func Int64ValueToPgtypeInt4(val *wrapperspb.Int64Value) pgtype.Int4 {
	if val == nil {
		return pgtype.Int4{}
	}
	return pgtype.Int4{Int32: int32(val.Value), Valid: true}
}

func Int32ValueToPgtypeInt4(val *wrapperspb.Int32Value) pgtype.Int4 {
	if val == nil {
		return pgtype.Int4{}
	}
	return pgtype.Int4{Int32: val.Value, Valid: true}
}

func PgtypeInt4ToInt32Value(val pgtype.Int4) *wrapperspb.Int32Value {
	if !val.Valid {
		return nil
	}
	return &wrapperspb.Int32Value{
		Value: val.Int32,
	}
}

func PgtypeInt4ToInt64Value(val pgtype.Int4) *wrapperspb.Int64Value {
	if !val.Valid {
		return nil
	}
	return &wrapperspb.Int64Value{
		Value: int64(val.Int32),
	}
}

func BoolValueToPgtypeBool(b *wrapperspb.BoolValue) pgtype.Bool {
	if b == nil {
		return pgtype.Bool{}
	}
	return pgtype.Bool{
		Bool:  b.Value,
		Valid: true,
	}
}

func PgtypeBoolToBoolValue(val pgtype.Bool) *wrapperspb.BoolValue {
	if !val.Valid {
		return nil
	}
	return &wrapperspb.BoolValue{
		Value: val.Bool,
	}
}

func PgtypeDateToTimePtr(val pgtype.Date) *timestamppb.Timestamp {
	if !val.Valid {
		return nil
	}

	return timestamppb.New(val.Time)
}

func TimePtrToPgtypeDate(val *timestamppb.Timestamp) pgtype.Date {
	if val == nil {
		return pgtype.Date{}
	}

	return pgtype.Date{
		Time:  val.AsTime(),
		Valid: true,
	}
}

func Float32ToPgtypeFloat8(val float32) pgtype.Float8 {
	return pgtype.Float8{
		Float64: float64(val),
		Valid:   true,
	}
}

func PgtypeFloat8ToFloat32(val pgtype.Float8) float32 {
	return float32(val.Float64)
}

func StringToStringValue(val string) *wrapperspb.StringValue {
	return &wrapperspb.StringValue{
		Value: val,
	}
}

func StringValueToString(val *wrapperspb.StringValue) string {
	return val.GetValue()
}

func TimeToPgtypeTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

func PgtypeInt8ToInt64(val pgtype.Int8) int64 {
	return val.Int64
}

func Int64ToPgtypeInt8(val int64) pgtype.Int8 {
	return pgtype.Int8{
		Valid: true,
		Int64: val,
	}
}
