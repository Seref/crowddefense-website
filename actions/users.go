package actions

import (
	"crowddefensewebsite/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

//UsersNew renders the users form
func UsersNew(c buffalo.Context) error {
	u := models.User{}
	c.Set("user", u)
	return c.Render(200, r.HTML("users/new.plush.html"))
}

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
		return c.Render(200, r.HTML("users/new.plush.html"))
	}

	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome to Crowddefense!")

	return c.Redirect(302, "/")
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Session().Set("redirectURL", c.Request().URL.String())

			err := c.Session().Save()
			if err != nil {
				return errors.WithStack(err)
			}

			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}

// Delete User
func AccountDeletionHandler(c buffalo.Context) error {

	currentUser := c.Value("current_user").(*models.User)

	redisClient.SAdd("to_be_deleted", currentUser.Username)

	// for _, statement := range []string{"state-gameplay", "state-gamelooks", "state-voting", "state-website"} {
	// 	val, err := redisClient.Get(statement).Result()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	c.Set(statement, val)
	// }

	return c.Render(200, r.Plain("ok"))
}

func AccountUndeletionHandler(c buffalo.Context) error {

	currentUser := c.Value("current_user").(*models.User)

	redisClient.SRem("to_be_deleted", currentUser.Username)

	// for _, statement := range []string{"state-gameplay", "state-gamelooks", "state-voting", "state-website"} {
	// 	val, err := redisClient.Get(statement).Result()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	c.Set(statement, val)
	// }

	return c.Render(200, r.Plain("ok"))
}
