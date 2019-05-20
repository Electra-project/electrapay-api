package server

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"github.com/Electra-project/electrapay-api/src/authenticators"
	"github.com/Electra-project/electrapay-api/src/controllers"
	"github.com/Electra-project/electrapay-api/src/middlewares"
	"os"
	"strconv"
)

func Router() *gin.Engine {

	authenticator := authenticators.Authenticator()

	router := gin.Default()
	router.Use(middlewares.CORS())
	router.Use(middlewares.ResponseHeaders())

	var version = os.Getenv("VERSION")
	v, err := strconv.Atoi(version)
	if err != nil {
		v = 1
	}
	vdir := fmt.Sprint("/v", v, "/")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "Electrapay API",
			"version": version,
		})
	})

	/**
	 * public routes
	 */
	accountController := new(controllers.AccountController)

	// register a new account - this will send an email with the login code
	router.POST(vdir+"/account/register/", accountController.Register)
	router.POST("/account/register/", accountController.Register)

	// login to create a JWT token
	router.POST(vdir+"/account/authenticate", authenticator.LoginHandler)
	router.POST("/account/authenticate", authenticator.LoginHandler)

	/**
	 * authenticated routes
	 */
	auth := router.Group("/")
	auth.Use(authenticator.MiddlewareFunc())
	{
		auth.GET(vdir+"/account/:accountid", accountController.Get)
		auth.GET("/account/:accountid", accountController.Get)
		auth.PUT(vdir+"/account/:accountid", accountController.Edit)
		auth.PUT("/account/:accountid", accountController.Edit)
		auth.POST(vdir+"/account/:accountid/close", accountController.Close)
		auth.POST("/account/:accountid/close", accountController.Close)
		auth.POST(vdir+"/account/:accountid/apikey", accountController.ApiKey)
		auth.POST("/account/:accountid/apikey", accountController.ApiKey)
		auth.POST(vdir+"/account/:accountid/suspend", accountController.Suspend)
		auth.POST("/account/:accountid/suspend", accountController.Suspend)
		auth.PUT(vdir+"/account/:accountid/address/:addressid", accountController.AddressEdit)

		auth.PUT("/account/:accountid/address/:addressid", accountController.AddressEdit)
		auth.POST(vdir+"/account/:accountid/address", accountController.AddressAdd)
		auth.POST("/account/:accountid/address", accountController.AddressAdd)
		auth.DELETE(vdir+"/account/:accountid/address/:addressid", accountController.AddressRemove)
		auth.DELETE("/account/:accountid/address/:addressid", accountController.AddressRemove)
		auth.PUT(vdir+"/account/:accountid/contact/:contactid", accountController.ContactEdit)
		auth.PUT("/account/:accountid/contact/:contactid", accountController.ContactEdit)
		auth.POST(vdir+"/account/:accountid/contact/", accountController.ContactAdd)
		auth.POST("/account/:accountid/contact/", accountController.ContactAdd)
		auth.DELETE(vdir+"/account/:accountid/contact/:contactid", accountController.ContactRemove)
		auth.DELETE("/account/:accountid/contact/:contactid", accountController.ContactRemove)
	}

	orderController := new(controllers.OrderController)
	router.POST(vdir+"/order/", orderController.New)
	router.POST("/order/", orderController.New)
	router.GET(vdir+"/order/:uuid", orderController.Get)
	router.GET("/order/:uuid/", orderController.Get)
	router.POST(vdir+"/order/:uuid/cancel", orderController.Cancel)
	router.POST("/order/:uuid/cancel", orderController.Cancel)
	router.POST(vdir+"/order/:uuid/reverse", orderController.Reverse)
	router.POST("/order/:uuid/reverse", orderController.Reverse)

	codeController := new(controllers.CodeController)
	router.GET(vdir+"/accounttype/", codeController.GetAccountType)
	router.GET("accounttype/", codeController.GetAccountType)
	router.GET(vdir+"/addresstype/", codeController.GetAddressType)
	router.GET("addresstype/", codeController.GetAddressType)
	router.GET(vdir+"/contacttype/", codeController.GetContactType)
	router.GET("contacttype/", codeController.GetContactType)
	router.GET(vdir+"/currencytype/", codeController.GetCurrencyType)
	router.GET("currencytype/", codeController.GetCurrencyType)
	router.GET("plugintype/", codeController.GetPluginType)
	router.GET(vdir+"/plugintype/", codeController.GetPluginType)
	router.GET(vdir+"/currency/", codeController.GetCurrency)
	router.GET("currency/", codeController.GetCurrency)
	router.GET(vdir+"/language/", codeController.GetLanguage)
	router.GET("language/", codeController.GetLanguage)
	router.GET(vdir+"/country/", codeController.GetCountry)
	router.GET("country/", codeController.GetCountry)
	router.GET(vdir+"/timezone/", codeController.GetTimeZone)
	router.GET("timezone/", codeController.GetTimeZone)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
