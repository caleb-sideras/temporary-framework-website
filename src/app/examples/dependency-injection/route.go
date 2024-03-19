package examples_dependencyinjection

import (
	"log"
	"net/http"
	"strconv"

	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Example(w http.ResponseWriter, r *http.Request, config utils.Config) templ.Component {

	newValue, err := config.Datastore.IncrementCounter(r.Context(), "counter")
	if err != nil {
		log.Printf("Error incrementing counter: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	return server.Counter(config.AppConfig, strconv.FormatInt(newValue, 10))
}

func Code_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/examples/_markdown/dependency_injection.md")
	if err != nil {
		panic(err)
	}

	return newTempl
}
