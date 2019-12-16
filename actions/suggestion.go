package actions

import (
	"crowddefensewebsite/models"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// SuggestionUpvote default implementation.
func SuggestionUpvote(c buffalo.Context) error {
	user := *c.Value("current_user").(*models.User)

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Suggestion
	suggestion := &models.Suggestion{}

	if err := tx.Find(suggestion, c.Param("suggestion_id")); err != nil {
		return c.Error(404, err)
	}

	_, alreadyUpvoted := suggestion.UpvotedBy[user.Username]
	_, alreadyDownvoted := suggestion.DownvotedBy[user.Username]

	if alreadyUpvoted {
		c.Flash().Add("success", T.Translate(c, "suggestion.upvoted.failure"))
		return c.Render(200, r.HTML("introduction.html"))
	}

	if alreadyDownvoted {
		delete(suggestion.DownvotedBy, user.Username)
	}

	suggestion.UpvotedBy[user.Username] = user.Username

	_, err := tx.ValidateAndUpdate(suggestion)
	if err != nil {
		return err
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "suggestion.upvoted.success"))
	// and redirect to the suggestions index page

	return c.Render(200, r.HTML("introduction.html"))
}

// SuggestionDownvote default implementation.
func SuggestionDownvote(c buffalo.Context) error {
	user := *c.Value("current_user").(*models.User)

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Suggestion
	suggestion := &models.Suggestion{}

	if err := tx.Find(suggestion, c.Param("suggestion_id")); err != nil {
		return c.Error(404, err)
	}

	_, alreadyUpvoted := suggestion.UpvotedBy[user.Username]
	_, alreadyDownvoted := suggestion.DownvotedBy[user.Username]

	if alreadyDownvoted {
		c.Flash().Add("success", T.Translate(c, "suggestion.downvoted.failure"))
		return c.Render(200, r.HTML("introduction.html"))
	}

	if alreadyUpvoted {
		delete(suggestion.UpvotedBy, user.Username)
	}

	suggestion.DownvotedBy[user.Username] = user.Username

	_, err := tx.ValidateAndUpdate(suggestion)
	if err != nil {
		return err
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "suggestion.downvoted.success"))
	// and redirect to the suggestions index page

	return c.Render(200, r.HTML("introduction.html"))
}
