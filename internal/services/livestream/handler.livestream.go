package livestream

import (
	"SDUI_Server/internal/model/dto"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo *Repository
	ui   *UserInterface
}

func NewHandler(repo *Repository, ui *UserInterface) *Handler {
	return &Handler{
		repo: repo,
		ui:   ui,
	}
}

func (h *Handler) Register(sdui fiber.Router, base fiber.Router) {
	base.Get("/member", h.GetMembersHandler)
	base.Get("member/livestream/:id", h.GetMemberLiveStreamURL)
	sdui.Get("/member", h.GetMembersUIHandler)
}

func (h *Handler) GetMemberLiveStreamURL(ctx *fiber.Ctx) error {
	roomid, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	url, err := h.repo.GetMemberLiveStreamURL(int64(roomid))
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.JSON(dto.MemberLiveStreamURLResponseDTO{
		Id:            int64(roomid),
		LiveStreamURL: url,
	})
}

func (h *Handler) GetMembersHandler(ctx *fiber.Ctx) error {
	members := h.repo.GetMembers()
	return ctx.JSON(members)
}

func (h *Handler) GetMembersUIHandler(ctx *fiber.Ctx) error {
	return nil
}
