package main

import (
	"SDUI_Server/internal"
	"SDUI_Server/internal/model/config"
	"SDUI_Server/internal/model/entity"
	"SDUI_Server/internal/services/item"
	"SDUI_Server/internal/services/livestream"
	"SDUI_Server/internal/services/signin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
		return
	}
	rsc, err := internal.NewRsc(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	app := fiber.New(fiber.Config{
		AppName: cfg.App,
	})
	sduiRoute := app.Group("/sdui")
	baseRoute := app.Group("/base")
	sduiRoute.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("from SDUI")
	})
	baseRoute.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("from Base")
	})

	// item route
	itemRepo, err := item.NewRepo(rsc.Db)
	if err != nil {
		log.Fatal(err)
		return
	}
	itemUI := item.NewUserInterface()
	itemHandler := item.NewHandler(itemUI, itemRepo)
	itemHandler.Register(sduiRoute, baseRoute)

	// sign in route
	signInRepo, err := signin.NewRepo(rsc.Db)
	if err != nil {
		log.Fatal(err)
		return
	}
	signInUI := signin.NewUserInterface()
	signInHandler := signin.NewHandler(signInUI, signInRepo)
	signInHandler.Register(sduiRoute, baseRoute)

	liveStreamUI := livestream.NewUserInterface()
	liveStreamRepo := livestream.NewRepository([]entity.MemberEntity{
		{
			Id:       1,
			Name:     "Callista Alifia",
			SubTitle: "Member",
			PhotoURL: "https://ucarecdn.com/16f47f20-4153-4942-bdd9-b84bc54db6e4/",
		},
	})
	liveStreamHandler := livestream.NewHandler(liveStreamRepo, liveStreamUI)
	liveStreamHandler.Register(sduiRoute, baseRoute)

	app.Listen(cfg.Port)
}
