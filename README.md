# bscript
A simple scripting language, similar to modern javascript.

Bscript features higher order functions, control flow commands, global variables and constants, etc. 

Browse the [tests](https://github.com/gabor-lbl/benji4000/tree/master/src/tests) for examples or see the [wiki](https://github.com/uzudil/benji4000/wiki) for more info. (Builtins and constants are configurable per project.) For a larger body of bscript code, see the [Curse of Svaltfen](https://github.com/uzudil/svaltfen).

## Using bscript

You can run the interpreter via `./runner`, or run a source file via `./runner -source file.b`. Program execution starts by calling the `main()` function.

## Embedding bscript in your Go code

Here is an example of starting the bscript interpreter:
```go
import "github.com/uzudil/bscript/bscript"

app := map[string]interface{}{
        "video": video,
        ...
}

// add some builtin functions and constants
bscript.AddBuiltin("doSomething", myGoFunction)
bscript.AddConstant("SPECIAL_VALUE", 42)

// source can be a file, or a directory (in which case every file there is loaded)
source := "./mycode/test.b"

// show the AST (useful for debugging)
showAst := false

// run the source file(s). Execution starts at main().
bscript.Run(source, showAst, nil, app)

// or, start the interpreter
bscript.Repl(app)
```

## bscript syntax highlighting
The vscode directory contains a plugin for syntax highlighting for .b files.
