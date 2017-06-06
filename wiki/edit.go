package wiki

import (
	"net/http"
	"text/template"
	"time"

	"github.com/zenazn/goji/web"

	"wiki/db"
)

// Edit is the edit endpoint of the Wiki
func (w *Wiki) Edit(c web.C, rw http.ResponseWriter, r *http.Request) {
	name := w.getPageName(c.URLParams["name"])

	w.DB().Update(func(tx *db.Tx) error {
		p, _ := tx.Page(name)

		y := BytesAsTime(p.Created)
		u, errr := time.Now().MarshalBinary()
		if errr != nil {
			panic(errr)
		}

		if y.Year() == 1 {
			p.Created = []byte(u)
			p.Save()
		}

		t, err := template.New("edit").Parse(editTpl)

		template.Must(t, err).Execute(rw, map[string]string{
			"Path": "/" + string(p.Name),
			"Text": string(p.Text),
		})

		return nil
	})
}
