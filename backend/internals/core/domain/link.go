package domain

type Link struct {
	ID           string `json:"id" validate:"required,uuid"`           // ID must be a valid UUID
	URL          string `json:"url" validate:"required,url"`           // URL must be a valid URL format
	ShortenedURL string `json:"shortened_url" validate:"required,url"` // Shortened URL must not be empty
}

func NewLink(id, url, shortenedURL string) (*Link, error) {
	link := &Link{
		ID:           id,
		URL:          url,
		ShortenedURL: shortenedURL,
	}

	return link, nil
}

func (l *Link) GetOriginalURL() string {
	return l.URL
}
