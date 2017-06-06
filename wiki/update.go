package wiki

import (
	"net/http"

	"github.com/goji/param"
	"github.com/zenazn/goji/web"

	"fmt"
	"time"

	"github.com/jg-l/wiki/db"
)

type formData struct {
	Text string `param:"text"`
}

// Update is the update endpoint of the Wiki
func (w *Wiki) Update(c web.C, rw http.ResponseWriter, r *http.Request) {
	name := w.getPageName(c.URLParams["name"])

	// Parse the POST body
	r.ParseForm()

	var fd formData
	param.Parse(r.Form, &fd)

	w.DB().Update(func(tx *db.Tx) error {
		gg, e := tx.Page(name)
		if e != nil {
			panic(e)
		}

		fmt.Println(gg.Name)
		v := new(time.Time)
		v.UnmarshalBinary(gg.Created)
		fmt.Println(v)
		t, err := time.Now().MarshalBinary()
		fmt.Println(err)
		p := db.Page{Tx: tx, Name: name}
		if v == nil {
			p.Created = []byte(t)
		} else {
			p.Created = gg.Created
		}
		p.Text = []byte(fd.Text)
		p.Modified = []byte(t)
		return p.Save()
	})

	path := "/" + string(name)

	if path == "/root" {
		path = "/"
	}

	http.Redirect(rw, r, path, 302)
}
