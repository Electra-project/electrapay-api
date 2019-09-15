package server

import (
	"github.com/Electra-project/electrapay-api/src/authenticators"
	"github.com/Electra-project/electrapay-api/src/controllers"
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/Electra-project/electrapay-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	apiauthenticator := authenticators.BasicAuth()

	router := gin.Default()
	router.Use(middlewares.CORS())
	router.Use(middlewares.ResponseHeaders())

	version := helpers.GetVersion()

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

	authController := new(controllers.AuthController)

	// Verify the email address for an account
	router.POST(version+"/auth/token", authController.Token)
	router.POST("/auth/token", authController.Token)

	// Verify the email address for an account
	router.GET(version+"/auth/verify/:email", accountController.AuthVerify)
	router.GET("/auth/verify/:email", accountController.AuthVerify)

	// Verify the email address for an account
	router.GET(version+"/auth/forgotpassword", accountController.ForgotPassword)
	router.GET("/auth/forgotpassword", accountController.ForgotPassword)

	// Set the password using the authorisation code
	router.POST(version+"/auth/setpassword", accountController.SetPassword)
	router.POST("/auth/setpassword", accountController.SetPassword)

	// register a new account - this will send an email with the authorisation code
	router.POST(version+"/account/register", accountController.Register)
	router.POST("/account/register", accountController.Register)

	/**
	 * authenticated routes
	 */
	auth := router.Group("/")
	auth.Use(authController.AuthenticationRequired)
	{
		auth.GET("/"+version+"/account/:accountid", accountController.Get)
		auth.GET("/account/:accountid", accountController.Get)
		auth.PUT("/"+version+"/account/:accountid", accountController.Edit)
		auth.PUT("/account/:accountid", accountController.Edit)
		auth.POST("/"+version+"/account/close/:accountid", accountController.Close)
		auth.POST("/account/close/:accountid", accountController.Close)
		auth.POST("/"+version+"/account/apikey/:accountid", accountController.ApiKey)
		auth.POST("/account/apikey/:accountid", accountController.ApiKey)
		auth.POST("/"+version+"/account/suspend/:accountid", accountController.Suspend)
		auth.POST("/account/suspend/:accountid", accountController.Suspend)

		auth.PUT("/account/:accountid/address/:addressid", accountController.AddressEdit)
		auth.PUT("/"+version+"/account/:accountid/address/:addressid", accountController.AddressEdit)
		//auth.POST(version+"/account/:accountid/addressnew/", accountController.AddressAdd)
		//auth.POST("/account/:accountid/address", accountController.AddressAdd)
		auth.DELETE("/"+version+"/account/:accountid/address/:addressid", accountController.AddressRemove)
		auth.DELETE("/account/:accountid/address/:addressid", accountController.AddressRemove)
		auth.PUT("/"+version+"/account/:accountid/contact/:contactid", accountController.ContactEdit)
		auth.PUT("/account/:accountid/contact/:contactid", accountController.ContactEdit)
		//auth.POST(version+"/account/:accountid/contact/new", accountController.ContactAdd)
		//auth.POST("/account/:accountid/contact/new", accountController.ContactAdd)
		auth.DELETE("/"+version+"/account/:accountid/contact/:contactid", accountController.ContactRemove)
		auth.DELETE("/account/:accountid/contact/:contactid", accountController.ContactRemove)
	}

	authapi := router.Group("/")
	authapi.Use(apiauthenticator)
	{

		orderController := new(controllers.OrderController)
		authapi.POST("/"+version+"/order", orderController.New)
		authapi.POST("/order", orderController.New)
		authapi.GET("/"+version+"/order/:uuid", orderController.Get)
		authapi.GET("/order/:uuid/", orderController.Get)
		authapi.POST("/"+version+"/order/:uuid/cancel", orderController.Cancel)
		authapi.POST("/order/:uuid/cancel", orderController.Cancel)
		authapi.POST("/"+version+"/order/:uuid/reverse", orderController.Reverse)
		authapi.POST("/order/:uuid/reverse", orderController.Reverse)
		authapi.GET("/paymentcategory/:accountid", orderController.PaymentCategory)
		authapi.GET("/"+version+"/paymentcategory/:accountid", orderController.PaymentCategory)
		authapi.GET("/allowedcurrency/:accountid", orderController.AllowedCurrency)
		authapi.GET("/"+version+"/allowedcurrency/:accountid", orderController.AllowedCurrency)
	}

	codeController := new(controllers.CodeController)
	router.GET("/"+version+"/accounttype", codeController.GetAccountType)
	router.GET("accounttype/", codeController.GetAccountType)
	router.GET("/"+version+"/addresstype/", codeController.GetAddressType)
	router.GET("addresstype/", codeController.GetAddressType)
	router.GET("/"+version+"/contacttype/", codeController.GetContactType)
	router.GET("contacttype/", codeController.GetContactType)
	router.GET("/"+version+"/currencytype/", codeController.GetCurrencyType)
	router.GET("currencytype/", codeController.GetCurrencyType)
	router.GET("plugintype/", codeController.GetPluginType)
	router.GET("/"+version+"/plugintype/", codeController.GetPluginType)
	router.GET("/"+version+"/currency/", codeController.GetCurrency)
	router.GET("currency/", codeController.GetCurrency)
	router.GET("/"+version+"/language/", codeController.GetLanguage)
	router.GET("language/", codeController.GetLanguage)
	router.GET("/"+version+"/country/", codeController.GetCountry)
	router.GET("country/", codeController.GetCountry)
	router.GET("/"+version+"/timezone/", codeController.GetTimeZone)
	router.GET("timezone/", codeController.GetTimeZone)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
