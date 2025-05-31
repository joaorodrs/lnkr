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

func (s *LinkService) CreateLink(URL string, title string) error {
	err := s.linkRepository.CreateLink(URL, title)
	if err != nil {
		return err
	}
	return nil
}

func (s *LinkService) GetLink(ID int) (domain.Link, error) {
	link, err := s.linkRepository.GetLink(ID)
	if err != nil {
		return domain.Link{}, err
	}
	if link.ID == "" || link.ID == uuid.Nil.String() {
		return domain.Link{}, errors.New("link not found")
	}
	return link, nil
}
