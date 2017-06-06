package wiki

import (
	"github.com/jg-l/wiki/db"
	"github.com/russross/blackfriday"
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
	"time"
)

// Show is the show endpoint of the Wiki
func (w *Wiki) All(c web.C, rw http.ResponseWriter, r *http.Request) {

	w.DB().View(func(tx *db.Tx) error {
		b := tx.Bucket([]byte("pages"))

		gg := b.Cursor()

		for k, v := gg.First(); k != nil; k, v = gg.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil

		if err != nil {
			p.Text = []byte(emptyPageString)
			//p.Created = []byte(y)
		}

		vars := map[string]interface{}{
			"Path":     "/" + string(p.Name) + "/edit",
			"Name":     BytesAsHTML(p.Name),
			"Text":     BytesAsHTML(ParsedMarkdown(p.Text)),
			"Created":  BytesAsTime(p.Created),
			"Modified": BytesAsTime(p.Modified),
		}

		t := template.Must(template.New("show").Parse(showTpl))
		t.Execute(rw, vars)

		return nil
	})
}

// RedirectToShow redirects to the show endpoint using a HTTP 302
func (w *Wiki) RedirectToShow(c web.C, rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/"+c.URLParams["name"], 302)
}

func BytesAsTime(b []byte) *time.Time {
	v := new(time.Time)
	v.UnmarshalBinary(b)
	return v
}

// BytesAsHTML returns the template bytes as HTML
func BytesAsHTML(b []byte) template.HTML {
	return template.HTML(string(b))
}

// ParsedMarkdown returns provided bytes parsed as Markdown
func ParsedMarkdown(b []byte) []byte {
	return blackfriday.MarkdownCommon(b)
}
