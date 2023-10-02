package main

import (
	"database/sql"
	"go-htmx-dashboard/handlers"
	"go-htmx-dashboard/internal/database"
	"go-htmx-dashboard/util"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

  _ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
  if err != nil {
    log.Fatal(err)
  }
	defer db.Close()
	handlerCfg := &handlers.HandlerConfig{
		Db: database.New(db),
	}
  
	viewEngine := html.New("views", ".html")
	app := fiber.New(fiber.Config{
		Views:       viewEngine,
		ViewsLayout: "layouts/base",
	})
	app.Static("/css", "./dist/css")
	app.Static("/js", "./dist/js")
	setupRoutes(app, handlerCfg)
  
  // go initialUserCreation(handlerCfg)

	log.Fatal(app.Listen(":6969"))
}

func initialUserCreation(handlerCfg *handlers.HandlerConfig) {
  user, err := util.CreateAUser(handlerCfg.Db, "woow@123.com", "woow@123")
  if err != nil {
    log.Printf("error: coulnd't create user %v", err)
  } else {
    log.Printf("user created %v", user)
  }
}
