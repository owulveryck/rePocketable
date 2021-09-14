package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"
	"time"

	"github.com/owulveryck/rePocketable/internal/db"
	"github.com/owulveryck/rePocketable/internal/pocket"
)

func initDatabase(signal chan os.Signal, cancelFunc context.CancelFunc, dstDir string) *db.Database {
	dbPath := filepath.Join(dstDir, ".db")
	database := db.NewDatabase()

	go func() {
		<-signal
		log.Println("bailing out")
		var content bytes.Buffer
		err := database.Write(&content)
		if err != nil {
			log.Println(err)
		}
		err = os.WriteFile(dbPath, content.Bytes(), 0766)
		if err != nil {
			log.Println(err)
		}
		log.Println("database written!")
		cancelFunc()
	}()
	f, err := os.Open(dbPath)
	if err != nil {
		log.Println(err)
		return database
	}
	err = database.Read(f)
	if err != nil {
		log.Println("Cannot initiate database: ", err)
	}
	f.Close()
	return database

}

func dumpContent(w io.Writer, dbPath string) error {
	database := db.NewDatabase()
	f, err := os.Open(dbPath)
	if err != nil {
		return err
	}
	err = database.Read(f)
	if err != nil {
		return err
	}
	f.Close()
	tabs := tabwriter.NewWriter(w, 1, 0, 4, ' ', 0)
	fmt.Fprintln(tabs, "ID\tTitle\tTime Added\tSource")
	database.Range(func(key interface{}, value interface{}) bool {
		item := value.(pocket.Item)
		fmt.Fprintf(tabs, "%v\t%v\t%v\t%v\n", item.ItemID, item.ResolvedTitle, time.Time(item.TimeAdded), item.ResolvedURL)
		return true
	})
	tabs.Flush()
	return nil
}

func dumpHTMLContent(w io.Writer, dbPath string) error {
	database := db.NewDatabase()
	f, err := os.Open(dbPath)
	if err != nil {
		return err
	}
	err = database.Read(f)
	if err != nil {
		return err
	}
	f.Close()
	fmt.Fprintln(w, `<!DOCTYPE html>
	<html lang="en-US">
	<head>
	    <title>rePocketable</title>
	    <meta charset="UTF-8">

	    <style type="text/css">
	    table, th, td {
		border: 1px solid black;
		table-layout: auto;
		width: 100%;
	    }
	    th {
		cursor: pointer;
	    }
	    </style>
	</head>
	<body>
	    <div>
		<table>
		    <thead>
			<tr>
			    <th>ID</th>
			    <th>Title</th>
			    <th>Time Added</th>
			    <th>Excerpt</th>
			    <th>Source</th>
			    <th>word count</th>
			</tr>
		    </thead>
		    <tbody>
	`)
	database.Range(func(key interface{}, value interface{}) bool {
		item := value.(pocket.Item)
		fmt.Fprintf(w, `<tr>
		<td><a href="%v.epub">%v</a></td>
		<td>%v</td>
		<td>%v</td>
		<td>%v</td>
		<td style="word-wrap:break-word;"><a href="%v">%v</a></td>
		<td>%v</td>
		</tr>`, item.ItemID,
			item.ItemID,
			item.ResolvedTitle,
			time.Time(item.TimeAdded).Format("2006/01/02"),
			item.Excerpt,
			item.ResolvedURL,
			item.ResolvedURL,
			item.WordCount)
		fmt.Fprintln(w, "")
		return true
	})
	fmt.Fprintln(w, `</tbody>
        </table>
    </div>
<script>
const getCellValue = (tr, idx) => tr.children[idx].innerText || tr.children[idx].textContent;

const comparer = (idx, asc) => (a, b) => ((v1, v2) => 
    v1 !== '' && v2 !== '' && !isNaN(v1) && !isNaN(v2) ? v1 - v2 : v1.toString().localeCompare(v2)
    )(getCellValue(asc ? a : b, idx), getCellValue(asc ? b : a, idx));

// do the work...
document.querySelectorAll('th').forEach(th => th.addEventListener('click', (() => {
    const table = th.closest('table');
    Array.from(table.querySelectorAll('tr:nth-child(n+2)'))
        .sort(comparer(Array.from(th.parentNode.children).indexOf(th), this.asc = !this.asc))
        .forEach(tr => table.appendChild(tr) );
})));
</script> 
</body>
</html>`)
	return nil
}
