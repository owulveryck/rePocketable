package epub

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
	"time"
)

type metaStruct struct {
	Title     string
	Author    string
	Website   string
	Build     string
	Published string
	Summary   string
}

var metaTmpl = template.Must(template.New("meta").Parse(meta))

const (
	meta = `
<div class="container">
<table class="table">
	<tbody>
		<tr>
			<th>Title</th>
			<td>{{ .Title }}</td>
		</tr>
		<tr>
			<th>Author</th>
			<td>{{ .Author }}</td>
		</tr>
		<tr>
			<th>Original</th>
			<td><a href="{{ .Website }}">{{ .Website }}</a></td>
		</tr>
		<tr>
			<th>Build time</th>
			<td>{{ .Build }}</td>
		</tr>
		<tr>
			<th>Published time</th>
			<td>{{ .Published }}</td>
		</tr>
		<tr>
			<th>Summary</th>
			<td>{{ .Summary }}</td>
		</tr>
	</tbody>
</table>
</div>
`

	cssMeta = `
.container {
	width: 100%;
}

.table {
	border: 1px solid;
	border-radius: 5px;
	width: 95%;
	margin: 0px auto;
	float: none;
}`
)

func (d *Document) createMeta() error {
	file, err := ioutil.TempFile("", "mystyle*.css")
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprint(file, cssMeta)
	//defer os.Remove(file.Name())
	css, err := d.Epub.AddCSS(file.Name(), "")
	if err != nil {
		return err
	}
	var metaInfo strings.Builder
	metaTmpl.Execute(&metaInfo, metaStruct{
		Author:  d.Author(),
		Title:   d.Title(),
		Summary: d.Description(),
		Build:   time.Now().Format("2006-02-01 15:04:05"),
		Website: d.item.ResolvedURL,
	})
	_, err = d.AddSection(metaInfo.String(), "meta", "", css)
	return err
}
