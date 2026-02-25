package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
)

var md = goldmark.New()

// ToHTML converts markdown source to HTML.
func ToHTML(source string) string {
	var buf bytes.Buffer
	if err := md.Convert([]byte(source), &buf); err != nil {
		return source
	}
	return buf.String()
}
