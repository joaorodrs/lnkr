package domain

type Link struct {
	ID           string `json:"id" validate:"required,uuid"`           // ID must be a valid UUID
	URL          string `json:"url" validate:"required,url"`           // URL must be a valid URL format
	Title        string `json:"title"`                                 // Title must not be empty
	ShortenedURL string `json:"shortened_url" validate:"required,url"` // Shortened URL must not be empty
}

func NewLink(id, url, title, shortenedURL string) (*Link, error) {
	link := &Link{
		ID:           id,
		URL:          url,
		Title:        title,
		ShortenedURL: shortenedURL,
	}

	return link, nil
}

func (l *Link) GetOriginalURL() string {
	return l.URL
}
