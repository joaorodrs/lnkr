package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
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
	// Implementation for creating a link
	return nil
}

func (h *LinkHandlers) GetLink(ctx *fiber.Ctx) error {
	// Implementation for getting a link
	return nil
}
