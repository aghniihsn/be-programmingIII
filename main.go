package main

import (
	"fmt"
	"inibackend/config"
	"inibackend/router"
	"log"
	"os"
	"strings"

	_ "inibackend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}
	} else {
		fmt.Println("File .env dimuat (local)")
	}
}

// @title TES SWAGGER PEMROGRAMAN III
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/indrariksa
// @contact.email indra@ulbi.ac.id

// @BasePath /
// @schemes http
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	app := fiber.New()

	//Logging Request di terminal
	app.Use(logger.New())

	//Basic Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigin(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE",
	}))

	//Route Mahasiswa
	router.SetupRoutes(app)

	//Handler 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Endpoint tidak ditemukan",
		})
	})

	//Baca PORT yang ada di .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" //Default port kalau tidak ada di .env
	}

	//untuk log cek konek di port mana
	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	} //Koneksi terputus
}

// package main

// import (
// 	"fmt"

// 	"aidanwoods.dev/go-paseto"
// )

// func main() {
// 	privateKey := paseto.NewV4AsymmetricSecretKey()
// 	publicKey := privateKey.Public()

// 	fmt.Println("Private Key (hex):", privateKey.ExportHex())
// 	fmt.Println("Public Key (hex):", publicKey.ExportHex())
// }
