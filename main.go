package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/tidwall/pretty"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	prettyJSON := pretty.Pretty(data)

	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println(string(pretty.Color(prettyJSON, nil)))
	} else {
		fmt.Println(string(prettyJSON))
	}
}
