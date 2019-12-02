package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"crowddefensewebsite/models"
)

// Forum default implementation.
func Forum(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)

	// q := tx.PaginateFromParams(c.Params())

	// c.Set("pagination", q.Paginator)

	suggestions := []models.Suggestion{}

	err := tx.All(&suggestions)

	if err != nil {
		fmt.Print("ERROR!\n")
		fmt.Printf("%v\n", err)
	} else {
		c.Set("suggestions", suggestions)
	}

	return c.Render(200, r.HTML("forum/forum.html"))
}
