package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BioData struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          int    `json:"age"`
	Birthday     string `json:"birthday"`
	Email        string `json:"email"`
	ContactNo    string `json:"contact_number"`
	ParentsNames string `json:"parents_names"`
	Address      string `json:"address"`
	Occupation   string `json:"occupation"`
}

var bioDataList []BioData
var idCounter = 1

func main() {
	app := fiber.New()

	// Get all bio data
	app.Get("/bio", func(c *fiber.Ctx) error {
		return c.JSON(bioDataList)
	})

	// Get bio data by ID
	app.Get("/bio/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for _, bio := range bioDataList {
			if strconv.Itoa(bio.ID) == id {
				return c.JSON(bio)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Bio data not found"})
	})

	// Create new bio data
	app.Post("/bio", func(c *fiber.Ctx) error {
		bio := new(BioData)
		if err := c.BodyParser(bio); err != nil {
			return err
		}
		bio.ID = idCounter
		idCounter++
		bioDataList = append(bioDataList, *bio)
		return c.JSON(bio)
	})

	// Update bio data by ID
	app.Put("/bio/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		newBio := new(BioData)
		if err := c.BodyParser(newBio); err != nil {
			return err
		}
		for i, bio := range bioDataList {
			if strconv.Itoa(bio.ID) == id {
				newBio.ID = bio.ID
				bioDataList[i] = *newBio
				return c.JSON(newBio)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Bio data not found"})
	})

	// Delete bio data by ID
	app.Delete("/bio/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, bio := range bioDataList {
			if strconv.Itoa(bio.ID) == id {
				bioDataList = append(bioDataList[:i], bioDataList[i+1:]...)
				return c.SendStatus(fiber.StatusNoContent)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Bio data not found"})
	})

	log.Fatal(app.Listen(":3000"))
}
