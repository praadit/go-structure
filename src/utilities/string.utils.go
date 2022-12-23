package utilities

import (
	"fmt"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func CreateSearchQuery(searchable []string, search string) (query string) {
	query = ""
	for i, v := range searchable {
		query += fmt.Sprintf("%s ilike '%%%s%%' ", v, search)
		if i < len(searchable) {
			query += "or "
		}
	}
	return
}
