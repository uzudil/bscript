# bscript
A simple scripting language, similar to modern javascript.

Bscript features higher order functions, control flow commands, global variables and constants, etc. 

## Code examples
Browse the [tests](https://github.com/uzudil/bscript/tree/master/src/tests) for examples.

## Larger projects using bscript:
- [Benji4000](https://github.com/uzudil/benji4000)
- [Curse of Svaltfen](https://github.com/uzudil/svaltfen)
- [Enalim](https://github.com/uzudil/enalim)

## Language details

## Using bscript

You can run the interpreter via `./runner`, or run a source file via `./runner -source file.b`. Program execution starts by calling the `main()` function.

## Embedding bscript in Go code

Here is an example of starting the bscript interpreter:
```go
import "github.com/uzudil/bscript/bscript"

// app context: objects that will be available in the builtin functions
app := map[string]interface{}{
        "myObject": o,
        "myOtherObject": oo,
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
```

To call out to bscript from inside your code:

```go
import "github.com/uzudil/bscript/bscript"

// Compile the bscript library. Note that you still need 
// an empty 'main' method.
_, ctx, err := bscript.Build(
	libraryPath, // file with the 'myFunc' function
	false,
	map[string]interface{}{},
)
if err != nil {
	panic(err)
}

// Create the command that calls bscript
command := &bscript.Command{}
err = bscript.CommandParser.ParseString("myFunc();", command)
if err != nil {
	panic(err)
}
```

Once this is set up, you can call out to bscript (for example from main loop):
```go
command.Evaluate(ctx)
```

## bscript syntax highlighting
The vscode directory contains a plugin for syntax highlighting for .b files.

