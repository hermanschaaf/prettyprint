package prettyprint

import "testing"

func TestPrettyPrint(t *testing.T) {
	html := `<html><head></head><body><a href="http://test.com">Test link</a><p><br/></p></body></html>`
	expected := `<html>
    <head>
    </head>
    <body>
        <a href="http://test.com">Test link</a>
        <p>
            <br/>
        </p>
    </body>
</html>`
	pretty, _ := Prettify(html, "    ")
	if pretty != expected {
		t.Errorf("Expected:\n\n%s\n\nbut got:\n\n%s", expected, pretty)
	}
}

func TestPrettyPrintWithNewlines(t *testing.T) {
	html := `<!doctype html><html><head>
<title>Website Title</title>
</head><body>
<div class="random-class">
<h1>I like pie</h1><p>It's true!</p></div>
</body></html>`
	expected := `<!doctype html>
<html>
    <head>
        <title>Website Title</title>
    </head>
    <body>
        <div class="random-class">
            <h1>I like pie</h1>
            <p>It's true!</p>
        </div>
    </body>
</html>`
	pretty, _ := Prettify(html, "    ")
	if pretty != expected {
		t.Errorf("Expected:\n\n%s\n\nbut got:\n\n%s", expected, pretty)
	}
}
