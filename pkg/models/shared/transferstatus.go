// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TransferStatus string

const (
	TransferStatusCreated   TransferStatus = "created"
	TransferStatusSubmitted TransferStatus = "submitted"
	TransferStatusFailed    TransferStatus = "failed"
)

func (e TransferStatus) ToPointer() *TransferStatus {
	return &e
}

func (e *TransferStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created":
		fallthrough
	case "submitted":
		fallthrough
	case "failed":
		*e = TransferStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransferStatus: %v", v)
	}
}