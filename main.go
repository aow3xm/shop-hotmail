package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aow3xm/shop-hotmail/config"
	"github.com/aow3xm/shop-hotmail/controller"
	"github.com/aow3xm/shop-hotmail/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	productService := service.NewProductService(cfg)
	productController := controller.NewProductController(productService)

	productController.Route(app)

	go func() {
		app.Listen(":3000")
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	fmt.Println("Graceful shutting down...")
	app.Shutdown()

}
