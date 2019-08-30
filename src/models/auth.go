package models

type UserVerify struct {
	Uuid                string `json:"uuid"`
	Status              string `json:"status"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
