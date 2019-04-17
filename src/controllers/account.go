package controllers

import "github.com/gin-gonic/gin"
import (
	"encoding/json"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"strconv"
	"strings"
)

type Error struct {
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type AccountEdit struct {
	Id                  int64    `json:"id"`
	Uuid                string   `json:"uuid"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Type                string   `json:"accounttype"`
	LogoURL             string   `json:"logourl"`
	LogoImg             string   `json:"logoimg"`
	Country             string   `json:"country"`
	Language            string   `json:"language"`
	Timezone            string   `json:"timezone"`
	CallbackURI         string   `json:"callbackurl"`
	Website             string   `json:"website"`
	Currencies          []string `json:"currencies"`
	WalletAddress       string   `json:"walletaddress"`
	WalletCurrency      string   `json:"walletcurrency"`
	VatNo               string   `json:"vatno"`
	DefaultVAT          int64    `json:"defaultvat"`
	Organisation        string   `json:"orgnisation"`
	PluginType          string   `json:"plugintype"`
	Status              string   `json:"status"`
	ResponseCode        string   `json:"responsecode"`
	ResponseDescription string   `json:"responsedescription"`
}

type Address struct {
	AccountId           int64  `json:"id"`
	Uuid                string `json:"uuid"`
	AddressType         string `json:"addresstype"`
	Address1            string `json:"address1"`
	Address2            string `json:"address2"`
	Address3            string `json:"address3"`
	Suburb              string `json:"suburb"`
	PostalCode          string `json:"postalcode"`
	City                string `json:"city"`
	Country             string `json:"country"`
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}

type Contact struct {
	AccountId           int64  `json:"id"`
	Uuid                string `json:"uuid"`
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

type AccountController struct{}

func (s AccountController) Get(c *gin.Context) {
	//API to retrieve account information
	// We get the authenticated user
	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_FIND"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = string(authenticatedAccount.Id)
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.Account
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	c.Header("X-Version", "1.0")
	c.JSON(200, account)

}

func (s AccountController) New(c *gin.Context) {
	//API to create a new account

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_NEW"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 2 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = "v1"
	}
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	queueinfo.RequestInfo = string(buf[0:num])
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	if queueinfo.ResponseCode != "00" {
		returnError := Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var account models.Account
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		c.Header("X-Version", "1.0")
		c.JSON(200, account)

	}
}

func (s AccountController) Register(c *gin.Context) {

	//API to Register an account to make it active
	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_REGISTER"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.Account
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	c.Header("X-Version", "1.0")
	c.JSON(200, account)
}

func (s AccountController) Edit(c *gin.Context) {

	//API to Edit account details

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_EDIT"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.Account
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	c.Header("X-Version", "1.0")
	c.JSON(200, account)

}

func (s AccountController) Close(c *gin.Context) {

	//API to Close account details

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.Account
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	c.Header("X-Version", "1.0")
	c.JSON(200, account)

}

func (s AccountController) AddressEdit(c *gin.Context) {

	//API to Edit account Address details

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var address Address
	addressbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(addressbyte, &address)

	c.Header("X-Version", "1.0")
	c.JSON(200, address)

}

func (s AccountController) AddressAdd(c *gin.Context) {

	//API to Edit account details

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var address Address
	addressbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(addressbyte, &address)

	c.Header("X-Version", "1.0")
	c.JSON(200, address)

}

func (s AccountController) AddressRemove(c *gin.Context) {

	//API to Delete account address details

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT"
	queueinfo.APIType = "DELETE"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var address Address
	addressbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(addressbyte, &address)

	c.Header("X-Version", "1.0")
	c.JSON(200, address)

}

func (s AccountController) ContactEdit(c *gin.Context) {

	//API to Edit account Contact details

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var contact Contact
	contactbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(contactbyte, &contact)

	c.Header("X-Version", "1.0")
	c.JSON(200, contact)

}

func (s AccountController) ContactAdd(c *gin.Context) {
	//API to Add account Contact details

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var contact Address
	contactbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(contactbyte, &contact)

	c.Header("X-Version", "1.0")
	c.JSON(200, contact)

}

func (s AccountController) ContactRemove(c *gin.Context) {

	//API to Remove account contact details

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT"
	queueinfo.APIType = "DELETE"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray) == 3 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = "v1"
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var contact Contact
	contactbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(contactbyte, &contact)

	c.Header("X-Version", "1.0")
	c.JSON(200, contact)

}
