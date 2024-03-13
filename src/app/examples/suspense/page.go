package examples_suspense

import (
	"log"

	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Page(dep interface{}) templ.Component {
	boss := dep.(utils.Hello)
	log.Println("MRBEAST:", boss)

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Suspense",
		ProjectURL:  "/docs/routing/suspense",
		ReadMeURL:   "/examples/suspense/example",
		VideoURL:    "/examples/suspense/code",
		InitialBody: server.Suspense(),
	})
}
