package main

import (
	"flag"

	"github.com/uzudil/bscript/bscript"
)

func main() {
	var source string
	flag.StringVar(&source, "source", "", "the bscript file to run")
	showAst := flag.Bool("ast", false, "print AST and not execute?")
	flag.Parse()

	if source != "" {
		bscript.Run(source, showAst, nil)
	} else {
		bscript.Repl()
	}

}
