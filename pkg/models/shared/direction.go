// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Direction - Direction of the money
type Direction string

const (
	DirectionCashIn  Direction = "cash_in"
	DirectionCashOut Direction = "cash_out"
)

func (e Direction) ToPointer() *Direction {
	return &e
}

func (e *Direction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cash_in":
		fallthrough
	case "cash_out":
		*e = Direction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Direction: %v", v)
	}
}
