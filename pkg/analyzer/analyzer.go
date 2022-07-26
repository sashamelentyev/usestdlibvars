package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// New returns new usestdlibvars analyzer.
func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "usestdlibvars",
		Doc:      "Detect the possibility to use constants/variables from the stdlib.",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
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
				basicLit := getBasicLit(v, 1, 0, token.INT)
				if basicLit == nil {
					return
				}

				checkHTTPStatusCode(pass, basicLit)

			case "NewRequest":
				basicLit := getBasicLit(v, 3, 0, token.STRING)
				if basicLit == nil {
					return
				}

				checkHTTPMethod(pass, basicLit)

			case "NewRequestWithContext":
				basicLit := getBasicLit(v, 4, 1, token.STRING)
				if basicLit == nil {
					return
				}

				checkHTTPMethod(pass, basicLit)
			}

		case *ast.BasicLit:
			currentVal := getBasicLiValue(v)

			checkTimeWeekday(pass, v.Pos(), currentVal)
			checkTimeMonth(pass, v.Pos(), currentVal)
			checkTimeLayout(pass, v.Pos(), currentVal)
			checkCryptoHash(pass, v.Pos(), currentVal)
		}
	})

	return nil, nil
}

func checkHTTPMethod(pass *analysis.Pass, basicLit *ast.BasicLit) {
	currentVal := getBasicLiValue(basicLit)

	newVal, ok := httpMethod[currentVal]
	if !ok {
		return
	}

	report(pass, basicLit.Pos(), newVal, currentVal)
}

func checkHTTPStatusCode(pass *analysis.Pass, basicLit *ast.BasicLit) {
	currentVal := getBasicLiValue(basicLit)

	newVal, ok := httpStatusCode[currentVal]
	if !ok {
		return
	}

	report(pass, basicLit.Pos(), newVal, currentVal)
}

func checkTimeWeekday(pass *analysis.Pass, pos token.Pos, currentVal string) {
	newVal, ok := timeWeekday[currentVal]
	if !ok {
		return
	}

	report(pass, pos, newVal, currentVal)
}

func checkTimeMonth(pass *analysis.Pass, pos token.Pos, currentVal string) {
	newVal, ok := timeMonth[currentVal]
	if !ok {
		return
	}

	report(pass, pos, newVal, currentVal)
}

func checkTimeLayout(pass *analysis.Pass, pos token.Pos, currentVal string) {
	newVal, ok := timeLayout[currentVal]
	if !ok {
		return
	}

	report(pass, pos, newVal, currentVal)
}

func checkCryptoHash(pass *analysis.Pass, pos token.Pos, currentVal string) {
	newVal, ok := cryptoHash[currentVal]
	if !ok {
		return
	}

	report(pass, pos, newVal, currentVal)
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

func getBasicLiValue(basicLit *ast.BasicLit) string {
	return strings.Trim(basicLit.Value, "\"")
}

func report(p *analysis.Pass, pos token.Pos, newVal, currentVal string) {
	p.Reportf(pos, `%q can be replaced by %s`, currentVal, newVal)
}
