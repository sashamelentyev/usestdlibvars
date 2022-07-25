package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// New returns new reusestdlibvars analyzer.
func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "reusestdlibvars",
		Doc:      "Detect possibility reuse variables from stdlib",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	imports := make(map[string]struct{})
	i.Preorder([]ast.Node{
		(*ast.ImportSpec)(nil),
	}, func(node ast.Node) {
		importSpec, ok := node.(*ast.ImportSpec)
		if !ok {
			return
		}
		imp := strings.Trim(importSpec.Path.Value, "\"")
		imports[imp] = struct{}{}
	})
	i.Preorder([]ast.Node{
		(*ast.BasicLit)(nil),
	}, func(node ast.Node) {
		basicLit, ok := node.(*ast.BasicLit)
		if !ok {
			return
		}
		key := strings.Trim(basicLit.Value, "\"")
		for imp := range imports {
			val, ok := _reuseStdlibVars[imp][key]
			if !ok {
				continue
			}
			pass.Reportf(
				basicLit.Pos(),
				`can use %s instead "%s"`,
				val,
				key,
			)
		}
	})
	return nil, nil
}

var _reuseStdlibVars = map[string]map[string]string{
	"time": {
		"Sunday":    "time.Sunday.String()",
		"Monday":    "time.Monday.String()",
		"Tuesday":   "time.Tuesday.String()",
		"Wednesday": "time.Wednesday.String()",
		"Thursday":  "time.Thursday.String()",
		"Friday":    "time.Friday.String()",
		"Saturday":  "time.Saturday.String()",

		"January":   "time.January.String()",
		"February":  "time.February.String()",
		"March":     "time.March.String()",
		"April":     "time.April.String()",
		"May":       "time.May.String()",
		"June":      "time.June.String()",
		"July":      "time.July.String()",
		"August":    "time.August.String()",
		"September": "time.September.String()",
		"October":   "time.October.String()",
		"November":  "time.November.String()",
		"December":  "time.December.String()",
	},
	"net/http": {
		"200": "http.StatusOK",
		"201": "http.StatusCreated",
		"204": "http.StatusNoContent",
		"400": "http.StatusBadRequest",
		"401": "http.StatusUnauthorized",
		"403": "http.StatusForbidden",
		"404": "http.StatusNotFound",
		"409": "http.StatusConflict",
		"500": "http.StatusInternalServerError",
	},
}
