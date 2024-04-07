package docs_routing_suspense

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Optimizing: Suspense | Temporary</title>",
	`<meta name="description" content="Suspense allows you to create fallback UI for specific route segments, and automatically load content as it becomes ready">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/routing/suspense">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Optimizing: Suspense | Temporary">`,
	`<meta property="og:description" content="Suspense allows you to create fallback UI for specific route segments, and automatically load content as it becomes ready.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/routing/suspense">`,
	`<meta name="twitter:title" content="Optimizing: Suspense | Temporary">`,
	`<meta name="twitter:description" content="Suspense allows you to create fallback UI for specific route segments, and automatically load content as it becomes ready.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/suspense.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
