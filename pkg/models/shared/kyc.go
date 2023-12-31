// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// KYCMethod - The KYC provider
type KYCMethod string

const (
	KYCMethodAlloy KYCMethod = "alloy"
)

func (e KYCMethod) ToPointer() *KYCMethod {
	return &e
}

func (e *KYCMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "alloy":
		*e = KYCMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KYCMethod: %v", v)
	}
}

type Kyc struct {
	// The KYC provider
	Method KYCMethod `json:"method"`
	// the token representing the user entity at your KYC provider
	Token string `json:"token"`
}

func (o *Kyc) GetMethod() KYCMethod {
	if o == nil {
		return KYCMethod("")
	}
	return o.Method
}

func (o *Kyc) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}
