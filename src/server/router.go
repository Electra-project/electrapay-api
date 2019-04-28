package server

import (
	"github.com/Electra-project/electrapay-api/src/controllers"
	"github.com/Electra-project/electrapay-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

// Router binds the routes to the controllers.
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "Electrapay API",
			"version": "1.0",
		})
	})

	accountController := new(controllers.AccountController)
	//public
	router.POST("/v1/account/", accountController.New)
	router.POST("/account/", accountController.New)
	router.POST("/v1/account/:accountid/register/", accountController.Register)
	router.POST("/account/:accountid/register/", accountController.Register)

	router.GET("/v1/account/:accountid", accountController.Get)
	router.GET("/account/:accountid", accountController.Get)
	router.PUT("/v1/account/:accountid", accountController.Edit)
	router.PUT("/account/:accountid", accountController.Edit)
	router.PUT("/v1/account/:accountid/close", accountController.Close)
	router.PUT("/account/:accountid/close", accountController.Close)
	router.PUT("/v1/account/:accountid/address/:addressid", accountController.AddressEdit)
	router.PUT("/account/:accountid/address/:addressid", accountController.AddressEdit)
	router.PUT("/v1/account/:accountid/address", accountController.AddressAdd)
	router.PUT("/account/:accountid/address", accountController.AddressAdd)
	router.DELETE("/v1/account/:accountid/address/:addressid", accountController.AddressRemove)
	router.DELETE("/account/:accountid/address/:addressid", accountController.AddressRemove)
	router.PUT("/v1/account/:accountid/contact/:contactid", accountController.ContactEdit)
	router.PUT("/account/:accountid/contact/:contactid", accountController.ContactEdit)
	router.PUT("/v1/account/:accountid/contact/", accountController.ContactAdd)
	router.PUT("/account/:accountid/contact/", accountController.ContactAdd)
	router.DELETE("/v1/account/:accountid/contact/:contactid", accountController.ContactRemove)
	router.DELETE("/account/:accountid/contact/:contactid", accountController.ContactRemove)

	orderController := new(controllers.OrderController)
	router.POST("/v1/order/", orderController.New)
	router.POST("/order/", orderController.New)
	router.GET("/v1/order/:uuid", orderController.Get)
	router.GET("/order/:uuid/", orderController.Get)
	router.POST("/v1/order/:uuid/cancel", orderController.Cancel)
	router.POST("/order/:uuid/cancel", orderController.Cancel)
	router.POST("/v1/order/:uuid/reverse", orderController.Reverse)
	router.POST("/order/:uuid/reverse", orderController.Reverse)

	codeController := new(controllers.CodeController)
	router.GET("/v1/accounttype/", codeController.GetAccountType)
	router.GET("accounttype/", codeController.GetAccountType)
	router.GET("/v1/addresstype/", codeController.GetAddressType)
	router.GET("addresstype/", codeController.GetAddressType)
	router.GET("/v1/contacttype/", codeController.GetContactType)
	router.GET("contacttype/", codeController.GetContactType)
	router.GET("/v1/currencytype/", codeController.GetCurrencyType)
	router.GET("plugintype/", codeController.GetCurrencyType)
	router.GET("/v1/plugintype/", codeController.GetPluginType)
	router.GET("currencytype/", codeController.GetPluginType)
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
