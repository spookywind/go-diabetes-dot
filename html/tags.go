package html

import (
	"strings"
)

func h1(s string, items ...string) string {
	return "<h1>" + s + "</h1>\n" + strings.Join(items, " ")
}

func h2(s string, items ...string) string {
	return "<h2>" + s + "</h2>\n" + strings.Join(items, " ")
}
