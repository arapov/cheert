// Package up
package up

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/arapov/cheert/lib/flight"

	"github.com/arapov/core/router"
)

// Load the routes.
func Load() {
	router.Post("/up", Index)
	router.Post("/up/submit", Submit)
}

// Index displays the home page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("up/index")

	re, _ := regexp.Compile(`rhos-dfg-[a-z-]*`)
	dfgs := re.FindAllStringSubmatch(r.FormValue("group"), -1)

	v.Vars["group"] = ""
	if len(dfgs) > 0 {

		var dfgComma string
		for _, dfg := range dfgs {
			log.Println(dfg)
			squad := strings.Index(dfg[0], "-squad")
			if squad != -1 {
				dfgComma += fmt.Sprintf("%s,", dfg[0][0:squad])
			} else {
				dfgComma += fmt.Sprintf("%s,", dfg[0])
			}

		}
		dfgComma = dfgComma[0 : len(dfgComma)-1]
		v.Vars["group"] = dfgComma
	}

	v.Vars["identity"] = r.FormValue("identity")
	v.Vars["uid"] = r.FormValue("uid")

	v.Render(w, r)
}

func Submit(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("up/submit")

	r.ParseForm()
	log.Println(r.Form)

	v.Render(w, r)
}
