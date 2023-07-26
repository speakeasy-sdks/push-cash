// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateUserRequest - the request format to create a user with Push
type CreateUserRequest struct {
	// the mailing address for the user
	Address string `json:"address"`
	// the email address for the user
	Email string `json:"email"`
	Kyc   Kyc    `json:"kyc"`
	// first and last name of user
	Name string `json:"name"`
	// the phone number for the user
	Phone string `json:"phone"`
	// the tag for the user, use this field for example to store your own internal identifier for a given user
	Tag string `json:"tag"`
}

func (o *CreateUserRequest) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *CreateUserRequest) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *CreateUserRequest) GetKyc() Kyc {
	if o == nil {
		return Kyc{}
	}
	return o.Kyc
}

func (o *CreateUserRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateUserRequest) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

func (o *CreateUserRequest) GetTag() string {
	if o == nil {
		return ""
	}
	return o.Tag
}
