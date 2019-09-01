package models

type UserVerify struct {
	EmailAddress        string `json:"emailaddress"`
	Status              string `json:"status"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
