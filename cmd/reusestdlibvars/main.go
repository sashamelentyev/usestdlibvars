package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/sashamelentyev/reusestdlibvars/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.New())
}
