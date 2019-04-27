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

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
