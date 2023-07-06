// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"push-cash/pkg/models/shared"
)

type GetUserRequest struct {
	// The push identifier for the user
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// successful operation
	User *shared.User
	// Error
	Error *shared.Error
}
