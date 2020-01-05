package actions

import "github.com/gobuffalo/buffalo"

// HistoryView default implementation.
func HistoryView(c buffalo.Context) error {
	return c.Render(200, r.HTML("history/view.html"))
}

