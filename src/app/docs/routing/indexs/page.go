package docs_index

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Routing: Index | Temporary</title>",
	`<meta name="description" content="Create your first index.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/indexs">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Routing: Index | Temporary">`,
	`<meta property="og:description" content="Create your first index.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/indexs">`,
	`<meta name="twitter:title" content="Routing: Index | Temporary">`,
	`<meta name="twitter:description" content="Create your first index.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/index.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
