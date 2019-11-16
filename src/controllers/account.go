package controllers

import "github.com/gin-gonic/gin"
import (
	"encoding/json"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"github.com/shopspring/decimal"
	"io/ioutil"
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

func (s AccountController) GetPersonalInformation(c *gin.Context) {
	//API to retrieve account information
	// We get the authenticated user

	version := helpers.GetVersion()

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_PERSONAL_FETCH"
	queueinfo.APIType = "GET"
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
	queueinfo, err := queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountPersonal
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

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
	queueinfo, err := queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountPayment
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

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
	queueinfo, err := queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountOrg
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

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
	queueinfo, err := queue.QueueProcess(queueinfo)
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

func (s AccountController) Suspend(c *gin.Context) {
	//API to Close account details

	version := helpers.GetVersion()
	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_SUSPEND"
	queueinfo.APIType = "POST"
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
	queueinfo, err := queue.QueueProcess(queueinfo)
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

func (s AccountController) ApiKey(c *gin.Context) {
	//API to generate a new APIKey

	version := helpers.GetVersion()
	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_APIKEY_RENEW"
	queueinfo.APIType = "POST"
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
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var apikey models.AccountAPIKey
	apikeybyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(apikeybyte, &apikey)

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
	queueinfo, err := queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountAddress
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

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
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var address models.Address
	addressbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(addressbyte, &address)

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
	queueinfo, err := queue.QueueProcess(queueinfo)

	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountContact
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

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
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var contact models.Contact
	contactbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(contactbyte, &contact)

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
		USDPrice, _ = decimal.NewFromString("13.45")
		ECAPrice, _ = decimal.NewFromString("11.45")
		BTCPrice, _ = decimal.NewFromString("12.45")
		WalletAmount, _ = decimal.NewFromString("12.45")
		account.USDPrice = USDPrice
		account.ECAPrice = ECAPrice
		account.BTCPrice = BTCPrice
		account.WalletBalance = WalletAmount
		account.WalletCurrency = "ECA"
		account.WalletAddress = "EVSXj6ExieGBtf4K7Fuw4mBpCwbffwBowm"
		account.ResponseCode = "00"
		account.ResponseDescription = "Success"

		c.JSON(200, account)
	} else {

		var queueinfo queue.Queue
		queueinfo.Category = "ACCOUNT_WALLET_BALANCE"
		queueinfo.APIType = "GET"
		URLArray := strings.Split(c.Request.RequestURI, "/")
		if URLArray[1] != "account" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = URLArray[3]
			queueinfo.Version = URLArray[1]
		}
		if URLArray[1] == "account" {
			queueinfo.APIURL = c.Request.RequestURI
			queueinfo.Parameters = URLArray[2]
			queueinfo.Version = version
		}
		queueinfo.RequestInfo = "{}"
		queueinfo, err := queue.QueueProcess(queueinfo)

		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var account models.AccountWallet
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		if account.ResponseCode != "00" {
			c.JSON(400, account)
		} else {
			c.JSON(200, account)
		}

	}

}

func (s AccountController) OrderSummary(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()

	if c.Request.Header.Get("mock") == "yes" {
		var order models.OrderSummary
		orderbyte := []byte(queueinfo.ResponseInfo)
		order.AwaitingPayment = 10
		order.PaymentReceived = 12
		order.Reversals = 20
		order.Settled = 10

		json.Unmarshal(orderbyte, &order)

		c.JSON(200, order)

	} else {

		queueinfo.Category = "ORDER_SUMMARY"
		queueinfo.APIType = "GET"
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
		queueinfo, err := queue.QueueProcess(queueinfo)
		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var order models.OrderSummary
		orderbyte := []byte(queueinfo.ResponseInfo)

		json.Unmarshal(orderbyte, &order)

		c.JSON(200, order)
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
			orderview.Reference = "#ord"
			orderview.Paymentcategory = "FOOD"
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
		queueinfo, err := queue.QueueProcess(queueinfo)
		if err != nil {
			c.AbortWithError(404, err)
			return
		}
		var orderlist []models.OrderView
		orderbyte := []byte(queueinfo.ResponseInfo)

		json.Unmarshal(orderbyte, &orderlist)

		c.JSON(200, orderlist)
	}

}
