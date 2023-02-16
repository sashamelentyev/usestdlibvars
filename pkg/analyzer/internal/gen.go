//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"embed"
	"go/format"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/sashamelentyev/usestdlibvars/pkg/analyzer/internal/mapping"
)

//go:embed template/*
var templateDir embed.FS

func main() {
	t := template.Must(
		template.New("template").
			Funcs(map[string]any{
				"quoteMeta": regexp.QuoteMeta,
				"toLower":   strings.ToLower,
			}).
			ParseFS(templateDir, "template/*.tmpl"),
	)

	operations := []struct {
		mapping      map[string]string
		packageName  string
		templateName string
		fileName     string
	}{
		{
			mapping:      mapping.CryptoHash,
			packageName:  "crypto_test",
			templateName: "test-template.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/crypto/crypto.go",
		},
		{
			mapping:      mapping.HTTPMethod,
			packageName:  "http_test",
			templateName: "test-httpmethod.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/http/method.go",
		},
		{
			mapping:      mapping.HTTPStatusCode,
			packageName:  "http_test",
			templateName: "test-httpstatus.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/http/status.go",
		},
		{
			mapping:      mapping.RPCDefaultPath,
			packageName:  "rpc_test",
			templateName: "test-template.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/rpc/rpc.go",
		},
		{
			mapping:      mapping.TimeWeekday,
			packageName:  "time_test",
			templateName: "test-template.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/time/weekday.go",
		},
		{
			mapping:      mapping.TimeMonth,
			packageName:  "time_test",
			templateName: "test-template.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/time/month.go",
		},
		{
			mapping:      mapping.TimeLayout,
			packageName:  "time_test",
			templateName: "test-template.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/time/layout.go",
		},
		{
			mapping:      mapping.HTTPMethod,
			packageName:  "http_test",
			templateName: "test-issue32.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/http/issue32.go",
		},
		{
			mapping:      mapping.SQLIsolationLevel,
			packageName:  "sql_test",
			templateName: "test-template.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/sql/isolationlevel.go",
		},
		{
			mapping:      mapping.TLSSignatureScheme,
			packageName:  "tls_test",
			templateName: "test-template.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/tls/signaturescheme.go",
		},
		{
			mapping:      mapping.ConstantKind,
			packageName:  "constant_test",
			templateName: "test-template.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/constant/kind.go",
		},
		{
			mapping:      mapping.HTTPStatusCode,
			packageName:  "http_test",
			templateName: "test-issue89.go.tmpl",
			fileName:     "pkg/analyzer/testdata/src/a/http/issue89.go",
		},
	}

	for _, operation := range operations {
		data := map[string]any{
			"PackageName": operation.packageName,
			"Mapping":     operation.mapping,
		}

		if err := execute(t, operation.templateName, data, operation.fileName); err != nil {
			log.Fatal(err)
		}
	}
}

func execute(t *template.Template, templateName string, data any, fileName string) error {
	var builder bytes.Buffer

	if err := t.ExecuteTemplate(&builder, templateName, data); err != nil {
		return err
	}

	sourceData, err := format.Source(builder.Bytes())
	if err != nil {
		return err
	}

	if err := os.WriteFile(fileName, sourceData, os.ModePerm); err != nil {
		return err
	}

	return nil
}
