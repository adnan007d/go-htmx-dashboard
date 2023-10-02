package handlers

import (
	"go-htmx-dashboard/util"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

func (cfg *HandlerConfig) UserLogin(c *fiber.Ctx) error {

	log.Println(c.Locals("htmx"))

	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := cfg.Db.GetUserByEmail(c.Context(), email)

	if err != nil {
		log.Printf("Failed to fetch user %v", err)

		c.Status(http.StatusBadRequest)
		if c.Locals("htmx") == true {
			return c.SendString("Invalid Email/Password")
		} else {
			c.Render("login", fiber.Map{
				"Email": email,
				"Error": "Invalid Email/Password",
			})
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		c.Status(http.StatusBadRequest)
		if c.Locals("htmx") == true {
			return c.SendString("Invalid Email/Password")
		} else {
			return c.Render("login", fiber.Map{
				"Email": email,
				"Error": "Invalid Email/Password",
			})
		}
	}

	// Valid Login
	tokenString, claims, err := util.GenerateJWTToken(user.ID.String())

	if err != nil {
		c.Status(http.StatusInternalServerError)
		if c.Locals("htmx") == true {
			return c.SendStatus(http.StatusInternalServerError)
		} else {
			return c.Render("login", fiber.Map{
				"Email": email,
				"Error": "Internal Server Error",
			})
		}
	}

	expirationTime, _ := claims.GetExpirationTime()

	c.Cookie(&fiber.Cookie{
		Name:     "bigpp_pass",
		Value:    tokenString,
		Expires:  expirationTime.Time, // 1 Day Validity
		Path:     "/",
		Secure:   true,
		HTTPOnly: true,
	})

	if c.Locals("htmx") == true {
		c.Response().Header.Set("HX-Redirect", "/")
		return nil
	} else {
		return c.Redirect("/")
	}
}

func UserLogout(c *fiber.Ctx) error {
	c.ClearCookie("bigpp_pass")

	if c.Locals("htmx") == true {
		c.Response().Header.Set("Hx-Location", "/login")
		return nil
	} else {
		return c.Redirect("/login")
	}
}
