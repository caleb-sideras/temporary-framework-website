package examples_suspense

import (
	"calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Temporary - Examples - Suspense</title>",
	`<meta name="description" content="Lazy Loads a Random Number">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/examples/supense">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Temporary - Examples - Suspense">`,
	`<meta property="og:description" content="Lazy Loads a Random Number">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/examples/supense">`,
	`<meta name="twitter:title" content="Temporary - Examples - Suspense">`,
	`<meta name="twitter:description" content="Lazy Loads a Random Number">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page() templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Suspense",
		ProjectURL:  "/docs/routing/suspense",
		ReadMeURL:   "/examples/suspense/example",
		VideoURL:    "/examples/suspense/code",
		InitialBody: server.Suspense(),
	})
}
