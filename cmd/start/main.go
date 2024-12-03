package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"compra/internal/app/api/routes"
	"compra/internal/app/infra/config/configEnv"
	"compra/internal/app/infra/config/db"
	"fmt"
	"github.com/joho/godotenv"
    
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// foi utilizado o fiber para criar a aplicação

func main() {  
	app := fiber.New() // cria o servidor

	app.Use(cors.New(cors.Config{ 
		AllowOrigins: "*", // Front-end permitindo o acesso
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error em het .env") // trata erro caso aconteça
	}

	initEnv := configEnv.NewConfig()

	// inicia a conexão com o banco de dados
	dataBase := db.InitDB(initEnv)

	// Configura as rotas da aplicação
	routes.SetupRoutes(app, dataBase)

	// inicia o servidor fiber
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", initEnv.Server.Ip, initEnv.Server.Port)))
}
