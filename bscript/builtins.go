package bscript

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func floatArgs(ctx *Context, count int, arg []interface{}) ([]float64, error) {
	if len(arg) < count {
		return nil, fmt.Errorf("%s Wrong number of arguments. Got %d instead of %d", ctx.Pos, count, len(arg))
	}
	r := make([]float64, count)
	for index, a := range arg[0:count] {
		f, ok := a.(float64)
		if !ok {
			return nil, fmt.Errorf("%s Argument %d should be a number (%v)", ctx.Pos, index, a)
		}
		r[index] = f
	}
	return r, nil
}

func intArgs(ctx *Context, count int, arg []interface{}) ([]int, error) {
	f, err := floatArgs(ctx, count, arg)
	if err != nil {
		return nil, err
	}
	r := make([]int, count)
	for index, value := range f {
		r[index] = int(value)
	}
	return r, nil
}

func stringArgs(ctx *Context, count int, arg []interface{}) ([]string, error) {
	if len(arg) < count {
		return nil, fmt.Errorf("%s Wrong number of arguments. Got %d instead of %d", ctx.Pos, count, len(arg))
	}
	r := make([]string, count)
	for index, a := range arg[0:count] {
		s, ok := a.(string)
		if !ok {
			return nil, fmt.Errorf("%s Argument %d should be a string (%v)", ctx.Pos, index, s)
		}
		r[index] = s
	}
	return r, nil
}

func print(ctx *Context, arg ...interface{}) (interface{}, error) {
	log.Println(EvalString(arg[0]))
	return nil, nil
}

func input(ctx *Context, arg ...interface{}) (interface{}, error) {
	log.Print(EvalString(arg[0]))
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return strings.TrimRight(text, "\n"), nil
}

func length(ctx *Context, arg ...interface{}) (interface{}, error) {
	a, ok := arg[0].(*[]interface{})
	if !ok {
		s, ok := arg[0].(string)
		if !ok {
			return nil, fmt.Errorf("%s argument to len() should be an array or a string", ctx.Pos)
		}
		return float64(len(s)), nil
	}
	return float64(len(*a)), nil
}

func split(ctx *Context, arg ...interface{}) (interface{}, error) {
	s, ok := arg[0].(string)
	if !ok {
		return nil, fmt.Errorf("%s argument 1 should be a string", ctx.Pos)
	}
	d, ok := arg[1].(string)
	if !ok {
		return nil, fmt.Errorf("%s argument 2 should be a string", ctx.Pos)
	}
	// a := strings.Split(s, d)
	a := regexp.MustCompile(d).Split(s, -1)
	arr := make([]interface{}, len(a))
	for i, aa := range a {
		arr[i] = aa
	}
	return &arr, nil
}

func substr(ctx *Context, arg ...interface{}) (interface{}, error) {
	s, ok := arg[0].(string)
	if !ok {
		return nil, fmt.Errorf("%s argument 1 to substr() should be a string", ctx.Pos)
	}
	index, ok := arg[1].(float64)
	if !ok {
		return nil, fmt.Errorf("%s argument 2 to substr() should be a number", ctx.Pos)
	}
	length := len(s)
	if len(arg) > 2 {
		f, ok := arg[2].(float64)
		if !ok {
			return nil, fmt.Errorf("%s argument 3 to substr() should be a number", ctx.Pos)
		}
		length = int(f)
	}
	start := int(math.Min(math.Max(index, 0), float64(len(s))))
	end := int(math.Min(math.Max(float64(start+length), 0), float64(len(s))))
	return string(s[start:end]), nil
}

func replace(ctx *Context, arg ...interface{}) (interface{}, error) {
	s, err := stringArgs(ctx, 3, arg)
	if err != nil {
		return nil, err
	}
	return strings.ReplaceAll(s[0], s[1], s[2]), nil
}

func keys(ctx *Context, arg ...interface{}) (interface{}, error) {
	m, ok := arg[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("argument to key() should be a map")
	}
	keys := make([]interface{}, len(m))
	index := 0
	for k := range m {
		keys[index] = k
		index++
	}
	return &keys, nil
}

func random(ctx *Context, arg ...interface{}) (interface{}, error) {
	return rand.Float64(), nil
}

func debug(ctx *Context, arg ...interface{}) (interface{}, error) {
	message, ok := arg[0].(string)
	if !ok {
		return nil, fmt.Errorf("argument to debug() should be a string")
	}
	ctx.debug(message)
	return nil, nil
}

func toAbs(ctx *Context, arg ...interface{}) (interface{}, error) {
	n, ok := arg[0].(float64)
	if !ok {
		return nil, fmt.Errorf("First argument should be a number")
	}
	return math.Abs(n), nil
}

func toMax(ctx *Context, arg ...interface{}) (interface{}, error) {
	f, err := floatArgs(ctx, 2, arg)
	if err != nil {
		return nil, err
	}
	return math.Max(f[0], f[1]), nil
}

func toMin(ctx *Context, arg ...interface{}) (interface{}, error) {
	f, err := floatArgs(ctx, 2, arg)
	if err != nil {
		return nil, err
	}
	return math.Min(f[0], f[1]), nil
}

func toInt(ctx *Context, arg ...interface{}) (interface{}, error) {
	n, ok := arg[0].(float64)
	if !ok {
		s, ok := arg[0].(string)
		if !ok {
			return nil, fmt.Errorf("First argument should be a number or a string")
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			i = 0
		}
		return float64(i), nil
	}
	return float64(int(n)), nil
}

func toRound(ctx *Context, arg ...interface{}) (interface{}, error) {
	n, ok := arg[0].(float64)
	if !ok {
		return nil, fmt.Errorf("First argument should be a number")
	}
	return float64(int(n)), nil
}

func typeof(ctx *Context, arg ...interface{}) (interface{}, error) {
	if arg[0] == nil {
		return "null", nil
	}
	_, ok := arg[0].(float64)
	if ok {
		return "number", nil
	}
	_, ok = arg[0].(string)
	if ok {
		return "string", nil
	}
	_, ok = arg[0].(bool)
	if ok {
		return "boolean", nil
	}
	_, ok = arg[0].(*Closure)
	if ok {
		return "function", nil
	}
	_, ok = arg[0].(map[string]interface{})
	if ok {
		return "map", nil
	}
	_, ok = arg[0].(*[]interface{})
	if ok {
		return "array", nil
	}
	return nil, fmt.Errorf("%s Unknown variable type", ctx.Pos)
}

func exit(ctx *Context, arg ...interface{}) (interface{}, error) {
	os.Exit(0)
	return nil, nil
}

func assert(ctx *Context, arg ...interface{}) (interface{}, error) {
	a := arg[0]
	b := arg[1]
	msg := "Incorrect value"
	if len(arg) > 2 {
		msg = arg[2].(string)
	}

	var res bool

	// for arrays, compare the values
	arr, ok := a.(*[]interface{})
	if ok {
		// array
		brr, ok := b.(*[]interface{})
		if !ok {
			res = true
		} else {
			if len(*arr) == len(*brr) {
				res = false
				for i := range *arr {
					if (*arr)[i] != (*brr)[i] {
						res = true
						break
					}
				}
			} else {
				res = true
			}
		}
	} else {
		// map
		amap, ok := a.(map[string]interface{})
		if ok {
			bmap, ok := b.(map[string]interface{})
			if !ok {
				res = true
			} else {
				if len(amap) == len(bmap) {
					res = false
					for k := range amap {
						if amap[k] != bmap[k] {
							res = true
							break
						}
					}
				} else {
					res = true
				}
			}
		} else {
			// default is to compare equality
			res = a != b
		}
	}

	if res {
		debug(ctx, fmt.Sprintf("Assertion failure: %s: %v != %v", msg, a, b))
		return nil, fmt.Errorf("%s Assertion failure: %s: %v != %v", ctx.Pos, msg, a, b)
	}
	return nil, nil
}

var builtins map[string]Builtin = map[string]Builtin{
	"print":   print,
	"input":   input,
	"len":     length,
	"keys":    keys,
	"substr":  substr,
	"split":   split,
	"replace": replace,
	"debug":   debug,
	"assert":  assert,
	"random":  random,
	"int":     toInt,
	"round":   toRound,
	"abs":     toAbs,
	"min":     toMin,
	"max":     toMax,
	"typeof":  typeof,
	"exit":    exit,
}

func Builtins() map[string]Builtin {
	return builtins
}

func AddBuiltin(name string, fx Builtin) {
	builtins[name] = fx
}

var constants map[string]interface{} = map[string]interface{}{}

func Constants() map[string]interface{} {
	return constants
}

func AddConstant(name string, constant interface{}) {
	constants[name] = constant
}
