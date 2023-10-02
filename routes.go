package main

import (
	"go-htmx-dashboard/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App, handlerCfg *handlers.HandlerConfig) {
	app.Use(handlerCfg.AuthRedirect)
	app.Use(handlers.CheckHtmx)

	app.Get("/", handlers.IndexPage)
	app.Get("/login", handlers.LoginPage)
	app.Post("/login", handlerCfg.UserLogin)
	app.Get("/logout", handlers.UserLogout)

	app.Get("/products", handlerCfg.ProductsPage)
	app.Post("/products", handlerCfg.AddProduct)

	app.Get("/add-product", handlers.GetAddProduct)
	app.Get("/edit-product/:id", handlerCfg.GetEditProduct)

	app.Post("/products/:id", handlerCfg.UpdateProduct)
	app.Delete("/products/:id", handlerCfg.DeleteProduct)
	app.Get("/delete-products/:id", handlerCfg.DeleteProduct)
}
