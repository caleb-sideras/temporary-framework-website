package examples_dynamicroutes

import (
	"net/http"

	"calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Temporary - Examples - Dynamic Routes</title>",
	`<meta name="description" content="Illutrates Dynamic Routes">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/examples/dynamic-routes">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Temporary - Examples - Dynamic Routes">`,
	`<meta property="og:description" content="Illutrates Dynamic Routes">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/examples/dynamic-routes">`,
	`<meta name="twitter:title" content="Temporary - Examples - Dynamic Routes">`,
	`<meta name="twitter:description" content="Illutrates Dynamic Routes">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page(w http.ResponseWriter, r *http.Request) templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Dynamic Routes",
		ProjectURL:  "/docs/routing/dynamic-routes",
		ReadMeURL:   "/examples/_replace-me_/dynamic-routes/example",
		VideoURL:    "/examples/_replace-me_/dynamic-routes/code",
		InitialBody: Example(w, r),
	})
}
