package main

import (
	"fmt"
	"./common"
)

type LogFunction = func(...interface{}) (int, error)

var logFunctions = make(map[*LogFunction]LogFunction)

func dummyDebugFunction (_ ...interface{}) (int,error) { return 0, nil}

func SetLogFunction(debugFunc LogFunction, newLogFunction LogFunction) {
	logFunctions[&debugFunc] = newLogFunction
}

func Destroy(debugFunc LogFunction) {
	delete(logFunctions, &debugFunc)
}

func Debug(namespace string) LogFunction {
	if !common.IsAllowed(namespace) {
		return dummyDebugFunction
	} else {
		var debugFunc LogFunction

		debugFunc = func (args ...interface{}) (int, error) {
			args = append([]interface{}{namespace}, args...)

			logFunction := logFunctions[&debugFunc]
			if logFunction != nil {
				return logFunction(args...)
			} else {
				return 0, nil
			}
		}

		SetLogFunction(debugFunc, fmt.Println)

		return debugFunc
	}
}

func Enable(pattern string) {
	common.Enable(pattern)
}

func Disable(pattern string) {
	common.Disable(pattern)
}

func main () {
	debugFunc := Debug("debug:test:allow:any")
	debugFunc("sadsad", 12321)

	//Destroy(debugFunc)

	fmt.Println( "debug func", (&debugFunc) )
	debugFunc( "debug func", &debugFunc )

	debugFunc = Debug("debug:test:allow:any1")
	fmt.Println( "debug func", &debugFunc )
}