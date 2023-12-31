// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/pkg/models/shared"
	"time"
)

type ListIntentsRequest struct {
	// Return intents created after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	CreatedAtAfter *time.Time `queryParam:"style=form,explode=true,name=created_at.after"`
	// Return intents created before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	CreatedAtBefore *time.Time `queryParam:"style=form,explode=true,name=created_at.before"`
	// the cursor for the next page of results to fetch
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// filter intents by status, can provide multiple values
	Status []shared.IntentStatus `queryParam:"style=form,explode=true,name=status"`
}

func (o *ListIntentsRequest) GetCreatedAtAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtAfter
}

func (o *ListIntentsRequest) GetCreatedAtBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtBefore
}

func (o *ListIntentsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListIntentsRequest) GetStatus() []shared.IntentStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// ListIntents200ApplicationJSON - successful operation
type ListIntents200ApplicationJSON struct {
	Data []shared.Intent `json:"data"`
	// Use cursor for paginating list endpoints in conjunction with the cursor request parameter.
	//
	NextCursor string `json:"next_cursor"`
}

func (o *ListIntents200ApplicationJSON) GetData() []shared.Intent {
	if o == nil {
		return []shared.Intent{}
	}
	return o.Data
}

func (o *ListIntents200ApplicationJSON) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

type ListIntentsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Error
	Error *shared.Error
	// successful operation
	ListIntents200ApplicationJSONObject *ListIntents200ApplicationJSON
}

func (o *ListIntentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListIntentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListIntentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListIntentsResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListIntentsResponse) GetListIntents200ApplicationJSONObject() *ListIntents200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListIntents200ApplicationJSONObject
}
