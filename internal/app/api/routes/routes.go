// Package routes defines the API routes for the application.
package routes

import (
	"compra/internal/app/api/model/product_model"
	"compra/internal/app/api/model/purchase_model"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupRoutes configures the API routes for the application.
//
// It sets up two routes: one for creating a new purchase (POST /purchase) and
// one for retrieving all purchases (GET /purchase).

func SetupRoutes(app *fiber.App, dataBase *gorm.DB) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Substitua pela URL do seu front-end
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Define the route for creating a new purchase.
	app.Post("/purchase", func(c *fiber.Ctx) error {
		// Define a slice to store the products in the request body.
		var products []product_model.Product

		// Parse the request body into the products slice.
		if err := c.BodyParser(&products); err != nil {
			// Return a 400 error if the request body is invalid.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error processing request"})
		}

		// Create a new purchase with the products from the request body.
		purchase := purchase_model.Purchase{
			Products: products,
		}

		// Save the purchase to the database.
		if err := dataBase.Create(&purchase).Error; err != nil {
			// Return a 500 error if there is an issue saving the purchase.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error saving purchase"})
		}

		// Return the created purchase with a 201 status code.
		return c.Status(fiber.StatusCreated).JSON(purchase)
	})

	// Define the route for retrieving all purchases.
	app.Get("/purchase", func(c *fiber.Ctx) error {
		// Define a slice to store the purchases from the database.
		var purchases []purchase_model.Purchase

		// Retrieve all purchases from the database, including their products.
		if err := dataBase.Preload("Products").Find(&purchases).Error; err != nil {
			// Return a 500 error if there is an issue retrieving the purchases.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving purchases"})
		}

		// Return the purchases with a 200 status code.
		return c.JSON(purchases)
	})

	// Define the route for retrieving a purchase by its ID.
	app.Get("/purchase/:id", func(c *fiber.Ctx) error {
		// Get the ID from the URL parameters
		id := c.Params("id")

		// Define a variable to store the purchase.
		var purchase purchase_model.Purchase

		// Retrieve the purchase by its ID, including its products.
		if err := dataBase.Preload("Products").First(&purchase, id).Error; err != nil {
			// Return a 404 error if the purchase is not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Purchase not found"})
		}

		// Return the found purchase with a 200 status code.
		return c.JSON(purchase)
	})
}
