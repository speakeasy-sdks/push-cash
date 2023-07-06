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

// ListIntents200ApplicationJSON - successful operation
type ListIntents200ApplicationJSON struct {
	Data []shared.Intent `json:"data"`
	// Use cursor for paginating list endpoints in conjunction with the cursor request parameter.
	//
	NextCursor string `json:"next_cursor"`
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
