package service

import (
  "fmt"
  "net/url"
  "regexp"
  "time"
)

const (
  SHORTENER_REGEX = `^[0-9a-zA-Z_]{6}$`
  SHORTENER_BASE  = 62
  SHORTENER_SET   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
  SHORTENER_MAX   = 56800235583
)

var (
  generateIndex  = 1
  shortenedUrls  = make(map[string]shortenURL)
  registeredUrls = make(map[string]string)
)

type shortenURL struct {
  URL           string
  StartDate     time.Time
  LastSeenDate  time.Time
  RedirectCount int
}

func CreateShortenURL(shortcode string, url string) error {
  if !IsValidShortcode(shortcode) {
    return fmt.Errorf("invalid short code: %s", shortcode)
  }

  if IsShortenURLExists(shortcode) {
    return fmt.Errorf("short code '%s' already used", shortcode)
  }

  shortenedUrls[shortcode] = shortenURL{
    URL:       url,
    StartDate: time.Now(),
  }

  return nil
}

func IsShortenURLExists(shortcode string) bool {
  _, ok := shortenedUrls[shortcode]
  return ok
}

func RegisterURL(shortcode string, url string) {
  registeredUrls[url] = shortcode
}

func GetRegisteredURL(url string) string {
  return registeredUrls[url]
}

func IsValidURL(str string) bool {
  if len(str) == 0 {
    return false
  }

  u, err := url.Parse(str)
  return err == nil && u.Scheme != "" && u.Host != ""
}

func IsValidShortcode(shortcode string) bool {
  re := regexp.MustCompile(SHORTENER_REGEX)
  return re.MatchString(shortcode)
}

func GetNextShortcode() string {
  var shortcode string

  for shortcode == "" || IsShortenURLExists(shortcode) {
    shortcode = GenerateShortcode(generateIndex)
    generateIndex++
  }

  return shortcode
}

func GenerateShortcode(num int) string {
  b := make([]byte, 0)

  for num > 0 {
    r := num % SHORTENER_BASE
    num /= SHORTENER_BASE
    b = append([]byte{SHORTENER_SET[int(r)]}, b...)
  }

  shortcode := string(b)
  for len(shortcode) < 6 {
    shortcode += "_"
  }

  return shortcode
}
