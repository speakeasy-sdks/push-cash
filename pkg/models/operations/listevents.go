// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/v2/pkg/models/shared"
	"push-cash/v2/pkg/utils"
	"time"
)

type ListEventsRequest struct {
	// Return events created after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	CreatedAtAfter *time.Time `queryParam:"style=form,explode=true,name=created_at.after"`
	// Return events created before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	CreatedAtBefore *time.Time `queryParam:"style=form,explode=true,name=created_at.before"`
	// the cursor for the next page of results to fetch
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListEventsRequest) GetCreatedAtAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtAfter
}

func (o *ListEventsRequest) GetCreatedAtBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtBefore
}

func (o *ListEventsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

// ListEventsResponseBody - Successful operation
type ListEventsResponseBody struct {
	Data []shared.Event `json:"data"`
	// Use cursor for paginating list endpoints in conjunction with the cursor request parameter.
	//
	NextCursor string `json:"next_cursor"`
}

func (o *ListEventsResponseBody) GetData() []shared.Event {
	if o == nil {
		return []shared.Event{}
	}
	return o.Data
}

func (o *ListEventsResponseBody) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

type ListEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	Error *shared.Error
	// Successful operation
	Object *ListEventsResponseBody
}

func (o *ListEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListEventsResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListEventsResponse) GetObject() *ListEventsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
