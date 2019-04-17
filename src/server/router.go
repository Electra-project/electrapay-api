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
	router.POST("/v1/account/:uuid/register/", accountController.Register)
	router.POST("/account/:uuid/register/", accountController.Register)

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

		auth.GET("/v1/account/:uuid", accountController.Get)
		auth.GET("/account/:uuid", accountController.Get)
		auth.PUT("/v1/account/:uuid", accountController.Edit)
		auth.PUT("/account/:uuid", accountController.Edit)
		auth.PUT("/v1/account/:uuid/close", accountController.Close)
		auth.PUT("/account/:uuid/close", accountController.Close)
		auth.PUT("/v1/account/:uuid/address/:addressid", accountController.AddressEdit)
		auth.PUT("/account/:uuid/address/:addressid", accountController.AddressEdit)
		auth.PUT("/v1/account/:uuid/address", accountController.AddressAdd)
		auth.PUT("/account/:uuid/address", accountController.AddressAdd)
		auth.DELETE("/v1/account/:uuid/address/:addressid", accountController.AddressRemove)
		auth.DELETE("/account/:uuid/address/:addressid", accountController.AddressRemove)
		auth.PUT("/v1/account/:uuid/contact/:contactid", accountController.ContactEdit)
		auth.PUT("/account/:uuid/contact/:contactid", accountController.ContactEdit)
		auth.PUT("/v1/account/:uuid/contact/", accountController.ContactAdd)
		auth.PUT("/account/:uuid/contact/", accountController.ContactAdd)
		auth.DELETE("/v1/account/:uuid/contact/:contactid", accountController.ContactRemove)
		auth.DELETE("/account/:uuid/contact/:contactid", accountController.ContactRemove)

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

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
