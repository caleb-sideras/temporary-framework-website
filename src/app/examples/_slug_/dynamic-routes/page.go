package examples_dynamicroutes

import (
	"log"
	"net/http"

	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Page(w http.ResponseWriter, r *http.Request, dep utils.Hello3) templ.Component {
	log.Println(dep.Hello2.Name)

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Dynamic Routes",
		ProjectURL:  "/docs/routing/dynamic-routes",
		ReadMeURL:   "/examples/_replace-me_/dynamic-routes/example",
		VideoURL:    "/examples/_replace-me_/dynamic-routes/code",
		InitialBody: Example(w, r),
	})
}
