package actions

import (
	"crowddefensewebsite/models"
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
)

// HistoryView default implementation.
func HistoryView(c buffalo.Context) error {

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

	ideas := &models.Ideas{}

	done := tx.Where("fullfilled = true")

	if err := done.All(ideas); err != nil {
		return err
	}

	// for _, idea := range *ideas {
	// 	idea.ID.String() = fmt.Sprint(idea.ID)
	// }

	c.Set("ideasdone", ideas)
	c.Set("uuid_to_string", func(u uuid.UUID) string { return fmt.Sprint(u) })

	return c.Render(200, r.HTML("history/view.html"))
}
