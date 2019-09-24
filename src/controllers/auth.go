package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
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
	mySigningKey := []byte(os.Getenv("JWTSECRET"))

	if err != nil || grantType == "" {
		c.JSON(http.StatusBadRequest, "malformed request")
		return
	}

	switch grantType {

	case "password":
		{
			var gt models.GrantTypePassword

			if err := helpers.DecodeJson(c, &gt); err != nil {
				c.JSON(http.StatusBadRequest, "malformed request")
				return
			}

			account, err := authenticate(gt.Email, gt.Password)

			if err != nil {
				c.JSON(http.StatusUnauthorized, "unauthorized")
				return
			}
			if account.ResponseCode != "00" {
				c.JSON(http.StatusUnauthorized, "unauthorized")
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
				c.JSON(http.StatusBadRequest, "malformed request")
				return
			}

			var claims jwt.StandardClaims

			tkn, err := jwt.ParseWithClaims(gt.RefreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					c.JSON(http.StatusUnauthorized, "unauthorized")
					return
				}
				c.JSON(http.StatusBadRequest, "malformed request")
				return
			}

			if !tkn.Valid {
				c.JSON(http.StatusUnauthorized, "unauthorized ")
				return
			}

			/*if err := s.db.Model(&user).Where("id = ?", claims.Subject).Select(); err != nil {
			  if strings.Contains(err.Error(), "no rows in result set") {
			    http.Error(w, "unauthorized", http.StatusUnauthorized)
			    return
			  }
			}*/
		}

	default:
		c.JSON(http.StatusBadRequest, "malformed request")
		return
	}

	tokenRes, err := generateTokenResponse(email, accountid)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "something went wrong. we are already investigating.")
		return
	}

	c.JSON(http.StatusOK, tokenRes)

}

func generateTokenResponse(email string, accountid string) (models.GrantTypeResponse, error) {

	mySigningKey := []byte(os.Getenv("JWTSECRET"))

	tokenExp := time.Now().Add(6 * time.Minute).Unix()
	refreshTokenExp := time.Now().Add(10 * time.Minute).Unix()

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
		fmt.Println(err.Error())
		return models.GrantTypeResponse{}, err
	}

	refreshTokenString, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err.Error())
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

func (s AuthController) AuthenticationRequired(c *gin.Context) {
	t, err := extractToken(c)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnauthorized, "invalid token.")
		c.Abort()
		return
	}

	mySigningKey := []byte(os.Getenv("JWTSECRET"))
	accessclaims := AccessClaims{}

	token, err := jwt.ParseWithClaims(t, &accessclaims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnauthorized, "invalid token.")
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*AccessClaims)

	if ok && token.Valid {
		authaccount := c.Param("accountid")
		if authaccount != claims.Accountid {
			c.JSON(http.StatusUnauthorized, "Invalid Account.")
			c.Abort()
			return
		}
	}
	return
}

func extractToken(c *gin.Context) (string, error) {
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(string(reqToken), "Bearer ")
	if len(splitToken) != 2 {
		return "", errors.New("wrong auth header format")
	}
	return strings.TrimSpace(splitToken[1]), nil
}
