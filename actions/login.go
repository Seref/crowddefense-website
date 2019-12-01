package actions

import "github.com/gobuffalo/buffalo"

// LoginHandler is a default handler to serve up
// a home page.
func Login(c buffalo.Context) error {
	return c.Render(200, r.HTML("auth/login.html"))
}
