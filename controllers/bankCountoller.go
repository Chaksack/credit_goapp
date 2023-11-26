package controllers

import (
	"math"
	"strconv"

	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
)

func AllBanks(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var bank []models.Bank

	database.Database.Db.Preload("Bank").Offset(offset).Limit(limit).Find(&bank)
	database.Database.Db.Model(&models.Bank{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": bank,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func CreateBank(c *fiber.Ctx) error {
	var bank models.Bank

	if err := c.BodyParser(&bank); err != nil {
		return err
	}

	database.Database.Db.Create(&bank)

	return c.JSON(bank)

}

func GetBank(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	bank := models.Bank{
		Id: uint(id),
	}
	database.Database.Db.Preload("User").Find(&bank)
	return c.JSON(bank)
}

func UpdateBank(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	bank := models.Bank{
		Id: uint(id),
	}
	if err := c.BodyParser(&bank); err != nil {
		return err
	}
	database.Database.Db.Model(&bank).Updates(bank)
	return c.JSON(bank)
}

func DeleteBank(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	bank := models.Bank{
		Id: uint(id),
	}

	database.Database.Db.Delete(&bank)
	return nil
}
