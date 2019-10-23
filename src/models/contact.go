package models

type Contact struct {
	AccountId           int64  `json:"accountid"`
	ContactId           int64  `json:"addressid"`
	ContactType         string `json:"contacttype"`
	ContactTitle        string `json:"contacttitle"`
	ContactFirstname    string `json:"contactfirstname"`
	ContactMiddlenames  string `json:"contactmiddlenames"`
	ContactLastname     string `json:"contactlastname"`
	ContactEmail        string `json:"contactemail"`
	ContactPhone        string `json:"contactphone"`
	ContactMobile       string `json:"contactmobile"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
