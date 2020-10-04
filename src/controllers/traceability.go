package controllers

import "github.com/gin-gonic/gin"
import (
	"encoding/json"
	"github.com/ruannelloyd/electrapay-api/src/helpers"
	"github.com/ruannelloyd/electrapay-api/src/models"
	"github.com/ruannelloyd/electrapay-api/src/queue"
	"io/ioutil"
	"strings"
)

type TraceabilityController struct{}

func (s TraceabilityController) SendContent(c *gin.Context) {

	//API to Edit account details
	version := helpers.GetVersion()
	var queueinfo queue.Queue
	queueinfo.Category = "TRACEABILITY_SENDCONTENT"
	queueinfo.QueueCategory = "traceability_queue"
	queueinfo.APIType = "PUT"
	t, err := extractToken(c)
	queueinfo.Token = t
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "traceability" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("accountid")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "traceability" {
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
		var traceability models.Traceability
		traceabilitybyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(traceabilitybyte, &traceability)

		if traceability.ResponseCode != "00" {
			c.JSON(400, traceability)
		} else {
			c.JSON(200, traceability)
		}
	}
}
