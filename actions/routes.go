package actions

import "github.com/gobuffalo/buffalo"

// Routes is a default handler to serve up
// a home page.
func Routes(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
