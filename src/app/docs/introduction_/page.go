package docs_introduction

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Getting Started: Introduction | Temporary</title>",
	`<meta name="description" content="Welcome to the Temporary docs.">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Getting Started: Introduction | Temporary">`,
	`<meta property="og:description" content="Welcome to the Temporary docs.">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs">`,
	`<meta name="twitter:title" content="Getting Started: Introduction | Temporary">`,
	`<meta name="twitter:description" content="Welcome to the Temporary docs.">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/introduction.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
