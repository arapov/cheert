// Package up
package up

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/arapov/cheert/lib/flight"
	"github.com/arapov/cheert/model/up"

	"github.com/arapov/core/router"
)

// Load the routes.
func Load() {
	router.Post("/up", Index)
	router.Post("/up/submit", Submit)
	router.Get("/up/submit", End)
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

	r.ParseForm()
	for k, v := range r.Form {
		log.Println(k, ">", v)
	}

	for i, value := range r.Form["uid"] {
		//log.Println(value, " praise ", r.Form["praise"][i])
		//log.Println(value, " plus ", r.Form["plus["+value+"]"])

		_, err := up.Create(c.DB, value, r.Form["from"][0], r.Form["plus["+value+"]"][0], r.Form["praise"][i])
		if err != nil {
			c.FlashErrorGeneric(err)
			return
		}

	}

}

func End(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("up/submit")

	v.Render(w, r)
}
