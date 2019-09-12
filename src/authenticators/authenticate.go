package authenticators

import (
	"encoding/json"
	"github.com/Electra-project/electrapay-api/src/models"
	"github.com/Electra-project/electrapay-api/src/queue"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var idKey = "id"
var identityKey = "email"
var nameKey = "name"
var descriptionKey = "description"
var websiteKey = "website"

func Authenticator() (middleware *jwt.GinJWTMiddleware) {
	// the jwt middleware
	//var ErrMissingLoginValues = errors.New("missing Username or Password")
	//var ErrFailedAuthentication = errors.New("incorrect Username or Password")

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("password"),
		Timeout:     time.Minute,
		MaxRefresh:  time.Minute * 5,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Account); ok {
				return jwt.MapClaims{
					idKey:          int64(v.Id),
					identityKey:    v.ContactEmail,
					nameKey:        v.Name,
					descriptionKey: v.Description,
					websiteKey:     v.Website,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.Account{
				Id:           int64(claims[idKey].(float64)),
				ContactEmail: claims[identityKey].(string),
				Name:         claims[nameKey].(string),
				Description:  claims[descriptionKey].(string),
				Website:      claims[websiteKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			var queueinfo queue.Queue
			queueinfo.Category = "ACCOUNT_AUTHENTICATE"
			queueinfo.APIType = "POST"
			queueinfo.Parameters = ""
			queueinfo.Version = "v1"
			queueinfo.RequestInfo = "{\"Email\": \"" + email + "\", \"Password\": \"" + password + "\"}"
			queueinfo, err := queue.QueueProcess(queueinfo)

			if err != nil {
				c.AbortWithError(404, err)
			}

			var account models.Account
			accountbyte := []byte(queueinfo.ResponseInfo)
			json.Unmarshal(accountbyte, &account)
			if account.ContactEmail == email {
				return &models.Account{
					Id:           account.Id,
					ContactEmail: account.ContactEmail,
					Name:         account.Name,
					Description:  account.Description,
					Website:      account.Website,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// We implement here a base authorization mechanism:
			// A given user with Email XXX can only send requests for that given Email.
			// If a user with UUID XXX sends a request for Email YYY, then the request
			// Will be rejected
			var requestEmail = c.Params.ByName("email")
			if requestEmail != "" {
				v, ok := data.(*models.Account)
				if ok && v.ContactEmail != requestEmail {
					return false
				}
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
