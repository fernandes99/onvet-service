package handlers

import (
	"log"
	petInterface "onvet/internal/app/pet/interfaces"
	"onvet/internal/app/user/interfaces"
	rep "onvet/internal/app/user/user_repository"

	"github.com/gofiber/fiber/v2"
)

func GetAll(fiberCtx *fiber.Ctx) error {
	var userList, err = rep.GetAll(fiberCtx)

	if err != nil {
		log.Println("failed to get all users", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    userList,
	})
}

func Create(fiberCtx *fiber.Ctx) error {
	body := interfaces.CreateUserPayload{}
	err := fiberCtx.BodyParser(&body)

	if err != nil {
		log.Println("failed to parse body", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = rep.Create(fiberCtx, body)
	if err != nil {
		log.Println("failed to create user", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func GetById(fiberCtx *fiber.Ctx) error {
	userId := fiberCtx.Params("id")
	user, err := rep.GetById(fiberCtx, userId)

	if err != nil {
		log.Println("failed to get user id", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func CreateUserAddress(fiberCtx *fiber.Ctx) error {
	userId := fiberCtx.Params("id")
	body := interfaces.CreateUserAddressPayload{}
	err := fiberCtx.BodyParser(&body)
	if err != nil {
		log.Println("failed to parse body", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	body.UserId = userId

	err = rep.CreateUserAddress(fiberCtx, body)
	if err != nil {
		log.Println("failed to create user addrress", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func GetUserAddresses(fiberCtx *fiber.Ctx) error {
	userId := fiberCtx.Params("id")
	userAddressList, err := rep.GetUserAddresses(fiberCtx, userId)

	if err != nil {
		log.Println("failed to get user addresses id", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    userAddressList,
	})
}

func CreateUserPet(fiberCtx *fiber.Ctx) error {
	userId := fiberCtx.Params("id")
	body := petInterface.CreatePetPayload{}
	err := fiberCtx.BodyParser(&body)
	if err != nil {
		log.Println("failed to parse body", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	body.UserId = userId

	err = rep.CreateUserPet(fiberCtx, body)
	if err != nil {
		log.Println("failed to create user pet", err)
		return fiberCtx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
