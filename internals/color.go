package internals

import "runtime"

var Reset = "\033[0m"
var Red = "\033[31m"
var Yellow = "\033[33m"
var Blue = "\033[34m"

func InitColors() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Yellow = ""
		Blue = ""
	}
}
