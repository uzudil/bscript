# bscript
A simple scripting language, similar to modern javascript. It can be used stand-alone or called from Go.

## Code examples
Browse the [tests](https://github.com/uzudil/bscript/tree/master/src/tests) for examples.

## Larger projects using bscript:
- [Benji4000](https://github.com/uzudil/benji4000)
- [Curse of Svaltfen](https://github.com/uzudil/svaltfen)
- [Enalim](https://github.com/uzudil/enalim)

## Language details

### Literals

- Basic types: numbers, strings, boolean, null. All numbers are floats but are cast to int if needed (array lookups, pixel coordinates, etc.) Literal strings are double-quoted and can span multiple lines.
- Arrays: `[1, "a"]` Arrays can contain any other type. To append to an array, add an element to an index past the length of the array.
- Maps: `{ "a": 1, "b": "c" }` or `{ a: 1, b: "c" }` Map keys are always strings. Map elements can be looked up by the `map["a"]` or `map.a` notation.
- Functions: `x => x + 1` Functions are first-class types.
- Numerical expressions support the usual arithmetic operations, plus `%` for modulo.

### Constants and variables

- Constants can only be declared outside of functions: `const PI = 3.14159;`
- Variables can be global (outside of functions), local (in a function) or function parameters: `x := 42;`. Variable scope starts in the local function and travels out toward the global scope.
- Variable assignment can also add, subtract, multiply, divide the current value. For example ` x :+ 10;` adds 10 to the current value.

### Functions
- Define named functions via `def square(n) { return n * n; }`.
- Function can take other functions as a parameter, or return a function.
- Function parameters can have default values, for example: `def add(a, b=1)`. You could call this via: `add(1,2)` or `add(1)`.
- Functions can be declared in another function.
- Anonymous functions are declared via: `x => x + 1` or `(x,y) => { return x + y; }`
- Functions are evaluated in the context of their closure. For example:
```
def x() {
   localVar := 1;
   return b => localVar + b;
}

fx := x();
print(fx(5));
```
will print `6`.
- map values can be functions that take a special first argument (usually called "self" or "this") that points to the map that contains them. For example: 
```
player := {
   isAlive: self => self.lives > 0,
   lives: 3,
};
print(player.isAlive());
```

### Control flow commands
- `if(x) { ... } if else(y) { ... } else { ... }` if statement works like you think it does. The `else` clauses are optional.
- BScript only has one type of loop: `while(x < 10) { ... }`
- Both `if` and `while` support the following conditional operations: `=`, `!=`, `<`, `<=`, `>`, `>=`.
- More complex boolean logic can be expressed with and: `&&` and or `||` operators. (These use short-circuit evaluation.)

### Other built-in commands
- To delete from an array or map, use `del a[10]`. In this case it would delete the 10th element in the array. The array/map reference can be arbitrarily complex. For example: `del game.enemies()[4].inventory[5]`
- `# this is a comment` Comments are allowed anywhere except inside literals (this may change soon.)

### Program execution
- Code execution starts by running the `main` method.
- Your program can be a single .b file, or a single directory containing any number of .b files.
- In the latter case, every file in the directory is parsed and joined to produce the program. 

### Built-in functions
Bscript code can use the following built-in functions. Their implementation falls into two categories:

These [standard functions](https://github.com/uzudil/bscript/blob/main/bscript/stdlib.go) are implemented in bscript:
- `array_map`: map array elements
- `array_join`: array to string
- `array_filter`: filter array elements
- `array_find`: find an array element
- `array_find_index`: find the index of an array element
- `array_foreach`: loop through an array
- `choose`: randomly select from an array
- `basic_sort`: sort an array in place
- `basic_sort_copy`: sort an array and return a new sorted array
- `array_reverse`: reverse an array
- `array_reduce`: turn an array into a single value
- `array_remove`: remove array elements
- `copy_array`: shallow clone an array
- `array_times`: fill an array
- `array_concat`: concat two arrays
- `array_flatten`: flatten an array of arrays
- `copy_map`: shallow clone a map
- `roll`: dice roll
- `normalize`: normalize a value to -1,0,1
- `endsWith`: check if a string ends with a suffix
- `startsWith`: check if a string starts with a prefix
- `asPercent`: turn a 0-1 value into an int percent
- `range`: loop through numbers
- `pow`: raise a number to a power

These [external functions](https://github.com/uzudil/bscript/blob/main/bscript/builtins.go) are implemented in go:
- `print`: print to console, example `print("x=" + x);`
- `input`: console input
- `len`: string and array length
- `keys`: get a map's keys
- `substr`: get a part of a string
- `split`: split a string
- `replace`: replace a part of a string
- `debug`: display the stack and variables
- `assert`: assertion for testing
- `random`: a random value from 0 - 1
- `int`: convert the number to int (floor)
- `round`: round a number to int
- `abs`: absolute value
- `min`: minimum of two values
- `max`: maximum of two values
- `typeof`: return a string that is the type of a variable
- `exit`: quit to console

## Building the bscript runner

`go build -o runner`

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

## Text editor support
The vscode directory contains a plugin for syntax highlighting for .b files. A more comprehensive Language Server version of the extension is planned.

