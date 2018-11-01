// Code generated by microgen. DO NOT EDIT.

package service

import (
	"context"
	service "github.com/devimteam/microgen/examples/usersvc/pkg/usersvc"
	log "github.com/go-kit/kit/log"
	"time"
)

//go:generate easyjson -all logging.microgen.go

var _ service.UserService = &loggingMiddleware{}

func LoggingMiddleware(logger log.Logger) func(service.UserService) service.UserService {
	return func(next service.UserService) service.UserService {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   service.UserService
}

func (M loggingMiddleware) CreateUser(arg_0 context.Context, arg_1 service.User) (res_0 string, res_1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateUser",
			"request", log_LoggingMiddleware_CreateUser_Request{User: arg_1},
			"response", log_LoggingMiddleware_CreateUser_Response{Id: res_0},
			"err", res_1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.CreateUser(arg_0, arg_1)
}

func (M loggingMiddleware) UpdateUser(arg_0 context.Context, arg_1 service.User) (res_0 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UpdateUser",
			"request", log_LoggingMiddleware_UpdateUser_Request{User: arg_1},
			"err", res_0,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.UpdateUser(arg_0, arg_1)
}

func (M loggingMiddleware) GetUser(arg_0 context.Context, arg_1 string) (res_0 service.User, res_1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetUser",
			"request", log_LoggingMiddleware_GetUser_Request{Id: arg_1},
			"response", log_LoggingMiddleware_GetUser_Response{User: res_0},
			"err", res_1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetUser(arg_0, arg_1)
}

func (M loggingMiddleware) FindUsers(arg_0 context.Context) (res_0 map[string]service.User, res_1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "FindUsers",
			"response", log_LoggingMiddleware_FindUsers_Response{
				LenResults: len(res_0),
				Results:    res_0,
			},
			"err", res_1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.FindUsers(arg_0)
}

func (M loggingMiddleware) CreateComment(arg_0 context.Context, arg_1 service.Comment) (res_0 string, res_1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateComment",
			"request", log_LoggingMiddleware_CreateComment_Request{Comment: arg_1},
			"response", log_LoggingMiddleware_CreateComment_Response{Id: res_0},
			"err", res_1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.CreateComment(arg_0, arg_1)
}

func (M loggingMiddleware) GetComment(arg_0 context.Context, arg_1 string) (res_0 service.Comment, res_1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetComment",
			"request", log_LoggingMiddleware_GetComment_Request{Id: arg_1},
			"response", log_LoggingMiddleware_GetComment_Response{Comment: res_0},
			"err", res_1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetComment(arg_0, arg_1)
}

func (M loggingMiddleware) GetUserComments(arg_0 context.Context, arg_1 string) (res_0 []service.Comment, res_1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetUserComments",
			"request", log_LoggingMiddleware_GetUserComments_Request{UserId: arg_1},
			"response", log_LoggingMiddleware_GetUserComments_Response{List: res_0},
			"err", res_1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetUserComments(arg_0, arg_1)
}

//easyjson:json
type (
	log_LoggingMiddleware_CreateUser_Request struct {
		User service.User
	}
	log_LoggingMiddleware_CreateUser_Response struct {
		Id string
	}
	log_LoggingMiddleware_UpdateUser_Request struct {
		User service.User
	}
	log_LoggingMiddleware_GetUser_Request struct {
		Id string
	}
	log_LoggingMiddleware_GetUser_Response struct {
		User service.User
	}
	log_LoggingMiddleware_FindUsers_Response struct {
		Results    map[string]service.User
		LenResults int `json:"len(Results)"`
	}
	log_LoggingMiddleware_CreateComment_Request struct {
		Comment service.Comment
	}
	log_LoggingMiddleware_CreateComment_Response struct {
		Id string
	}
	log_LoggingMiddleware_GetComment_Request struct {
		Id string
	}
	log_LoggingMiddleware_GetComment_Response struct {
		Comment service.Comment
	}
	log_LoggingMiddleware_GetUserComments_Request struct {
		UserId string
	}
	log_LoggingMiddleware_GetUserComments_Response struct {
		List []service.Comment
	}
)
