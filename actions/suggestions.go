package actions

import (
	"crowddefensewebsite/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// SuggestionsNew default implementation.
func SuggestionsNew(c buffalo.Context) error {
	s := models.Suggestion{}

	c.Set("suggestion", s)

	return c.Render(200, r.HTML("forum/submitidea.html"))
}

// SuggestionsCreate default implementation.
func SuggestionsCreate(c buffalo.Context) error {
	s := &models.Suggestion{}

	if err := c.Bind(s); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := s.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("suggestion", s)
		c.Set("errors", verrs)
		return c.Render(200, r.HTML("forum/submitidea.plush.html"))
	}

	c.Flash().Add("success", "Suggestion submitted")

	return c.Redirect(302, "/forum")
}
