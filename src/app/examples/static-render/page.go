package examples_staticrender

import (
	"calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Temporary - Examples - Static Render</title>",
	`<meta name="description" content="Shows Rendered Markdown">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/examples/static-render">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Temporary - Examples - Static Render">`,
	`<meta property="og:description" content="Shows Rendered Markdown">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/examples/static-render">`,
	`<meta name="twitter:title" content="Temporary - Examples - Static Render">`,
	`<meta name="twitter:description" content="Shows Rendered Markdown">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Static Render",
		ProjectURL:  "/docs/routing/static-pages-and-routes",
		ReadMeURL:   "/examples/static-render/example",
		VideoURL:    "/examples/static-render/code",
		InitialBody: Example_(),
	})
}
