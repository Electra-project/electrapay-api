package models

type UserVerify struct {
	Email               string `json:"email"`
	Status              string `json:"status"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
