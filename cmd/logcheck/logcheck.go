package main

import (
	"github.com/s-khechnev/logcheck/pkg/golinters/logcheck"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(logcheck.Analyzers...)
}
