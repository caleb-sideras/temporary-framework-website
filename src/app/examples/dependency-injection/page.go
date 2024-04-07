package examples_dependencyinjection

import (
	"net/http"

	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Temporary - Examples - Dependency Injection</title>",
	`<meta name="description" content="DB injection with viewer counter">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/examples/dependency-injection">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Temporary - Examples - Dependency Injection">`,
	`<meta property="og:description" content="DB injection with viewer counter">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/examples/dependency-injection">`,
	`<meta name="twitter:title" content="Temporary - Examples - Dependency Injection">`,
	`<meta name="twitter:description" content="DB injection with viewer counter">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page(w http.ResponseWriter, r *http.Request, conf utils.Config) templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Dependency Injection",
		ProjectURL:  "/docs/routing/dependency-injection",
		ReadMeURL:   "/examples/dependency-injection/example",
		VideoURL:    "/examples/dependency-injection/code",
		InitialBody: Example(w, r, conf),
	})
}
