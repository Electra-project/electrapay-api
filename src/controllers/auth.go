package controllers

import (
	"encoding/json"
	"errors"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type AccessClaims struct {
	Accountid string `json:"accountid"`
	jwt.StandardClaims
}

type AuthController struct{}

func (s AuthController) Token(c *gin.Context) {

	grantType, err := checkGrantType(c)
	var email string
	var accountid string
	var response models.Error
	mySigningKey := []byte(os.Getenv("JWTSECRET"))

	if err != nil || grantType == "" {
		response.ResponseCode = "AUTH001"
		response.ResponseDescription = "Grant Type not Specified"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	switch grantType {

	case "password":
		{
			var gt models.GrantTypePassword

			if err := helpers.DecodeJson(c, &gt); err != nil {
				response.ResponseCode = "AUTH002"
				response.ResponseDescription = "Malformed Request"
				c.JSON(http.StatusBadRequest, response)
				return
			}

			account, err := authenticate(gt.Email, gt.Password)

			if err != nil {
				response.ResponseCode = "AUTH003"
				response.ResponseDescription = "Unable to Authorise. Please try again later"
				c.JSON(http.StatusBadRequest, response)
				return
			}
			if account.ResponseCode != "00" {
				response.ResponseCode = account.ResponseCode
				response.ResponseDescription = account.ResponseDescription
				c.JSON(http.StatusBadRequest, response)
				return

			}
			email = account.ContactEmail
			accountid = strconv.FormatInt(account.Id, 10)
		}

	case "refresh_token":
		{
			var gt models.GrantTypeRefreshToken

			err := helpers.DecodeJson(c, &gt)

			if err != nil {
				response.ResponseCode = "AUTH004"
				response.ResponseDescription = "Malformed Request"
				c.JSON(http.StatusUnauthorized, response)
				return
			}

			var claims jwt.StandardClaims

			tkn, err := jwt.ParseWithClaims(gt.RefreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					response.ResponseCode = "AUTH005"
					response.ResponseDescription = "Invalid Signature"
					c.JSON(http.StatusUnauthorized, response)
					return
				}
				response.ResponseCode = "AUTH006"
				response.ResponseDescription = "Token Expired"
				c.JSON(http.StatusUnauthorized, response)
				return
			}

			if !tkn.Valid {
				response.ResponseCode = "AUTH007"
				response.ResponseDescription = "Invalid Token"
				c.JSON(http.StatusUnauthorized, response)
				return
			}

			mySigningKey := []byte(os.Getenv("JWTSECRET"))
			accessclaims := AccessClaims{}

			t, err := extractToken(c)

			at, err := jwt.ParseWithClaims(t, &accessclaims, func(token *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})
			aclaims, _ := at.Claims.(*AccessClaims)
			accountid = aclaims.Accountid
			email = aclaims.Subject

		}

	default:
		response.ResponseCode = "AUTH008"
		response.ResponseDescription = "Malformed Request"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	tokenRes, err := generateTokenResponse(email, accountid)
	if err != nil {
		response.ResponseCode = "AUTH010"
		response.ResponseDescription = "Failed to Generate a token"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	c.JSON(http.StatusOK, tokenRes)

}

func generateTokenResponse(email string, accountid string) (models.GrantTypeResponse, error) {

	mySigningKey := []byte(os.Getenv("JWTSECRET"))

	tokenExp := time.Now().Add(20 * time.Minute).Unix()
	refreshTokenExp := time.Now().Add(24 * time.Hour).Unix()

	accessclaims := AccessClaims{accountid,
		jwt.StandardClaims{
			Subject:   email,
			ExpiresAt: tokenExp,
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessclaims)

	rClaims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: refreshTokenExp,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return models.GrantTypeResponse{}, err
	}

	refreshTokenString, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		return models.GrantTypeResponse{}, err
	}

	return models.GrantTypeResponse{TokenType: "Bearer", ExpiresIn: tokenExp, AccessToken: tokenString, RefreshToken: refreshTokenString}, nil
}

func checkGrantType(c *gin.Context) (string, error) {
	jsonMap := make(map[string]string)
	if err := helpers.DecodeJson(c, &jsonMap); err != nil {
		return "", err
	}
	return jsonMap["grant_type"], nil
}

func authenticate(email string, password string) (models.Account, error) {
	var account models.Account

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNT_AUTHENTICATE"
	queueinfo.APIType = "POST"
	queueinfo.Parameters = ""
	queueinfo.Version = "v1"
	queueinfo.RequestInfo = "{\"Email\": \"" + email + "\", \"Password\": \"" + password + "\"}"
	queueinfo, err := queue.QueueProcess(queueinfo)

	if err != nil {
		return account, err
	}

	accountbyte := []byte(queueinfo.ResponseInfo)
	json.Unmarshal(accountbyte, &account)
	if account.ContactEmail == email {
		return account, nil
	}
	return account, nil

}

func (s AuthController) AccountAuthenticationRequired(c *gin.Context) {
	t, err := extractToken(c)
	var response models.Error

	if err != nil {
		response.ResponseCode = "AUTH100"
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
		response.ResponseCode = "AUTH101"
		response.ResponseDescription = "Invalid Token"
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*AccessClaims)

	if ok && token.Valid {
		authaccount := c.Param("accountid")
		if authaccount != claims.Accountid {
			response.ResponseCode = "AUTH103"
			response.ResponseDescription = "Invalid Account Identified"
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}
	}
	return
}

func extractToken(c *gin.Context) (string, error) {
	var response models.Error
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(string(reqToken), "Bearer ")
	if len(splitToken) != 2 {
		response.ResponseCode = "AUTH300"
		response.ResponseDescription = "Invalid Header"
		c.JSON(http.StatusBadRequest, response)

		return "", errors.New("Invalid Header")
	}
	return strings.TrimSpace(splitToken[1]), nil
}

func (s AuthController) ForgotPassword(c *gin.Context) {
	//API to set the user password
	var queueinfo queue.Queue
	queueinfo.Category = "AUTH_FORGOTPASSWORD"
	queueinfo.APIType = "POST"
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
		c.JSON(400, returnError)
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

func (s AuthController) SetPassword(c *gin.Context) {
	//API to set the user password
	version := helpers.GetVersion()

	var queueinfo queue.Queue
	queueinfo.Category = "AUTH_SETPASSWORD"
	queueinfo.APIType = "POST"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if URLArray[1] != "auth" {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if URLArray[1] == "auth" {
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
		c.JSON(400, returnError)
	} else {
		var user models.UserVerify
		userbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(userbyte, &user)

		c.JSON(200, user)
	}
}

func (s AuthController) AuthVerify(c *gin.Context) {
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
		c.JSON(400, returnError)
	} else {
		var user models.UserVerify
		userbyte := []byte(queueinfo.ResponseInfo)
		json.Unmarshal(userbyte, &user)

		c.JSON(200, user)
	}
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

}
