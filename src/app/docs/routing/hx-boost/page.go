package docs_routing_hxboost

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Routing: HX-BOOST | Temporary</title>",
	`<meta name="description" content="Understand how Temporary overrides HX-BOOST.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/hx-boost">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Routing: HX-BOOST | Temporary">`,
	`<meta property="og:description" content="Understand how Temporary overrides HX-BOOST.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/hx-boost">`,
	`<meta name="twitter:title" content="Routing: HX-BOOST | Temporary">`,
	`<meta name="twitter:description" content="Understand how Temporary overrides HX-BOOST.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/hx_boost.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
