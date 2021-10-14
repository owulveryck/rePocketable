package epub

import (
	"html/template"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/vincent-petithory/dataurl"
)

type metaStruct struct {
	Title     string
	Author    string
	Website   string
	Build     string
	Published string
	Modified  string
	Summary   string
	Tags      []string
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
			<th>Published time</th>
			<td>{{ .Published }}</td>
		</tr>
		<tr>
			<th>Modified time</th>
			<td>{{ .Modified }}</td>
		</tr>
		<tr>
			<th>Summary</th>
			<td>{{ .Summary }}</td>
		</tr>
		<tr>
			<th>Tags</th>
			<td>{{ .Tags }}</td>
		</tr>
		<tr>
			<th>Epub build time</th>
			<td>{{ .Build }}</td>
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
	dataurl.EncodeBytes([]byte(cssMeta))
	//defer os.Remove(file.Name())
	css, err := d.Epub.AddCSS(dataurl.EncodeBytes([]byte(cssMeta)), uuid.Must(uuid.NewV4()).String()+".css")
	if err != nil {
		return err
	}
	var metaInfo strings.Builder
	mi := metaStruct{
		Author:  d.Author(),
		Title:   d.Title(),
		Summary: d.Description(),
		Build:   time.Now().Format("2006-02-01 15:04:05"),
		Website: d.item.ResolvedURL,
	}
	if d.OG != nil && d.OG.Article != nil {
		if d.OG.Article.PublishedTime != nil {
			mi.Published = d.OG.Article.PublishedTime.Format("2006-02-01 15:04:05")
		}
		if d.OG.Article.ModifiedTime != nil {
			mi.Modified = d.OG.Article.ModifiedTime.Format("2006-02-01 15:04:05")
		}
		mi.Tags = d.OG.Article.Tags
	}
	metaTmpl.Execute(&metaInfo, mi)
	_, err = d.AddSection(metaInfo.String(), "meta", "", css)
	return err
}
