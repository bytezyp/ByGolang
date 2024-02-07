package main

import (
	"fmt"
	"github.com/beevik/etree"
)

func main() {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)

	people := doc.CreateElement("people")
	people.CreateComment("these are all known people")

	jon := people.CreateElement("people")
	jon.CreateAttr("name", "jon")

	sally := people.CreateElement("people")
	sally.CreateAttr("name", "jon")

	doc.Indent(2)
	//doc.WriteTo(os.Stdout)
	bytes, _ := doc.WriteToBytes()
	fmt.Println(bytes)
}
