package actions

import "github.com/gobuffalo/buffalo"

// Forum default implementation.
func Forum(c buffalo.Context) error {
	return c.Render(200, r.HTML("forum/forum.html"))
}