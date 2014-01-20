package prettyprint

import "testing"

func TestPrettyPrint(t *testing.T) {
	html := `<html><head></head><body><a href="http://test.com">Test link</a><p><br/></p></body></html>`
	expected := `<html>
    <head>
    </head>
    <body>
        <a href="http://test.com">
            Test link
        </a>
        <p>
            <br/>
        </p>
    </body>
</html>`
	pretty, _ := PrettyPrint(html, "    ")
	if pretty != expected {
		t.Errorf("Expected:\n\n%s\n\nbut got:\n\n%s", expected, pretty)
	}
}
