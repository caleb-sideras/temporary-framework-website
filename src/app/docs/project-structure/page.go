package docs_projectstructure

import (
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Getting Started: Project Structure | Temporary</title>",
	`<meta name="description" content="A list of folders and files conventions in a Temporary project">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/docs/project-structure">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Getting Started: Project Structure | Temporary">`,
	`<meta property="og:description" content="A list of folders and files conventions in a Temporary project">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/docs/project-structure">`,
	`<meta name="twitter:title" content="Getting Started: Project Structure | Temporary">`,
	`<meta name="twitter:description" content="A list of folders and files conventions in a Temporary project">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/docs/_markdown/project_structure.md")

	if err != nil {
		panic(err)
	}

	return newTempl
}
