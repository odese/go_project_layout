package controllers

import (
	// logic ".../backend/pkg/businessLogic"
	// ".../backend/pkg/domain/models"
)


// CreateUser godoc
// @Summary CreateUser
// @Description	Logics:
// @Description		- Check if mail exists
// @Description		- Set custom fields
// @Description		- Insert into ptt bucket
// @Description		- Send verification mail
// @Tags userController
// @Accept json
// @Produce json
// @Param userReq body dto.User true "User Request"
// @Param param1 path string true "param1"
// @Param param2 query string false "param2"
// @Success 200 {object} models.User "Created User Struct"
// @Failure 400 {string} return code
// @Failure 404 {string} return code
// @Failure 500 {string} return code
// @Router /api/user/create/{param1} [post]
// func CreateUser(c xxx.Ctx) error {
	// var userReq models.User
	// param1 := c.Params("param1")
	// param2 := c.Query("param2")
	// if err := c.BodyParser(&userReq); err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": err.Error()})
	// }

	// res, err := logic.CreateUser(userReq)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Error on Creating User: " + err.Error()})
	// } else {
	// 	return c.Status(fiber.StatusOK).JSON(res)
	// }
// }
