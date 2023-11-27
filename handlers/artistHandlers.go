package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)

func (h *Handler) CreateArtistHandler(c *fiber.Ctx) error {
	var artist db.Artist

	if err := c.BodyParser(&artist); err != nil {
		log.Fatalf(err.Error())
		return c.Status(400).JSON(err.Error())
	}

	newArtist, err := h.Queries.CreateArtist(c.Context(), db.CreateArtistParams{Name: artist.Name, Birthday: artist.Birthday})

	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(newArtist)
}

func (h *Handler) GetArtistsHandler(c *fiber.Ctx) error {
	artists, err := h.Queries.GetArtists(c.Context())
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(artists)
}

func (h *Handler) GetArtistHandler(c *fiber.Ctx) error {
	artistId, err := c.ParamsInt("id")
	if err != nil {
		log.Fatalf("An error has occurred:\n%s", err.Error())
		return c.SendStatus(500)
	}

	artist, err := h.Queries.GetArtist(c.Context(), int32(artistId))
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(artist)
}

func (h *Handler) UpdateArtistHandler(c *fiber.Ctx) error {
	var artist db.Artist

	if err := c.BodyParser(&artist); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updatedArtist, err := h.Queries.UpdateArtist(c.Context(), db.UpdateArtistParams{ID: artist.ID, Name: artist.Name, Birthday: artist.Birthday})
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(updatedArtist)
}

func (h *Handler) DeleteArtistHandler(c *fiber.Ctx) error {
	artistId, err := c.ParamsInt("id")
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	err = h.Queries.DeleteArtist(c.Context(), int32(artistId))
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	return c.Status(200).JSON("Artist deleted successfully")
}