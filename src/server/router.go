package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Electra-project/electrapay-api/src/authenticators"
	"github.com/Electra-project/electrapay-api/src/controllers"
	"github.com/Electra-project/electrapay-api/src/middlewares"
)

func Router() *gin.Engine {

	authenticator := authenticators.Authenticator()

	router := gin.Default()
	router.Use(middlewares.CORS())
	router.Use(middlewares.ResponseHeaders())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "Electrapay API",
			"version": "1.0",
		})
	})

	/**
	 * public routes
	 */
	accountController := new(controllers.AccountController)

	// account
	router.POST("/v1/account/:accountid/register/", accountController.Register)
	router.POST("/account/:accountid/register/", accountController.Register)

	// login
	router.POST("/v1/authenticate", authenticator.LoginHandler)
	router.POST("/authenticate", authenticator.LoginHandler)

	/**
	 * authenticated routes
	 */
	auth := router.Group("/")
	auth.Use(authenticator.MiddlewareFunc())
	{
		auth.POST("/v1/account/", accountController.New)
		auth.POST("/account/", accountController.New)

		auth.GET("/v1/account/:accountid", accountController.Get)
		auth.GET("/account/:accountid", accountController.Get)
		auth.PUT("/v1/account/:accountid", accountController.Edit)
		auth.PUT("/account/:accountid", accountController.Edit)
		auth.PUT("/v1/account/:accountid/close", accountController.Close)
		auth.PUT("/account/:accountid/close", accountController.Close)
		auth.PUT("/v1/account/:accountid/address/:addressid", accountController.AddressEdit)
		auth.PUT("/account/:accountid/address/:addressid", accountController.AddressEdit)
		auth.PUT("/v1/account/:accountid/address", accountController.AddressAdd)
		auth.PUT("/account/:accountid/address", accountController.AddressAdd)
		auth.DELETE("/v1/account/:accountid/address/:addressid", accountController.AddressRemove)
		auth.DELETE("/account/:accountid/address/:addressid", accountController.AddressRemove)
		auth.PUT("/v1/account/:accountid/contact/:contactid", accountController.ContactEdit)
		auth.PUT("/account/:accountid/contact/:contactid", accountController.ContactEdit)
		auth.PUT("/v1/account/:accountid/contact/", accountController.ContactAdd)
		auth.PUT("/account/:accountid/contact/", accountController.ContactAdd)
		auth.DELETE("/v1/account/:accountid/contact/:contactid", accountController.ContactRemove)
		auth.DELETE("/account/:accountid/contact/:contactid", accountController.ContactRemove)

		orderController := new(controllers.OrderController)
		auth.POST("/v1/order/", orderController.New)
		auth.POST("/order/", orderController.New)
		auth.GET("/v1/order/:uuid", orderController.Get)
		auth.GET("/order/:uuid/", orderController.Get)
		auth.POST("/v1/order/:uuid/cancel", orderController.Cancel)
		auth.POST("/order/:uuid/cancel", orderController.Cancel)
		auth.POST("/v1/order/:uuid/reverse", orderController.Reverse)
		auth.POST("/order/:uuid/reverse", orderController.Reverse)
	}

	codeController := new(controllers.CodeController)
	router.GET("/v1/accounttype/", codeController.GetAccountType)
	router.GET("accounttype/", codeController.GetAccountType)
	router.GET("/v1/addresstype/", codeController.GetAddressType)
	router.GET("addresstype/", codeController.GetAddressType)
	router.GET("/v1/contacttype/", codeController.GetContactType)
	router.GET("contacttype/", codeController.GetContactType)
	router.GET("/v1/currencytype/", codeController.GetCurrencyType)
	router.GET("currencytype/", codeController.GetCurrencyType)
	router.GET("plugintype/", codeController.GetPluginType)
	router.GET("/v1/plugintype/", codeController.GetPluginType)
	router.GET("/v1/currency/", codeController.GetCurrency)
	router.GET("currency/", codeController.GetCurrency)
	router.GET("/v1/language/", codeController.GetLanguage)
	router.GET("language/", codeController.GetLanguage)
	router.GET("/v1/country/", codeController.GetCountry)
	router.GET("country/", codeController.GetCountry)
	router.GET("/v1/timezone/", codeController.GetTimeZone)
	router.GET("timezone/", codeController.GetTimeZone)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
