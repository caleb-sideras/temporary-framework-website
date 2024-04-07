package examples_todo

import (
	"calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

var Metadata []string = []string{

	// HTML Tags
	"<title>Examples: Todo | Temporary</title>",
	`<meta name="description" content="Simple Todo List">`,

	// OG Tags
	`<meta property="og:url" content="https://temporary-framework.org/examples/todo">`,
	`<meta property="og:type" content="website">`,
	`<meta property="og:title" content="Examples: Todo | Temporary">`,
	`<meta property="og:description" content="Simple Todo List">`,
	`<meta property="og:image" content="https://temporary-framework.org/static/assets/temporary.png">`,

	// Twitter Tags
	`<meta name="twitter:card" content="summary_large_image">`,
	`<meta property="twitter:domain" content="temporary-framework.org">`,
	`<meta property="twitter:url" content="https://temporary-framework.org/examples/todo">`,
	`<meta name="twitter:title" content="Examples: Todo | Temporary">`,
	`<meta name="twitter:description" content="Simple Todo List">`,
	`<meta name="twitter:image" content="https://temporary-framework.org/static/assets/temporary.png">`,
}

func Page_() templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Todo",
		ProjectURL:  "",
		ReadMeURL:   "/examples/todo/example",
		VideoURL:    "/examples/todo/code",
		InitialBody: server.Todo(),
	})
}
