package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/alwialdi9/be-jajanskuy/internal/models"
	"github.com/gofiber/fiber/v2"
)

const TimeFormat = "2006-01-02 15:04:05"

func InitController(c *fiber.Ctx) error {
	// Initiate all controllers here
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginRequest request

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	userData := models.User{
		ID:   1,
		Name: "Alwi",
		// Email:     "alwi@mail.com",
		Age:       25,
		LastLogin: time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"),
	}

	userData.Email = loginRequest.Email

	type response struct {
		Status       string      `json:"status"`
		Data         models.User `json:"data"`
		ExpiredLogin bool        `json:"expired_login"`
	}

	login, err := time.Parse(TimeFormat, userData.LastLogin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Failed to parse time")
	}
	day, _ := strconv.Atoi(os.Getenv("EXPIRED_TIME"))

	if login.Before(time.Now().AddDate(0, 0, -1*day)) {
		return c.Status(fiber.StatusOK).JSON(response{
			Status:       "success",
			Data:         userData,
			ExpiredLogin: true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response{
		Status:       "success",
		Data:         userData,
		ExpiredLogin: false,
	})
}
