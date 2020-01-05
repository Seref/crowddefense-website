package actions

import "github.com/gobuffalo/buffalo"

// LegalDataprotection default implementation.
func LegalDataProtection(c buffalo.Context) error {
	return c.Render(200, r.HTML("legal/dataprotection.html"))
}

// LegalContactinformation default implementation.
func LegalContactInformation(c buffalo.Context) error {
	return c.Render(200, r.HTML("legal/contactinformation.html"))
}

