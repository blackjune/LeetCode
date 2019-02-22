package solution

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	var count int
	book := make(map[string]map[string]struct{})
	for _, email := range emails {
		localName, domain := parseEmail(email)
		if _, ok := book[domain]; !ok {
			book[domain] = make(map[string]struct{})
		}
		if _, ok := book[domain][localName]; !ok {
			count++
			book[domain][localName] = struct{}{}
		}
	}
	return count
}

func parseEmail(email string) (localName string, domain string) {
	parts := strings.Split(email, "@")
	domain = parts[1]
	if idx := strings.Index(parts[0], "+"); idx >= 0 {
		localName = parts[0][0:idx]
	} else {
		localName = parts[0]
	}
	strings.ToLower
	localName = strings.Replace(localName, ".", "", -1)
	return
}
