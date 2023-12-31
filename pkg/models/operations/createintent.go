// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/pkg/models/shared"
)

type CreateIntentRequest struct {
	CreateIntentRequest *shared.CreateIntentRequest `request:"mediaType=application/json"`
	// The idempotency key for the request
	XIdempotencyKey string `header:"style=simple,explode=false,name=X-Idempotency-Key"`
}

func (o *CreateIntentRequest) GetCreateIntentRequest() *shared.CreateIntentRequest {
	if o == nil {
		return nil
	}
	return o.CreateIntentRequest
}

func (o *CreateIntentRequest) GetXIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.XIdempotencyKey
}

type CreateIntentResponse struct {
	ContentType string
	// successful operation
	Intent      *shared.Intent
	StatusCode  int
	RawResponse *http.Response
	// Error
	Error *shared.Error
}

func (o *CreateIntentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIntentResponse) GetIntent() *shared.Intent {
	if o == nil {
		return nil
	}
	return o.Intent
}

func (o *CreateIntentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIntentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateIntentResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
