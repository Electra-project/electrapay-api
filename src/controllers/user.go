package controllers

import "github.com/gin-gonic/gin"
import (
	"encoding/json"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type UserController struct{}

func (s UserController) UserAuthenticationRequired(c *gin.Context) {
	t, err := extractToken(c)
	var response models.Error

	if err != nil {
		response.ResponseCode = "AUTH200"
		response.ResponseDescription = "Invalid Token"
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	mySigningKey := []byte(os.Getenv("JWTSECRET"))
	accessclaims := AccessClaims{}

	token, err := jwt.ParseWithClaims(t, &accessclaims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		response.ResponseCode = "AUTH201"
		response.ResponseDescription = "Invalid Token"
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*AccessClaims)

	if ok && token.Valid {
		email := c.Param("email")
		if email != claims.Subject {
			response.ResponseCode = "AUTH202"
			response.ResponseDescription = "Invalid Account Identified"
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
	}
	return
}

func (s UserController) Get(c *gin.Context) {
	//API to retrieve User information

	var queueinfo queue.Queue
	queueinfo.Category = "USER_FETCH"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	version := helpers.GetVersion()

	if URLArray[1] != "user" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("email")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "user" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("email")
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)

	if queueinfo.ResponseCode != "00" {
		returnError := models.Error{}
		returnError.ResponseCode = queueinfo.ResponseCode
		returnError.ResponseDescription = queueinfo.ResponseDescription
		c.JSON(400, returnError)
	} else {
		var user models.UserInfo
		userbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(userbyte, &user)

		if user.ResponseCode != "00" {
			c.JSON(400, user)
		} else {
			c.JSON(200, user)
		}
	}
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

}

func (s UserController) Edit(c *gin.Context) {
	//API to edit user information

	var queueinfo queue.Queue
	queueinfo.Category = "USER_EDIT"
	queueinfo.APIType = "PUT"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	version := helpers.GetVersion()

	if URLArray[1] != "user" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("email")
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "user" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = c.Param("email")
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
		var user models.UserInfo
		userbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(userbyte, &user)

		if user.ResponseCode != "00" {
			c.JSON(400, user)
		} else {
			c.JSON(200, user)
		}
	}
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

}
