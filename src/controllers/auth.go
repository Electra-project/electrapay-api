package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type AuthController struct{}

func (s AuthController) Token(c *gin.Context) {

	grantType, err := checkGrantType(c)
	var email string

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
			email = account.ContactEmail
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

			tkn, err := jwt.ParseWithClaims(gt.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return os.Getenv("JWTSECRET"), nil
			})

			fmt.Println(err)
			fmt.Println(tkn)

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

	tokenRes, err := generateTokenResponse(email)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "something went wrong. we are already investigating.")
		return
	}

	c.JSON(http.StatusOK, tokenRes)

}

func generateTokenResponse(email string) (models.GrantTypeResponse, error) {

	tokenExp := time.Now().Add(6 * time.Minute).Unix()
	refreshTokenExp := time.Now().Add(10 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"exp": tokenExp,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"exp": refreshTokenExp,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		fmt.Println(err.Error())
		return models.GrantTypeResponse{}, err
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(email))
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
