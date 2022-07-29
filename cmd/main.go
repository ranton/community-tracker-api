package main

import (
	"github.com/VncntDzn/community-tracker-api/pkg/admin"
	"github.com/VncntDzn/community-tracker-api/pkg/cities"
	"github.com/VncntDzn/community-tracker-api/pkg/login"
	"github.com/VncntDzn/community-tracker-api/pkg/member_skills"
	"github.com/VncntDzn/community-tracker-api/pkg/people"
	"github.com/VncntDzn/community-tracker-api/pkg/projects"
	"github.com/VncntDzn/community-tracker-api/pkg/people_skills"
	"github.com/VncntDzn/community-tracker-api/pkg/people_details"

	"github.com/VncntDzn/community-tracker-api/pkg/common/db"
	"github.com/VncntDzn/community-tracker-api/pkg/community"
	"github.com/VncntDzn/community-tracker-api/pkg/community_managers"
	"github.com/VncntDzn/community-tracker-api/pkg/community_members"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	myDB := db.Init()
	app.Use(cors.New())

	community.RegisterRoutes(app, myDB)
	cities.RegisterRoutes(app, myDB)
	community_managers.RegisterRoutes(app, myDB)
	people.RegisterRoutes(app, myDB)
	community_members.RegisterRoutes(app, myDB)
	projects.RegisterRoutes(app, myDB)
	member_skills.RegisterRoutes(app, myDB)
	admin.RegisterRoutes(app, myDB)
	login.RegisterRoutes(app, myDB)
	people_skills.RegisterRoutes(app, myDB)
	peopleDetails.RegisterRoutes(app, myDB)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(":8000")
	})

	app.Listen(":8000")
}
