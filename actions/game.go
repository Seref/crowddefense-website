package actions

import "github.com/gobuffalo/buffalo"

// GameGameHandler default implementation.
func GameHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("game/game.html"))
}

