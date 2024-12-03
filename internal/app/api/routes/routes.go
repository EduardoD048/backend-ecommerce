// Package routes defines the API routes for the application.
package routes

import (
	"compra/internal/app/api/model/product_model"
	"compra/internal/app/api/model/purchase_model"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func SetupRoutes(app *fiber.App, dataBase *gorm.DB) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // permissão pro front pode receber requisições
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// definindo a rota para adicionar um produto 
	app.Post("/purchase", func(c *fiber.Ctx) error {
		var products []product_model.Product

		// faz o parse do corpo da requisição para a variável products
		if err := c.BodyParser(&products); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error processing request"}) // trata erro
		}

		// cria uma nova compra 
		purchase := purchase_model.Purchase{
			Products: products,
		}

		// Salva no banco 
		if err := dataBase.Create(&purchase).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error saving purchase"}) // trata erro 
		}

		// retorna 201
		return c.Status(fiber.StatusCreated).JSON(purchase)
	})

	// rota para puxar todas as compras feitas 
	app.Get("/purchase", func(c *fiber.Ctx) error {
		var purchases []purchase_model.Purchase

		// Retorna toda as compras feitas do banco
		if err := dataBase.Preload("Products").Find(&purchases).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error retrieving purchases"})
		}

		// 201
		return c.JSON(purchases)
	})

	app.Get("/purchase/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var purchase purchase_model.Purchase

	     // retorna a compra feita com o id passado
		if err := dataBase.Preload("Products").First(&purchase, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Purchase not found"})
		} 

		return c.JSON(purchase)
	})

    // rota delete
	app.Delete("/purchase", func(c *fiber.Ctx) error {
		if err := dataBase.Exec("DELETE FROM purchases").Error; err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error deleting purchases",
			})
		}
	
		return c.JSON(fiber.Map{
			"message": "All purchases have been successfully deleted",
		})
	})
}
