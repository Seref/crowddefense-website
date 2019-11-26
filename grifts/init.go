package grifts

import (
	"github.com/medienprojekt-19-20/crowddefense-website/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
