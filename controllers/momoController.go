package controllers

import (
	"math"
	"strconv"

	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
)

func AllMomos(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var momos []models.Momo

	database.Database.Db.Preload("User").Offset(offset).Limit(limit).Find(&momos)
	database.Database.Db.Model(&models.Momo{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": momos,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func CreateMomo(c *fiber.Ctx) error {
	var momo models.Momo

	if err := c.BodyParser(&momo); err != nil {
		return err
	}

	database.Database.Db.Create(&momo)

	return c.JSON(momo)

}

func GetMomo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	momo := models.Momo{
		Id: uint(id),
	}
	database.Database.Db.Preload("User").Find(&momo)
	return c.JSON(momo)
}

func UpdateMomo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	momo := models.Momo{
		Id: uint(id),
	}
	if err := c.BodyParser(&momo); err != nil {
		return err
	}
	database.Database.Db.Model(&momo).Updates(momo)
	return c.JSON(momo)
}

func DeleteMomo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	momo := models.Momo{
		Id: uint(id),
	}

	database.Database.Db.Delete(&momo)
	return nil
}
