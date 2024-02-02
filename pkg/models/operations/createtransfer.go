// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/v3/pkg/models/shared"
)

type CreateTransferRequest struct {
	CreateTransferRequest *shared.CreateTransferRequest `request:"mediaType=application/json"`
	// The idempotency key for the request
	XIdempotencyKey string `header:"style=simple,explode=false,name=X-Idempotency-Key"`
}

func (o *CreateTransferRequest) GetCreateTransferRequest() *shared.CreateTransferRequest {
	if o == nil {
		return nil
	}
	return o.CreateTransferRequest
}

func (o *CreateTransferRequest) GetXIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.XIdempotencyKey
}

type CreateTransferResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	Transfer *shared.Transfer
	// Error
	Error *shared.Error
}

func (o *CreateTransferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTransferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTransferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTransferResponse) GetTransfer() *shared.Transfer {
	if o == nil {
		return nil
	}
	return o.Transfer
}

func (o *CreateTransferResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
