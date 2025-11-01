package controller

import (
	"github.com/aow3xm/shop-hotmail/constants"
	"github.com/aow3xm/shop-hotmail/model"
	"github.com/aow3xm/shop-hotmail/service"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (c *ProductController) Route(app *fiber.App) {
	g := app.Group("/product")

	g.Get("/stock", c.GetStockController)
	g.Get("/purchase", c.PurchaseProductController)
}

func (c *ProductController) validateKioskID(kioskID string) bool {
	return kioskID == constants.HOTMAIL_KIOSK_ID || kioskID == constants.OUTLOOK_KIOSK_ID
}

func (c *ProductController) GetStockController(ctx *fiber.Ctx) error {
	var request model.GetStockRequest
	if err := ctx.QueryParser(&request); err != nil {
		return ctx.Status(400).SendString("Invalid input")
	}
	if ok := c.validateKioskID(request.KioskID); !ok {
		return ctx.Status(400).SendString("Invalid kiosk id")
	}
	stock, err := c.productService.GetStockService(request)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.Status(fiber.StatusOK).JSON(stock)
}

func (c *ProductController) PurchaseProductController(ctx *fiber.Ctx) error {
	var request model.PurchaseRequest
	if err := ctx.QueryParser(&request); err != nil {
		return ctx.Status(400).SendString("Invalid input")
	}
	if ok := c.validateKioskID(request.KioskID); !ok {
		return ctx.Status(400).SendString("Invalid kiosk id")
	}
	products, err := c.productService.PurchaseService(request)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.Status(200).JSON(products)
}
