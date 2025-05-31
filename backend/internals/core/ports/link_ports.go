package ports

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/linker/internals/core/domain"
)

type LinkService interface {
	CreateLink(URL string) error
	GetLink(hash string) (domain.Link, error)
	GetAllLinks() ([]domain.Link, error)
}

type LinkRepository interface {
	CreateLink(URL string) error
	GetLink(hash string) (domain.Link, error)
	GetAllLinks() ([]domain.Link, error)
}

type LinkHandlers interface {
	CreateLink(ctx *fiber.Ctx) error
	GetLink(ctx *fiber.Ctx) error
	GetAllLinks(ctx *fiber.Ctx) error
}
