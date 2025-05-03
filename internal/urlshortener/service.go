package urlshortener

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/furkankarayel/URL_Shortener/internal/cache"
	"github.com/furkankarayel/URL_Shortener/internal/utils"
)

type URLService struct {
	db       *sql.DB
	urlCache *cache.URLCache
}

func NewURLService(db *sql.DB, urlCache *cache.URLCache) *URLService {
	return &URLService{
		db:       db,
		urlCache: urlCache,
	}
}

func (s *URLService) shortenURL(originalURL string) (string, error) {
	if err := validateInput(originalURL); err != nil {
		return "", err
	}

	cachedURL, found := s.findShortCodeEntry(originalURL)
	if found {
		return cachedURL, nil
	}

	shortCode := utils.GenerateHybridString()

	query := `INSERT INTO urls (original_url, short_code, created_at) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, originalURL, shortCode, time.Now())
	if err != nil {
		return "", err
	}

	s.urlCache.Save(originalURL, shortCode)

	return "/url/" + shortCode, nil
}

func (s *URLService) getLongURL(shortCode string) (string, error) {

	longURL, found := s.findLongURLEntry(shortCode)
	if found {
		return longURL, nil
	}

	return longURL, nil
}

func (s *URLService) findShortCodeEntry(originalURL string) (string, bool) {

	shortCode := ""
	found := false
	shortCode, found = s.urlCache.FindValue(originalURL)
	if found {
		return shortCode, true
	}
	query := `SELECT short_code FROM urls WHERE original_url = $1`
	err := s.db.QueryRow(query, originalURL).Scan(&shortCode)
	if err != nil {
		log.Println("Error during findEntry" + err.Error())
		return "", false
	}

	return shortCode, true
}

func (s *URLService) findLongURLEntry(shortCode string) (string, bool) {
	longURL := ""
	found := false
	longURL, found = s.urlCache.Get(shortCode)
	if found {
		return longURL, true
	}
	query := `SELECT original_url FROM urls WHERE short_code = $1`
	err := s.db.QueryRow(query, shortCode).Scan(&longURL)
	if err != nil {
		log.Println("Error during findLongURLEntry" + err.Error())
		return "", false
	}

	return longURL, true
}

func validateInput(inputURL string) error {
	if inputURL == "" {
		return errors.New("URL cannot be empty")
	}

	return nil
}
