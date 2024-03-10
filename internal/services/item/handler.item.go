package item

import (
	"SDUI_Server/internal/model/dto"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	ui   *UserInterface
	repo *Repository
}

func NewHandler(ui *UserInterface, repo *Repository) *Handler {
	return &Handler{
		ui:   ui,
		repo: repo,
	}
}

func (h *Handler) Register(sdui fiber.Router, base fiber.Router) {
	// sdui
	sdui.Get("/item", h.GetListOfItemInterfaceHandler)
	sdui.Get("/item/:id", h.GetItemDetailInterfaceHandler)
	// base
	base.Get("/item", h.GetListOfItemsHandler)
	base.Get("/item/:id", h.GetItemDetailHandler)
}

func (h *Handler) GetListOfItemsHandler(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	items, _ := h.repo.GetItems(c.Context(), uint64(page), uint64(limit))
	res := make([]dto.ItemDTO, 0)
	for _, it := range items {
		res = append(res, dto.ItemDTO{
			Id:          it.Id,
			Title:       it.Title,
			Price:       it.Price,
			Description: it.Description,
			Category:    it.Category,
			Image:       it.Image,
			Rating: dto.RatingDTO{
				Rate:  it.Rate,
				Count: it.Count,
			},
		})
	}
	return c.JSON(res)
}

func (h *Handler) GetItemDetailHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).SendString("cannot find item")
	}
	item, err := h.repo.GetItemById(c.Context(), int64(id))
	if err != nil {
		return c.Status(404).SendString("cannot find item")
	}
	return c.JSON(dto.ItemDTO{
		Id:          item.Id,
		Title:       item.Title,
		Price:       item.Price,
		Description: item.Description,
		Category:    item.Category,
		Image:       item.Image,
		Rating: dto.RatingDTO{
			Rate:  item.Rate,
			Count: item.Count,
		},
	})
}
