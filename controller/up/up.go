// Package up
package up

import (
	"net/http"

	"github.com/arapov/cheert/lib/flight"

	"github.com/arapov/core/router"
)

// Load the routes.
func Load() {
	router.Post("/up", Index)
}

// Index displays the home page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	v := c.View.New("up/index")
	v.Vars["identity"] = r.FormValue("identity")
	v.Vars["uid"] = r.FormValue("uid")
	v.Vars["group"] = r.FormValue("group")

	v.Render(w, r)
}
