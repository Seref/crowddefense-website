package actions

import (
	"crowddefensewebsite/models"
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// GameGameHandler default implementation.
func GameHandler(c buffalo.Context) error {

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	pq := &models.Prequestionnaire{}

	err := tx.Where("username = ?", strings.ToLower(strings.TrimSpace(c.Value("current_user").(*models.User).Username))).First(pq)

	if err != nil {
		log.Print(err)
		return c.Redirect(302, "/prequestionnaires/new")
	}

	return c.Render(200, r.HTML("game/game.html"))
}
