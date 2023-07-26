// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type UpdateUserRequestStatus string

const (
	UpdateUserRequestStatusActive    UpdateUserRequestStatus = "active"
	UpdateUserRequestStatusSuspended UpdateUserRequestStatus = "suspended"
)

func (e UpdateUserRequestStatus) ToPointer() *UpdateUserRequestStatus {
	return &e
}

func (e *UpdateUserRequestStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "suspended":
		*e = UpdateUserRequestStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateUserRequestStatus: %v", v)
	}
}

type UpdateUserRequest struct {
	Status UpdateUserRequestStatus `json:"status"`
}

func (o *UpdateUserRequest) GetStatus() UpdateUserRequestStatus {
	if o == nil {
		return UpdateUserRequestStatus("")
	}
	return o.Status
}
