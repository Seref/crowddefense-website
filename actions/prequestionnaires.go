package actions

import (
	"crowddefensewebsite/models"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Prequestionnaire)
// DB Table: Plural (prequestionnaires)
// Resource: Plural (Prequestionnaires)
// Path: Plural (/prequestionnaires)
// View Template Folder: Plural (/templates/prequestionnaires/)

// PrequestionnairesResource is the resource for the Prequestionnaire model
type PrequestionnairesResource struct {
	buffalo.Resource
}

// List gets all Prequestionnaires. This function is mapped to the path
// GET /prequestionnaires
func (v PrequestionnairesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	prequestionnaires := &models.Prequestionnaires{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Prequestionnaires from the DB
	if err := q.All(prequestionnaires); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, prequestionnaires))
}

// Show gets the data for one Prequestionnaire. This function is mapped to
// the path GET /prequestionnaires/{prequestionnaire_id}
func (v PrequestionnairesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Prequestionnaire
	prequestionnaire := &models.Prequestionnaire{}

	// To find the Prequestionnaire the parameter prequestionnaire_id is used.
	if err := tx.Find(prequestionnaire, c.Param("prequestionnaire_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, prequestionnaire))
}

// New renders the form for creating a new Prequestionnaire.
// This function is mapped to the path GET /prequestionnaires/new
func (v PrequestionnairesResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Prequestionnaire{}))
}

// Create adds a Prequestionnaire to the DB. This function is mapped to the
// path POST /prequestionnaires
func (v PrequestionnairesResource) Create(c buffalo.Context) error {
	// Allocate an empty Prequestionnaire
	prequestionnaire := &models.Prequestionnaire{}

	// Bind prequestionnaire to the html form elements
	if err := c.Bind(prequestionnaire); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(prequestionnaire)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		fmt.Printf("%+v", verrs.Errors)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, prequestionnaire))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "prequestionnaire.created.success"))
	// and redirect to the prequestionnaires index page
	return c.Render(201, r.Auto(c, prequestionnaire))
}

// Edit renders a edit form for a Prequestionnaire. This function is
// mapped to the path GET /prequestionnaires/{prequestionnaire_id}/edit
func (v PrequestionnairesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Prequestionnaire
	prequestionnaire := &models.Prequestionnaire{}

	if err := tx.Find(prequestionnaire, c.Param("prequestionnaire_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, prequestionnaire))
}

// Update changes a Prequestionnaire in the DB. This function is mapped to
// the path PUT /prequestionnaires/{prequestionnaire_id}
func (v PrequestionnairesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Prequestionnaire
	prequestionnaire := &models.Prequestionnaire{}

	if err := tx.Find(prequestionnaire, c.Param("prequestionnaire_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Prequestionnaire to the html form elements
	if err := c.Bind(prequestionnaire); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(prequestionnaire)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, prequestionnaire))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "prequestionnaire.updated.success"))
	// and redirect to the prequestionnaires index page
	return c.Render(200, r.Auto(c, prequestionnaire))
}

// Destroy deletes a Prequestionnaire from the DB. This function is mapped
// to the path DELETE /prequestionnaires/{prequestionnaire_id}
func (v PrequestionnairesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Prequestionnaire
	prequestionnaire := &models.Prequestionnaire{}

	// To find the Prequestionnaire the parameter prequestionnaire_id is used.
	if err := tx.Find(prequestionnaire, c.Param("prequestionnaire_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(prequestionnaire); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "prequestionnaire.destroyed.success"))
	// Redirect to the prequestionnaires index page
	return c.Render(200, r.Auto(c, prequestionnaire))
}