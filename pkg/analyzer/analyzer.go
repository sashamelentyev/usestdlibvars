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
	stdlibVars(pass, i,
		_timeWeekdayVars,
		_timeMonthVars,
		_timeParseLayoutVars,
		_crypto,
	)
	writeHeader(pass, i)
	newRequest(pass, i)
	newRequestWithContext(pass, i)
	return nil, nil
}

func stdlibVars(pass *analysis.Pass, i *inspector.Inspector, dictionaries ...map[string]string) {
	filter := []ast.Node{
		(*ast.BasicLit)(nil),
	}
	i.Preorder(filter, func(node ast.Node) {
		basicLit, ok := node.(*ast.BasicLit)
		if !ok {
			return
		}
		oldVal := strings.Trim(basicLit.Value, "\"")
		for _, dict := range dictionaries {
			newVal, ok := dict[oldVal]
			if !ok {
				continue
			}
			report(pass, basicLit.Pos(), newVal, oldVal)
		}
	})
}

func writeHeader(pass *analysis.Pass, i *inspector.Inspector) {
	callExpr(pass, i, func(p *analysis.Pass, ce *ast.CallExpr) {
		selectorExpr, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		if selectorExpr.Sel.Name != "WriteHeader" {
			return
		}
		if len(ce.Args) != 1 {
			return
		}
		basicLit, ok := ce.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}
		if basicLit.Kind != token.INT {
			return
		}
		oldVal := strings.Trim(basicLit.Value, "\"")
		newVal, ok := _httpStatusCodeVars[oldVal]
		if !ok {
			return
		}
		report(p, basicLit.Pos(), newVal, oldVal)
	})
}

func newRequest(pass *analysis.Pass, i *inspector.Inspector) {
	callExpr(pass, i, func(p *analysis.Pass, ce *ast.CallExpr) {
		selectorExpr, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		if selectorExpr.Sel.Name != "NewRequest" {
			return
		}
		if len(ce.Args) != 3 {
			return
		}
		basicLit, ok := ce.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}
		if basicLit.Kind != token.STRING {
			return
		}
		oldVal := strings.Trim(basicLit.Value, "\"")
		newVal, ok := _httpMethodVars[oldVal]
		if !ok {
			return
		}
		report(p, basicLit.Pos(), newVal, oldVal)
	})
}

func newRequestWithContext(pass *analysis.Pass, i *inspector.Inspector) {
	callExpr(pass, i, func(p *analysis.Pass, ce *ast.CallExpr) {
		selectorExpr, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		if selectorExpr.Sel.Name != "NewRequestWithContext" {
			return
		}
		if len(ce.Args) != 4 {
			return
		}
		basicLit, ok := ce.Args[1].(*ast.BasicLit)
		if !ok {
			return
		}
		if basicLit.Kind != token.STRING {
			return
		}
		oldVal := strings.Trim(basicLit.Value, "\"")
		newVal, ok := _httpMethodVars[oldVal]
		if !ok {
			return
		}
		report(p, basicLit.Pos(), newVal, oldVal)
	})
}

func callExpr(pass *analysis.Pass, i *inspector.Inspector, fn func(p *analysis.Pass, ce *ast.CallExpr)) {
	filter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	i.Preorder(filter, func(node ast.Node) {
		ce, ok := node.(*ast.CallExpr)
		if !ok {
			return
		}
		fn(pass, ce)
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
	_httpMethodVars = map[string]string{
		"GET":     "http.MethodGet",
		"HEAD":    "http.MethodHead",
		"POST":    "http.MethodPost",
		"PUT":     "http.MethodPut",
		"PATCH":   "http.MethodPatch",
		"DELETE":  "http.MethodDelete",
		"CONNECT": "http.MethodConnect",
		"OPTIONS": "http.MethodOptions",
		"TRACE":   "http.MethodTrace",
	}
	_httpStatusCodeVars = map[string]string{
		"100": "http.StatusContinue",
		"101": "http.StatusSwitchingProtocols",
		"102": "http.StatusProcessing",
		"103": "http.StatusEarlyHints",

		"200": "http.StatusOK",
		"201": "http.StatusCreated",
		"202": "http.StatusAccepted",
		"203": "http.StatusNonAuthoritativeInfo",
		"204": "http.StatusNoContent",
		"205": "http.StatusResetContent",
		"206": "http.StatusPartialContent",
		"207": "http.StatusMultiStatus",
		"208": "http.StatusAlreadyReported",
		"226": "http.StatusIMUsed",

		"300": "http.StatusMultipleChoices",
		"301": "http.StatusMovedPermanently",
		"302": "http.StatusFound",
		"303": "http.StatusSeeOther",
		"304": "http.StatusNotModified",
		"305": "http.StatusUseProxy",
		"307": "http.StatusTemporaryRedirect",
		"308": "http.StatusPermanentRedirect",

		"400": "http.StatusBadRequest",
		"401": "http.StatusUnauthorized",
		"402": "http.StatusPaymentRequired",
		"403": "http.StatusForbidden",
		"404": "http.StatusNotFound",
		"405": "http.StatusMethodNotAllowed",
		"406": "http.StatusNotAcceptable",
		"407": "http.StatusProxyAuthRequired",
		"408": "http.StatusRequestTimeout",
		"409": "http.StatusConflict",
		"410": "http.StatusGone",
		"411": "http.StatusLengthRequired",
		"412": "http.StatusPreconditionFailed",
		"413": "http.StatusRequestEntityTooLarge",
		"414": "http.StatusRequestURITooLong",
		"415": "http.StatusUnsupportedMediaType",
		"416": "http.StatusRequestedRangeNotSatisfiable",
		"417": "http.StatusExpectationFailed",
		"418": "http.StatusTeapot",
		"421": "http.StatusMisdirectedRequest",
		"422": "http.StatusUnprocessableEntity",
		"423": "http.StatusLocked",
		"424": "http.StatusFailedDependency",
		"425": "http.StatusTooEarly",
		"426": "http.StatusUpgradeRequired",
		"428": "http.StatusPreconditionRequired",
		"429": "http.StatusTooManyRequests",
		"431": "http.StatusRequestHeaderFieldsTooLarge",
		"451": "http.StatusUnavailableForLegalReasons",

		"500": "http.StatusInternalServerError",
		"501": "http.StatusNotImplemented",
		"502": "http.StatusBadGateway",
		"503": "http.StatusServiceUnavailable",
		"504": "http.StatusGatewayTimeout",
		"505": "http.StatusHTTPVersionNotSupported",
		"506": "http.StatusVariantAlsoNegotiates",
		"507": "http.StatusInsufficientStorage",
		"508": "http.StatusLoopDetected",
		"510": "http.StatusNotExtended",
		"511": "http.StatusNetworkAuthenticationRequired",
	}
)

var (
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

var (
	_crypto = map[string]string{
		"MD4":         "crypto.MD4.String()",
		"MD5":         "crypto.MD5.String()",
		"SHA-1":       "crypto.SHA1.String()",
		"SHA-224":     "crypto.SHA224.String()",
		"SHA-256":     "crypto.SHA256.String()",
		"SHA-384":     "crypto.SHA384.String()",
		"SHA-512":     "crypto.SHA512.String()",
		"MD5+SHA1":    "crypto.MD5SHA1.String()",
		"RIPEMD-160":  "crypto.RIPEMD160.String()",
		"SHA3-224":    "crypto.SHA3_224.String()",
		"SHA3-256":    "crypto.SHA3_256.String()",
		"SHA3-384":    "crypto.SHA3_384.String()",
		"SHA3-512":    "crypto.SHA3_512.String()",
		"SHA-512/224": "crypto.SHA512_224.String()",
		"SHA-512/256": "crypto.SHA512_256.String()",
		"BLAKE2s-256": "crypto.BLAKE2s_256.String()",
		"BLAKE2b-256": "crypto.BLAKE2b_256.String()",
		"BLAKE2b-384": "crypto.BLAKE2b_384.String()",
		"BLAKE2b-512": "crypto.BLAKE2b_512.String()",
	}
)
