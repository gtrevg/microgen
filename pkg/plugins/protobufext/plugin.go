package protobufext

import (
	"reflect"
	"time"

	"github.com/cv21/microgen/pkg/plugins"
)

func init() {
	plugins.RegisterProtobufTypeBinding(protoext{})
}

const (
	googleProtobuf             = "google.protobuf."
	googleProtobufStringValue  = googleProtobuf + "StringValue"
	googleProtobufBoolValue    = googleProtobuf + "BoolValue"
	googleProtobufInt64Value   = googleProtobuf + "Int64Value"
	googleProtobufUInt64Value  = googleProtobuf + "UInt64Value"
	googleProtobufInt32Value   = googleProtobuf + "Int32Value"
	googleProtobufUInt32Value  = googleProtobuf + "UInt32Value"
	googleProtobufFloat64Value = googleProtobuf + "DoubleValue"
	googleProtobufFloat32Value = googleProtobuf + "FloatValue"
	googleProtobufTimestamp    = googleProtobuf + "Timestamp"

	importGoogleProtobuf          = "google/protobuf/"
	importGoogleProtobufWrappers  = importGoogleProtobuf + "wrappers.proto"
	importGoogleProtobufTimestamp = importGoogleProtobuf + "timestamp.proto"

	currentPackageImport = "github.com/cv21/microgen/pkg/plugins/protobufext"
)

type protoext struct{}

func (protoext) ProtobufType(origType reflect.Type) (pbType reflect.Type, ok bool) {
	return origType, false
}

func (protoext) MarshalLayout(origType reflect.Type) (marshalLayout string, requiredImport *string, ok bool) {
	switch origType {
	case stringPType:
		return "P_String_ToProtobuf(%s)", sp(currentPackageImport), true
	case uint64PType:
		return "P_UInt64_ToProtobuf(%s)", sp(currentPackageImport), true
	case uint32PType:
		return "P_UInt32_ToProtobuf(%s)", sp(currentPackageImport), true
	case boolPType:
		return "P_Bool_ToProtobuf(%s)", sp(currentPackageImport), true
	case intPType, int32PType, int64PType,
		uintPType,
		timeType, timePType,
		float32PType, float64PType:
		return "%s", nil, true
	default:
		return "", nil, false
	}
}

func (protoext) UnmarshalLayout(origType reflect.Type) (unmarshalLayout string, requiredImport *string, ok bool) {
	switch origType {
	case stringPType:
		return "P_String_FromProtobuf(%s)", sp(currentPackageImport), true
	case uint64PType:
		return "P_UInt64_FromProtobuf(%s)", sp(currentPackageImport), true
	case uint32PType:
		return "P_UInt32_FromProtobuf(%s)", sp(currentPackageImport), true
	case boolPType:
		return "P_Bool_FromProtobuf(%s)", sp(currentPackageImport), true
	case intPType, int32PType, int64PType,
		uintPType,
		timeType, timePType,
		float32PType, float64PType:
		return "%s", nil, true
	default:
		return "", nil, false
	}
}

func (protoext) ProtoBinding(origType reflect.Type) (fieldType string, requiredImport *string, ok bool) {
	switch origType {
	case stringPType:
		return googleProtobufStringValue, sp(importGoogleProtobufWrappers), true
	case boolPType:
		return googleProtobufBoolValue, sp(importGoogleProtobufWrappers), true
	case intPType:
		return googleProtobufInt64Value, sp(importGoogleProtobufWrappers), true
	case int32PType:
		return googleProtobufInt32Value, sp(importGoogleProtobufWrappers), true
	case int64PType:
		return googleProtobufInt64Value, sp(importGoogleProtobufWrappers), true
	case uintPType:
		return googleProtobufUInt64Value, sp(importGoogleProtobufWrappers), true
	case uint32PType:
		return googleProtobufUInt32Value, sp(importGoogleProtobufWrappers), true
	case uint64PType:
		return googleProtobufUInt64Value, sp(importGoogleProtobufWrappers), true
	case float64PType:
		return googleProtobufFloat64Value, sp(importGoogleProtobufWrappers), true
	case float32PType:
		return googleProtobufFloat32Value, sp(importGoogleProtobufWrappers), true
	case timeType:
		return googleProtobufTimestamp, sp(importGoogleProtobufTimestamp), true
	case timePType:
		return googleProtobufTimestamp, sp(importGoogleProtobufTimestamp), true
	default:
		return "", nil, false
	}
}

func sp(s string) *string {
	return &s
}

var (
	stringType = reflect.TypeOf(new(string)).Elem()

	stringPType  = reflect.TypeOf(new(*string)).Elem()
	boolPType    = reflect.TypeOf(new(*bool)).Elem()
	intPType     = reflect.TypeOf(new(*int)).Elem()
	int32PType   = reflect.TypeOf(new(*int32)).Elem()
	int64PType   = reflect.TypeOf(new(*int64)).Elem()
	uintPType    = reflect.TypeOf(new(*uint)).Elem()
	uint32PType  = reflect.TypeOf(new(*uint32)).Elem()
	uint64PType  = reflect.TypeOf(new(*uint64)).Elem()
	timeType     = reflect.TypeOf(new(time.Time)).Elem()
	timePType    = reflect.TypeOf(new(*time.Time)).Elem()
	float32PType = reflect.TypeOf(new(*float32)).Elem()
	float64PType = reflect.TypeOf(new(*float64)).Elem()
)
