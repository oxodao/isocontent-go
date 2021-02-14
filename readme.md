# Isocontent

Isocontent is a simple library that lets use normalize the content you store in your database to be easily displayed wherever you want.

Since an example is worth a thousand words...

Let's say you're building a website, this website has a WYSIWYG for comments. How do you handle that ? With Isocontent you can simply get the HTML from the user, normalize it to store in the database, and do the reverse operation to show it back again.

Why the extra hassle ? 

Doing so will enable you to display it on different places: 
- Maybe you want to put it on your website
- Maybe your website is a React app: this library will generate a JSON representation to send through your API and display it automatically
- Then you're building a mobile app: this library will handle the translation in React Native components easily stylable

Adding other renderer is quite an easy task, and the added benefit is that Isocontent works in a whitelist-based fashion: any tags not handled (such as script tags) will not be rendered nor stored in database, this add an extra layer of security for your inputs.

## Getting started - Golang

Parsing HTML and generating a JSON representation
```go
import (
    "fmt"
    isogo "github.com/oxodao/isocontent-go"
)

func main() {
    isocontent := isogo.New()
    parsedContent, err := isocontent.Parse(`<p>This is a sample text</p>`, "html")
    if err != nil {
        panic(err)
    }
    
    jsonContent, err := isocontent.Render(parsedContent, "json")
    if err != nil {
        panic(err)
    }

    // Result: [{"type":"block","block_type":"paragraph","children":[{"type":"text","value":"This is a sample text"}]}]
    fmt.Println(jsonContent)
}
```

### Currently supported parsers / renderers

#### Parsers
- html

#### Renderers
- html
- json

### Implementing your own parser / renderer

You need to create a struct that implements one of the following interface:
```go
type Parser interface {
	Parse(*builder.Builder, interface{}) error
	SupportsFormat(string) bool
}


type Renderer interface {
	Render([]AST.Node) (interface{}, error)
	SupportsFormat(string) bool
}
```
You can use the already implemented one as an example to help you: [Parser](https://github.com/oxodao/isocontent-go/blob/master/parser/DOMParser.go) / [Renderer](https://github.com/oxodao/isocontent-go/blob/master/renderer/JsonRenderer.go)
Then you simply register it 
```go
isocontent := isogo.New()
isocontent.RegisterParser(myParser{})
isocontent.RegisterRenderer(myRenderer{})

isocontent.Parse(input, "my_parser")
isocontent.Render(input, "my_renderer")
```

## Libraries available

Here are the currently available libraries that takes care of Isocontent

- [Isocontent](https://github.com/un-zero-un/Isocontent) - PHP library with Symfony / Doctrine bridges
- [isocontent-go](https://github.com/oxodao/isocontent-go) - Golang port of the PHP library
- [isocontent-js](https://github.com/un-zero-un/isocontent-js/tree/master/packages/isocontent) - Vanilla Javascript library
- [react-isocontent-dom](https://github.com/un-zero-un/isocontent-js/tree/master/packages/react-isocontent-dom) - React library for in-browser / electron usage
- [react-isocontent-native](https://github.com/un-zero-un/isocontent-js/tree/master/packages/react-isocontent-native) - React Native library for mobile app development
