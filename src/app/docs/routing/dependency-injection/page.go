package docs_routing_dependencyinjecton

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Routing: Dependency Injection | Temporary</title>",
	`<meta name="description" content="Learn how to inject dependencies that are accessible in your handlers.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/dependency-injection">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Routing: Dependency Injection | Temporary">`,
	`<meta property="og:description" content="Learn how to inject dependencies that are accessible in your handlers.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/dependency-injection">`,
	`<meta name="twitter:title" content="Routing: Dependency Injection | Temporary">`,
	`<meta name="twitter:description" content="Learn how to inject dependencies that are accessible in your handlers.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/dependency_injection.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
