package models

type UserInfo struct {
	Email               string `json:"email"`
	Firstname           string `json:"firstname"`
	Lastname            string `json:"lastname"`
	Avatar              string `json:"avatar"`
	Status              string `json:"status"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
