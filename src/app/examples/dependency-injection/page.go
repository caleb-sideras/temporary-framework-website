package examples_dependencyinjection

import (
	"net/http"

	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Page(w http.ResponseWriter, r *http.Request, conf utils.Config) templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Dependency Injection",
		ProjectURL:  "",
		ReadMeURL:   "/examples/dependency-injection/example",
		VideoURL:    "/examples/dependency-injection/code",
		InitialBody: Example(w, r, conf),
	})
}
