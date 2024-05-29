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
	base.Get("/member/:id", h.GetMemberDetailHandler)
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
	members, err := h.repo.GetMembers()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	res := make([]dto.MemberResponseDTO, 0)
	for _, mem := range members {
		res = append(res, dto.MemberResponseDTO{
			Id:        mem.Id,
			Name:      mem.Name,
			SubTitle:  mem.SubTitle,
			PhotoURL:  mem.PhotoURL,
			BornDate:  mem.BornDate.Time.Format("02 January 2006"),
			Horoscope: mem.Horoscope,
			Height:    mem.Height,
			Jiko:      mem.Jiko,
		})
	}
	return ctx.JSON(res)
}

func (h *Handler) GetMembersUIHandler(ctx *fiber.Ctx) error {
	members, err := h.repo.GetMembers()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	body := h.ui.CreateMemberListChildInterface(members)
	return ctx.JSON(dto.SDUIResponseDTO{
		Title:  "Members",
		Header: nil,
		Body:   body,
	})
}

func (h *Handler) GetMemberDetailHandler(ctx *fiber.Ctx) error {
	roomid, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	member, err := h.repo.GetMemberDetail(int64(roomid))
	if err != nil {
		return ctx.Status(404).SendString(err.Error())
	}
	return ctx.JSON(dto.MemberResponseDTO{
		Id:        member.Id,
		Name:      member.Name,
		SubTitle:  member.SubTitle,
		PhotoURL:  member.PhotoURL,
		BornDate:  member.BornDate.Time.Format("02 January 2006"),
		Horoscope: member.Horoscope,
		Height:    member.Height,
		Jiko:      member.Jiko,
	})
}
