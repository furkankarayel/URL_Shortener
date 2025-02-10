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

	log.Println("shortenURL called with " + originalURL)
	if err := validateInput(originalURL); err != nil {
		return "", err
	}

	cachedURL, found := s.urlCache.Get(originalURL)
	if found {
		log.Println("URL found in cache during shortenURL" + cachedURL)
		return cachedURL, nil
	}
	log.Println("URL not found in cache during shortenURL")

	shortCode := utils.GenerateHybridString()
	log.Println(shortCode)

	query := `INSERT INTO urls (original_url, short_code, created_at) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, originalURL, shortCode, time.Now())
	if err != nil {
		return "", err
	}
	log.Println("URL inserted into database during shortenURL" + originalURL + " " + shortCode)

	s.urlCache.Save(originalURL, shortCode)

	return shortCode, nil
}

func (s *URLService) getLongURL(shortCode string) (string, error) {
	log.Println("getLongURL called with " + shortCode)

	longURL, found := s.urlCache.Get(shortCode)
	if found {
		log.Println("URL found in cache during getOriginalURL" + longURL)
		return longURL, nil
	}

	query := `SELECT original_url FROM urls WHERE short_code = $1`
	var originalURL string
	err := s.db.QueryRow(query, shortCode).Scan(&originalURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("URL not found")
		}
		return "", err
	}

	log.Println("URL found in database during getOriginalURL" + originalURL)

	return originalURL, nil
}

func validateInput(inputURL string) error {
	if inputURL == "" {
		return errors.New("URL cannot be empty")
	}
	// Add more validation logic as needed
	return nil
}
