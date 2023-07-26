// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/pkg/models/shared"
)

type GetEventRequest struct {
	// The ID of the event to retrieve
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetEventRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetEventResponse struct {
	ContentType string
	// Successful operation
	Event       *shared.Event
	StatusCode  int
	RawResponse *http.Response
	// Error
	Error *shared.Error
}

func (o *GetEventResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEventResponse) GetEvent() *shared.Event {
	if o == nil {
		return nil
	}
	return o.Event
}

func (o *GetEventResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEventResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEventResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
