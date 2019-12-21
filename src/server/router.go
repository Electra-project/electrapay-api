package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ruannelloyd/electrapay-api/src/authenticators"
	"github.com/ruannelloyd/electrapay-api/src/controllers"
	"github.com/ruannelloyd/electrapay-api/src/helpers"
	"github.com/ruannelloyd/electrapay-api/src/middlewares"
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
	userController := new(controllers.UserController)

	// Create a login token
	router.POST(version+"/auth/token", authController.Token)
	router.POST("/auth/token", authController.Token)

	// Send email for forgot password
	router.POST(version+"/auth/forgotpassword/:email", authController.ForgotPassword)
	router.POST("/auth/forgotpassword/:email", authController.ForgotPassword)

	// Set the password using the jwt token sent with expiry of 24hour
	router.POST(version+"/auth/setpassword", authController.SetPassword)
	router.POST("/auth/setpassword", authController.SetPassword)

	// register a new account - this will send an email with the authorisation code
	router.POST(version+"/account/register", accountController.Register)
	router.POST("/account/register", accountController.Register)

	/**
	 * authenticated routes
	 */

	authUser := router.Group("/")
	authUser.Use(userController.UserAuthenticationRequired)
	{
		authUser.GET("/"+version+"/user/:email", userController.Get)
		authUser.GET("/user/:email", userController.Get)
		authUser.GET("/"+version+"/user/:email/avatar", userController.GetAvatar)
		authUser.GET("/user/:email/avatar", userController.GetAvatar)
		authUser.PUT("/"+version+"/user/:email/avatar", userController.EditAvatar)
		authUser.PUT("/user/:email/avatar", userController.EditAvatar)
	}

	auth := router.Group("/")
	auth.Use(authController.AccountAuthenticationRequired)
	{
		auth.GET("/"+version+"/account/details/:accountid", accountController.GetAccount)
		auth.GET("/account/details/:accountid", accountController.GetAccount)

		auth.GET("/"+version+"/account/logo/:accountid", accountController.GetAccountLogo)
		auth.GET("/account/logo/:accountid", accountController.GetAccountLogo)
		auth.PUT("/"+version+"/account/logo/:accountid", accountController.EditAccountLogo)
		auth.PUT("/account/logo/:accountid", accountController.EditAccountLogo)

		auth.GET("/"+version+"/account/personalinformation/:accountid", accountController.GetPersonalInformation)
		auth.GET("/account/personalinformation/:accountid", accountController.GetPersonalInformation)
		auth.PUT("/"+version+"/account/personalinformation/:accountid", accountController.EditPersonalInformation)
		auth.PUT("/account/personalinformation/:accountid", accountController.EditPersonalInformation)

		auth.GET("/"+version+"/account/paymentdetails/:accountid", accountController.GetPaymentDetails)
		auth.GET("/account/paymentdetails/:accountid", accountController.GetPaymentDetails)
		auth.PUT("/"+version+"/account/paymentdetails/:accountid", accountController.EditPaymentDetails)
		auth.PUT("/account/paymentdetails/:accountid", accountController.EditPaymentDetails)

		auth.GET("/"+version+"/account/orgdetails/:accountid", accountController.GetOrganizationDetails)
		auth.GET("/account/orgdetails/:accountid", accountController.GetOrganizationDetails)
		auth.PUT("/"+version+"/account/orgdetails/:accountid", accountController.EditOrganizationDetails)
		auth.PUT("/account/orgdetails/:accountid", accountController.EditOrganizationDetails)

		auth.POST("/"+version+"/account/close/:accountid", accountController.Close)
		auth.POST("/account/close/:accountid", accountController.Close)
		auth.POST("/"+version+"/account/apikey/:accountid", accountController.ApiKey)
		auth.POST("/account/apikey/:accountid", accountController.ApiKey)
		auth.POST("/"+version+"/account/suspend/:accountid", accountController.Suspend)
		auth.POST("/account/suspend/:accountid", accountController.Suspend)

		auth.GET("/account/address/:accountid/:addresstype", accountController.AddressFetch)
		auth.GET("/"+version+"/account/address/:accountid/:addresstype", accountController.AddressFetch)
		auth.PUT("/account/address/:accountid", accountController.AddressEdit)
		auth.PUT("/"+version+"/account/address/:accountid", accountController.AddressEdit)
		auth.POST(version+"/account/address/:accountid/", accountController.AddressAdd)
		auth.POST("/account/address/:accountid", accountController.AddressAdd)
		auth.DELETE("/"+version+"/account/address/:accountid/:addressid", accountController.AddressRemove)
		auth.DELETE("/account/address/:accountid/:addressid", accountController.AddressRemove)

		auth.GET("/account/contact/:accountid/:contacttype", accountController.ContactFetch)
		auth.GET("/"+version+"/account/contact/:accountid/:contacttype", accountController.ContactFetch)
		auth.PUT("/"+version+"/account/contact/:accountid/:contactid", accountController.ContactEdit)
		auth.PUT("/account/contact/:accountid/:contactid", accountController.ContactEdit)
		auth.POST(version+"/account/contact/:accountid", accountController.ContactAdd)
		auth.POST("/account/contact/:accountid", accountController.ContactAdd)
		auth.DELETE("/"+version+"/account/contact/:accountid/:contactid", accountController.ContactRemove)
		auth.DELETE("/account/contact/:accountid/:contactid", accountController.ContactRemove)

		auth.GET("/account/rules/:accountid", accountController.RulesFetch)
		auth.GET("/"+version+"/account/rules/:accountid", accountController.RulesFetch)
		auth.PUT("/account/rules/:accountid", accountController.RulesEdit)
		auth.PUT("/"+version+"/account/rules/:accountid", accountController.RulesEdit)

		auth.GET("/"+version+"/account/balance/:accountid", accountController.AccountBalance)
		auth.GET("/account/balance/:accountid/", accountController.AccountBalance)

		auth.GET("/"+version+"/account/ordersummary/:accountid/:frequency", accountController.OrderSummary)
		auth.GET("/account/ordersummary/:accountid/:frequency", accountController.OrderSummary)

		auth.GET("/"+version+"/account/orderlist/:accountid", accountController.OrderList)
		auth.GET("/account/orderlist/:accountid", accountController.OrderList)

		auth.GET("/"+version+"/account/orderlist/:accountid/:maxlimit", accountController.OrderListMax)
		auth.GET("/account/orderlist/:accountid/:maxlimit", accountController.OrderListMax)

		auth.GET("/"+version+"/account/activitylist/:accountid", accountController.ActivityList)
		auth.GET("/account/activitylist/:accountid", accountController.ActivityList)
	}

	authapi := router.Group("/")
	authapi.Use(apiauthenticator)
	{

		orderController := new(controllers.OrderController)
		authapi.POST("/"+version+"/order", orderController.New)
		authapi.POST("/order", orderController.New)
		authapi.GET("/"+version+"/order/:uuid", orderController.Get)
		authapi.GET("/order/:uuid", orderController.Get)
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
