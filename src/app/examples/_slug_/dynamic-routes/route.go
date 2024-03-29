package examples_dynamicroutes

import (
	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"net/http"
)

func Example(w http.ResponseWriter, r *http.Request) templ.Component {

	vars := mux.Vars(r)
	slug := vars["slug"]

	return server.Slug(slug)
}

func Code_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/examples/_markdown/slug.md")
	if err != nil {
		panic(err)
	}

	return newTempl
}
