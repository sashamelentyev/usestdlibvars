package analyzer

import (
	"flag"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	TimeWeekdayFlag    = "time-weekday"
	TimeMonthFlag      = "time-month"
	TimeLayoutFlag     = "time-layout"
	CryptoHashFlag     = "crypto-hash"
	HTTPMethodFlag     = "http-method"
	HTTPStatusCodeFlag = "http-status-code"
)

// New returns new usestdlibvars analyzer.
func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "usestdlibvars",
		Doc:      "A linter that detect the possibility to use variables/constants from the Go standard library.",
		Run:      run,
		Flags:    flags(),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func flags() flag.FlagSet {
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.Bool(HTTPMethodFlag, true, "suggest the use of http.MethodXX")
	flags.Bool(HTTPStatusCodeFlag, true, "suggest the use of http.StatusXX")
	flags.Bool(TimeWeekdayFlag, false, "suggest the use of time.Weekday")
	flags.Bool(TimeMonthFlag, false, "suggest the use of time.Month")
	flags.Bool(TimeLayoutFlag, false, "suggest the use of time.Layout")
	flags.Bool(CryptoHashFlag, false, "suggest the use of crypto.Hash")
	return *flags
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	filter := []ast.Node{
		(*ast.BasicLit)(nil),
		(*ast.CallExpr)(nil),
	}

	insp.Preorder(filter, func(n ast.Node) {
		switch v := n.(type) {
		case *ast.CallExpr:
			selectorExpr, ok := v.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}

			switch selectorExpr.Sel.Name {
			case "WriteHeader":
				if !lookupFlag(pass, HTTPStatusCodeFlag) {
					return
				}

				basicLit := getBasicLit(v, 1, 0, token.INT)
				if basicLit == nil {
					return
				}

				checkHTTPStatusCode(pass, basicLit)

			case "NewRequest":
				if !lookupFlag(pass, HTTPMethodFlag) {
					return
				}

				basicLit := getBasicLit(v, 3, 0, token.STRING)
				if basicLit == nil {
					return
				}

				checkHTTPMethod(pass, basicLit)

			case "NewRequestWithContext":
				if !lookupFlag(pass, HTTPMethodFlag) {
					return
				}

				basicLit := getBasicLit(v, 4, 1, token.STRING)
				if basicLit == nil {
					return
				}

				checkHTTPMethod(pass, basicLit)
			}

		case *ast.BasicLit:
			currentVal := getBasicLitValue(v)

			if lookupFlag(pass, TimeWeekdayFlag) {
				checkTimeWeekday(pass, v.Pos(), currentVal)
			}

			if lookupFlag(pass, TimeMonthFlag) {
				checkTimeMonth(pass, v.Pos(), currentVal)
			}

			if lookupFlag(pass, TimeLayoutFlag) {
				checkTimeLayout(pass, v.Pos(), currentVal)
			}

			if lookupFlag(pass, CryptoHashFlag) {
				checkCryptoHash(pass, v.Pos(), currentVal)
			}
		}
	})

	return nil, nil
}

func lookupFlag(pass *analysis.Pass, name string) bool {
	return pass.Analyzer.Flags.Lookup(name).Value.(flag.Getter).Get().(bool)
}

func checkHTTPMethod(pass *analysis.Pass, basicLit *ast.BasicLit) {
	currentVal := getBasicLitValue(basicLit)

	newVal, ok := httpMethod[currentVal]
	if ok {
		report(pass, basicLit.Pos(), newVal, currentVal)
	}
}

func checkHTTPStatusCode(pass *analysis.Pass, basicLit *ast.BasicLit) {
	currentVal := getBasicLitValue(basicLit)

	newVal, ok := httpStatusCode[currentVal]
	if ok {
		report(pass, basicLit.Pos(), newVal, currentVal)
	}
}

func checkTimeWeekday(pass *analysis.Pass, pos token.Pos, currentVal string) {
	newVal, ok := timeWeekday[currentVal]
	if ok {
		report(pass, pos, newVal, currentVal)
	}
}

func checkTimeMonth(pass *analysis.Pass, pos token.Pos, currentVal string) {
	newVal, ok := timeMonth[currentVal]
	if ok {
		report(pass, pos, newVal, currentVal)
	}
}

func checkTimeLayout(pass *analysis.Pass, pos token.Pos, currentVal string) {
	newVal, ok := timeLayout[currentVal]
	if ok {
		report(pass, pos, newVal, currentVal)
	}
}

func checkCryptoHash(pass *analysis.Pass, pos token.Pos, currentVal string) {
	newVal, ok := cryptoHash[currentVal]
	if ok {
		report(pass, pos, newVal, currentVal)
	}
}

// getBasicLit gets the *ast.BasicLit of a function argument.
//
// - count: expected number of argument in function
// - idx: index of the argument to get the *ast.BasicLit
// - typ: argument type
func getBasicLit(ce *ast.CallExpr, count, idx int, typ token.Token) *ast.BasicLit {
	if len(ce.Args) != count {
		return nil
	}

	basicLit, ok := ce.Args[idx].(*ast.BasicLit)
	if !ok {
		return nil
	}

	if basicLit.Kind != typ {
		return nil
	}

	return basicLit
}

func getBasicLitValue(basicLit *ast.BasicLit) string {
	return strings.Trim(basicLit.Value, "\"")
}

func report(p *analysis.Pass, pos token.Pos, newVal, currentVal string) {
	p.Reportf(pos, `%q can be replaced by %s`, currentVal, newVal)
}
