package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)

func (h *Handler) CreateArtistHandler(c *fiber.Ctx) error {
	var artist db.Artist

	if err := c.BodyParser(&artist); err != nil {
		log.Fatalf(err.Error())
		return c.Status(400).JSON(err.Error())
	}

	newArtist, err := h.queries.CreateArtist(c.Context(), db.CreateArtistParams{Name: artist.Name, Birthday: artist.Birthday})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(200).JSON(newArtist)
}

func (h *Handler) GetArtistsHandler(c *fiber.Ctx) error {
	artists, err := h.queries.GetArtists(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(artists)
}

func (h *Handler) GetArtistHandler(c *fiber.Ctx) error {
	artistId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	artist, err := h.queries.GetArtist(c.Context(), int32(artistId))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(200).JSON(artist)
}

func (h *Handler) UpdateArtistHandler(c *fiber.Ctx) error {
	var artist db.Artist

	if err := c.BodyParser(&artist); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updatedArtist, err := h.queries.UpdateArtist(c.Context(), db.UpdateArtistParams{ID: artist.ID, Name: artist.Name, Birthday: artist.Birthday})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(updatedArtist)
}

func (h *Handler) DeleteArtistHandler(c *fiber.Ctx) error {
	artistId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	err = h.queries.DeleteArtist(c.Context(), int32(artistId))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON("Artist deleted successfully")
}