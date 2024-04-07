package docs_optimizing_lazyloading

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Optimizing: Lazy Loading | Temporary</title>",
	`<meta name="description" content="Lazy load imported libraries and custom Javascript to improve your application's loading performance.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/optimizing/lazy-loading">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Optimizing: Lazy Loading | Temporary">`,
	`<meta property="og:description" content="Lazy load imported libraries and custom Javascript to improve your application's loading performance.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/optimizing/lazy-loading">`,
	`<meta name="twitter:title" content="Optimizing: Lazy Loading | Temporary">`,
	`<meta name="twitter:description" content="Lazy load imported libraries and custom Javascript to improve your application's loading performance.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/lazy_loading.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
