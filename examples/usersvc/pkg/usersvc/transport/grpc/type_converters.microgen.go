// Code generated by microgen. DO NOT EDIT.

package grpc

import (
	"context"
	pb "github.com/devimteam/microgen/examples/usersvc/pb"
	service "github.com/devimteam/microgen/examples/usersvc/pkg/usersvc"
	transport "github.com/devimteam/microgen/examples/usersvc/pkg/usersvc/transport"
)

func __CreateComment_Request_ToProtobuf(ctx context.Context, value transport.CreateComment_Request) (pb.CreateComment_Request, error) {
	_Comment, err := __usersvcComment_ToProtobuf(value.Comment)
	if err != nil {
		return nil, err
	}
	return pb.CreateComment_Request{Comment: _Comment}, nil
}
func __CreateComment_Request_FromProtobuf(ctx context.Context, value pb.CreateComment_Request) (transport.CreateComment_Request, error) {
	_Comment, err := __usersvcComment_FromProtobuf(value.Comment)
	if err != nil {
		return nil, err
	}
	return transport.CreateComment_Request{Comment: _Comment}, nil
}
func __CreateUser_Request_ToProtobuf(ctx context.Context, value transport.CreateUser_Request) (pb.CreateUser_Request, error) {
	_User, err := __usersvcUser_ToProtobuf(value.User)
	if err != nil {
		return nil, err
	}
	return pb.CreateUser_Request{User: _User}, nil
}
func __CreateUser_Request_FromProtobuf(ctx context.Context, value pb.CreateUser_Request) (transport.CreateUser_Request, error) {
	_User, err := __usersvcUser_FromProtobuf(value.User)
	if err != nil {
		return nil, err
	}
	return transport.CreateUser_Request{User: _User}, nil
}
func __FindUsers_Response_ToProtobuf(ctx context.Context, value transport.FindUsers_Response) (pb.FindUsers_Response, error) {
	_Results, err := _SP_usersvcUser_ToProtobuf(value.Results)
	if err != nil {
		return nil, err
	}
	return pb.FindUsers_Response{Results: _Results}, nil
}
func __FindUsers_Response_FromProtobuf(ctx context.Context, value pb.FindUsers_Response) (transport.FindUsers_Response, error) {
	_Results, err := _SP_usersvcUser_FromProtobuf(value.Results)
	if err != nil {
		return nil, err
	}
	return transport.FindUsers_Response{Results: _Results}, nil
}
func __GetComment_Response_ToProtobuf(ctx context.Context, value transport.GetComment_Response) (pb.GetComment_Response, error) {
	_Comment, err := __usersvcComment_ToProtobuf(value.Comment)
	if err != nil {
		return nil, err
	}
	return pb.GetComment_Response{Comment: _Comment}, nil
}
func __GetComment_Response_FromProtobuf(ctx context.Context, value pb.GetComment_Response) (transport.GetComment_Response, error) {
	_Comment, err := __usersvcComment_FromProtobuf(value.Comment)
	if err != nil {
		return nil, err
	}
	return transport.GetComment_Response{Comment: _Comment}, nil
}
func __GetUserComments_Response_ToProtobuf(ctx context.Context, value transport.GetUserComments_Response) (pb.GetUserComments_Response, error) {
	_List, err := _S_usersvcComment_ToProtobuf(value.List)
	if err != nil {
		return nil, err
	}
	return pb.GetUserComments_Response{List: _List}, nil
}
func __GetUserComments_Response_FromProtobuf(ctx context.Context, value pb.GetUserComments_Response) (transport.GetUserComments_Response, error) {
	_List, err := _S_usersvcComment_FromProtobuf(value.List)
	if err != nil {
		return nil, err
	}
	return transport.GetUserComments_Response{List: _List}, nil
}
func __GetUser_Response_ToProtobuf(ctx context.Context, value transport.GetUser_Response) (pb.GetUser_Response, error) {
	_User, err := __usersvcUser_ToProtobuf(value.User)
	if err != nil {
		return nil, err
	}
	return pb.GetUser_Response{User: _User}, nil
}
func __GetUser_Response_FromProtobuf(ctx context.Context, value pb.GetUser_Response) (transport.GetUser_Response, error) {
	_User, err := __usersvcUser_FromProtobuf(value.User)
	if err != nil {
		return nil, err
	}
	return transport.GetUser_Response{User: _User}, nil
}
func __UpdateUser_Request_ToProtobuf(ctx context.Context, value transport.UpdateUser_Request) (pb.UpdateUser_Request, error) {
	_User, err := __usersvcUser_ToProtobuf(value.User)
	if err != nil {
		return nil, err
	}
	return pb.UpdateUser_Request{User: _User}, nil
}
func __UpdateUser_Request_FromProtobuf(ctx context.Context, value pb.UpdateUser_Request) (transport.UpdateUser_Request, error) {
	_User, err := __usersvcUser_FromProtobuf(value.User)
	if err != nil {
		return nil, err
	}
	return transport.UpdateUser_Request{User: _User}, nil
}
func _SP_usersvcUser_ToProtobuf(ctx context.Context, value []*service.User) ([]*pb.User, error) {
	if value == nil {
		return nil, nil
	}
	var err error
	converted := make([]*pb.User, len(value))
	for i := range value {
		converted[i], err = _P_usersvcUser_ToProtobuf(value[i])
		if err != nil {
			return nil, err
		}
	}
	return converted, nil
}
func _SP_usersvcUser_FromProtobuf(ctx context.Context, value []*pb.User) ([]*service.User, error) {
	if value == nil {
		return nil, nil
	}
	var err error
	converted := make([]*service.User, len(value))
	for i := range value {
		converted[i], err = _P_usersvcUser_FromProtobuf(value[i])
		if err != nil {
			return nil, err
		}
	}
	return converted, nil
}
func _S_usersvcComment_ToProtobuf(ctx context.Context, value []service.Comment) ([]pb.Comment, error) {
	if value == nil {
		return nil, nil
	}
	var err error
	converted := make([]pb.Comment, len(value))
	for i := range value {
		converted[i], err = __usersvcComment_ToProtobuf(value[i])
		if err != nil {
			return nil, err
		}
	}
	return converted, nil
}
func _S_usersvcComment_FromProtobuf(ctx context.Context, value []pb.Comment) ([]service.Comment, error) {
	if value == nil {
		return nil, nil
	}
	var err error
	converted := make([]service.Comment, len(value))
	for i := range value {
		converted[i], err = __usersvcComment_FromProtobuf(value[i])
		if err != nil {
			return nil, err
		}
	}
	return converted, nil
}
func __usersvcComment_ToProtobuf(ctx context.Context, value service.Comment) (pb.Comment, error) {
	_Id := value.Id
	_Text := value.Text
	return pb.Comment{
		Id:   _Id,
		Text: _Text,
	}, nil
}
func __usersvcComment_FromProtobuf(ctx context.Context, value pb.Comment) (service.Comment, error) {
	_Id := value.Id
	_Text := value.Text
	return service.Comment{
		Id:   _Id,
		Text: _Text,
	}, nil
}
func __usersvcUser_ToProtobuf(ctx context.Context, value service.User) (pb.User, error) {
	_Id := value.Id
	_Name := value.Name
	_Gender := value.Gender
	_Comments, err := _S_usersvcComment_ToProtobuf(value.Comments)
	if err != nil {
		return nil, err
	}
	return pb.User{
		Comments: _Comments,
		Gender:   _Gender,
		Id:       _Id,
		Name:     _Name,
	}, nil
}
func __usersvcUser_FromProtobuf(ctx context.Context, value pb.User) (service.User, error) {
	_Id := value.Id
	_Name := value.Name
	_Gender := value.Gender
	_Comments, err := _S_usersvcComment_FromProtobuf(value.Comments)
	if err != nil {
		return nil, err
	}
	return service.User{
		Comments: _Comments,
		Gender:   _Gender,
		Id:       _Id,
		Name:     _Name,
	}, nil
}
func _P_usersvcUser_ToProtobuf(ctx context.Context, value *service.User) (*pb.User, error) {
	if value == nil {
		return nil, nil
	}
	return __usersvcUser_ToProtobuf(*value)
}
func _P_usersvcUser_FromProtobuf(ctx context.Context, value *pb.User) (*service.User, error) {
	if value == nil {
		return nil, nil
	}
	return __usersvcUser_FromProtobuf(*value)
}
