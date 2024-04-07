package docs_routing_staticpagesandroutes

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Optimizing: Rendering | Temporary</title>",
	`<meta name="description" content="Leverage Rendering to Boost Performance.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/optimizing/rendering">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Optimizing: Rendering | Temporary">`,
	`<meta property="og:description" content="Leverage Rendering to Boost Performance.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/optimizing/rendering">`,
	`<meta name="twitter:title" content="Optimizing: Rendering | Temporary">`,
	`<meta name="twitter:description" content="Leverage Rendering to Boost Performance.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/rendering.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
