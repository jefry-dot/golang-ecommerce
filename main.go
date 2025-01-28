package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jefry-dot/golang-ecommer/controllers"
	"github.com/jefry-dot/golang-ecommer/database"
	"github.com/jefry-dot/golang-ecommer/models"
	"github.com/jefry-dot/golang-ecommer/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserData(database.Client, "User"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authenticate(app))

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}