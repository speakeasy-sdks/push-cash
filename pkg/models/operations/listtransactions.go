// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/v3/pkg/models/shared"
	"push-cash/v3/pkg/utils"
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

func (l ListTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
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

// ListTransactionsResponseBody - Successful operation
type ListTransactionsResponseBody struct {
	Data []shared.Transaction `json:"data"`
	// Use cursor for paginating list endpoints in conjunction with the cursor request parameter.
	//
	NextCursor string `json:"next_cursor"`
}

func (o *ListTransactionsResponseBody) GetData() []shared.Transaction {
	if o == nil {
		return []shared.Transaction{}
	}
	return o.Data
}

func (o *ListTransactionsResponseBody) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

type ListTransactionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	Error *shared.Error
	// Successful operation
	Object *ListTransactionsResponseBody
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

func (o *ListTransactionsResponse) GetObject() *ListTransactionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
