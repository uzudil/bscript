package bscript

import (
	"fmt"
	"os"
	"strings"
)

const syntaxError = "?Syntax error"

func processCommand(ctx *Context, cmds string) (bool, error) {
	cmd := strings.Split(cmds, " ")
	switch {
	case cmd[0] == "exit":
		ctx.Builtins["print"](ctx, "Goodbye.")
		os.Exit(0)
		return true, nil
	case cmd[0] == "debug":
		ctx.Builtins["debug"](ctx, "State in repl:")
		return true, nil
	case cmd[0] == "run":
		var err error
		if len(cmd) > 1 {
			_, err = Run(cmd[1], false, ctx, ctx.App)
		} else if ctx.Program != nil {
			_, err = ctx.Program.Evaluate(ctx)
		} else {
			err = fmt.Errorf("No program loaded")
		}
		return true, err
	case cmd[0] == "load":
		_, err := Load(cmd[1], false, ctx)
		return true, err
	case cmd[0] == "help":
		ctx.Builtins["print"](ctx, "bscript Repl commands:")
		ctx.Builtins["print"](ctx, "exit - quit to shell")
		ctx.Builtins["print"](ctx, "run [<filename>] - if filename is given, load and run the program specified by filename. Without a filename: run program currently in memory.")
		ctx.Builtins["print"](ctx, "load <filename> - load the program specified by filename")
		ctx.Builtins["print"](ctx, "help - print this help")
		ctx.Builtins["print"](ctx, "debug - print stack and closures")
		return true, nil
	default:
		return false, nil
	}
}

// Repl is an interactive command interpreter
func Repl(app map[string]interface{}) {
	ctx := CreateContext(nil)
	ctx.App = app

	ctx.Builtins["print"](ctx, "")
	ctx.Builtins["print"](ctx, "     **** Benji4000 bscript v1 ****")
	ctx.Builtins["print"](ctx, "")
	for {
		ctx.Builtins["print"](ctx, "Ready.")
		command, err := ctx.Builtins["input"](ctx, "")
		if err != nil {
			ctx.Builtins["print"](ctx, syntaxError)
		}

		ast := &Command{}
		handled, err := processCommand(ctx, command.(string))
		if err != nil {
			ctx.Builtins["print"](ctx, fmt.Sprintf("%s: %s", syntaxError, err))
		} else if !handled {
			value, err := ParseString(command.(string), ast, ctx)
			if err != nil {
				ctx.Builtins["print"](ctx, fmt.Sprintf("%s: %s", syntaxError, err))
			}
			if value != nil {
				ctx.Builtins["print"](ctx, fmt.Sprintf("%v", value))
			}
		}

		ctx.Builtins["print"](ctx, "")
	}
}
