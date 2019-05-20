package models

type ApiKey struct {
	AccountId           int64  `json:"id"`
	Uuid                string `json:"uuid"`
	WalletAddress       string `json:"walletaddress"`
	ApiKey              string `json:"apikey"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
