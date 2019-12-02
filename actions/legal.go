package actions

import "github.com/gobuffalo/buffalo"

// Legal default implementation.
func LegalLegal(c buffalo.Context) error {
	return c.Render(200, r.HTML("legal/legal.html"))
}

// LegalDataprotection default implementation.
func LegalDataprotection(c buffalo.Context) error {
	return c.Render(200, r.HTML("legal/dataprotection.html"))
}
