package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	"crowddefensewebsite/models"

	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	csrf "github.com/gobuffalo/mw-csrf"
	i18n "github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/packr/v2"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_crowddefensewebsite_session",
		})

		// Automatically redirect to SSL
		// app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		csrf := csrf.New
		app.Use(csrf)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)

		app.Middleware.Skip(Authorize, HomeHandler, LegalDataProtection, LegalContactInformation, StatisticsHandler)

		//AuthMiddlewares
		app.Use(SetCurrentUser)
		app.Use(Authorize)

		//Routes for Auth
		auth := app.Group("/auth")
		auth.GET("/", AuthLanding)
		auth.GET("/new", AuthNew)
		auth.POST("/", AuthCreate)
		auth.DELETE("/", AuthDestroy)

		//Routes for User registration
		users := app.Group("/users")
		users.GET("/new", UsersNew)
		users.POST("/", UsersCreate)
		users.Middleware.Remove(Authorize)


		app.Resource("/ideas", IdeasResource{})
		app.GET("/allideas", IdeasViewall)
		app.PUT("/ideas/{idea_id}/upvote", UpvoteIdea)
		app.PUT("/ideas/{idea_id}/downvote", DownvoteIdea)

		app.GET("/game", GameHandler)
		
		legal := app.Group("/legal")
		legal.GET("/dataprotection", LegalDataProtection)
		legal.GET("/contactinformation", LegalContactInformation)
		
		app.GET("/history", HistoryView)

		app.Resource("/prequestionnaires", PrequestionnairesResource{})
		app.Resource("/postquestionnaires", PostquestionnairesResource{})

		app.POST("/statistics", StatisticsHandler)
		app.Middleware.Skip(csrf, StatisticsHandler)

		auth.Middleware.Skip(Authorize, AuthLanding, AuthNew, AuthCreate)
		
		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("app:locales", "../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
