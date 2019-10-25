package models

type Address struct {
	AccountId           int64  `json:"accountid"`
	AddressId           int64  `json:"addressid"`
	DefaultAddress      string `json:"defaultaddress"`
	AddressType         string `json:"addresstype"`
	Address1            string `json:"address1"`
	Address2            string `json:"address2"`
	Address3            string `json:"address3"`
	Suburb              string `json:"suburb"`
	PostalCode          string `json:"postalcode"`
	City                string `json:"city"`
	Country             string `json:"country"`
	Language            string `json:"language"`
	Timezone            string `json:"timezone"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
