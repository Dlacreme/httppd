// Package authlevel adds an AuthLevel variable to the view template.
package authlevel

import (
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/view"
)

// Modify sets AuthLevel in the template to auth if the user is authenticated.
// Sets AuthLevel to anon if not authenticated.
func Modify(w http.ResponseWriter, r *http.Request, v *view.Info) {
	c := flight.Context(w, r)

	// Set the AuthLevel to auth if the user is logged in
	if c.Sess.Values["UserId"] != nil {
		v.Vars["AuthLevel"] = "auth"
	} else {
		v.Vars["AuthLevel"] = "anon"
	}
}
