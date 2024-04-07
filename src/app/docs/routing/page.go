package docs_routing

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Routing: Fundamentals | Temporary</title>",
	`<meta name="description" content="Learn the fundamentals of routing.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Routing: Fundamentals | Temporary">`,
	`<meta property="og:description" content="Learn the fundamentals of routing.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing">`,
	`<meta name="twitter:title" content="Routing: Fundamentals | Temporary">`,
	`<meta name="twitter:description" content="Learn the fundamentals of routing.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/routing.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
