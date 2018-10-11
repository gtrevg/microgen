package plugins

import (
	"bytes"
	"encoding/json"
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	"github.com/devimteam/microgen/gen"
	"github.com/devimteam/microgen/internal"
	"github.com/devimteam/microgen/pkg/microgen"
	"github.com/devimteam/microgen/pkg/plugins/pkg"
	"github.com/pkg/errors"
	"github.com/vetcher/go-astra/types"
)

const (
	grpcKitPlugin = "go-kit-grpc"
)

type grpcGokitPlugin struct{}

type grpcGokitConfig struct {
	Path         string
	Protobuf     string
	TransportPkg string
	Client       struct {
		DefaultAddr string
		Trace       bool
	}
	Server struct {
		Trace bool
	}
}

func (p *grpcGokitPlugin) Generate(ctx microgen.Context, args json.RawMessage) (microgen.Context, error) {
	cfg := grpcGokitConfig{}
	if len(args) > 0 {
		err := json.Unmarshal(args, &cfg)
		if err != nil {
			return ctx, err
		}
	}
	if cfg.Protobuf == "" {
		return ctx, errors.New("argument 'protobuf' is required")
	}
	if cfg.TransportPkg == "" {
		cfg.TransportPkg = "transport"
	}
	if cfg.Path == "" {
		cfg.Path = "transport/grpc"
	}
	resolvedPkgPath, err := gen.GetPkgPath(cfg.TransportPkg, true)
	if err != nil {
		return ctx, err
	}
	cfg.TransportPkg = resolvedPkgPath

	ctx, err = p.client(ctx, cfg)
	if err != nil {
		return ctx, err
	}
	ctx, err = p.server(ctx, cfg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func (p *grpcGokitPlugin) client(ctx microgen.Context, cfg grpcGokitConfig) (microgen.Context, error) {
	const filename = "client.microgen.go"
	ImportAliasFromSources = true
	pluginPackagePath, err := gen.GetPkgPath(filepath.Join(cfg.Path, filename), false)
	if err != nil {
		return ctx, err
	}
	pkgName, err := gen.PackageName(pluginPackagePath, "")
	if err != nil {
		return ctx, err
	}
	f := NewFilePathName(pluginPackagePath, pkgName)
	f.ImportAlias(ctx.SourcePackageImport, serviceAlias)
	f.ImportAlias(cfg.Protobuf, "pb")
	f.ImportAlias(pkg.GoKitGRPC, "grpckit")
	f.HeaderComment(ctx.FileHeader)

	f.Func().Id("NewGRPCClient").
		ParamsFunc(func(p *Group) {
			p.Id("conn").Op("*").Qual(pkg.GoogleGRPC, "ClientConn")
			p.Id("addr").Id("string")
			p.Id("opts").Op("...").Qual(pkg.GoKitGRPC, "ClientOption")
		}).Qual(cfg.TransportPkg, _Endpoints_).
		BlockFunc(func(g *Group) {
			if cfg.Client.DefaultAddr != "" {
				g.If(Id("addr").Op("==").Lit("")).Block(
					Id("addr").Op("=").Lit(cfg.Client.DefaultAddr),
				)
			}
			g.Return().Qual(cfg.TransportPkg, _Endpoints_).Values(DictFunc(func(d Dict) {
				for _, fn := range ctx.Interface.Methods {
					if !ctx.AllowedMethods[fn.Name] {
						continue
					}
					client := &Statement{}
					client.Qual(pkg.GoKitGRPC, "NewClient").Call(
						Line().Id("conn"), Id("addr"), Lit(fn.Name),
						Line().Id(join_("_Encode", fn.Name, _Request_)),
						Line().Id(join_("_Decode", fn.Name, _Response_)),
						Line().Add(p.protobufReplyType(ctx, cfg, fn)).Values(),
						Line().Id("opts...").Line(),
					).Dot("Endpoint").Call()
					d[Id(join_(fn.Name, "Endpoint"))] = client
				}
			}))
		})
	f.Line().Func().Id("ClientOptionsBuilder").Params(
		Id("opts").Op("[]").Qual(pkg.GoKitGRPC, "ClientOption"),
		Id("fns...").Func().Params(Op("[]").Qual(pkg.GoKitGRPC, "ClientOption")).Params(Op("[]").Qual(pkg.GoKitGRPC, "ClientOption")),
	).Params(Op("[]").Qual(pkg.GoKitGRPC, "ClientOption")).Block(
		For().Id("i := range fns").Block(
			Id("opts = fns[i](opts)"),
		),
		Return(Id("opts")),
	)

	if cfg.Client.Trace {
		f.Line().Func().Id("TracingClientOptions").Params(
			Id("tracer").Qual(pkg.OpenTracing, "Tracer"),
			Id("logger").Qual(pkg.GoKitLog, "Logger"),
		).Params(
			Func().Params(Op("[]").Qual(pkg.GoKitGRPC, "ClientOption")).Params(Op("[]").Qual(pkg.GoKitGRPC, "ClientOption")),
		).Block(
			Return().Func().Params(Id("opts").Op("[]").Qual(pkg.GoKitGRPC, "ClientOption")).Params(Op("[]").Qual(pkg.GoKitGRPC, "ClientOption")).Block(
				Return().Append(Id("opts"), Qual(pkg.GoKitGRPC, "ClientBefore").Call(
					Line().Qual(pkg.GoKitOpenTracing, "ContextToGRPC").Call(Id("tracer"), Id("logger")).Op(",").Line(),
				)),
			),
		)
	}
	/*if cfg.Server.Trace {
		f.Func().Id("TraceServer").Params(
			Id("tracer").Qual(pkg.OpenTracing, "Tracer"),
		).Params(
			Func().Params(Id("endpoints").Id(_Endpoints_)).Params(Id(_Endpoints_)),
		).Block(
			Return().Func().Params(Id("endpoints").Id(_Endpoints_)).Params(Id(_Endpoints_)).
				BlockFunc(func(body *Group) {
					body.Return(Id(_Endpoints_).Values(DictFunc(func(d Dict) {
						for _, signature := range ctx.Interface.Methods {
							if ctx.AllowedMethods[signature.Name] {
								// CreateComment_Endpoint:   latency(dur, "CreateComment")(endpoints.CreateComment_Endpoint),
								d[Id(join_(signature.Name, "Endpoint"))] = Qual(pkg.GoKitOpenTracing, "TraceServer").Call(Id("tracer"),
									Lit(signature.Name)).Call(Id("endpoints").Dot(join_(signature.Name, "Endpoint")))
							}
						}
					})))
				}),
		)
	}*/

	outfile := microgen.File{
		Name: grpcKitPlugin,
		Path: filepath.Join(cfg.Path, filename),
	}
	var b bytes.Buffer
	err = f.Render(&b)
	if err != nil {
		return ctx, err
	}
	outfile.Content = b.Bytes()
	ctx.Files = append(ctx.Files, outfile)
	return ctx, nil
}

func (p *grpcGokitPlugin) server(ctx microgen.Context, cfg grpcGokitConfig) (microgen.Context, error) {
	const filename = "server.microgen.go"
	ImportAliasFromSources = true
	pluginPackagePath, err := gen.GetPkgPath(filepath.Join(cfg.Path, filename), false)
	if err != nil {
		return ctx, err
	}
	pkgName, err := gen.PackageName(pluginPackagePath, "")
	if err != nil {
		return ctx, err
	}
	f := NewFilePathName(pluginPackagePath, pkgName)
	f.ImportAlias(ctx.SourcePackageImport, serviceAlias)
	f.HeaderComment(ctx.FileHeader)

	if cfg.Server.Trace {
		f.Func().Id("TraceServer").Params(
			Id("tracer").Qual(pkg.OpenTracing, "Tracer"),
		).Params(
			Func().Params(Id("endpoints").Id(_Endpoints_)).Params(Id(_Endpoints_)),
		).Block(
			Return().Func().Params(Id("endpoints").Id(_Endpoints_)).Params(Id(_Endpoints_)).
				BlockFunc(func(body *Group) {
					body.Return(Id(_Endpoints_).Values(DictFunc(func(d Dict) {
						for _, signature := range ctx.Interface.Methods {
							if ctx.AllowedMethods[signature.Name] {
								// CreateComment_Endpoint:   latency(dur, "CreateComment")(endpoints.CreateComment_Endpoint),
								d[Id(join_(signature.Name, "Endpoint"))] = Qual(pkg.GoKitOpenTracing, "TraceServer").Call(Id("tracer"),
									Lit(signature.Name)).Call(Id("endpoints").Dot(join_(signature.Name, "Endpoint")))
							}
						}
					})))
				}),
		)
	}

	outfile := microgen.File{
		Name: grpcKitPlugin,
		Path: filepath.Join(cfg.Path, filename),
	}
	var b bytes.Buffer
	err = f.Render(&b)
	if err != nil {
		return ctx, err
	}
	outfile.Content = b.Bytes()
	ctx.Files = append(ctx.Files, outfile)
	return ctx, nil
}

func (p *grpcGokitPlugin) protobufReplyType(ctx microgen.Context, cfg grpcGokitConfig, fn *types.Function) Code {
	results := internal.RemoveErrorIfLast(fn.Results)
	if len(results) == 0 {
		return Qual(pkg.EmptyProtobuf, "Empty")
	}
	if len(results) == 1 {
		str, ok := findCustomBinding(results[0].Type)
		if ok {
			return Id(str)
		}
	}
	return Qual(cfg.Protobuf, fn.Name+_Response_)
}

type ProtobufTypeBinder func(types.Type) (string, bool)

var protobufBindings = make([]ProtobufTypeBinder, 0)

func RegisterProtobufTypeBinding(fn ProtobufTypeBinder) {
	protobufBindings = append(protobufBindings, fn)
}

func findCustomBinding(t types.Type) (string, bool) {
	n := len(protobufBindings)
	for i := 0; i < n; i++ {
		if s, ok := protobufBindings[n-i-1](t); ok {
			return s, true
		}
	}
	return "", false
}