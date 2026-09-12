package main

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
)

// A simple svg with three shapes, two of which have the id=replace* as our
// target to replace, in place with the same dimensions
const source = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>

<svg
   width="210mm"
   height="297mm"
   viewBox="0 0 210 297"
   version="1.1"
   id="svg1"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:xlink="http://www.w3.org/1999/xlink"
   xmlns:svg="http://www.w3.org/2000/svg">
  <defs
     id="defs1" />
  <g
     id="layer1">
    <ellipse
       style="fill:#000000;stroke:#25ff00;stroke-width:4.99999"
       id="path1"
       cx="105"
       cy="65.001778"
       rx="48.971531"
       ry="33.998219" />
    <rect
       style="fill:#000000;stroke:#25ff00;stroke-width:4.99999"
       id="replace1"
       width="91.953735"
       height="81.736649"
       x="59.023132"
       y="169.46263" />
    <rect
       style="fill:#000000;stroke:#25ff00;stroke-width:4.99999"
       id="replace2"
       width="22.19573"
       height="24.309608"
       x="153.96085"
       y="113.79716" />
  </g>
</svg>`

// The image is normally readed from a file or base64 encoded from an in memory
// representation.
const img = `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&#10;`

func main() {
	// etree actually works with xml files, svg are xml files with specific
	// elements.
	svg := etree.NewDocument()

	// Read from either a string or a file normally.
	if err := svg.ReadFromString(source); err != nil {
		// Don't panic in actual projects, this is an example.
		panic(err)
	}

	// Save the svg to a file to be able to see the before/after replacement.
	svg.WriteToFile("before.svg")

	// Search for all elements
	path := etree.MustCompilePath("//*")
	repl := svg.FindElementsPath(path)

	// I manually copy paste the image string from a svg generated by inkscape
	// it added some newlines &#10 so i strip them here, if you actually read
	// an image from a file you would not need this
	cleanImg := strings.ReplaceAll(img, "&#10;", "")
	for i, c := range repl {
		if strings.HasPrefix(c.SelectAttrValue("id", "nil"), "replace") {
			c.Parent().CreateChild("image", func(e *etree.Element) {
				e.CreateAttr("id", fmt.Sprintf("qr%d", i))
				e.CreateAttr("x", c.SelectAttrValue("x", "nil"))
				e.CreateAttr("y", c.SelectAttrValue("y", "nil"))
				e.CreateAttr("width", c.SelectAttrValue("width", "nil"))
				e.CreateAttr("height", c.SelectAttrValue("height", "nil"))
				e.CreateAttr("preserveAspectRatio", "none")
				e.CreateAttr("xlink:href", "data:image/png;base64,"+cleanImg)
			})
			// Get the `Parent()` of the current element, and remove its child at the
			// `Index()` of the current element, which is to say, remove the current
			// element
			c.Parent().RemoveChildAt(c.Index())
		}
	}

	svg.WriteToFile("after.svg")
}
