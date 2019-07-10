package main

import "github.com/musenwill/experimentgo/sample"

var (
	// Version should be a git tag
	Version = ""
	// BuildTime in example: 2019-06-29T23:23:09+0800
	BuildTime = ""
)

// check Version and BuildTime
func init() {
	if len(Version) <= 0 {
		panic("Version unset, expect be set with a git tag")
	}
	if len(BuildTime) <= 0 {
		panic("BuildTime unset, expect to be like '2019-06-29T23:23:09+0800'")
	}
}

func main() {
	sample.Once()
}
