package models

type UserInfo struct {
	Email               string           `json:"email"`
	Firstname           string           `json:"firstname"`
	Lastname            string           `json:"lastname"`
	Avatar              string           `json:"avatar"`
	Status              string           `json:"status"`
	AccountInfo         []AccountDisplay `json:"accountinfo"`
	ResponseCode        string           `json:"responsecode"`
	ResponseDescription string           `json:"responsedescription"`
}

type AccountDisplay struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Website             string `json:"website"`
	WalletAddress       string `json:"walletaddress"`
	WalletCurrency      string `json:"walletcurrency"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
