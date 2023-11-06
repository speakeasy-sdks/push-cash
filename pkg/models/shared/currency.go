// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Currency associated with the amount
type Currency string

const (
	CurrencyUsd Currency = "USD"
)

func (e Currency) ToPointer() *Currency {
	return &e
}

func (e *Currency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USD":
		*e = Currency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Currency: %v", v)
	}
}
