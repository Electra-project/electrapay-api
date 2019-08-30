package models

import (
	"github.com/shopspring/decimal"
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
