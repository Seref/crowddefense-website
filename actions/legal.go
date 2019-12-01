package actions

import "github.com/gobuffalo/buffalo"

// Legal default implementation.
func Legal(c buffalo.Context) error {
	return c.Render(200, r.HTML("legal/legal.html"))
}

