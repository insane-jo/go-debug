package common

import (
	"os"
	"strings"
	"github.com/ryanuber/go-glob"
)

var debugEnv string = ""

var patternsForDebug = make(map[string]bool)

func init() {
	debugEnv = os.Getenv("DEBUG")
	allPatterns := strings.Split(debugEnv, ",")

	for _, pattern := range allPatterns {
		if (pattern[0] == '-') {
			patternsForDebug[pattern[1:]] = false
		} else {
			patternsForDebug[pattern] = true
		}
	}
}

//Check if namespace is in allowed pattern and not in disallowed
func IsAllowed (namespace string) bool {
	isInAllowed := false

	for key, value := range patternsForDebug {
		namespaceIsInPattern := glob.Glob(key, namespace)
		if namespaceIsInPattern {
			if !value {
				return false
			} else {
				isInAllowed = true
			}
		}
	}

	return isInAllowed
}

func Enable(pattern string) {
	patternsForDebug[pattern] = true
}

func Disable(pattern string) {
	patternsForDebug[pattern] = false
}