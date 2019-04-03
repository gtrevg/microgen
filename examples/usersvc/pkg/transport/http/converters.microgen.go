// Code generated by microgen 1.0.0-beta. DO NOT EDIT.

// Please, do not change functions names!
package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	transport "github.com/cv21/microgen/examples/usersvc/pkg/transport"
	"io/ioutil"
	"net/http"
	"path"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func CommonHTTPResponseEncoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func _Decode_CreateUser_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_UpdateUser_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_GetUser_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_FindUsers_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.FindUsersRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_CreateComment_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_GetComment_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_GetUserComments_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetUserCommentsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_CreateUser_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.CreateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_UpdateUser_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.UpdateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_GetUser_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_FindUsers_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.FindUsersResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_CreateComment_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.CreateCommentResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_GetComment_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetCommentResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_GetUserComments_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetUserCommentsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Encode_CreateUser_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "create-user")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_UpdateUser_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "update-user")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetUser_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "get-user")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_FindUsers_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "find-users")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_CreateComment_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "create-comment")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetComment_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "get-comment")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetUserComments_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "get-user-comments")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_CreateUser_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_UpdateUser_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_GetUser_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_FindUsers_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_CreateComment_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_GetComment_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_GetUserComments_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
