package main

import (
	"fmt"
	"log"
	"strconv"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/hojin-kr/go-http-game-server/cmd/docs"
	EtcModel "github.com/hojin-kr/go-http-game-server/cmd/model/etc"
	ProfileModel "github.com/hojin-kr/go-http-game-server/cmd/model/profile"
	SocialModel "github.com/hojin-kr/go-http-game-server/cmd/model/social"
	UserModel "github.com/hojin-kr/go-http-game-server/cmd/model/user"
)

// @title go-http-game-server API
// @version 1.0.0
// @description go-http-game-server API
// @host localhost:3000
// @BasePath /
func main() {
	log.Println(uuid.New())
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // defaut

	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		log.Println("🥇 First handler")
		return c.Next()
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// APIs for go-http-game-server Games

	// User APIs

	// GET GetUserByUUID
	v1.Get("/user/:uuid", func(c *fiber.Ctx) error {
		uuid := c.Params("uuid")
		user := UserModel.GetByUUID(uuid)
		return c.JSON(user)
	})

	// GET GetRecoveryCodeByUserID
	v1.Get("/user/:user_id/recovery", func(c *fiber.Ctx) error {
		userId, _ := strconv.Atoi(c.Params("user_id"))
		recovery := UserModel.GetRecoveryCodeByUserID(userId)
		return c.JSON(recovery)
	})

	// POST RecoeveryUser
	// Change UUID Using Recovery Code
	v1.Post("/user/recovery", func(c *fiber.Ctx) error {
		req := UserModel.Recovery{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		user := UserModel.RecoveryUser(req.Code)
		return c.JSON(user)
	})

	// Profile APIs

	// GET GetProfileByID
	v1.Get("/profile/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		profile := ProfileModel.Get(id)
		return c.JSON(profile)
	})

	// POST UpdateProfileNicnameByID
	v1.Post("/profile", func(c *fiber.Ctx) error {
		req := ProfileModel.Profile{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		// check nickname exists
		if ProfileModel.CheckNicknameExists(req.Nickname) {
			return c.Status(400).JSON(fiber.Map{
				"error": "nickname exists",
			})
		}
		profile := ProfileModel.Update(req.UserId, req.Nickname)
		return c.JSON(profile)
	})

	// Etc APIs

	// GET GetEtcByUserID
	v1.Get("/etc/:user_id/:key", func(c *fiber.Ctx) error {
		userId, _ := strconv.Atoi(c.Params("user_id"))
		key := c.Params("key")
		etc := EtcModel.Get(userId, key)
		return c.JSON(etc)
	})

	// GET GetAllEtcByUserID
	v1.Get("/etc/:user_id", func(c *fiber.Ctx) error {
		userId, _ := strconv.Atoi(c.Params("user_id"))
		etcs := EtcModel.GetAll(userId)
		return c.JSON(etcs)

	})

	// POST UpdateEtcByUserID
	v1.Post("/etc", func(c *fiber.Ctx) error {
		req := EtcModel.Etc{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		etc := EtcModel.Update(req.UserID, req.Key, req.Value)
		return c.JSON(etc)
	})

	// Social APIs

	// GET GetSocialCountByTargetIDAndType
	v1.Get("/social/:target_id/:type", func(c *fiber.Ctx) error {
		targetId, _ := strconv.Atoi(c.Params("target_id"))
		socialType := c.Params("type")
		count := SocialModel.GetCountByTargetIDAndType(targetId, socialType)
		return c.JSON(count)
	})

	// GET GetSocialByTargetIDAndTypeLimitOffset
	v1.Get("/social/:target_id/:type/:limit/:offset", func(c *fiber.Ctx) error {
		targetId, _ := strconv.Atoi(c.Params("target_id"))
		socialType := c.Params("type")
		limit, _ := strconv.Atoi(c.Params("limit"))
		offset, _ := strconv.Atoi(c.Params("offset"))
		socials := SocialModel.GetByTargetIDAndTypeLimitOffset(targetId, socialType, limit, offset)
		return c.JSON(socials)
	})

	// GET GetSocialByUserIDAndTypeLimitOffset
	v1.Get("/social/:user_id/:type/:limit/:offset", func(c *fiber.Ctx) error {
		userId, _ := strconv.Atoi(c.Params("user_id"))
		socialType := c.Params("type")
		limit, _ := strconv.Atoi(c.Params("limit"))
		offset, _ := strconv.Atoi(c.Params("offset"))
		socials := SocialModel.GetByUserIDAndTypeLimitOffset(userId, socialType, limit, offset)
		return c.JSON(socials)
	})

	// POST InsertSocial
	v1.Post("/social", func(c *fiber.Ctx) error {
		req := SocialModel.SocialRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		social := SocialModel.Insert(req.UserID, req.TargetID, req.Type, req.Vars)
		return c.JSON(social)
	})

	// POST DeleteSocialByID
	v1.Post("/social/delete", func(c *fiber.Ctx) error {
		req := SocialModel.SocialDeleteRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		SocialModel.DeleteByID(req.ID)
		// return success
		return c.JSON(fiber.Map{
			"success": true,
		})
	})

	// GET /api/user/id
	app.Get("/api/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		//  select to postgres db
		log.Println(" select to postgres db" + id)
		msg := fmt.Sprintf("✋ %s", c.Params("id"))
		return c.SendString(msg) // => ✋ register
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	log.Fatal(app.Listen(":3000"))
}
