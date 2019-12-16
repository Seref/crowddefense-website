package actions

import (
	"crowddefensewebsite/models"
	"fmt"
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Suggestion)
// DB Table: Plural (suggestions)
// Resource: Plural (Suggestions)
// Path: Plural (/suggestions)
// View Template Folder: Plural (/templates/suggestions/)

// SuggestionsResource is the resource for the Suggestion model
type SuggestionsResource struct {
	buffalo.Resource
}

// List gets all Suggestions. This function is mapped to the path
// GET /suggestions
func (v SuggestionsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	for _, statement := range []string{"state-gameplay", "state-gamelooks", "state-voting", "state-website"} {
		val, err := redisClient.Get(statement).Result()
		if err != nil {
			log.Print(err)
		}
		c.Set(statement, val)
	}

	suggestions := &models.Suggestions{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	q.Order("JSON_LENGTH(upvoted_by) - JSON_LENGTH(downvoted_by) DESC")

	// Retrieve all Suggestions from the DB
	if err := q.All(suggestions); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, suggestions))
}

// Show gets the data for one Suggestion. This function is mapped to
// the path GET /suggestions/{suggestion_id}
func (v SuggestionsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Suggestion
	suggestion := &models.Suggestion{}

	// To find the Suggestion the parameter suggestion_id is used.
	if err := tx.Find(suggestion, c.Param("suggestion_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, suggestion))
}

// New renders the form for creating a new Suggestion.
// This function is mapped to the path GET /suggestions/new
func (v SuggestionsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Suggestion{}))
}

// Create adds a Suggestion to the DB. This function is mapped to the
// path POST /suggestions
func (v SuggestionsResource) Create(c buffalo.Context) error {
	// Allocate an empty Suggestion
	suggestion := &models.Suggestion{}

	// Bind suggestion to the html form elements
	if err := c.Bind(suggestion); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	suggestion.SuggestedBy = c.Value("current_user").(*models.User).Username

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(suggestion)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, suggestion))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "suggestion.created.success"))
	// and redirect to the suggestions index page
	return c.Render(201, r.Auto(c, suggestion))
}

// Edit renders a edit form for a Suggestion. This function is
// mapped to the path GET /suggestions/{suggestion_id}/edit
func (v SuggestionsResource) Edit(c buffalo.Context) error {
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

	return c.Render(200, r.Auto(c, suggestion))
}

// Update changes a Suggestion in the DB. This function is mapped to
// the path PUT /suggestions/{suggestion_id}
func (v SuggestionsResource) Update(c buffalo.Context) error {
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

	// Bind Suggestion to the html form elements
	if err := c.Bind(suggestion); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(suggestion)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, suggestion))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "suggestion.updated.success"))
	// and redirect to the suggestions index page
	return c.Render(200, r.Auto(c, suggestion))
}

// Upvote functionality
func (v SuggestionsResource) Upvote(c buffalo.Context) error {
	user := c.Value("current_user").(models.User)

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

	if alreadyUpvoted {
		/*TODO hadle this*/
	} else {
		suggestion.UpvotedBy[user.Username] = user.Username
	}

	// // Bind Suggestion to the html form elements
	// if err := c.Bind(suggestion); err != nil {
	// 	return err
	// }

	_, err := tx.ValidateAndUpdate(suggestion)
	if err != nil {
		return err
	}

	// if verrs.HasAny() {
	// 	// Make the errors available inside the html template
	// 	c.Set("errors", verrs)

	// 	// Render again the edit.html template that the user can
	// 	// correct the input.
	// 	return c.Render(422, r.Auto(c, suggestion))
	// }

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "suggestion.updated.success"))
	// and redirect to the suggestions index page
	return c.Render(200, r.Auto(c, suggestion))
}

// Destroy deletes a Suggestion from the DB. This function is mapped
// to the path DELETE /suggestions/{suggestion_id}
func (v SuggestionsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Suggestion
	suggestion := &models.Suggestion{}

	// To find the Suggestion the parameter suggestion_id is used.
	if err := tx.Find(suggestion, c.Param("suggestion_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(suggestion); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "suggestion.destroyed.success"))
	// Redirect to the suggestions index page
	return c.Render(200, r.Auto(c, suggestion))
}