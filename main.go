package main

import (
	"log"
	"os"

	"github.com/DieFleitas/ecommerce-akhil/controllers"
	"github.com/DieFleitas/ecommerce-akhil/database"
	"github.com/DieFleitas/ecommerce-akhil/middleware"
	"github.com/DieFleitas/ecommerce-akhil/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
