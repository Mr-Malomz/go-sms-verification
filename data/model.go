package data

type Data struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

type VerifyData struct {
	User *Data  `json:"user,omitempty"`
	Code string `json:"code,omitempty" validate:"required"`
}
