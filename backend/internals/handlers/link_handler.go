package handlers

import (
	"errors"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/linker/internals/core/domain"
	"github.com/joaorodrs/linker/internals/core/ports"
	. "github.com/joaorodrs/linker/internals/helpers"
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
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Malparsed body JSON."})
	}

	err := h.linkService.CreateLink(p.URL)

	if errors.Is(err, ErrInvalidPayload) {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	} else if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"msg": "Link created successfully"})
}

func (h *LinkHandlers) GetLink(ctx *fiber.Ctx) error {
	link, err := h.linkService.GetLink(ctx.Params("hash"))

	if errors.Is(err, ErrNotFound) {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": err})
	} else if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	return ctx.Redirect(link.URL, http.StatusPermanentRedirect)
}

func (h *LinkHandlers) GetAllLinks(ctx *fiber.Ctx) error {
	links, err := h.linkService.GetAllLinks()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	return ctx.JSON(links)
}
