package docs_routing_dynamicrouting

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Routing: Dynamic Routes | Temporary</title>",
	`<meta name="description" content="Dynamic Routes can be used to programmatically generate route segments from dynamic data.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/dynamic-routes">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Routing: Dynamic Routes | Temporary">`,
	`<meta property="og:description" content="Dynamic Routes can be used to programmatically generate route segments from dynamic data.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/dynamic-routes">`,
	`<meta name="twitter:title" content="Routing: Dynamic Routes | Temporary">`,
	`<meta name="twitter:description" content="Dynamic Routes can be used to programmatically generate route segments from dynamic data.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/dynamic_routes.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
