package ports

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/linker/internals/core/domain"
)

type LinkService interface {
	CreateLink(URL string, title string) error
	GetLink(ID int) (domain.Link, error)
}

type LinkRepository interface {
	CreateLink(URL string, title string) error
	GetLink(ID int) (domain.Link, error)
}

type LinkHandlers interface {
	CreateLink(ctx *fiber.Ctx) error
	GetLink(ctx *fiber.Ctx) error
}
