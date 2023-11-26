package routes

import (
	"strconv"
	"xactscore/controllers"
	"xactscore/middleware"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to xactscore backend. Xactscore application is up and running!!! \n health checker passed.")
	})
	// app.Get("/swagger/*", swagger.Handler)

	//controllers endpoints
	app.Post("/api/register", controllers.Register)
	app.Post("/api/registerstaff", controllers.StaffRegister)
	app.Post("/api/registerbusiness", controllers.BusinessRegister)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/loginstaff", controllers.StaffLogin)
	app.Post("/api/loginBusiness", controllers.BusinessLogin)

	app.Use(middleware.Isauthenticated)

	app.Put("api/users/info", controllers.UpdateInfo)
	app.Put("api/users/password", controllers.UpdatePassword)
	app.Put("api/business/info", controllers.UpdateInfo)
	app.Put("api/business/password", controllers.UpdatePassword)
	app.Put("api/staff/info", controllers.UpdateInfo)
	app.Put("api/staff/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)
	app.Get("api/staff", controllers.Staff)
	app.Get("api/business", controllers.Business)
	app.Post("api/logout", controllers.Logout)

	app.Get("api/staffs", controllers.AllStaff)
	app.Post("api/staffs", controllers.CreateStaff)
	app.Get("api/staffs/:id", controllers.GetStaff)
	app.Put("api/staffs/:id", controllers.UpdateStaff)
	app.Delete("api/staffs/:id", controllers.DeleteStaff)

	app.Get("api/users", controllers.AllUsers)
	app.Post("api/users", controllers.CreateUser)
	app.Get("api/users/:id", controllers.GetUser)
	app.Put("api/users/:id", controllers.UpdateUser)
	app.Delete("api/users/:id", controllers.DeleteUser)

	app.Get("api/business", controllers.AllBusiness)
	app.Post("api/business", controllers.CreateBusiness)
	app.Get("api/business/:id", controllers.GetBusiness)
	app.Put("api/business/:id", controllers.UpdateBusiness)
	app.Delete("api/business/:id", controllers.DeleteBusiness)

	app.Get("api/roles", controllers.AllRoles)
	app.Post("api/roles", controllers.CreateRole)
	app.Get("api/roles/:id", controllers.GetRole)
	app.Put("api/roles/:id", controllers.UpdateRole)
	app.Delete("api/roles/:id", controllers.DeleteRole)

	app.Get("api/permissions", controllers.AllPermissions)
	app.Post("api/permissions", controllers.CreatePermission)
	app.Get("api/permissions/:id", controllers.GetPermission)
	app.Put("api/permissions/:id", controllers.UpdatePermission)
	app.Delete("api/permissions/:id", controllers.DeletePermission)

	//accounts
	//bank accounts
	app.Get("api/banks", controllers.AllBanks)
	app.Post("api/bank", controllers.CreateBank)
	app.Get("api/banks/:id", controllers.GetBank)
	app.Put("api/banks/:id", controllers.UpdateBank)
	app.Delete("api/banks/:id", controllers.DeleteBank)

	//momo accounts
	app.Get("api/momos", controllers.AllMomos)
	app.Post("api/momos", controllers.CreateMomo)
	app.Get("api/momos/:id", controllers.GetMomo)
	app.Put("api/momos/:id", controllers.UpdateMomo)
	app.Delete("api/momos/:id", controllers.DeleteMomo)

	//credits
	// app.Get("api/credits", controllers.credits)
	// app.Post("api/credits", controllers.CreateCredit)
	// app.Get("api/permissions/:id", controllers.GetCredit)
	// app.Put("api/permissions/:id", controllers.UpdateCredit)

	app.Get("/api/other-app-endpoint", func(c *fiber.Ctx) error {
		rawScoreStr := c.Query("rawScore")
		scorePercentageStr := c.Query("scorePercentage")

		rawScore, errRaw := strconv.Atoi(rawScoreStr)
		scorePercentage, errPercentage := strconv.Atoi(scorePercentageStr)

		if errRaw != nil || errPercentage != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input. rawScore and scorePercentage must be integers.",
			})
		}

		creditData := models.Credit{
			RawScore:        rawScore,
			ScorePercentage: scorePercentage,
		}

		return c.JSON(creditData)

	})

	//reports
}
