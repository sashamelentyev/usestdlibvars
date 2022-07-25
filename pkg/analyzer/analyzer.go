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
	stdlibVars(pass, i,
		_timeWeekdayVars,
		_timeMonthVars,
		_timeParseLayoutVars,
	)
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
	_timeParseLayoutVars = map[string]string{
		"01/02 03:04:05PM '06 -0700":          "time.Layout",
		"Mon Jan _2 15:04:05 2006":            "time.ANSIC",
		"Mon Jan _2 15:04:05 MST 2006":        "time.UnixDate",
		"Mon Jan 02 15:04:05 -0700 2006":      "time.RubyDate",
		"02 Jan 06 15:04 MST":                 "time.RFC822",
		"02 Jan 06 15:04 -0700":               "time.RFC822Z",
		"Monday, 02-Jan-06 15:04:05 MST":      "time.RFC850",
		"Mon, 02 Jan 2006 15:04:05 MST":       "time.RFC1123",
		"Mon, 02 Jan 2006 15:04:05 -0700":     "time.RFC1123Z",
		"2006-01-02T15:04:05Z07:00":           "time.RFC3339",
		"2006-01-02T15:04:05.999999999Z07:00": "time.RFC3339Nano",
		"3:04PM":                              "time.Kitchen",
		"Jan _2 15:04:05":                     "time.Stamp",
		"Jan _2 15:04:05.000":                 "time.StampMilli",
		"Jan _2 15:04:05.000000":              "time.StampMicro",
		"Jan _2 15:04:05.000000000":           "time.StampNano",
	}
)
