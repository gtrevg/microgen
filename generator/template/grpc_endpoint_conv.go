package template

import (
	"fmt"
	"path/filepath"

	"github.com/devimteam/microgen/generator/write_strategy"
	"github.com/devimteam/microgen/util"
	"github.com/vetcher/godecl/types"
	. "github.com/vetcher/jennifer/jen"
)

var (
	defaultProtoTypes = []string{"string", "bool", "byte", "int64", "uint64", "float64", "int32", "uint32", "float32"}
	goToProtoTypesMap = map[string]string{
		"uint": "uint64",
		"int":  "int64",
	}
	defaultGolangTypes = []string{"string", "bool", "int", "uint", "byte", "int64", "uint64", "float64", "int32", "uint32", "float32"}
)

type gRPCEndpointConverterTemplate struct {
	Info             *GenerationInfo
	requestEncoders  []*types.Function
	requestDecoders  []*types.Function
	responseEncoders []*types.Function
	responseDecoders []*types.Function
	state            WriteStrategyState
}

func NewGRPCEndpointConverterTemplate(info *GenerationInfo) Template {
	return &gRPCEndpointConverterTemplate{
		Info: info.Duplicate(),
	}
}

func requestDecodeName(f *types.Function) string {
	return "Decode" + f.Name + "Request"
}

func responseDecodeName(f *types.Function) string {
	return "Decode" + f.Name + "Response"
}

func requestEncodeName(f *types.Function) string {
	return "Encode" + f.Name + "Request"
}

func responseEncodeName(f *types.Function) string {
	return "Encode" + f.Name + "Response"
}

// Renders converter file.
//
//		// This file was automatically generated by "microgen" utility.
//		// Please, do not edit.
//		package protobuf
//
//		import (
//			context "context"
//			svc "github.com/devimteam/microgen/example/svc"
//			stringsvc "gitlab.devim.team/protobuf/stringsvc"
//		)
//
//		func EncodeCountRequest(_ context.Context, request interface{}) (interface{}, error) {
//			req := request.(*svc.CountRequest)
//			return &stringsvc.CountRequest{
//				Symbol: req.Symbol,
//				Text:   req.Text,
//			}, nil
//		}
//
//		func EncodeCountResponse(_ context.Context, response interface{}) (interface{}, error) {
//			resp := response.(*svc.CountResponse)
//			respPositions, err := IntListToProto(resp.Positions)
//			if err != nil {
//				return nil, err
//			}
//			return &stringsvc.CountResponse{
//				Count:     int64(resp.Count),
//				Positions: respPositions,
//			}, nil
//		}
//
//		func DecodeCountRequest(_ context.Context, request interface{}) (interface{}, error) {
//			req := request.(*stringsvc.CountRequest)
//			return &svc.CountRequest{
//				Symbol: string(req.Symbol),
//				Text:   string(req.Text),
//			}, nil
//		}
//
//		func DecodeCountResponse(_ context.Context, response interface{}) (interface{}, error) {
//			resp := response.(*stringsvc.CountResponse)
//			respPositions, err := ProtoToIntList(resp.Positions)
//			if err != nil {
//				return nil, err
//			}
//			return &svc.CountResponse{
//				Count:     int(resp.Count),
//				Positions: respPositions,
//			}, nil
//		}
//
func (t *gRPCEndpointConverterTemplate) Render() write_strategy.Renderer {
	f := &Statement{}

	for _, signature := range t.requestEncoders {
		f.Line().Add(t.encodeRequest(signature))
	}
	for _, signature := range t.responseEncoders {
		f.Line().Add(t.encodeResponse(signature))
	}
	for _, signature := range t.requestDecoders {
		f.Line().Add(t.decodeRequest(signature))
	}
	for _, signature := range t.responseDecoders {
		f.Line().Add(t.decodeResponse(signature))
	}

	if t.state == AppendStrat {
		return f
	}

	file := NewFile("protobuf")
	file.PackageComment(FileHeader)
	file.PackageComment(`Please, do not change functions names!`)
	file.Add(f)

	return file
}

// Returns FieldTypeToProto.
func typeToProto(field *types.Type, depth int) string {
	methodName := util.ToUpperFirst(field.Name)
	if field.IsPointer {
		methodName += "Ptr"
	}
	if field.IsArray {
		methodName += "List"
	}
	if field.IsMap {
		methodName += "Map"
		m := field.Map()
		methodName += typeToProto(&m.Key, 1) + typeToProto(&m.Value, 1)
	}
	if depth == 0 {
		methodName += "ToProto"
	}
	return methodName
}

// Returns ProtoToFieldType.
func protoToType(field *types.Type, depth int) string {
	methodName := ""
	if depth == 0 {
		methodName += "ProtoTo"
	}
	methodName += util.ToUpperFirst(field.Name)
	if field.IsPointer {
		methodName += "Ptr"
	}
	if field.IsArray {
		methodName += "List"
	}
	if field.IsMap {
		methodName += "Map"
		m := field.Map()
		methodName += protoToType(&m.Key, 1) + protoToType(&m.Value, 1)
	}
	return methodName
}

func (gRPCEndpointConverterTemplate) DefaultPath() string {
	return "./transport/converter/protobuf/endpoint_converters.go"
}

func (t *gRPCEndpointConverterTemplate) Prepare() error {
	if t.Info.ProtobufPackage == "" {
		return fmt.Errorf("protobuf package is empty")
	}
	for _, fn := range t.Info.Iface.Methods {
		t.requestDecoders = append(t.requestDecoders, fn)
		t.requestEncoders = append(t.requestEncoders, fn)
		t.responseDecoders = append(t.responseDecoders, fn)
		t.responseEncoders = append(t.responseEncoders, fn)
	}
	return nil
}

func (t *gRPCEndpointConverterTemplate) ChooseStrategy() (write_strategy.Strategy, error) {
	if err := util.TryToOpenFile(t.Info.AbsOutPath, t.DefaultPath()); t.Info.Force || err != nil {
		t.state = FileStrat
		return write_strategy.NewCreateFileStrategy(t.Info.AbsOutPath, t.DefaultPath()), nil
	}
	file, err := util.ParseFile(filepath.Join(t.Info.AbsOutPath, t.DefaultPath()))
	if err != nil {
		return nil, err
	}

	RemoveAlreadyExistingFunctions(file.Functions, &t.requestEncoders, requestEncodeName)
	RemoveAlreadyExistingFunctions(file.Functions, &t.requestDecoders, requestDecodeName)
	RemoveAlreadyExistingFunctions(file.Functions, &t.responseEncoders, responseEncodeName)
	RemoveAlreadyExistingFunctions(file.Functions, &t.responseDecoders, responseDecodeName)

	t.state = AppendStrat
	return write_strategy.NewAppendToFileStrategy(t.Info.AbsOutPath, t.DefaultPath()), nil
}

func RemoveAlreadyExistingFunctions(existing []types.Function, generating *[]*types.Function, nameFormer func(*types.Function) string) {
	x := (*generating)[:0]
	for _, fn := range *generating {
		if f := util.FindFunctionByName(existing, nameFormer(fn)); f == nil {
			x = append(x, fn)
		}
	}
	*generating = x
}

// Renders type conversion (if need) to default protobuf types.
//		req.Symbol
// or
//		int(resp.Count)
// or
//		structNamePositions
// based on field type
// Second result means can field converts to default protobuf type.
func golangTypeToProto(structName string, field *types.Variable) (*Statement, bool) {
	if field.Type.IsArray || field.Type.IsPointer {
		return Id(structName + util.ToUpperFirst(field.Name)), false
	} else if isDefaultProtoField(field) {
		return Id(structName).Dot(util.ToUpperFirst(field.Name)), true
	} else if newType, ok := goToProtoTypesMap[field.Type.Name]; ok {
		newField := &types.Type{
			Name:      newType,
			IsArray:   field.Type.IsArray,
			Import:    field.Type.Import,
			IsPointer: field.Type.IsPointer,
		}
		return fieldType(newField).Call(Id(structName).Dot(util.ToUpperFirst(field.Name))), true
	}
	return Id(structName + util.ToUpperFirst(field.Name)), false
}

// Renders type conversion to default golang types.
// 		int(resp.Count)
// or
// 		structNamePositions
// based on field type
// Second result means can field converts to golang type.
func protoTypeToGolang(structName string, field *types.Variable) (*Statement, bool) {
	if field.Type.IsArray || field.Type.IsPointer {
		return Id(structName + util.ToUpperFirst(field.Name)), false
	} else if isDefaultGolangField(field) {
		return fieldType(&field.Type).Call(Id(structName).Dot(util.ToUpperFirst(field.Name))), true
	}
	return Id(structName + util.ToUpperFirst(field.Name)), false
}

func isDefaultProtoField(field *types.Variable) bool {
	return util.IsInStringSlice(field.Type.Name, defaultProtoTypes)
}

func isDefaultGolangField(field *types.Variable) bool {
	return util.IsInStringSlice(field.Type.Name, defaultGolangTypes)
}

// Render custom type converting and error checking
//
//		structNamePositions, err := ProtoToIntList(structName.Positions)
//		if err != nil {
//			return nil, err
//		}
//
func (t *gRPCEndpointConverterTemplate) convertCustomType(structName, converterName string, field *types.Variable) *Statement {
	return List(Id(structName+util.ToUpperFirst(field.Name)), Err()).
		Op(":=").
		Add(
			Id(converterName).
				Call(Id(structName).
					Dot(util.ToUpperFirst(field.Name))),
		).
		Line().If(Err().Op("!=").Nil()).Block(
		Return().List(Nil(), Err()),
	)
}

// Renders function for encoding request, golang type converts to proto type.
//
//		func EncodeCountRequest(_ context.Context, request interface{}) (interface{}, error) {
//			req := request.(*svc.CountRequest)
//			return &stringsvc.CountRequest{
//				Symbol: req.Symbol,
//				Text:   req.Text,
//			}, nil
//		}
//
func (t *gRPCEndpointConverterTemplate) encodeRequest(signature *types.Function) *Statement {
	methodParams := removeContextIfFirst(signature.Args)
	return Line().Func().Id(requestEncodeName(signature)).Params(Op("_").Qual(PackagePathContext, "Context"), Id("request").Interface()).Params(Interface(), Error()).BlockFunc(
		func(group *Group) {
			if len(methodParams) > 0 {
				group.Id("req").Op(":=").Id("request").Assert(Op("*").Qual(t.Info.ServiceImportPath, requestStructName(signature)))
				for _, field := range methodParams {
					if _, ok := golangTypeToProto("", &field); !ok {
						group.Add(t.convertCustomType("req", typeToProto(&field.Type, 0), &field))
					}
				}
			}
			group.Return().List(Op("&").Qual(t.Info.ProtobufPackage, requestStructName(signature)).Values(DictFunc(func(dict Dict) {
				for _, field := range methodParams {
					req, _ := golangTypeToProto("req", &field)
					dict[structFieldName(&field)] = Line().Add(req)
				}
			})), Nil())
		},
	).Line()
}

// Renders function for encoding response, golang type converts to proto type.
//
//		func EncodeCountResponse(_ context.Context, response interface{}) (interface{}, error) {
//			resp := response.(*svc.CountResponse)
//			respPositions, err := IntListToProto(resp.Positions)
//			if err != nil {
//				return nil, err
//			}
//			return &stringsvc.CountResponse{
//				Count:     int64(resp.Count),
//				Positions: respPositions,
//			}, nil
//		}
//
func (t *gRPCEndpointConverterTemplate) encodeResponse(signature *types.Function) *Statement {
	methodResults := removeContextIfFirst(signature.Results)
	return Line().Func().Id(responseEncodeName(signature)).Call(Op("_").Qual(PackagePathContext, "Context"), Id("response").Interface()).Params(Interface(), Error()).BlockFunc(
		func(group *Group) {
			if len(methodResults) > 0 {
				group.Id("resp").Op(":=").Id("response").Assert(Op("*").Qual(t.Info.ServiceImportPath, responseStructName(signature)))
				for _, field := range methodResults {
					if _, ok := golangTypeToProto("", &field); !ok {
						group.Add(t.convertCustomType("resp", typeToProto(&field.Type, 0), &field))
					}
				}
			}
			group.Return().List(Op("&").Qual(t.Info.ProtobufPackage, responseStructName(signature)).Values(DictFunc(func(dict Dict) {
				for _, field := range methodResults {
					resp, _ := golangTypeToProto("resp", &field)
					dict[structFieldName(&field)] = Line().Add(resp)
				}
			})), Nil())
		},
	).Line()
}

// Renders function for decoding request, proto type converts to golang type.
//
//		func DecodeCountRequest(_ context.Context, request interface{}) (interface{}, error) {
//			req := request.(*stringsvc.CountRequest)
//			return &svc.CountRequest{
//				Symbol: string(req.Symbol),
//				Text:   string(req.Text),
//			}, nil
//		}
//
func (t *gRPCEndpointConverterTemplate) decodeRequest(signature *types.Function) *Statement {
	methodParams := removeContextIfFirst(signature.Args)
	return Line().Func().Id(requestDecodeName(signature)).Call(Op("_").Qual(PackagePathContext, "Context"), Id("request").Interface()).Params(Interface(), Error()).BlockFunc(
		func(group *Group) {
			if len(methodParams) > 0 {
				group.Id("req").Op(":=").Id("request").Assert(Op("*").Qual(t.Info.ProtobufPackage, requestStructName(signature)))
				for _, field := range methodParams {
					if _, ok := protoTypeToGolang("", &field); !ok {
						group.Add(t.convertCustomType("req", protoToType(&field.Type, 0), &field))
					}
				}
			}
			group.Return().List(Op("&").Qual(t.Info.ServiceImportPath, requestStructName(signature)).Values(DictFunc(func(dict Dict) {
				for _, field := range methodParams {
					req, _ := protoTypeToGolang("req", &field)
					dict[structFieldName(&field)] = Line().Add(req)
				}
			})), Nil())
		},
	).Line()
}

// Renders function for decoding response, proto type converts to golang type.
//
//		func DecodeCountResponse(_ context.Context, response interface{}) (interface{}, error) {
//			resp := response.(*stringsvc.CountResponse)
//			respPositions, err := ProtoToIntList(resp.Positions)
//			if err != nil {
//				return nil, err
//			}
//			return &svc.CountResponse{
//				Count:     int(resp.Count),
//				Positions: respPositions,
//			}, nil
//		}
//
func (t *gRPCEndpointConverterTemplate) decodeResponse(signature *types.Function) *Statement {
	methodResults := removeContextIfFirst(signature.Results)
	return Line().Func().Id(responseDecodeName(signature)).Call(Op("_").Qual(PackagePathContext, "Context"), Id("response").Interface()).Params(Interface(), Error()).BlockFunc(
		func(group *Group) {
			if len(methodResults) > 0 {
				group.Id("resp").Op(":=").Id("response").Assert(Op("*").Qual(t.Info.ProtobufPackage, responseStructName(signature)))
				for _, field := range methodResults {
					if _, ok := protoTypeToGolang("", &field); !ok {
						group.Add(t.convertCustomType("resp", protoToType(&field.Type, 0), &field))
					}
				}
			}
			group.Return().List(Op("&").Qual(t.Info.ServiceImportPath, responseStructName(signature)).Values(DictFunc(func(dict Dict) {
				for _, field := range methodResults {
					resp, _ := protoTypeToGolang("resp", &field)
					dict[structFieldName(&field)] = Line().Add(resp)
				}
			})), Nil())
		},
	).Line()
}
