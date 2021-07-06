package docs

import (
	"bytes"
	"io"
	"text/tabwriter"

	"github.com/kelseyhightower/envconfig"
)

const (
	mdFormat = `| KEY	| TYPE	| DEFAULT	| REQUIRED	| DESCRIPTION	|
| 	| 	| 	| 	| 	|
{{range .}}| {{ usage_key . }}	| {{usage_type .}}	| {{usage_default .}}	| {{usage_required .}}	| {{usage_description .}}	|
{{end}}`
)

func Usage(prefix string, spec interface{}, output io.Writer) {
	var b bytes.Buffer
	tabs := tabwriter.NewWriter(&b, 1, 0, 4, ' ', 0)
	envconfig.Usagef(prefix, spec, tabs, mdFormat)
	tabs.Flush()
	// replace the second line's spaces by dash
	bytes := b.Bytes()
	lines := 0
	for i, b := range bytes {
		if b == '\n' {
			lines++
		}
		if lines == 1 && b == ' ' {
			bytes[i] = '-'
		}
	}
	io.Copy(output, &b)
}
