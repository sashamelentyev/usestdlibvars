package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/sashamelentyev/usestdlibvars/pkg/analyzer"
)

//go:generate go run pkg/analyzer/internal/gen.go

func main() {
	singlechecker.Main(analyzer.New())
}
