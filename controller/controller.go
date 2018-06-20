// Package controller loads the routes for each of the controllers.
package controller

import (
	"github.com/arapov/cheert/controller/about"
	"github.com/arapov/cheert/controller/home"
	"github.com/arapov/cheert/controller/static"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	//	debug.Load()
	//	register.Load()
	//	login.Load()
	home.Load()
	static.Load()
	//	status.Load()
	//	notepad.Load()
}
