package analyzer

import (
	"go/ast"
	"go/token"
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

func run(pass *analysis.Pass) (interface{}, error) {
	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	writeHeaderCase(pass, i)
	stdlibVars(pass, i, _timeWeekdayVars, _timeMonthVars)
	return nil, nil
}

func writeHeaderCase(pass *analysis.Pass, i *inspector.Inspector) {
	filter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	i.Preorder(filter, func(node ast.Node) {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return
		}
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		if selectorExpr.Sel.Name != "WriteHeader" {
			return
		}
		if len(callExpr.Args) > 1 {
			return
		}
		basicLit, ok := callExpr.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}
		if basicLit.Kind != token.INT {
			return
		}
		oldVal := strings.Trim(basicLit.Value, "\"")
		newVal, ok := _httpStatusCodesVars[oldVal]
		if !ok {
			return
		}
		report(pass, basicLit.Pos(), newVal, oldVal)
	})
}

func stdlibVars(pass *analysis.Pass, i *inspector.Inspector, dicts ...map[string]string) {
	filter := []ast.Node{
		(*ast.BasicLit)(nil),
	}
	i.Preorder(filter, func(node ast.Node) {
		basicLit, ok := node.(*ast.BasicLit)
		if !ok {
			return
		}
		oldVal := strings.Trim(basicLit.Value, "\"")
		for _, dict := range dicts {
			newVal, ok := dict[oldVal]
			if !ok {
				continue
			}
			report(pass, basicLit.Pos(), newVal, oldVal)
		}
	})
}

func report(pass *analysis.Pass, pos token.Pos, newVal, oldVal string) {
	pass.Reportf(
		pos,
		`can use %s instead "%s"`,
		newVal,
		oldVal,
	)
}

var (
	_httpStatusCodesVars = map[string]string{
		"200": "http.StatusOK",
		"201": "http.StatusCreated",
		"204": "http.StatusNoContent",
		"400": "http.StatusBadRequest",
		"401": "http.StatusUnauthorized",
		"403": "http.StatusForbidden",
		"404": "http.StatusNotFound",
		"409": "http.StatusConflict",
		"500": "http.StatusInternalServerError",
	}
	_timeWeekdayVars = map[string]string{
		"Sunday":    "time.Sunday.String()",
		"Monday":    "time.Monday.String()",
		"Tuesday":   "time.Tuesday.String()",
		"Wednesday": "time.Wednesday.String()",
		"Thursday":  "time.Thursday.String()",
		"Friday":    "time.Friday.String()",
		"Saturday":  "time.Saturday.String()",
	}
	_timeMonthVars = map[string]string{
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
	}
)
