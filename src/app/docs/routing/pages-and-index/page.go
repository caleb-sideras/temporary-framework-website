package docs_routing_pagesandindex

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Routing: Pages and Indexs | Temporary</title>",
	`<meta name="description" content="Create your first page and shared index.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/pages-and-index">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Routing: Pages and Indexs | Temporary">`,
	`<meta property="og:description" content="Create your first page and shared index.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/pages-and-index">`,
	`<meta name="twitter:title" content="Routing: Pages and Indexs | Temporary">`,
	`<meta name="twitter:description" content="Create your first page and shared index.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/pages_and_index.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
