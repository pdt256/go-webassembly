package jdoc

import (
	"fmt"
	"syscall/js"
)

type document struct {
	doc js.Value
}

func New() *document {
	return &document{
		doc: js.Global().Get("document"),
	}
}

func (d *document) Print(message string) {
	d.doc.Call("write", "<pre>"+message+"</pre>")
}

func (d *document) Printf(format string, v ...interface{}) {
	d.Print(fmt.Sprintf(format, v...))
}
