package main

import (
	"flag"
	"log"
	"os"

	"github.com/uzudil/bscript/bscript"
)

func main() {
	var source, output string
	flag.StringVar(&source, "source", "", "the bscript file to run")
	flag.StringVar(&output, "output", "", "where to save compiled code")
	showAst := flag.Bool("ast", false, "print AST and not execute?")
	flag.Parse()

	if source != "" {
		if output != "" {
			if err := bscript.Compile(source, output); err != nil {
				log.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		} else {
			_, err := bscript.Run(source, *showAst, nil, nil)
			if err != nil {
				log.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		}
	} else {
		bscript.Repl(nil)
	}
}
