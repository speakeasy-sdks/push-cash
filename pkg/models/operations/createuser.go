// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/pkg/models/shared"
)

type CreateUserRequest struct {
	CreateUserRequest shared.CreateUserRequest `request:"mediaType=application/json"`
	// The idempotency key for the request
	XIdempotencyKey string `header:"style=simple,explode=false,name=X-Idempotency-Key"`
}

func (o *CreateUserRequest) GetCreateUserRequest() shared.CreateUserRequest {
	if o == nil {
		return shared.CreateUserRequest{}
	}
	return o.CreateUserRequest
}

func (o *CreateUserRequest) GetXIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.XIdempotencyKey
}

type CreateUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// User created successfully
	User *shared.User
	// Error
	Error *shared.Error
}

func (o *CreateUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUserResponse) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *CreateUserResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
