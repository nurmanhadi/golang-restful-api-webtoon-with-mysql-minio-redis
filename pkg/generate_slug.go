package pkg

import (
	"regexp"
	"strings"
)

func GenerateSlug(title string) string {
	s := strings.ToLower(title)
	reg := regexp.MustCompile(`[^a-z0-9\s-]`)
	s = reg.ReplaceAllString(s, "")
	s = strings.Join(strings.Fields(s), "-")
	return s
}
