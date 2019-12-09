package models

type UserInfo struct {
	Id                  int64            `json:"id"`
	Email               string           `json:"email"`
	Title               string           `json:"title"`
	Firstname           string           `json:"firstname"`
	Middlename          string           `json:"middlename"`
	Lastname            string           `json:"lastname"`
	Telephone           string           `json:"telephone"`
	Mobile              string           `json:"mobile"`
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
