package controllers

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/kooroshh/fiber-boostrap/app/repository"
	"github.com/kooroshh/fiber-boostrap/pkg/response"
)

func GetHistory(ctx fiber.Ctx) error {
	resp, err := repository.GetAllMessage(ctx)
	if err != nil {
		log.Println(err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "terjadi kesalahan pada server", nil)
	}
	return response.SendSuccessResponse(ctx, resp)
}
