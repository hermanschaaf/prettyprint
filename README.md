Go PrettyPrint HTML
===================

A simple pretty printer for HTML, written in Go.

#### Installation

```shell
go get github.com/hermanschaaf/prettyprint
```

#### Usage

```go
import (
    "github.com/hermanschaaf/prettyprint"
    "fmt"
)

func main() {
    html := `<html><head></head><body><a href="http://test.com">Test link</a><p><br/></p></body></html>`
    // second parameter is the indent to use 
    // (can be anything: spaces, a tab, etc.)
    pretty, err := prettyprint.Prettify(html, "    ")
    if err != nil {
        // do something in case of error
    }
    // print the prettified html:
    fmt.Println(pretty)
}
```

This will have the following output:

```html
<html>
    <head>
    </head>
    <body>
        <a href="http://test.com">Test link</a>
        <p>
            <br/>
        </p>
    </body>
</html>
```

#### Known issues

 - Poor support for self-closing tags, like HTML5 `<br>`. For now, only explicitly self-closing tags like `<br/>` will print correctly. 
 - No special cases for `<pre>` and `<code>` tags
 - Comments may or may not work, they are more or less untested
