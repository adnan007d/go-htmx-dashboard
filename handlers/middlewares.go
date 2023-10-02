package handlers

import (
	"go-htmx-dashboard/util"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var UnProtectedRoutes = []string{"/login"}

// Redirect to Login if not authenticated or
// if authenticated redirects to home page from login page
func (cfg *HandlerConfig) AuthRedirect(c *fiber.Ctx) error {

	var isUnProtectedRoute bool

	for _, path := range UnProtectedRoutes {
		if strings.HasPrefix(c.Path(), path) {
			isUnProtectedRoute = true
		}
	}

	cookie := string(c.Request().Header.Cookie("bigpp_pass"))
	claims := &util.JWTClaims{}

	_, err := jwt.ParseWithClaims(cookie, claims, func(_ *jwt.Token) (interface{}, error) {
		key := []byte(os.Getenv("JWT_SECRET"))
		return key, nil
	})

	if err != nil {
		log.Printf("error: %v", err)
		if isUnProtectedRoute {
			return c.Next()
		} else {
			c.Response().Header.Set("HX-Location", "/login")
			return c.Redirect("/login")
		}
	}

	if c.Path() == "/login" {
		c.Response().Header.Set("HX-Location", "/")
		return c.Redirect("/")
	}

	return c.Next()
}

func CheckHtmx(c *fiber.Ctx) error {
	c.Locals("htmx", c.GetReqHeaders()["Hx-Request"] == "true")
	return c.Next()
}
