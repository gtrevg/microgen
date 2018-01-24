// This file was automatically generated by "microgen 0.7.0b" utility.
// Please, do not change functions names!
package protobuf

import (
	context "context"
	generated "github.com/devimteam/microgen/example/generated"
	protobuf "github.com/devimteam/microgen/example/protobuf"
)

func EncodeUppercaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*generated.UppercaseRequest)
	reqStringsMap, err := MapStringStringToProto(req.StringsMap)
	if err != nil {
		return nil, err
	}
	return &protobuf.UppercaseRequest{StringsMap: reqStringsMap}, nil
}

func EncodeCountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*generated.CountRequest)
	return &protobuf.CountRequest{
		Symbol: req.Symbol,
		Text:   req.Text,
	}, nil
}

func EncodeTestCaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*generated.TestCaseRequest)
	reqComments, err := ListPtrEntityCommentToProto(req.Comments)
	if err != nil {
		return nil, err
	}
	return &protobuf.TestCaseRequest{Comments: reqComments}, nil
}

func EncodeUppercaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*generated.UppercaseResponse)
	return &protobuf.UppercaseResponse{Ans: resp.Ans}, nil
}

func EncodeCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*generated.CountResponse)
	respPositions, err := ListIntToProto(resp.Positions)
	if err != nil {
		return nil, err
	}
	return &protobuf.CountResponse{
		Count:     int64(resp.Count),
		Positions: respPositions,
	}, nil
}

func EncodeTestCaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*generated.TestCaseResponse)
	respTree, err := MapStringIntToProto(resp.Tree)
	if err != nil {
		return nil, err
	}
	return &protobuf.TestCaseResponse{Tree: respTree}, nil
}

func DecodeUppercaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protobuf.UppercaseRequest)
	reqStringsMap, err := ProtoToMapStringString(req.StringsMap)
	if err != nil {
		return nil, err
	}
	return &generated.UppercaseRequest{StringsMap: reqStringsMap}, nil
}

func DecodeCountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protobuf.CountRequest)
	return &generated.CountRequest{
		Symbol: string(req.Symbol),
		Text:   string(req.Text),
	}, nil
}

func DecodeTestCaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protobuf.TestCaseRequest)
	reqComments, err := ProtoToListPtrEntityComment(req.Comments)
	if err != nil {
		return nil, err
	}
	return &generated.TestCaseRequest{Comments: reqComments}, nil
}

func DecodeUppercaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*protobuf.UppercaseResponse)
	return &generated.UppercaseResponse{Ans: string(resp.Ans)}, nil
}

func DecodeCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*protobuf.CountResponse)
	respPositions, err := ProtoToListInt(resp.Positions)
	if err != nil {
		return nil, err
	}
	return &generated.CountResponse{
		Count:     int(resp.Count),
		Positions: respPositions,
	}, nil
}

func DecodeTestCaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*protobuf.TestCaseResponse)
	respTree, err := ProtoToMapStringInt(resp.Tree)
	if err != nil {
		return nil, err
	}
	return &generated.TestCaseResponse{Tree: respTree}, nil
}
