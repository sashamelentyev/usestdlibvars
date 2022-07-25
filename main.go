package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/sashamelentyev/usestdlibvars/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.New())
}
