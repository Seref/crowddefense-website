package actions

import "github.com/gobuffalo/buffalo"

//IntroductionHandler is a default handler to serve up the Crowdjump2 introduction page.
func Home(c buffalo.Context) error {
	return c.Render(200, r.HTML("introduction.html"))
}
