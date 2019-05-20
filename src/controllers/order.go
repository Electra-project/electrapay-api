package controllers

import (
	"encoding/json"
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

type OrderController struct{}

func (s OrderController) New(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ORDER_NEW"
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
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[2]
		queueinfo.Version = URLArray[3]
	}
	if len(URLArray) == 3 {
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

func (s OrderController) Cancel(c *gin.Context) {

	var queueinfo queue.Queue
	version := helpers.GetVersion()

	queueinfo.Category = "ORDER_CANCEL"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[2]
		queueinfo.Version = URLArray[3]
	}
	if len(URLArray) == 3 {
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
	if len(URLArray) == 4 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = URLArray[2]
		queueinfo.Version = URLArray[3]
	}
	if len(URLArray) == 3 {
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
