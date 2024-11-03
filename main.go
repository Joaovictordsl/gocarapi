package main

import (
	"Desktop/test/test"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Carro struct {
	Id     int    `json:"id"`
	Marca  string `json:"marca"`
	Modelo string `json:"modelo"`
	Preco  string `json:"preco"`
}

func main() {

	connStr := "postgresql://postgres:12345678@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	queries := test.New(db)

	defer db.Close()

	app := fiber.New()

	app.Post("/carros", func(c *fiber.Ctx) error {
		var carro Carro
		if err := c.BodyParser(&carro); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data",
			})
		}

		newCarro, err := queries.CreateCarro(c.Context(), test.CreateCarroParams{
			Marca:  carro.Marca,
			Modelo: carro.Modelo,
			Preco:  carro.Preco,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(newCarro)
	})

	app.Get("/carros", func(c *fiber.Ctx) error {
		carros, err := queries.ListCarros(c.Context())

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(carros)

	})

	app.Get("/carro", func(c *fiber.Ctx) error {
		carro, err := queries.GetCarro(c.Context(), 1)

		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(carro)

	})

	app.Get("/carros/marca/:marca", func(c *fiber.Ctx) error {
		marca := c.Params("marca")

		carros, err := queries.GetCarroByMarca(c.Context(), marca)

		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(carros)

	})

	app.Listen(":3000")
}
