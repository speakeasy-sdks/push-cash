// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/pkg/models/shared"
)

type GetBalanceResponse struct {
	// Account balance retrieved successfully
	AccountBalance *shared.AccountBalance
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	// Error
	Error *shared.Error
}

func (o *GetBalanceResponse) GetAccountBalance() *shared.AccountBalance {
	if o == nil {
		return nil
	}
	return o.AccountBalance
}

func (o *GetBalanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBalanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBalanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetBalanceResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
