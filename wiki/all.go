package wiki

import (
	"fmt"
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

	})
}
