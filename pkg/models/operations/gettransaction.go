// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/pkg/models/shared"
)

type GetTransactionRequest struct {
	// The ID of the transaction to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTransactionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTransactionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful operation
	Transaction *shared.Transaction
	// Error
	Error *shared.Error
}

func (o *GetTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransactionResponse) GetTransaction() *shared.Transaction {
	if o == nil {
		return nil
	}
	return o.Transaction
}

func (o *GetTransactionResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
