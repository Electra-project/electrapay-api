package authenticators

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			c.AbortWithError(401, errors.New("Authorization is missing"))
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || !authenticateUser(pair[0], pair[1]) {
			c.AbortWithError(401, errors.New("Validation Failed"))
		}
		c.Next()
	}
}

func authenticateUser(username, password string) bool {
	var queueinfo queue.Queue

	version := helpers.GetVersion()
	queueinfo.Category = "API_AUTHENTICATE"
	queueinfo.APIType = "POST"
	queueinfo.Parameters = ""
	queueinfo.Version = version
	queueinfo.RequestInfo = "{\"id\": \"" + username + "\", \"apikey\": \"" + password + "\"}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		return false
	}
	var account models.Account
	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)
	if strconv.FormatInt(account.Id, 10) == username {
		return true
	} else {
		return false
	}
}
