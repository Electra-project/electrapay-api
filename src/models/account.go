package models

import "github.com/shopspring/decimal"

type Account struct {
	Id                  int64    `json:"id"`
	Uuid                string   `json:"uuid"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Type                string   `json:"accounttype"`
	LogoURL             string   `json:"logourl"`
	LogoImg             string   `json:"logoimg"`
	Address1            string   `json:"address1"`
	Address2            string   `json:"address2"`
	Address3            string   `json:"address3"`
	Suburb              string   `json:"suburb"`
	PostalCode          string   `json:"postalcode"`
	City                string   `json:"city"`
	Country             string   `json:"country"`
	Language            string   `json:"language"`
	Timezone            string   `json:"timezone"`
	CallbackURI         string   `json:"callbackurl"`
	Website             string   `json:"website"`
	Currencies          []string `json:"currencies"`
	WalletAddress       string   `json:"walletaddress"`
	WalletCurrency      string   `json:"walletcurrency"`
	ContactTitle        string   `json:"contacttitle"`
	ContactFirstname    string   `json:"contactfirstname"`
	ContactMiddlenames  string   `json:"contactmiddlenames"`
	ContactLastname     string   `json:"contactlastname"`
	ContactEmail        string   `json:"contactemail"`
	ContactPhone        string   `json:"contactphone"`
	ContactMobile       string   `json:"contactmobile"`
	VatNo               string   `json:"vatno"`
	DefaultVAT          int64    `json:"defaultvat"`
	Organisation        string   `json:"orgnisation"`
	PluginType          string   `json:"plugintype"`
	Status              string   `json:"status"`
	ResponseCode        string   `json:"responsecode"`
	ResponseDescription string   `json:"responsedescription"`
}

type AccountNew struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	Type                string `json:"accounttype"`
	Country             string `json:"country"`
	Language            string `json:"language"`
	Timezone            string `json:"timezone"`
	WalletAddress       string `json:"walletaddress"`
	WalletCurrency      string `json:"walletcurrency"`
	ContactFirstname    string `json:"contactfirstname"`
	ContactLastname     string `json:"contactlastname"`
	ContactEmail        string `json:"contactemail"`
	Status              string `json:"status"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountPersonal struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Type                string `json:"accounttype"`
	LogoURL             string `json:"logourl"`
	LogoImg             string `json:"logoimg"`
	Status              string `json:"status"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountPayment struct {
	Id                  int64    `json:"id"`
	CallbackURI         string   `json:"callbackurl"`
	Website             string   `json:"website"`
	Currencies          []string `json:"currencies"`
	WalletAddress       string   `json:"walletaddress"`
	WalletCurrency      string   `json:"walletcurrency"`
	PluginType          string   `json:"plugintype"`
	ResponseCode        string   `json:"responsecode"`
	ResponseDescription string   `json:"responsedescription"`
}

type AccountOrg struct {
	Id                  int64  `json:"id"`
	VatNo               string `json:"vatno"`
	DefaultVAT          int64  `json:"defaultvat"`
	Organisation        string `json:"organisation"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountAPIKey struct {
	Id                  int64  `json:"id"`
	Uuid                string `json:"uuid"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	ApiKey              string `json:"apikey"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountAddress struct {
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

type AccountContact struct {
	AccountId           int64  `json:"accountid"`
	ContactId           int64  `json:"contactid"`
	ContactType         string `json:"contactype"`
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

type AccountWallet struct {
	WalletAddress       string `json:"walletaddress"`
	WalletCurrency      string `json:"walletcurrency"`
	WalletBalance       decimal.Decimal
	ECAPrice            decimal.Decimal
	BTCPrice            decimal.Decimal
	USDPrice            decimal.Decimal
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
