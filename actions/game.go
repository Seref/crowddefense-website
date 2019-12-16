package actions

import "github.com/gobuffalo/buffalo"

// Game default implementation.
func Game(c buffalo.Context) error {
	return c.Render(200, r.HTML("game/game.html"))
}
