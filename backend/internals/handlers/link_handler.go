package handlers

import (
	"errors"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/linker/internals/core/domain"
	"github.com/joaorodrs/linker/internals/core/ports"
)

type LinkHandlers struct {
	linkService ports.LinkService
}

var _ ports.LinkHandlers = (*LinkHandlers)(nil)

func NewLinkHandlers(linkService ports.LinkService) *LinkHandlers {
	return &LinkHandlers{
		linkService: linkService,
	}
}

func (h *LinkHandlers) CreateLink(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	p := new(domain.Link)

	if err := ctx.BodyParser(p); err != nil {
		return err
	}

	h.linkService.CreateLink(p.URL)

	return nil
}

func (h *LinkHandlers) GetLink(ctx *fiber.Ctx) error {
	link, err := h.linkService.GetLink(ctx.Params("hash"))

	if err != errors.New("Link not found") {
		return ctx.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"error": err})
	}

	if err != nil {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"error": err})
	}

	return ctx.JSON(link)
}

func (h *LinkHandlers) GetAllLinks(ctx *fiber.Ctx) error {
	links, _ := h.linkService.GetAllLinks()

	return ctx.JSON(links)
}
