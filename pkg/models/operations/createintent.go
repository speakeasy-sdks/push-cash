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

type CreateIntentResponse struct {
	ContentType string
	// successful operation
	Intent      *shared.Intent
	StatusCode  int
	RawResponse *http.Response
	// Error
	Error *shared.Error
}
