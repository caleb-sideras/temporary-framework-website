package docs_installation

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Getting Started: Installation | Temporary</title>",
	`<meta name="description" content="Create a new Temporary Project.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/installation">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Getting Started: Installation | Temporary">`,
	`<meta property="og:description" content="Create a new Temporary Project.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/installation">`,
	`<meta name="twitter:title" content="Getting Started: Installation | Temporary">`,
	`<meta name="twitter:description" content="Create a new Temporary Project.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/installation.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
