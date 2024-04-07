package docs_pages

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Routing: Page | Temporary</title>",
	`<meta name="description" content="Create your first page.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/pages">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Routing: Page | Temporary">`,
	`<meta property="og:description" content="Create your first page.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/pages">`,
	`<meta name="twitter:title" content="Routing: Page | Temporary">`,
	`<meta name="twitter:description" content="Create your first page.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/pages.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
