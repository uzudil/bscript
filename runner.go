package main

import (
	"flag"
	"os"

	"github.com/uzudil/bscript/bscript"
)

func main() {
	var source string
	flag.StringVar(&source, "source", "", "the bscript file to run")
	showAst := flag.Bool("ast", false, "print AST and not execute?")
	flag.Parse()

	if source != "" {
		_, err := bscript.Run(source, *showAst, nil, nil)
		if err != nil {
			os.Exit(1)
		}
	} else {
		bscript.Repl(nil)
	}
}
