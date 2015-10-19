package prettyprint

import (
	"golang.org/x/net/html"
	"strings"
)

// Prettify takes a string with unformatted HTML content and formats
// it prettily with indentation supplied through the indent argument.
//
// For example, it takes a string of html like this:
//
//     <html><head></head><body><p>Hello</p></body></html>
//
// and formats it into this:
//
// <html>
//     <head>
//     </head>
//     <body>
//         <a href="http://test.com">Test link</a>
//         <p>
//             <br/>
//         </p>
//     </body>
// </html>
//
// if the provided indent parameter is a string with four spaces ("    ").
//
// It returns the prettified string :pretty:, and possibly an error :e:
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
