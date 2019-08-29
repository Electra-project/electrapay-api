package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/queue"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type OrderNew struct {
	AccountId        int64           `json:"accountid"`
	ShortDescription string          `json:"shortdescription"`
	LongDescription  string          `json:"longdescription"`
	Reference        string          `json:"reference"`
	Paymentcategory  string          `json:"paymentcategory"`
	OrderCurrency    string          `json:"ordercurrency"`
	OrderAmount      decimal.Decimal `json:"orderamount"`
}

type OrderQuery struct {
	AccountId int64  `json: "accountid"`
	Reference string `json: "reference"`
	Uuid      string `json: "uuid"`
	NodeId    int64  `json: "nodeid"`
}

type Order struct {
	OrderId                  int64           `json:"id"`
	Uuid                     string          `json:"uuid"`
	AccountId                int64           `json:"accountid"`
	ShortDescription         string          `json:"shortdescription"`
	LongDescription          string          `json:"longdescription"`
	Reference                string          `json:"reference"`
	Paymentcategory          string          `json:"paymentcategory"`
	OrderCurrency            string          `json:"ordercurrency"`
	OrderAmount              decimal.Decimal `json:"orderamount"`
	QuoteCurrency            string          `json:"quotecurrency"`
	QuoteAmount              decimal.Decimal `json:"quoteamount"`
	QuoteTranFee             decimal.Decimal `json:"quotetranfee"`
	QuoteFeeAmount           decimal.Decimal `json:"quotetranfeeamount"`
	QRCode                   string          `json:"qrcode"`
	OrderToken               string          `json:"ordertoken"`
	WalletAddress            string          `json:"walletaddress"`
	OrderReceivedDate        time.Time       `json:"orderreceivedate"`
	OrderReceivedTransaction string          `json:"orderreceivetransaction"`
	OrderFinalTransaction    string          `json:"orderfinaltransaction"`
	OrderReversalTransaction string          `json:"orderreversaltransaction"`
	OrderQuoteSubmittedDate  time.Time       `json:"orderquotesubmitteddate"`
	OrderReceivedPaymentDate time.Time       `json:"orderreceivedpaymentdate"`
	OrderFinalPaymentDate    time.Time       `json:"orderfinalpaymentdate"`
	OrderStatus              string          `json:"orderstatus"`
	ResponseCode             string          `json:"responsecode"`
	ResponseDescription      string          `json:"responsedescription"`
}

type PaymentCategory struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type AllowedCurrency struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type OrderController struct{}

func (s OrderController) New(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ORDER_NEW"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] == "order" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] != "order" {
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

	var order Order
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &order)

	c.Header("X-Version", "1.0")
	c.JSON(200, order)

}

func (s OrderController) Get(c *gin.Context) {

	var queueinfo queue.Queue

	queueinfo.Category = "ORDER_FIND"
	queueinfo.APIType = "GET"
	version := helpers.GetVersion()
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] == "order" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] != "order" {
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

	var order Order
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &order)

	c.Header("X-Version", "1.0")
	c.JSON(200, order)

}

func GetOrderNode(orderquery, version string) (orderqueryresult OrderQuery) {

	var queueinfo queue.Queue

	queueinfo.Category = "ORDER_FIND_NODE"
	queueinfo.APIType = "GET"
	queueinfo.APIURL = ""
	queueinfo.Parameters = ""
	queueinfo.Version = version
	queueinfo.RequestInfo = orderquery
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		return
	}

	var orderqueryresponse OrderQuery
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
	if URLArray[3] != "cancel" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[3] == "cancel" {
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

	var order Order
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &order)

	c.Header("X-Version", "1.0")
	c.JSON(200, order)

}

func (s OrderController) Reverse(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()

	queueinfo.Category = "ORDER_REVERSE"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[3] != "reverse" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[3]
		queueinfo.Version = URLArray[1]
	}
	if URLArray[3] == "reverse" {
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

	var order Order
	orderbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(orderbyte, &order)

	c.Header("X-Version", "1.0")
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

	var paymentcategories []PaymentCategory
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &paymentcategories)

	c.Header("X-Version", "1.0")
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

	var allowedcurrencies []AllowedCurrency
	queueResult := queueinfo.ResponseInfo
	fmt.Println(queueResult)
	json.Unmarshal([]byte(queueResult), &allowedcurrencies)

	c.Header("X-Version", "1.0")
	c.JSON(200, allowedcurrencies)

}
