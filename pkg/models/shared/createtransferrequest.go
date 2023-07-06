// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateTransferRequest struct {
	// amount (in cents) for the transaction
	Amount int64 `json:"amount"`
	// Currency associated with the amount
	Currency Currency `json:"currency"`
	// Direction of the money
	Direction Direction `json:"direction"`
}
