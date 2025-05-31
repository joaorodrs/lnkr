package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/joaorodrs/linker/internals/core/domain"
	"github.com/joaorodrs/linker/internals/core/ports"
)

type LinkService struct {
	linkRepository ports.LinkRepository
}

var _ ports.LinkService = (*LinkService)(nil)

func NewLinkService(linkRepository ports.LinkRepository) *LinkService {
	return &LinkService{
		linkRepository: linkRepository,
	}
}

func (s *LinkService) CreateLink(URL string) error {
	err := s.linkRepository.CreateLink(URL)
	if err != nil {
		return err
	}
	return nil
}

func (s *LinkService) GetLink(hash string) (domain.Link, error) {
	link, err := s.linkRepository.GetLink(hash)
	if err != nil {
		return domain.Link{}, err
	}
	if link.ID == "" || link.ID == uuid.Nil.String() {
		return domain.Link{}, errors.New("Link not found")
	}
	return link, nil
}

func (s *LinkService) GetAllLinks() ([]domain.Link, error) {
	links, _ := s.linkRepository.GetAllLinks()
	return links, nil
}
