package main

import (
	"os"

	"github.com/bootlab-dev/llm100x-tester/internal/stages"
	tester_utils "github.com/bootlab-dev/tester-utils"
)

func main() {
	definition := stages.GetDefinition()
	os.Exit(tester_utils.Run(os.Args[1:], definition))
}
