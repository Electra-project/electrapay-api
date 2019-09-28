package controllers

import (
	"encoding/json"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"github.com/gin-gonic/gin"
	"strings"
)

type OrderController struct{}

func (s OrderController) New(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ORDER_NEW"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "order" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "order" {
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

	var order models.Order
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &order)

	c.JSON(200, order)

}

func (s OrderController) Get(c *gin.Context) {

	var queueinfo queue.Queue

	queueinfo.Category = "ORDER_FIND"
	queueinfo.APIType = "GET"
	version := helpers.GetVersion()
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "order" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("uuid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "order" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("uuid")
		queueinfo.Version = version
	}

	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var order models.Order
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &order)

	c.JSON(200, order)

}

func GetOrderNode(orderuuid string, accountid int64, orderreference string, version string) (orderqueryresult models.OrderQuery) {

	var queueinfo queue.Queue
	var orderquery models.OrderQuery

	orderquery.Uuid = orderuuid
	orderquery.AccountId = accountid
	orderquery.Reference = orderreference

	queueinfo.Category = "ORDER_FIND_NODE"
	queueinfo.APIType = "GET"
	queueinfo.APIURL = ""
	queueinfo.Parameters = ""
	queueinfo.Version = version
	str, err := json.Marshal(orderquery)
	queueinfo.RequestInfo = string(str)
	queueinfo, err = queue.QueueProcess(queueinfo)
	if err != nil {
		return
	}

	var orderqueryresponse models.OrderQuery
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &orderqueryresponse)

	return orderqueryresponse

}

func (s OrderController) Cancel(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()

	queueinfo.Category = "ORDER_CANCEL"
	queueinfo.APIType = "PUT"

	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "order" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "order" {
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

	var order models.Order
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &order)

	c.JSON(200, order)

}

func (s OrderController) Reverse(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()

	queueinfo.Category = "ORDER_REVERSE"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "order" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "order" {
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

	var order models.Order
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &order)

	c.JSON(200, order)

}

func (s OrderController) PaymentCategory(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()

	queueinfo.Category = "ORDER_PAYMENTCATEGORY"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "paymentcategory" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "paymentcategory" {
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

	var paymentcategories []models.PaymentCategory
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &paymentcategories)

	c.JSON(200, paymentcategories)

}

func (s OrderController) AllowedCurrency(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()

	queueinfo.Category = "ORDER_ALLOWEDCURRENCY"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "allowedcurrency" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "allowedcurrency" {
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

	var allowedcurrencies []models.AllowedCurrency
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &allowedcurrencies)

	c.JSON(200, allowedcurrencies)

}
