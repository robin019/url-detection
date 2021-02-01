package router

import (
	"log"

	"github.com/robin019/url-detection/src/utils/config"

	"github.com/robin019/url-detection/src/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/robin019/url-detection/src/handler"
)

func Route() {
	app := fiber.New()
	app.Use(middleware.Recover) //recover from a panic thrown by any handler
	app.Use(cors.New())         //enable CORS (if web browser support is needed)

	api := app.Group("/api")

	v1 := api.Group("/v1")
	{
		v1.Get("/malicious_url", handler.UrlCheck)
	}

	log.Fatal(app.Listen(":" + config.Server.Port))
}
