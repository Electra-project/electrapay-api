package controllers

import "github.com/gin-gonic/gin"
import (
	"encoding/json"
	"fmt"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"strconv"
	"strings"
)

type AccountController struct{}

func (s AccountController) AuthVerify(c *gin.Context) {
	//API to verify the status of a user

	var queueinfo queue.Queue
	queueinfo.Category = "AUTH_VERIFY"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	version := helpers.GetVersion()

	if URLArray[1] != "auth" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[4]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "auth" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)

	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var user models.UserVerify
		userbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(userbyte, &user)

		c.JSON(200, user)
	}
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

}

func (s AccountController) SetPassword(c *gin.Context) {
	//API to set the user password
	version := helpers.GetVersion()

	var queueinfo queue.Queue
	queueinfo.Category = "AUTH_SETPASSWORD"
	queueinfo.APIType = "PUT"
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
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	queueinfo.RequestInfo = string(buf[0:num])
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var user models.UserVerify
		userbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(userbyte, &user)

		c.JSON(200, user)
	}

}

func (s AccountController) Get(c *gin.Context) {
	//API to retrieve account information
	// We get the authenticated user
	user, _ := c.Get("uuid")
	version := helpers.GetVersion()
	var authenticatedAccount = user.(*models.Account)

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_FETCH"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	fmt.Print(len(URLArray))
	fmt.Print(URLArray[1])
	fmt.Print(URLArray[2])
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var account models.AccountAPIKey
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)

	c.JSON(200, account)
}

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
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	queueinfo.RequestInfo = string(buf[0:num])
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var account models.Account
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		c.JSON(200, account)
	}
}

func (s AccountController) Edit(c *gin.Context) {

	//API to Edit account details
	version := helpers.GetVersion()
	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_EDIT"
	queueinfo.APIType = "PUT"
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
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	queueinfo.RequestInfo = string(buf[0:num])
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var account models.AccountEdit
		accountbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(accountbyte, &account)

		c.JSON(200, account)
	}
}

func (s AccountController) Close(c *gin.Context) {
	//API to Close account details
	version := helpers.GetVersion()
	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)

	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_CLOSE"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
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

	c.JSON(200, account)
}

func (s AccountController) Suspend(c *gin.Context) {
	//API to Close account details

	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)
	version := helpers.GetVersion()
	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_SUSPEND"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
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

	c.JSON(200, account)
}

func (s AccountController) ApiKey(c *gin.Context) {
	//API to generate a new APIKey

	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)
	version := helpers.GetVersion()
	var queueinfo queue.Queue

	queueinfo.Category = "ACCOUNT_APIKEY_RENEW"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
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

	c.JSON(200, apikey)
}

func (s AccountController) AddressEdit(c *gin.Context) {

	//API to Edit account Address details

	var queueinfo queue.Queue
	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_ADDRESS_EDIT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id)) + "," + URLArray[4]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id)) + "," + URLArray[3]
		queueinfo.Version = version
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
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var address models.Address
		addressbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(addressbyte, &address)

		c.JSON(200, address)
	}
}

func (s AccountController) AddressAdd(c *gin.Context) {

	//API to Add account address details

	var queueinfo queue.Queue
	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_ADDRESS_NEW"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = version
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
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var address models.Address
		addressbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(addressbyte, &address)

		c.JSON(200, address)
	}
}

func (s AccountController) AddressRemove(c *gin.Context) {

	//API to Delete account address details
	var queueinfo queue.Queue
	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_ADDRESS_DELETE"
	queueinfo.APIType = "DELETE"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id)) + "," + URLArray[4]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id)) + "," + URLArray[4]
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

	c.JSON(200, address)
}

func (s AccountController) ContactEdit(c *gin.Context) {

	//API to Edit account Contact details

	var queueinfo queue.Queue
	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_CONTACT_EDIT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id)) + "," + URLArray[4]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id)) + "," + URLArray[3]
		queueinfo.Version = version
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
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var contact models.Contact
		contactbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(contactbyte, &contact)

		c.JSON(200, contact)
	}
}

func (s AccountController) ContactAdd(c *gin.Context) {

	//API to Add account contact details

	var queueinfo queue.Queue
	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_CONTACT_NEW"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id))
		queueinfo.Version = version
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
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.Header("X-Version", "1.0")
		c.JSON(200, returnError)
	} else {
		var contact models.Contact
		contactbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(contactbyte, &contact)

		c.JSON(200, contact)
	}
}

func (s AccountController) ContactRemove(c *gin.Context) {

	//API to Delete contact address details
	var queueinfo queue.Queue
	user, _ := c.Get("uuid")
	var authenticatedAccount = user.(*models.Account)
	version := helpers.GetVersion()
	queueinfo.Category = "ACCOUNT_CONTACT_DELETE"
	queueinfo.APIType = "DELETE"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id)) + "," + URLArray[4]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "account" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = strconv.Itoa(int(authenticatedAccount.Id)) + "," + URLArray[4]
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

	c.JSON(200, contact)
}
