package docs_optimizing_caching

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Optimizing: Caching | Temporary</title>",
	`<meta name="description" content="Understand how Temporary caches your responses.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/optimizing/caching">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Optimizing: Caching | Temporary">`,
	`<meta property="og:description" content="Understand how Temporary caches your responses.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/optimizing/caching">`,
	`<meta name="twitter:title" content="Optimizing: Caching | Temporary">`,
	`<meta name="twitter:description" content="Understand how Temporary caches your responses.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/caching.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
