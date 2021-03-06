package epub

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/mtex"
	"github.com/vincent-petithory/dataurl"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func renderTex(w io.Writer, expr string) error {
	defer func() {
		if r := recover(); r != nil {
			err := renderTexOnline(w, expr)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	fnts := liberationFonts()
	dst := drawimg.NewRenderer(w)
	err := mtex.Render(dst, "$"+expr+"$", size, dpi, fnts)
	return err
}

func renderTexOnline(w io.Writer, expr string) error {
	u, _ := url.Parse("https://latex.codecogs.com/png.latex")
	q := url.QueryEscape(expr)
	u.RawQuery = strings.Replace(q, "+", "%20", -1)
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad return code for LaTeX generation of %v: %v", expr, res.Status)

	}
	defer res.Body.Close()
	_, err = io.Copy(w, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func hasMathJax(n *html.Node) bool {
	begin := regexp.MustCompile(`(?m)^\\begin`)
	beginIndex := begin.FindStringIndex(n.Data)
	end := regexp.MustCompile(`(?m)^\\end.*`)
	endIndices := end.FindAllStringIndex(n.Data, -1)
	if len(endIndices) > 0 {
		endIndex := endIndices[len(endIndices)-1]
		if len(beginIndex) != 0 && len(endIndex) != 0 {
			return true
		}
	}
	return len(mathJax.FindAllString(n.Data, -1)) > 0
}

func processMathTex(n *html.Node, inline bool) error {
	var currentFormula []byte
	completeText := n.Data
	var allFormulas [][]byte
	if inline {
		allFormulas = mathJax.FindAll([]byte(completeText), -1)
	} else {
		allFormulas = extractTex(completeText)
	}
	images := make([]*html.Node, len(allFormulas))
	var i int
	var remaining string
	var delete bool
	for i, currentFormula = range allFormulas {
		expr := strings.TrimFunc(string(currentFormula), func(r rune) bool {
			return r == '$'
		})
		var b bytes.Buffer
		err := renderTex(&b, expr)
		if err != nil {
			return err
		}
		if b.Bytes() == nil {
			return errors.New("no content")
		}
		dataURL := dataurl.New(b.Bytes(), "image/png")
		content, err := dataURL.MarshalText()
		if err != nil {
			return err
		}
		images[i] = &html.Node{
			Type:      html.ElementNode,
			DataAtom:  atom.Img,
			Data:      "img",
			Namespace: n.Namespace,
			Attr: []html.Attribute{
				{
					Key: "src",
					Val: string(content),
				},
				{
					Key: "style",
					Val: "width:0.77em; height:0.77em;",
				},
			},
		}
		delete = true
		if remaining == "" {
			remaining = n.Data
		}
		data := strings.SplitN(remaining, string(currentFormula), 2)
		if len(data) > 1 {
			remaining = data[1]
		}
		firstpart := &html.Node{
			Type:      n.Type,
			DataAtom:  n.DataAtom,
			Data:      data[0],
			Namespace: n.Namespace,
			Attr:      n.Attr,
		}
		n.Parent.InsertBefore(&html.Node{
			Type:      html.CommentNode,
			DataAtom:  atom.Content,
			Data:      expr,
			Namespace: n.Namespace,
		}, n)
		n.Parent.InsertBefore(images[i], n)
		n.Parent.InsertBefore(firstpart, images[i])
	}
	if remaining != "" {
		n.Parent.InsertBefore(&html.Node{
			Type:      n.Type,
			DataAtom:  n.DataAtom,
			Data:      remaining,
			Namespace: n.Namespace,
			Attr:      n.Attr,
		}, n)
	}
	if delete {
		n.Data = ""
	}
	return nil
}

func extractTex(input string) [][]byte {
	scanner := bufio.NewScanner(strings.NewReader(input))
	beginEquation := []byte(`\begin`)
	endEquation := []byte(`\end`)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if len(data) >= len(beginEquation) {
			if string(data[0:len(beginEquation)]) == string(beginEquation) {
				return len(beginEquation), beginEquation, nil
			}
		}
		if len(data) >= len(endEquation) {
			if string(data[0:len(endEquation)]) == string(endEquation) {
				var b strings.Builder
				var i int
				for i = 0; i < len(data); i++ {
					b.WriteByte(data[i])
					if rune(data[i]) == '}' {
						break
					}
				}
				if i == len(data) {
					return b.Len(), []byte(b.String()), io.EOF
				}
				return b.Len(), []byte(b.String()), nil
			}
		}
		return 1, data[0:1], nil
	}
	scanner.Split(split)
	inEquation := 0
	var b strings.Builder
	output := make([][]byte, 0)
	for scanner.Scan() {
		current := scanner.Text()
		if current == string(beginEquation) {
			inEquation++
		}
		if inEquation > 0 {
			b.WriteString(current)
		}
		if strings.HasPrefix(current, string(endEquation)) {
			inEquation--
		}
		if inEquation == 0 && b.Len() > 0 {
			output = append(output, []byte(b.String()))
			b.Reset()
		}
	}

	return output
}
