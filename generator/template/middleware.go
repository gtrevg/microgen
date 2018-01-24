package template

import (
	. "github.com/dave/jennifer/jen"
	"github.com/devimteam/microgen/generator/write_strategy"
)

const (
	MiddlewareTypeName = "Middleware"
)

type middlewareTemplate struct {
	Info *GenerationInfo
}

func NewMiddlewareTemplate(info *GenerationInfo) Template {
	return &middlewareTemplate{
		Info: info,
	}
}

// Render middleware decorator
//
//		// This file was automatically generated by "microgen" utility.
//		// Please, do not edit.
//		package middleware
//
//		import svc "github.com/devimteam/microgen/example/svc"
//
//		type Middleware func(svc.StringService) svc.StringService
//
func (t *middlewareTemplate) Render() write_strategy.Renderer {
	f := NewFile("middleware")
	f.PackageComment(t.Info.FileHeader)
	f.PackageComment(`Please, do not edit.`)
	f.Comment("Service middleware").
		Line().Type().Id(MiddlewareTypeName).Func().Call(Qual(t.Info.ServiceImportPath, t.Info.Iface.Name)).Qual(t.Info.ServiceImportPath, t.Info.Iface.Name)
	return f
}

func (middlewareTemplate) DefaultPath() string {
	return "./middleware/middleware.go"
}

func (middlewareTemplate) Prepare() error {
	return nil
}

func (t *middlewareTemplate) ChooseStrategy() (write_strategy.Strategy, error) {
	return write_strategy.NewCreateFileStrategy(t.Info.AbsOutPath, t.DefaultPath()), nil
}
