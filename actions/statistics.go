package actions

import (
	"crowddefensewebsite/models"
	"encoding/json"
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// StatisticsHandler default implementation.
func StatisticsHandler(c buffalo.Context) error {
	g := &models.GameDataPoint{}

	if err := c.Bind(g); err != nil {
		return errors.WithStack(err)
	}

	val, err := redisClient.Get("game-currentversion").Result()
	if err != nil {
		log.Print(err)
	}

	g.Version = val

	tx := c.Value("tx").(*pop.Connection)
	
	_, err = g.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Print(c.Params().Get("Test"))

	test, _ := json.Marshal(struct{ success bool }{success: true})

	return c.Render(200, r.JSON(test))
}
