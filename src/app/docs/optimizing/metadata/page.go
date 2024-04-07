package docs_routing_metadata

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Optimizing: Metadata | Temporary</title>",
	`<meta name="description" content="Use the Metadata API to define metadata in any index or page.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/metadata">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Optimizing: Metadata | Temporary">`,
	`<meta property="og:description" content="Use the Metadata API to define metadata in any index or page.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/metadata">`,
	`<meta name="twitter:title" content="Optimizing: Metadata | Temporary">`,
	`<meta name="twitter:description" content="Use the Metadata API to define metadata in any index or page.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/metadata.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
