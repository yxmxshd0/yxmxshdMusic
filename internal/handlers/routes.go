package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"yxmxshdMusic/internal/databases/repos"
	"yxmxshdMusic/internal/models"
)

type Routes struct {
	mongo repos.MongoRepos
	ctx   *context.Context
}

func New(mongo repos.MongoRepos, ctx *context.Context) *Routes {
	return &Routes{mongo: mongo, ctx: ctx}
}

func (route *Routes) RegisterRoutes(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error { return c.JSON("Hello, World!") })
	app.Get("/music", route.HandlerRegisterMusic)
}

func (route *Routes) HandlerRegisterMusic(c *fiber.Ctx) error {
	var body models.DocumentsToSend

	err := c.BodyParser(&body)

	if err != nil {
		customError := models.Errors{
			PersError: err.Error(),
			Message:   "Error parsing body",
			Status:    fiber.StatusBadRequest,
		}
		return c.JSON(customError)
	}
	err = route.mongo.SendDocument(route.ctx, body)
	if err != nil {
		customError := models.Errors{
			PersError: err.Error(),
			Message:   "Error sending document",
			Status:    fiber.StatusInternalServerError,
		}
		return c.JSON(customError)
	}

	c.Status(fiber.StatusOK)
	return c.JSON("Send music successfully")
}
