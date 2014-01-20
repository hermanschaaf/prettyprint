package prettyprint

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"strings"
)

func PrettyPrint(raw string, indent string) (pretty string, e error) {
	r := strings.NewReader(raw)
	z := html.NewTokenizer(r)
	pretty = ""
	depth := 0
	for {
		tt := z.Next()

		if tt == html.EndTagToken {
			depth -= 1
		}

		for i := 0; i < depth; i++ {
			pretty += indent
		}
		pretty += z.Token().String() + "\n"

		fmt.Println(tt)
		// last token
		if tt == html.ErrorToken {
			break
		} else if tt == html.StartTagToken {
			depth += 1
		}

	}
	return strings.Trim(pretty, "\n"), nil
}
