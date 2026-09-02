package main

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

const xmlExample = `<nesting>
   <parent>
     <child>
       <plant id="27">
         <name>Coffee</name>
         <origin>Ethiopia</origin>
         <origin>Brazil</origin>
       </plant>
       <plant id="81">
         <name>Tomato</name>
         <origin>Mexico</origin>
         <origin>California</origin>
       </plant>
       <rect
       style="fill:none;stroke:#25ff00;stroke-width:0"
       id="replace"
       width="302.36221"
       height="302.36221"
       x="128.50394"
       y="245.6693" />
     </child>
   </parent>
</nesting>`

func main() {
	svg := etree.NewDocument()
	if err := svg.ReadFromString(xmlExample); err != nil {
		panic(err)
	}
	svg.WriteTo(os.Stdout)
	fmt.Println()
	fmt.Println()
	fmt.Println("Afterwards:")
	fmt.Println()
	fmt.Println()

	// This is the know attribute inside the svg we will replace
	path := etree.MustCompilePath("//[@id='replace']")
	repl := svg.Root().FindElementsPath(path)
	for i, c := range repl {
		svg.Root().CreateChild(fmt.Sprintf("replaced%d", i), func(e *etree.Element) {
			e.CreateAttr("id", fmt.Sprintf("qr%d", i))
			e.CreateAttr("x", c.SelectAttrValue("x", "nil"))
			e.CreateAttr("y", c.SelectAttrValue("y", "nil"))
			e.CreateAttr("width", c.SelectAttrValue("width", "nil"))
			e.CreateAttr("height", c.SelectAttrValue("height", "nil"))
			e.CreateAttr("preserveAspectRatio", "none")
			// this Attr is incomplete, it need the actual image, but for simplicity
			// i will leave it as is
			e.CreateAttr("xlink:href", "data:image/jpeg;base64,")
		})
		// Get the `Parent()` of the current element, and remove its child at the
		// `Index()` of the current element
		c.Parent().RemoveChildAt(c.Index())
	}
	svg.WriteTo(os.Stdout)
}
