package handlers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)

func (h *Handler) GetAlbumHandler(c *fiber.Ctx) error {
	albumId, err := c.ParamsInt("id")
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}
	album, err := h.Queries.GetAlbum(c.Context(), int32(albumId))
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(album)
}

func (h *Handler) GetAlbumsHandler(c *fiber.Ctx) error {
	albums, err := h.Queries.ListAlbums(c.Context())
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}
	jsonRes, err := json.Marshal(albums)
	if err != nil {
		log.Fatalf("An error occured:\n%s", err)
		return c.SendStatus(500)
	}
	return c.SendString(string(jsonRes))
}

func (h *Handler) CreateAlbumHandler(c *fiber.Ctx) error {
	var album db.Album

	if err := c.BodyParser(&album); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	newAlbum, err := h.Queries.CreateAlbum(c.Context(), db.CreateAlbumParams{Title: album.Title, Artist: album.Artist, Price: album.Price})
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(newAlbum)
}

func (h *Handler) UpdateAlbumHandlder(c *fiber.Ctx) error {
	var album db.Album

	if err := c.BodyParser(&album); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updatedAlbum, err := h.Queries.UpdateAlbum(c.Context(), db.UpdateAlbumParams{ID: album.ID, Title: album.Title, Artist: album.Artist, Price: album.Price})
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(updatedAlbum)
}

func (h *Handler) DeleteAlbumHandler(c *fiber.Ctx) error {
	albumId, err := c.ParamsInt("id")
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	err = h.Queries.DeleteAlbum(c.Context(), int32(albumId))
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	return c.Status(200).JSON("Album deleted successfully")
}