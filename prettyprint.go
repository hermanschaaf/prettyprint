package prettyprint

import (
	"code.google.com/p/go.net/html"
	"strings"
)

func Prettify(raw string, indent string) (pretty string, e error) {
	r := strings.NewReader(raw)
	z := html.NewTokenizer(r)
	pretty = ""
	depth := 0
	prevToken := html.CommentToken
	for {
		tt := z.Next()
		tokenString := string(z.Raw())

		// strip away newlines
		if tt == html.TextToken {
			stripped := strings.Trim(tokenString, "\n")
			if len(stripped) == 0 {
				continue
			}
		}

		if tt == html.EndTagToken {
			depth -= 1
		}

		if tt != html.TextToken {
			if prevToken != html.TextToken {
				pretty += "\n"
				for i := 0; i < depth; i++ {
					pretty += indent
				}
			}
		}

		pretty += tokenString

		// last token
		if tt == html.ErrorToken {
			break
		} else if tt == html.StartTagToken {
			depth += 1
		}
		prevToken = tt
	}
	return strings.Trim(pretty, "\n"), nil
}
