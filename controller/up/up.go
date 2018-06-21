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
	if c.Sess.Values["id"] != nil {
		v.Vars["first_name"] = c.Sess.Values["first_name"]
	}

	v.Render(w, r)
}
