package docs_routing_templ

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Routing: Templ | Temporary</title>",
	`<meta name="description" content="Understand how Temporary uses Templ.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/templ">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Routing: Templ | Temporary">`,
	`<meta property="og:description" content="Understand how Temporary uses Templ.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/templ">`,
	`<meta name="twitter:title" content="Routing: Templ | Temporary">`,
	`<meta name="twitter:description" content="Understand how Temporary uses Templ.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/templ.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
