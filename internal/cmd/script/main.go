package main

import (
	"SDUI_Server/internal"
	"SDUI_Server/internal/model/config"
	"SDUI_Server/internal/model/entity"
	"SDUI_Server/internal/services/item"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jaswdr/faker"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Error(err)
		return
	}
	imgs, err := config.LoadImages()
	if err != nil {
		log.Error(err)
		return
	}
	ctx := context.Background()
	rsc, err := internal.NewRsc(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	itemRepo, err := item.NewRepo(rsc.Db)
	if err != nil {
		log.Fatal(err)
		return
	}
	fake := faker.New()
	for _, img := range imgs {
		item := entity.ItemEntity{
			Title:       fake.Payment().CreditCardNumber(),
			Price:       fake.RandomFloat(2, 1, 99),
			Description: fake.Car().Model(),
			Category:    fake.Car().FuelType(),
			Image:       img.Url,
			Rate:        fake.Float64(1, 0, 5),
			Count:       fake.IntBetween(0, 99),
		}
		err := itemRepo.AddItem(ctx, item)
		if err != nil {
			log.Error(err)
		}
	}
	fmt.Println("done")
}
