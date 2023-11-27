package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)

func (h *Handler) CreatePurchaseHandler (c *fiber.Ctx) error {
	var purchase db.Purchase

	if err := c.BodyParser(&purchase); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	tx, err := h.Conn.Begin(c.Context())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer tx.Rollback(c.Context())

	qtx := h.Queries.WithTx(tx)

	user, err := qtx.GetUser(c.Context(), purchase.User)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	album, err := qtx.GetAlbum(c.Context(), purchase.Album)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	
	updateUser, err := qtx.UpdateUser(c.Context(), db.UpdateUserParams{ID: user.ID, Username: user.Username, Email: user.Email, Balance: user.Balance - album.Price})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	newPurchase, err := qtx.CreatePurchase(c.Context(), db.CreatePurchaseParams{User: user.ID, Album: album.ID, Date: purchase.Date})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if updateUser.Balance < 0 {
		fmt.Println("The user does not have enough money to purchase this album")
		c.Status(400).JSON("The user does not have enough money to purchase this album")
		return tx.Rollback(c.Context())
	}

	c.Status(200).JSON(newPurchase)
	return tx.Commit(c.Context())
}

func (h *Handler) GetPurchasesHandler(c *fiber.Ctx) error {
	purchases, err := h.Queries.GetPurchases(c.Context())
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(purchases)
}