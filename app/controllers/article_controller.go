package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/themuks/blog-api/database"
)

func GetArticles(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errorCode": fiber.StatusInternalServerError,
			"message":   err.Error(),
		})
	}

	articles, err := db.GetArticles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errorCode": fiber.StatusInternalServerError,
			"message":   err.Error(),
		})
	}

	return c.JSON(articles)
}

func GetArticle(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errorCode": fiber.StatusInternalServerError,
			"message":   err.Error(),
		})
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorCode": fiber.StatusBadRequest,
			"message":   "ID must be a valid integer",
		})
	}

	article, err := db.GetArticleById(int(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errorCode": fiber.StatusInternalServerError,
			"message":   err.Error(),
		})
	}

	return c.JSON(article)
}

func AddArticle(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"errorCode": fiber.StatusNotImplemented,
		"message":   "Not implemented",
	})
}

func UpdateArticle(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"errorCode": fiber.StatusNotImplemented,
		"message":   "Not implemented",
	})
}

func DeleteArticle(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"errorCode": fiber.StatusNotImplemented,
		"message":   "Not implemented",
	})
}

func GetArticleComments(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"errorCode": fiber.StatusNotImplemented,
		"message":   "Not implemented",
	})
}
