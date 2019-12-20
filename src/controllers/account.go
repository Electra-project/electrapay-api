package controllers

import "github.com/gin-gonic/gin"
import (
	"encoding/json"
	"github.com/ruannelloyd/electrapay-api/src/helpers"
	"github.com/ruannelloyd/electrapay-api/src/models"
	"github.com/ruannelloyd/electrapay-api/src/queue"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type AccountController struct{}

func (s AccountController) Register(c *gin.Context) {
	//API to register a new account

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_REGISTER"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err := queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var account models.Account
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		if account.ResponseCode != "00" {
			c.JSON(400, account)
		} else {
			c.JSON(200, account)
		}
	}
}

func (s AccountController) GetAccount(c *gin.Context) {
	//API to retrieve account information
	// We get the authenticated user

	version := helpers.GetVersion()
	t, err := extractToken(c)

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_FETCH"
	queueinfo.APIType = "GET"
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.Account
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	if account.ResponseCode != "00" {
		c.JSON(400, account)
	} else {
		c.JSON(200, account)
	}
}

func (s AccountController) GetAccountLogo(c *gin.Context) {
	//API to retrieve account logo information
	// We get the authenticated user

	version := helpers.GetVersion()
	t, err := extractToken(c)

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_LOGO_FETCH"
	queueinfo.APIType = "GET"
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	logo := queueinfo.ResponseInfo

	if len(logo) < 1 {
		c.JSON(400, "")
	} else {
		c.JSON(200, logo)
	}
}

func (s AccountController) EditAccountLogo(c *gin.Context) {

	//API to Edit account details
	version := helpers.GetVersion()
	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_LOGO_EDIT"
	queueinfo.APIType = "PUT"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(200, returnError)
	}
}

func (s AccountController) GetPersonalInformation(c *gin.Context) {
	//API to retrieve account information
	// We get the authenticated user

	version := helpers.GetVersion()

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_PERSONAL_FETCH"
	queueinfo.APIType = "GET"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountPersonal
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	if queueinfo.ResponseCode != "00" {
		account.ResponseCode = queueinfo.ResponseCode
		account.ResponseDescription = queueinfo.ResponseDescription
	}

	if account.ResponseCode != "00" {
		c.JSON(400, account)
	} else {
		c.JSON(200, account)
	}
}

func (s AccountController) EditPersonalInformation(c *gin.Context) {

	//API to Edit account details
	version := helpers.GetVersion()
	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_PERSONAL_EDIT"
	queueinfo.APIType = "PUT"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var account models.AccountPersonal
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		if account.ResponseCode != "00" {
			c.JSON(400, account)
		} else {
			c.JSON(200, account)
		}
	}
}

func (s AccountController) GetPaymentDetails(c *gin.Context) {
	// API to retrieve account information - Payment Details
	// We get the authenticated user

	version := helpers.GetVersion()

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_PAYMENT_FETCH"
	queueinfo.APIType = "GET"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountPayment
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	if queueinfo.ResponseCode != "00" {
		account.ResponseCode = queueinfo.ResponseCode
		account.ResponseDescription = queueinfo.ResponseDescription
	}

	if account.ResponseCode != "00" {
		c.JSON(400, account)
	} else {
		c.JSON(200, account)
	}
}

func (s AccountController) EditPaymentDetails(c *gin.Context) {

	//API to Edit account details
	version := helpers.GetVersion()
	var queueinfo queue.Queue
	t, err := extractToken(c)
	queueinfo.Token = t
	queueinfo.Category = "ACCOUNT_PAYMENT_EDIT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var account models.AccountPayment
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		if account.ResponseCode != "00" {
			c.JSON(400, account)
		} else {
			c.JSON(200, account)
		}
	}
}

func (s AccountController) GetOrganizationDetails(c *gin.Context) {
	//API to retrieve account information - organization Details
	// We get the authenticated user

	version := helpers.GetVersion()

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_ORG_FETCH"
	queueinfo.APIType = "GET"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountOrg
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	if queueinfo.ResponseCode != "00" {
		account.ResponseCode = queueinfo.ResponseCode
		account.ResponseDescription = queueinfo.ResponseDescription
	}

	if account.ResponseCode != "00" {
		c.JSON(400, account)
	} else {
		c.JSON(200, account)
	}
}

func (s AccountController) EditOrganizationDetails(c *gin.Context) {

	//API to Edit account details
	version := helpers.GetVersion()
	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_ORG_EDIT"
	queueinfo.APIType = "PUT"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var account models.AccountOrg
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		if account.ResponseCode != "00" {
			c.JSON(400, account)
		} else {
			c.JSON(200, account)
		}
	}
}

func (s AccountController) Close(c *gin.Context) {
	//API to Close account details
	version := helpers.GetVersion()

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_CLOSE"
	queueinfo.APIType = "POST"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.Account
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	if queueinfo.ResponseCode != "00" {
		account.ResponseCode = queueinfo.ResponseCode
		account.ResponseDescription = queueinfo.ResponseDescription
	}

	if account.ResponseCode != "00" {
		c.JSON(400, account)
	} else {
		c.JSON(200, account)
	}
}

func (s AccountController) Suspend(c *gin.Context) {
	//API to Close account details

	version := helpers.GetVersion()
	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_SUSPEND"
	queueinfo.APIType = "POST"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.Account
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	if queueinfo.ResponseCode != "00" {
		account.ResponseCode = queueinfo.ResponseCode
		account.ResponseDescription = queueinfo.ResponseDescription
	}

	if account.ResponseCode != "00" {
		c.JSON(400, account)
	} else {
		c.JSON(200, account)
	}
}

func (s AccountController) ApiKey(c *gin.Context) {
	//API to generate a new APIKey

	version := helpers.GetVersion()
	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_APIKEY_RENEW"
	queueinfo.APIType = "POST"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var apikey models.AccountAPIKey
	apikeybyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(apikeybyte, &apikey)

	if queueinfo.ResponseCode != "00" {
		apikey.ResponseCode = queueinfo.ResponseCode
		apikey.ResponseDescription = queueinfo.ResponseDescription
	}

	if apikey.ResponseCode != "00" {
		c.JSON(400, apikey)
	} else {
		c.JSON(200, apikey)
	}
}

func (s AccountController) AddressEdit(c *gin.Context) {

	//API to Edit account Address details

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_ADDRESS_EDIT"
	queueinfo.APIType = "PUT"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var address models.Address
		addressbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(addressbyte, &address)

		if address.ResponseCode != "00" {
			c.JSON(400, address)
		} else {
			c.JSON(200, address)
		}
	}
}

func (s AccountController) AddressFetch(c *gin.Context) {

	//API to Edit account Address details

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_ADDRESS_FETCH"
	queueinfo.APIType = "GET"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("addresstype")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("addresstype")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountAddress
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	if queueinfo.ResponseCode != "00" {
		account.ResponseCode = queueinfo.ResponseCode
		account.ResponseDescription = queueinfo.ResponseDescription
	}

	if account.ResponseCode != "00" {
		c.JSON(400, account)
	} else {
		c.JSON(200, account)
	}
}

func (s AccountController) AddressAdd(c *gin.Context) {

	//API to Add account address details

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_ADDRESS_NEW"
	queueinfo.APIType = "POST"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var address models.Address
		addressbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(addressbyte, &address)

		if address.ResponseCode != "00" {
			c.JSON(400, address)
		} else {
			c.JSON(200, address)
		}
	}
}

func (s AccountController) AddressRemove(c *gin.Context) {

	//API to Delete account address details
	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_ADDRESS_DELETE"
	queueinfo.APIType = "DELETE"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("addressid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("addressid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var address models.Address
	addressbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(addressbyte, &address)

	if queueinfo.ResponseCode != "00" {
		address.ResponseCode = queueinfo.ResponseCode
		address.ResponseDescription = queueinfo.ResponseDescription
	}

	if address.ResponseCode != "00" {
		c.JSON(400, address)
	} else {
		c.JSON(200, address)
	}
}

func (s AccountController) ContactFetch(c *gin.Context) {

	//API to Edit account Address details

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_CONTACT_FETCH"
	queueinfo.APIType = "GET"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("contacttype")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("contacttype")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountContact
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	if queueinfo.ResponseCode != "00" {
		account.ResponseCode = queueinfo.ResponseCode
		account.ResponseDescription = queueinfo.ResponseDescription
	}

	if account.ResponseCode != "00" {
		c.JSON(400, account)
	} else {
		c.JSON(200, account)
	}
}

func (s AccountController) ContactEdit(c *gin.Context) {

	//API to Edit account Contact details

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_CONTACT_EDIT"
	queueinfo.APIType = "PUT"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("contactid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("contactid")
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var contact models.Contact
		contactbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(contactbyte, &contact)

		if contact.ResponseCode != "00" {
			c.JSON(400, contact)
		} else {
			c.JSON(200, contact)
		}
	}
}

func (s AccountController) ContactAdd(c *gin.Context) {

	//API to Add account contact details

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_CONTACT_NEW"
	queueinfo.APIType = "POST"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = version
	}
	x, _ := ioutil.ReadAll(c.Request.Body)
	queueinfo.RequestInfo = string(x)
	queueinfo, err = queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var contact models.Contact
		contactbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(contactbyte, &contact)

		if contact.ResponseCode != "00" {
			c.JSON(400, contact)
		} else {
			c.JSON(200, contact)
		}
	}
}

func (s AccountController) ContactRemove(c *gin.Context) {

	//API to Delete contact address details
	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_CONTACT_DELETE"
	queueinfo.APIType = "DELETE"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("contactid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("contactid")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err = queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var contact models.Contact
	contactbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(contactbyte, &contact)

	if queueinfo.ResponseCode != "00" {
		contact.ResponseCode = queueinfo.ResponseCode
		contact.ResponseDescription = queueinfo.ResponseDescription
	}

	if contact.ResponseCode != "00" {
		c.JSON(400, contact)
	} else {
		c.JSON(200, contact)
	}
}

func (s AccountController) AccountBalance(c *gin.Context) {
	//API to retrieve account Balance information
	// We get the authenticated user

	version := helpers.GetVersion()

	if c.Request.Header.Get("mock") == "yes" {
		var account models.AccountWallet
		var USDPrice, ECAPrice, BTCPrice, WalletAmount decimal.Decimal
		USDPrice, _ = decimal.NewFromString("0.0002793411")
		ECAPrice, _ = decimal.NewFromString("1")
		BTCPrice, _ = decimal.NewFromString("0.00000003")
		WalletAmount, _ = decimal.NewFromString("802.25")
		account.USDPrice = USDPrice
		account.ECAPrice = ECAPrice
		account.BTCPrice = BTCPrice
		account.WalletBalance = WalletAmount
		account.WalletFiat = account.WalletBalance.Mul(account.BTCPrice).Round(2)
		account.WalletCurrency = "ECA"
		account.WalletAddress = "EVSXj6ExieGBtf4K7Fuw4mBpCwbffwBowm"
		account.ResponseCode = "00"
		account.ResponseDescription = "Success"

		c.JSON(200, account)
	} else {

		var queueinfo queue.Queue
		queueinfo.Category = "ACCOUNT_WALLET_BALANCE"
		queueinfo.APIType = "GET"
		t, err := extractToken(c)
		queueinfo.Token = t
		URLArray := strings.Split(c.Request.RequestURI, "/")
		if URLArray[1] != "account" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid")
			queueinfo.Version = URLArray[1]
		}
		if URLArray[1] == "account" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid")
			queueinfo.Version = version
		}
		queueinfo.RequestInfo = "{}"
		queueinfo, err = queue.QueueProcess(queueinfo)

		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var account models.AccountWallet
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		if queueinfo.ResponseCode != "00" {
			account.ResponseCode = queueinfo.ResponseCode
			account.ResponseDescription = queueinfo.ResponseDescription
		}

		if account.ResponseCode == "00" {
			c.JSON(200, account)
		} else {
			c.JSON(400, account)
		}

	}

}

func (s AccountController) RulesFetch(c *gin.Context) {
	//API to retrieve account Rules information
	// We get the authenticated user

	version := helpers.GetVersion()

	if c.Request.Header.Get("mock") == "yes" {
		var accountrules []models.AccountRule
		var accountrule models.AccountRule
		var rulecompulsory models.RuleCompulsory
		var ruleparameter models.RuleParameter

		rulecompulsory.Value = true
		rulecompulsory.Condition = ""

		accountrule.Code = "001"
		accountrule.Display = "Pay {{parameter1}} % to the Electra Foundation"
		accountrule.Description = "Enter a percentage of the order to donate back to the Foundation"
		accountrule.Compulsory = rulecompulsory

		ruleparameter.Parameter = "parameter1"
		ruleparameter.Type = "integer"
		ruleparameter.Validation = "parameter1 >=2 && parameter1 <=100"
		ruleparameter.Options = ""
		ruleparameter.Description = "Percentage"
		ruleparameter.Value = "2"
		accountrule.Parameters = append(accountrule.Parameters, ruleparameter)

		accountrules = append(accountrules, accountrule)

		c.JSON(200, accountrules)
	} else {

		var queueinfo queue.Queue
		queueinfo.Category = "ACCOUNT_RULES_FETCH"
		queueinfo.APIType = "GET"
		t, err := extractToken(c)
		queueinfo.Token = t
		URLArray := strings.Split(c.Request.RequestURI, "/")
		if URLArray[1] != "account" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid")
			queueinfo.Version = URLArray[1]
		}
		if URLArray[1] == "account" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid")
			queueinfo.Version = version
		}
		queueinfo.RequestInfo = "{}"
		queueinfo, err = queue.QueueProcess(queueinfo)

		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var rulelist []models.AccountRule
		rulelistbyte := []byte(queueinfo.ResponseInfo)

		json.Unmarshal(rulelistbyte, &rulelist)
		if len(rulelist) == 0 {
			empty := []string{}
			c.JSON(200, empty)
		} else {
			c.JSON(200, rulelist)
		}

	}

}

func (s AccountController) OrderSummary(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()

	if c.Request.Header.Get("mock") == "yes" {
		var order models.OrderSummary
		var orderseries models.OrderSeries

		orderbyte := []byte(queueinfo.ResponseInfo)
		order.AwaitingPayment = 101
		order.PaymentReceived = 5034
		order.Reversals = 10
		order.Settled = 4098
		orderseries.Name = "Total Ordered"
		num1, _ := decimal.NewFromString("6001.20")
		num2, _ := decimal.NewFromString("5012.22")
		num3, _ := decimal.NewFromString("4005.45")
		num4, _ := decimal.NewFromString("5678.90")
		num5, _ := decimal.NewFromString("8809.01")
		num6, _ := decimal.NewFromString("9908.44")
		num7, _ := decimal.NewFromString("7560.20")
		num8, _ := decimal.NewFromString("9078.42")
		num9, _ := decimal.NewFromString("4598.33")
		num10, _ := decimal.NewFromString("8908.90")
		orderseries.Data = append(orderseries.Data, num1)
		orderseries.Data = append(orderseries.Data, num2)
		orderseries.Data = append(orderseries.Data, num3)
		orderseries.Data = append(orderseries.Data, num4)
		orderseries.Data = append(orderseries.Data, num5)
		orderseries.Data = append(orderseries.Data, num6)
		orderseries.Data = append(orderseries.Data, num7)
		orderseries.Data = append(orderseries.Data, num8)
		orderseries.Data = append(orderseries.Data, num9)
		orderseries.Data = append(orderseries.Data, num10)
		orderseries.Total, _ = decimal.NewFromString("69561.07")
		order.Series = append(order.Series, orderseries)

		orderseries = models.OrderSeries{}

		orderseries.Name = "Total Settled"
		num1, _ = decimal.NewFromString("5402.20")
		num2, _ = decimal.NewFromString("5001.22")
		num3, _ = decimal.NewFromString("3967.45")
		num4, _ = decimal.NewFromString("5452.90")
		num5, _ = decimal.NewFromString("7098.01")
		num6, _ = decimal.NewFromString("9902.44")
		num7, _ = decimal.NewFromString("7342.19")
		num8, _ = decimal.NewFromString("8034.56")
		num9, _ = decimal.NewFromString("4498.23")
		num10, _ = decimal.NewFromString("8309.43")
		orderseries.Data = append(orderseries.Data, num1)
		orderseries.Data = append(orderseries.Data, num2)
		orderseries.Data = append(orderseries.Data, num3)
		orderseries.Data = append(orderseries.Data, num4)
		orderseries.Data = append(orderseries.Data, num5)
		orderseries.Data = append(orderseries.Data, num6)
		orderseries.Data = append(orderseries.Data, num7)
		orderseries.Data = append(orderseries.Data, num8)
		orderseries.Data = append(orderseries.Data, num9)
		orderseries.Data = append(orderseries.Data, num10)
		orderseries.Total, _ = decimal.NewFromString("65008.63")
		order.Series = append(order.Series, orderseries)

		ordertimeline := models.OrderTimeline{}

		ordertimeline.Name = "per month"
		string1 := "Jan 2019"
		string2 := "Feb 2019"
		string3 := "Mar 2019"
		string4 := "Apr 2019"
		string5 := "May 2019"
		string6 := "Jun 2019"
		string7 := "Jul 2019"
		string8 := "Aug 2019"
		string9 := "Sept 2019"
		string10 := "Oct 2019"

		ordertimeline.Data = append(ordertimeline.Data, string1)
		ordertimeline.Data = append(ordertimeline.Data, string2)
		ordertimeline.Data = append(ordertimeline.Data, string3)
		ordertimeline.Data = append(ordertimeline.Data, string4)
		ordertimeline.Data = append(ordertimeline.Data, string5)
		ordertimeline.Data = append(ordertimeline.Data, string6)
		ordertimeline.Data = append(ordertimeline.Data, string7)
		ordertimeline.Data = append(ordertimeline.Data, string8)
		ordertimeline.Data = append(ordertimeline.Data, string9)
		ordertimeline.Data = append(ordertimeline.Data, string10)
		order.Timeline = ordertimeline

		json.Unmarshal(orderbyte, &order)

		c.JSON(200, order)

	} else {

		queueinfo.Category = "ORDER_SUMMARY"
		queueinfo.APIType = "GET"
		t, err := extractToken(c)
		queueinfo.Token = t
		URLArray := strings.Split(c.Request.RequestURI, "/")
		if URLArray[1] != "order" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("frequency")
			queueinfo.Version = URLArray[1]
		}
		if URLArray[1] == "order" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("frequency")
			queueinfo.Version = version
		}
		queueinfo.RequestInfo = "{}"
		queueinfo, err = queue.QueueProcess(queueinfo)
		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var order models.OrderSummary
		orderbyte := []byte(queueinfo.ResponseInfo)

		json.Unmarshal(orderbyte, &order)

		if queueinfo.ResponseCode != "00" {
			order.ResponseCode = queueinfo.ResponseCode
			order.ResponseDescription = queueinfo.ResponseDescription
		}

		if order.ResponseCode == "00" {
			c.JSON(200, order)
		} else {
			c.JSON(400, order)
		}
	}
}

func (s AccountController) OrderList(c *gin.Context) {

	if c.Request.Header.Get("mock") == "yes" {
		var amount decimal.Decimal
		var i int64
		amount, _ = decimal.NewFromString("10.00")
		var queueinfo queue.Queue

		var orderlist []models.OrderView
		orderbyte := []byte(queueinfo.ResponseInfo)

		for i = 0; i < 10; i++ {

			var orderview models.OrderView
			orderview.OrderId = i
			orderview.Reference = strings.Join([]string{"ord#", strconv.FormatInt(i+1, 10)}, "")
			orderview.Paymentcategory = "ElectraPay Donation"
			orderview.OrderCurrency = "USD"
			orderview.OrderAmount = amount
			orderview.QuoteCurrency = "ECA"
			orderview.QuoteTotal = amount
			orderview.OrderDate = time.Now()
			orderview.OrderQuoteSubmittedDate = time.Now()
			orderview.OrderReceivedPaymentDate = time.Now()
			orderview.OrderSettled = true
			orderview.OrderStatus = "SETTLED"

			orderlist = append(orderlist, orderview)
		}

		json.Unmarshal(orderbyte, &orderlist)

		c.JSON(200, orderlist)

	} else {

		var queueinfo queue.Queue
		version := helpers.GetVersion()

		queueinfo.Category = "ORDER_LIST"
		queueinfo.APIType = "GET"
		t, err := extractToken(c)
		queueinfo.Token = t
		URLArray := strings.Split(c.Request.RequestURI, "/")
		if URLArray[1] != "order" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid")
			queueinfo.Version = URLArray[1]
		}
		if URLArray[1] == "order" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid")
			queueinfo.Version = version
		}
		queueinfo.RequestInfo = "{}"
		queueinfo, err = queue.QueueProcess(queueinfo)
		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var orderlist []models.OrderView
		orderbyte := []byte(queueinfo.ResponseInfo)

		json.Unmarshal(orderbyte, &orderlist)
		if len(orderlist) == 0 {
			empty := []string{}
			c.JSON(200, empty)
		} else {
			c.JSON(200, orderlist)
		}

	}

}

func (s AccountController) OrderListMax(c *gin.Context) {

	if c.Request.Header.Get("mock") == "yes" {
		var amount decimal.Decimal
		var i int64
		amount, _ = decimal.NewFromString("10.00")
		var queueinfo queue.Queue

		var orderlist []models.OrderView
		orderbyte := []byte(queueinfo.ResponseInfo)

		for i = 0; i < 10; i++ {

			var orderview models.OrderView
			orderview.OrderId = i
			orderview.Reference = strings.Join([]string{"ord#", strconv.FormatInt(i+1, 10)}, "")
			orderview.Paymentcategory = "ElectraPay Donation"
			orderview.OrderCurrency = "USD"
			orderview.OrderAmount = amount
			orderview.QuoteCurrency = "ECA"
			orderview.QuoteTotal = amount
			orderview.OrderDate = time.Now()
			orderview.OrderQuoteSubmittedDate = time.Now()
			orderview.OrderReceivedPaymentDate = time.Now()
			orderview.OrderSettled = true
			orderview.OrderStatus = "SETTLED"

			orderlist = append(orderlist, orderview)
		}

		json.Unmarshal(orderbyte, &orderlist)

		c.JSON(200, orderlist)

	} else {

		var queueinfo queue.Queue
		version := helpers.GetVersion()

		queueinfo.Category = "ORDER_LIST_MAX"
		queueinfo.APIType = "GET"
		t, err := extractToken(c)
		queueinfo.Token = t
		URLArray := strings.Split(c.Request.RequestURI, "/")
		if URLArray[1] != "order" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("maxlimit")
			queueinfo.Version = URLArray[1]
		}
		if URLArray[1] == "order" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid") + "?" + c.Param("maxlimit")
			queueinfo.Version = version
		}
		queueinfo.RequestInfo = "{}"
		queueinfo, err = queue.QueueProcess(queueinfo)
		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var orderlist []models.OrderView
		orderbyte := []byte(queueinfo.ResponseInfo)

		json.Unmarshal(orderbyte, &orderlist)

		if len(orderlist) == 0 {
			empty := []string{}
			c.JSON(200, empty)
		} else {
			c.JSON(200, orderlist)
		}
	}

}

func (s AccountController) ActivityList(c *gin.Context) {

	if c.Request.Header.Get("mock") == "yes" {
		var queueinfo queue.Queue
		var i int64

		var activitylist []models.AccountActivity
		activitybyte := []byte(queueinfo.ResponseInfo)

		for i = 0; i < 10; i++ {

			var activity models.AccountActivity
			activity.ActivityDate = time.Now()
			activity.Category = "account"
			activity.Description = "Update account address details"
			activity.UserName = "John Doe"

			activitylist = append(activitylist, activity)
		}

		json.Unmarshal(activitybyte, &activitylist)

		c.JSON(200, activitylist)

	} else {

		var queueinfo queue.Queue
		version := helpers.GetVersion()

		queueinfo.Category = "ACTIVITY_LIST"
		queueinfo.APIType = "GET"
		t, err := extractToken(c)
		queueinfo.Token = t
		URLArray := strings.Split(c.Request.RequestURI, "/")
		if URLArray[1] != "account" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid")
			queueinfo.Version = URLArray[1]
		}
		if URLArray[1] == "account" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = c.Param("accountid")
			queueinfo.Version = version
		}
		queueinfo.RequestInfo = "{}"
		queueinfo, err = queue.QueueProcess(queueinfo)
		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var activitylist []models.AccountActivity
		activitybyte := []byte(queueinfo.ResponseInfo)

		json.Unmarshal(activitybyte, &activitylist)

		if len(activitylist) == 0 {
			empty := []string{}
			c.JSON(200, empty)
		} else {
			c.JSON(200, activitylist)
		}
	}

}
