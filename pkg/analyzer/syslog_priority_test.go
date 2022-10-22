//go:build !windows && !plan9

package analyzer_test

func init() {
	pkgs = append(pkgs, "a/syslog")
}
