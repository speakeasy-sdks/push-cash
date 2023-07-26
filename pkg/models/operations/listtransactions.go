// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/pkg/models/shared"
	"time"
)

type ListTransactionsRequest struct {
	// Return transactions created after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	CreatedAtAfter *time.Time `queryParam:"style=form,explode=true,name=created_at.after"`
	// Return transactions created before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	CreatedAtBefore *time.Time `queryParam:"style=form,explode=true,name=created_at.before"`
	// the cursor for the next page of results to fetch
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// filter transactions by status
	Status *shared.TransactionStatus `queryParam:"style=form,explode=true,name=status"`
}

func (o *ListTransactionsRequest) GetCreatedAtAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtAfter
}

func (o *ListTransactionsRequest) GetCreatedAtBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtBefore
}

func (o *ListTransactionsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListTransactionsRequest) GetStatus() *shared.TransactionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// ListTransactions200ApplicationJSON - Successful operation
type ListTransactions200ApplicationJSON struct {
	Data []shared.Transaction `json:"data"`
	// Use cursor for paginating list endpoints in conjunction with the cursor request parameter.
	//
	NextCursor string `json:"next_cursor"`
}

func (o *ListTransactions200ApplicationJSON) GetData() []shared.Transaction {
	if o == nil {
		return []shared.Transaction{}
	}
	return o.Data
}

func (o *ListTransactions200ApplicationJSON) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

type ListTransactionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Error
	Error *shared.Error
	// Successful operation
	ListTransactions200ApplicationJSONObject *ListTransactions200ApplicationJSON
}

func (o *ListTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTransactionsResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListTransactionsResponse) GetListTransactions200ApplicationJSONObject() *ListTransactions200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListTransactions200ApplicationJSONObject
}
