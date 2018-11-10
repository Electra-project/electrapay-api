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
	router.POST("/v1/account/:uuid/register/", accountController.Register)
	router.POST("/account/:uuid/register/", accountController.Register)

	router.GET("/v1/account/:uuid", accountController.Get)
	router.GET("/account/:uuid", accountController.Get)
	router.PUT("/v1/account/:uuid", accountController.Edit)
	router.PUT("/account/:uuid", accountController.Edit)
	router.PUT("/v1/account/:uuid/close", accountController.Close)
	router.PUT("/account/:uuid/close", accountController.Close)
	router.PUT("/v1/account/:uuid/address/:addressid", accountController.AddressEdit)
	router.PUT("/account/:uuid/address/:addressid", accountController.AddressEdit)
	router.PUT("/v1/account/:uuid/address", accountController.AddressAdd)
	router.PUT("/account/:uuid/address", accountController.AddressAdd)
	router.DELETE("/v1/account/:uuid/address/:addressid", accountController.AddressRemove)
	router.DELETE("/account/:uuid/address/:addressid", accountController.AddressRemove)
	router.PUT("/v1/account/:uuid/contact/:contactid", accountController.ContactEdit)
	router.PUT("/account/:uuid/contact/:contactid", accountController.ContactEdit)
	router.PUT("/v1/account/:uuid/contact/", accountController.ContactAdd)
	router.PUT("/account/:uuid/contact/", accountController.ContactAdd)
	router.DELETE("/v1/account/:uuid/contact/:contactid", accountController.ContactRemove)
	router.DELETE("/account/:uuid/contact/:contactid", accountController.ContactRemove)

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
