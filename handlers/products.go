package handlers

import (
	"go-htmx-dashboard/internal/database"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (cfg *HandlerConfig) ProductsPage(c *fiber.Ctx) error {
	const LIMIT = 12

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	count, err := cfg.Db.GetProductsCount(c.Context())
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	totalPages := int(math.Ceil(float64(count) / float64(LIMIT)))

	switch {
	case page < 1:
		page = 1
	case page >= totalPages:
		page = totalPages
	default:
	}

	products, err := cfg.Db.GetAllProducts(c.Context(), database.GetAllProductsParams{
		Limit:  LIMIT,
		Offset: int32((page - 1)) * LIMIT,
	})
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	data := fiber.Map{
		"Products":    products,
		"Page":        page,
		"HasPrevPage": page > 1,
		"PrevPage":    page - 1,
		"HasNextPage": page < totalPages,
		"NextPage":    page + 1,
	}

	if c.Locals("htmx") == true {
		return c.Render("product-content", data)
	} else {
		return c.Render("products", data)
	}

}

func GetAddProduct(c *fiber.Ctx) error {
	if c.Locals("htmx") == true {
		return c.Render("add-edit-product-content", nil)
	} else {
		return c.Render("add-edit-product", nil)
	}
}

func (cfg *HandlerConfig) GetEditProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	parsedUUID, err := uuid.Parse(id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.SendString("not a valid id")
	}

	product, err := cfg.Db.GetAProduct(c.Context(), parsedUUID)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.SendString("not a valid id")
	}

	if c.Locals("htmx") == true {
		return c.Render("add-edit-product-content", product)
	} else {
		return c.Render("add-edit-product", product)
	}
}

func (cfg *HandlerConfig) AddProduct(c *fiber.Ctx) error {

	name := c.FormValue("name")
	description := c.FormValue("description")

	if len(strings.TrimSpace(name)) == 0 {
		c.Status(http.StatusBadRequest)
		return c.SendString("name cannot be empty")
	}

	if len(strings.TrimSpace(description)) == 0 {
		c.Status(http.StatusBadRequest)
		return c.SendString("description cannot be empty")
	}

	_, err := cfg.Db.CreateProduct(c.Context(), database.CreateProductParams{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	})

	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Redirect("/products")
}

func (cfg *HandlerConfig) UpdateProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	name := c.FormValue("name")
	description := c.FormValue("description")

	if len(strings.TrimSpace(name)) == 0 {
		c.Status(http.StatusBadRequest)
		return c.SendString("name cannot be empty")
	}

	if len(strings.TrimSpace(description)) == 0 {
		c.Status(http.StatusBadRequest)
		return c.SendString("description cannot be empty")
	}

	parsedUUID, err := uuid.Parse(id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.SendString("not a valid id")
	}

	_, err = cfg.Db.UpdatedProduct(c.Context(), database.UpdatedProductParams{
		ID:          parsedUUID,
		Name:        name,
		Description: description,
	})

	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Redirect("/products")

}

func (cfg *HandlerConfig) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	parsedUUID, err := uuid.Parse(id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.SendString("not a valid id")
	}

	err = cfg.Db.DeleteProduct(c.Context(), parsedUUID)

	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	c.Status(http.StatusOK)

	if c.Locals("htmx") == true {
		return nil
	} else {
		return c.Redirect("/products")
	}
}
